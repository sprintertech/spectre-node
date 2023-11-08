// The Licensed Work is (c) 2023 Sygma
// SPDX-License-Identifier: LGPL-3.0-only

package handlers

import (
	"context"
	"math/big"

	"github.com/attestantio/go-eth2-client/api"
	apiv1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/rs/zerolog/log"
	evmMessage "github.com/sygmaprotocol/spectre-node/chains/evm/message"
	"github.com/sygmaprotocol/sygma-core/relayer/message"
)

type SyncCommitteeFetcher interface {
	SyncCommittee(ctx context.Context, opts *api.SyncCommitteeOpts) (*api.Response[*apiv1.SyncCommittee], error)
}

type RotateHandler struct {
	domainID uint8
	domains  []uint8
	msgChan  chan []*message.Message

	prover Prover

	syncCommitteeFetcher SyncCommitteeFetcher
	currentSyncCommittee *api.Response[*apiv1.SyncCommittee]
}

func NewRotateHandler(domainID uint8, domains []uint8, msgChan chan []*message.Message, syncCommitteeFetcher SyncCommitteeFetcher, prover Prover) *RotateHandler {
	return &RotateHandler{
		syncCommitteeFetcher: syncCommitteeFetcher,
		prover:               prover,
		domainID:             domainID,
		domains:              domains,
		msgChan:              msgChan,
	}
}

// HandleEvents checks if there is a new sync committee and sends a rotate message
// if there is
func (h *RotateHandler) HandleEvents(startBlock *big.Int, endBlock *big.Int) error {
	syncCommittee, err := h.syncCommitteeFetcher.SyncCommittee(context.Background(), &api.SyncCommitteeOpts{
		State: "finalized",
	})
	if err != nil {
		return err
	}
	if syncCommittee.Data.String() == h.currentSyncCommittee.Data.String() {
		return nil
	}

	log.Info().Uint8("domainID", h.domainID).Msgf("Rotating committee")

	stepProof, err := h.prover.StepProof(endBlock)
	if err != nil {
		return err
	}
	rotateProof, err := h.prover.RotateProof(endBlock)
	if err != nil {
		return err
	}

	for _, domain := range h.domains {
		log.Debug().Uint8("domainID", h.domainID).Msgf("Sending rotate message to domain %d", domain)
		h.msgChan <- []*message.Message{
			evmMessage.NewEvmRotateMessage(h.domainID, domain, evmMessage.RotateData{
				RotateInput: evmMessage.RotateInput{},
				RotateProof: rotateProof,
				StepProof:   stepProof,
				StepInput:   evmMessage.SyncStepInput{},
			}),
		}
	}

	return nil
}
