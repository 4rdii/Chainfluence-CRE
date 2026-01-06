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
	NATIVEETHADDRESS     func() (common.Address, error)
	CampaignCounter      func() (*big.Int, error)
	Campaigns            func(CampaignsInput) (CampaignsOutput, error)
	ExpectedAuthor       func() (common.Address, error)
	ExpectedWorkflowId   func() ([32]byte, error)
	ExpectedWorkflowName func() ([10]byte, error)
	ForwarderAddress     func() (common.Address, error)
	GetCampaign          func(GetCampaignInput) (AdEscrowCampaign, error)
	IsExpired            func(IsExpiredInput) (bool, error)
	Owner                func() (common.Address, error)
	WhitelistedTokens    func(WhitelistedTokensInput) (bool, error)
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
		string(abi.Methods["NATIVE_ETH_ADDRESS"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.NATIVEETHADDRESS == nil {
				return nil, errors.New("NATIVE_ETH_ADDRESS method not mocked")
			}
			result, err := mock.NATIVEETHADDRESS()
			if err != nil {
				return nil, err
			}
			return abi.Methods["NATIVE_ETH_ADDRESS"].Outputs.Pack(result)
		},
		string(abi.Methods["campaignCounter"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.CampaignCounter == nil {
				return nil, errors.New("campaignCounter method not mocked")
			}
			result, err := mock.CampaignCounter()
			if err != nil {
				return nil, err
			}
			return abi.Methods["campaignCounter"].Outputs.Pack(result)
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
				result.Token,
				result.Amount,
				result.ContentText,
				result.MinViews,
				result.CampaignDuration,
				result.Deadline,
				result.State,
			)
		},
		string(abi.Methods["expectedAuthor"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.ExpectedAuthor == nil {
				return nil, errors.New("expectedAuthor method not mocked")
			}
			result, err := mock.ExpectedAuthor()
			if err != nil {
				return nil, err
			}
			return abi.Methods["expectedAuthor"].Outputs.Pack(result)
		},
		string(abi.Methods["expectedWorkflowId"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.ExpectedWorkflowId == nil {
				return nil, errors.New("expectedWorkflowId method not mocked")
			}
			result, err := mock.ExpectedWorkflowId()
			if err != nil {
				return nil, err
			}
			return abi.Methods["expectedWorkflowId"].Outputs.Pack(result)
		},
		string(abi.Methods["expectedWorkflowName"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.ExpectedWorkflowName == nil {
				return nil, errors.New("expectedWorkflowName method not mocked")
			}
			result, err := mock.ExpectedWorkflowName()
			if err != nil {
				return nil, err
			}
			return abi.Methods["expectedWorkflowName"].Outputs.Pack(result)
		},
		string(abi.Methods["forwarderAddress"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.ForwarderAddress == nil {
				return nil, errors.New("forwarderAddress method not mocked")
			}
			result, err := mock.ForwarderAddress()
			if err != nil {
				return nil, err
			}
			return abi.Methods["forwarderAddress"].Outputs.Pack(result)
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
		string(abi.Methods["isExpired"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.IsExpired == nil {
				return nil, errors.New("isExpired method not mocked")
			}
			inputs := abi.Methods["isExpired"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := IsExpiredInput{
				CampaignId: values[0].(*big.Int),
			}

			result, err := mock.IsExpired(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["isExpired"].Outputs.Pack(result)
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
		string(abi.Methods["whitelistedTokens"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.WhitelistedTokens == nil {
				return nil, errors.New("whitelistedTokens method not mocked")
			}
			inputs := abi.Methods["whitelistedTokens"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := WhitelistedTokensInput{
				Arg0: values[0].(common.Address),
			}

			result, err := mock.WhitelistedTokens(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["whitelistedTokens"].Outputs.Pack(result)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}
