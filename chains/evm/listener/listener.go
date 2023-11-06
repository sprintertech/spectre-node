// The Licensed Work is (c) 2023 Sygma
// SPDX-License-Identifier: LGPL-3.0-only

package listener

import (
	"context"
	"math/big"
	"time"

	"github.com/attestantio/go-eth2-client/api"
	apiv1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type EventHandler interface {
	HandleEvents(startBlock *big.Int, endBlock *big.Int) error
}

type BeaconProvider interface {
	Finality(ctx context.Context, opts *api.FinalityOpts) (*api.Response[*apiv1.Finality], error)
	SignedBeaconBlock(ctx context.Context, opts *api.SignedBeaconBlockOpts) (*api.Response[*spec.VersionedSignedBeaconBlock], error)
}

type BlockStorer interface {
	StoreBlock(epoch *big.Int, domainID uint8) error
}

type EVMListener struct {
	beaconProvider BeaconProvider

	eventHandlers []EventHandler

	blockstore BlockStorer

	domainID      uint8
	retryInterval time.Duration
	blockInterval *big.Int

	log zerolog.Logger
}

// NewEVMListener creates an EVMListener that listens to deposit events on chain
// and calls event handler when one occurs
func NewEVMListener(
	beaconProvider BeaconProvider,
	eventHandlers []EventHandler,
	domainID uint8,
	retryInterval time.Duration,
	blockInterval *big.Int) *EVMListener {
	logger := log.With().Uint8("domainID", domainID).Logger()
	return &EVMListener{
		log:            logger,
		beaconProvider: beaconProvider,
		eventHandlers:  eventHandlers,
		domainID:       domainID,
		retryInterval:  retryInterval,
		blockInterval:  blockInterval,
	}
}

// ListenToEvents waits for new finality checkpoints and calls event handlers
// with the finalized epoch block range
func (l *EVMListener) ListenToEvents(ctx context.Context, epoch *big.Int) {
	latestCheckpoint := "0x0000000000000000000000000000000000000000000000000000000000000000"
loop:
	for {
		select {
		case <-ctx.Done():
			return
		default:
			finalityCheckpoint, err := l.beaconProvider.Finality(ctx, &api.FinalityOpts{
				State: "finalized",
			})
			if err != nil {
				l.log.Warn().Err(err).Msgf("Unable to fetch finalized checkpoint")
				time.Sleep(l.retryInterval)
				continue
			}
			if finalityCheckpoint.Data.Finalized.Root.String() == latestCheckpoint {
				time.Sleep(l.retryInterval)
				continue
			}

			justifiedRoot, err := l.beaconProvider.SignedBeaconBlock(ctx, &api.SignedBeaconBlockOpts{
				Block: finalityCheckpoint.Data.Justified.Root.String(),
			})
			if err != nil {
				l.log.Warn().Err(err).Msgf("Unable to fetch justified root")
				time.Sleep(l.retryInterval)
				continue
			}
			endBlock := big.NewInt(int64(justifiedRoot.Data.Capella.Message.Body.ExecutionPayload.BlockNumber))
			startBlock := new(big.Int).Sub(endBlock, big.NewInt(l.blockInterval.Int64()))

			l.log.Debug().Msgf("Fetching evm events for block range %s-%s", startBlock, endBlock)

			for _, handler := range l.eventHandlers {
				err := handler.HandleEvents(startBlock, endBlock)
				if err != nil {
					l.log.Warn().Err(err).Msgf("Unable to handle events")
					continue loop
				}
			}
			l.log.Debug().Msgf("Handled events for block range %s-%s", startBlock, endBlock)

			latestCheckpoint = finalityCheckpoint.Data.Finalized.Root.String()
		}
	}
}