// Code generated — DO NOT EDIT.

//go:build !wasip1

package escrow_v2

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

// EscrowV2Mock is a mock implementation of EscrowV2 for testing.
type EscrowV2Mock struct {
	MAXFEEBPS            func() (*big.Int, error)
	NATIVEETHADDRESS     func() (common.Address, error)
	Deals                func(DealsInput) (DealsOutput, error)
	ExpectedAuthor       func() (common.Address, error)
	ExpectedWorkflowId   func() ([32]byte, error)
	ExpectedWorkflowName func() ([10]byte, error)
	FeeRecipient         func() (common.Address, error)
	ForwarderAddress     func() (common.Address, error)
	GetDeal              func(GetDealInput) (AdEscrowV2Campaign, error)
	IsExpired            func(IsExpiredInput) (bool, error)
	Owner                func() (common.Address, error)
	PlatformFeeBps       func() (*big.Int, error)
	WhitelistedTokens    func(WhitelistedTokensInput) (bool, error)
}

// NewEscrowV2Mock creates a new EscrowV2Mock for testing.
func NewEscrowV2Mock(address common.Address, clientMock *evmmock.ClientCapability) *EscrowV2Mock {
	mock := &EscrowV2Mock{}

	codec, err := NewCodec()
	if err != nil {
		panic("failed to create codec for mock: " + err.Error())
	}

	abi := codec.(*Codec).abi
	_ = abi

	funcMap := map[string]func([]byte) ([]byte, error){
		string(abi.Methods["MAX_FEE_BPS"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.MAXFEEBPS == nil {
				return nil, errors.New("MAX_FEE_BPS method not mocked")
			}
			result, err := mock.MAXFEEBPS()
			if err != nil {
				return nil, err
			}
			return abi.Methods["MAX_FEE_BPS"].Outputs.Pack(result)
		},
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
		string(abi.Methods["deals"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Deals == nil {
				return nil, errors.New("deals method not mocked")
			}
			inputs := abi.Methods["deals"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := DealsInput{
				Arg0: values[0].(*big.Int),
			}

			result, err := mock.Deals(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["deals"].Outputs.Pack(
				result.Advertiser,
				result.Influencer,
				result.Token,
				result.Amount,
				result.ContentHash,
				result.MinViews,
				result.CampaignDuration,
				result.Deadline,
				result.State,
				result.PlatformFee,
				result.InfluencerAccepted,
				result.ChannelName,
				result.TweetUrl,
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
		string(abi.Methods["feeRecipient"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.FeeRecipient == nil {
				return nil, errors.New("feeRecipient method not mocked")
			}
			result, err := mock.FeeRecipient()
			if err != nil {
				return nil, err
			}
			return abi.Methods["feeRecipient"].Outputs.Pack(result)
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
		string(abi.Methods["getDeal"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetDeal == nil {
				return nil, errors.New("getDeal method not mocked")
			}
			inputs := abi.Methods["getDeal"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetDealInput{
				DealId: values[0].(*big.Int),
			}

			result, err := mock.GetDeal(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getDeal"].Outputs.Pack(result)
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
				DealId: values[0].(*big.Int),
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
		string(abi.Methods["platformFeeBps"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.PlatformFeeBps == nil {
				return nil, errors.New("platformFeeBps method not mocked")
			}
			result, err := mock.PlatformFeeBps()
			if err != nil {
				return nil, err
			}
			return abi.Methods["platformFeeBps"].Outputs.Pack(result)
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
