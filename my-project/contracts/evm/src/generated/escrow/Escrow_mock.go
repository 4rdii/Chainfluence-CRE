// Code generated — DO NOT EDIT.

//go:build !wasip1

package escrow

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	evmmock "github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/mock"
)

var (
	_ = errors.New
	_ = fmt.Errorf
	_ = big.NewInt
	_ = common.Big1
)

// EscrowMock is a mock implementation of Escrow for testing.
type EscrowMock struct {
	AuthorizedWithdrawer func() (common.Address, error)
	Campaigns            func(CampaignsInput) (CampaignsOutput, error)
	GetCampaign          func(GetCampaignInput) (AdEscrowCampaign, error)
	GetNextCampaignId    func() (*big.Int, error)
	Owner                func() (common.Address, error)
}

// NewEscrowMock creates a new EscrowMock for testing.
func NewEscrowMock(address common.Address, clientMock *evmmock.ClientCapability) *EscrowMock {
	mock := &EscrowMock{}

	codec, err := NewCodec()
	if err != nil {
		panic("failed to create codec for mock: " + err.Error())
	}

	abi := codec.(*Codec).abi
	_ = abi

	funcMap := map[string]func([]byte) ([]byte, error){
		string(abi.Methods["authorizedWithdrawer"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.AuthorizedWithdrawer == nil {
				return nil, errors.New("authorizedWithdrawer method not mocked")
			}
			result, err := mock.AuthorizedWithdrawer()
			if err != nil {
				return nil, err
			}
			return abi.Methods["authorizedWithdrawer"].Outputs.Pack(result)
		},
		string(abi.Methods["campaigns"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Campaigns == nil {
				return nil, errors.New("campaigns method not mocked")
			}
			inputs := abi.Methods["campaigns"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := CampaignsInput{
				Arg0: values[0].(*big.Int),
			}

			result, err := mock.Campaigns(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["campaigns"].Outputs.Pack(
				result.Advertiser,
				result.Influencer,
				result.Amount,
				result.ContentText,
				result.MinViews,
				result.Deadline,
				result.Fulfilled,
				result.Withdrawn,
			)
		},
		string(abi.Methods["getCampaign"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetCampaign == nil {
				return nil, errors.New("getCampaign method not mocked")
			}
			inputs := abi.Methods["getCampaign"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetCampaignInput{
				CampaignId: values[0].(*big.Int),
			}

			result, err := mock.GetCampaign(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getCampaign"].Outputs.Pack(result)
		},
		string(abi.Methods["getNextCampaignId"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetNextCampaignId == nil {
				return nil, errors.New("getNextCampaignId method not mocked")
			}
			result, err := mock.GetNextCampaignId()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getNextCampaignId"].Outputs.Pack(result)
		},
		string(abi.Methods["owner"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Owner == nil {
				return nil, errors.New("owner method not mocked")
			}
			result, err := mock.Owner()
			if err != nil {
				return nil, err
			}
			return abi.Methods["owner"].Outputs.Pack(result)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}
