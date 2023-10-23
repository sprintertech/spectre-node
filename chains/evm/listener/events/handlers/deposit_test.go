// The Licensed Work is (c) 2023 Sygma
// SPDX-License-Identifier: LGPL-3.0-only

package handlers_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/suite"
	"github.com/sygmaprotocol/spectre-node/chains/evm/listener/events"
	"github.com/sygmaprotocol/spectre-node/chains/evm/listener/events/handlers"
	"github.com/sygmaprotocol/spectre-node/mock"
	"github.com/sygmaprotocol/sygma-core/relayer/message"
	"go.uber.org/mock/gomock"
)

func readFromChannel(msgChan chan []*message.Message) ([]*message.Message, error) {
	select {
	case msgs := <-msgChan:
		return msgs, nil
	default:
		return make([]*message.Message, 0), fmt.Errorf("no message sent")
	}
}

func SliceTo32Bytes(in []byte) [32]byte {
	var res [32]byte
	copy(res[:], in)
	return res
}

type DepositHandlerTestSuite struct {
	suite.Suite

	depositHandler *handlers.DepositEventHandler

	msgChan          chan []*message.Message
	mockEventFetcher *mock.MockEventFetcher
}

func TestRunConfigTestSuite(t *testing.T) {
	suite.Run(t, new(DepositHandlerTestSuite))
}

func (s *DepositHandlerTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.mockEventFetcher = mock.NewMockEventFetcher(ctrl)
	s.msgChan = make(chan []*message.Message, 1)
	s.depositHandler = handlers.NewDepositEventHandler(
		s.msgChan,
		s.mockEventFetcher,
		common.HexToAddress("0xb0b13f0109ef097C3Aa70Fb543EA4942114A845d"))
}

func (s *DepositHandlerTestSuite) Test_HandleEvents_FetchingDepositsFails() {
	startBlock := big.NewInt(0)
	endBlock := big.NewInt(4)
	s.mockEventFetcher.EXPECT().FetchEventLogs(
		context.Background(),
		gomock.Any(),
		string(events.DepositSig),
		startBlock,
		endBlock,
	).Return(nil, fmt.Errorf("Error"))

	err := s.depositHandler.HandleEvents(startBlock, endBlock)
	s.NotNil(err)

	_, err = readFromChannel(s.msgChan)
	s.NotNil(err)
}

func (s *DepositHandlerTestSuite) Test_HandleEvents_NoEvents_MessageNotSent() {
	startBlock := big.NewInt(0)
	endBlock := big.NewInt(4)
	s.mockEventFetcher.EXPECT().FetchEventLogs(
		context.Background(),
		gomock.Any(),
		string(events.DepositSig),
		startBlock,
		endBlock,
	).Return(make([]types.Log, 0), nil)

	err := s.depositHandler.HandleEvents(startBlock, endBlock)
	s.Nil(err)

	_, err = readFromChannel(s.msgChan)
	s.NotNil(err)
}

func (s *DepositHandlerTestSuite) Test_HandleEvents_ValidDeposit() {
	validDepositData, _ := hex.DecodeString("000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000")
	invalidDepositData := []byte("invalid")

	startBlock := big.NewInt(0)
	endBlock := big.NewInt(4)
	s.mockEventFetcher.EXPECT().FetchEventLogs(
		context.Background(),
		gomock.Any(),
		string(events.DepositSig),
		startBlock,
		endBlock,
	).Return([]types.Log{
		{
			Data: invalidDepositData,
		},
		{
			Data: validDepositData,
			Topics: []common.Hash{
				{},
				common.HexToHash("0xd68eb9b5E135b96c1Af165e1D8c4e2eB0E1CE4CD"),
			},
		},
	}, nil)

	err := s.depositHandler.HandleEvents(startBlock, endBlock)
	s.Nil(err)

	msgs, err := readFromChannel(s.msgChan)
	s.Nil(err)
	s.Equal(len(msgs), 1)
}