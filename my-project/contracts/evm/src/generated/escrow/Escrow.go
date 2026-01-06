// Code generated — DO NOT EDIT.

package escrow

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb2 "github.com/smartcontractkit/chainlink-protos/cre/go/sdk"
	"github.com/smartcontractkit/chainlink-protos/cre/go/values/pb"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Sprintf
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = emptypb.Empty{}
	_ = pb.NewBigIntFromInt
	_ = pb2.AggregationType_AGGREGATION_TYPE_COMMON_PREFIX
	_ = bindings.FilterOptions{}
	_ = evm.FilterLogTriggerRequest{}
	_ = cre.ResponseBufferTooSmall
	_ = rpc.API{}
	_ = json.Unmarshal
	_ = reflect.Bool
)

var EscrowMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_forwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_whitelistedTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"NATIVE_ETH_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"campaignCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"campaigns\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"advertiser\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"influencer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contentText\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"minViews\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"campaignDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumAdEscrow.CampaignState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deliveryAction\",\"inputs\":[{\"name\":\"campaignId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"influencer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contentText\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"minViews\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"campaignDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"campaignId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"expectedAuthor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"expectedWorkflowId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"expectedWorkflowName\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forwarderAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCampaign\",\"inputs\":[{\"name\":\"campaignId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structAdEscrow.Campaign\",\"components\":[{\"name\":\"advertiser\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"influencer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contentText\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"minViews\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"campaignDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumAdEscrow.CampaignState\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isExpired\",\"inputs\":[{\"name\":\"campaignId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onReport\",\"inputs\":[{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeTokenFromWhitelist\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedAuthor\",\"inputs\":[{\"name\":\"_author\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedWorkflowId\",\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedWorkflowName\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setForwarderAddress\",\"inputs\":[{\"name\":\"_forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"whitelistToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"whitelistedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CampaignDeposited\",\"inputs\":[{\"name\":\"campaignId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"advertiser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"influencer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CampaignExpired\",\"inputs\":[{\"name\":\"campaignId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"advertiser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeliveryActionCalled\",\"inputs\":[{\"name\":\"campaignId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundsWithdrawn\",\"inputs\":[{\"name\":\"campaignId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"influencer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenRemovedFromWhitelist\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenWhitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidAuthor\",\"inputs\":[{\"name\":\"received\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWorkflowId\",\"inputs\":[{\"name\":\"received\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidWorkflowName\",\"inputs\":[{\"name\":\"received\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"},{\"name\":\"expected\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// Structs
type AdEscrowCampaign struct {
	Advertiser       common.Address
	Influencer       common.Address
	Token            common.Address
	Amount           *big.Int
	ContentText      string
	MinViews         *big.Int
	CampaignDuration uint64
	Deadline         *big.Int
	State            uint8
}

// Contract Method Inputs
type CampaignsInput struct {
	Arg0 *big.Int
}

type DeliveryActionInput struct {
	CampaignId *big.Int
}

type DepositInput struct {
	Token            common.Address
	Influencer       common.Address
	Amount           *big.Int
	ContentText      string
	MinViews         *big.Int
	ExpiryDeadline   *big.Int
	CampaignDuration uint64
}

type GetCampaignInput struct {
	CampaignId *big.Int
}

type IsExpiredInput struct {
	CampaignId *big.Int
}

type OnReportInput struct {
	Metadata []byte
	Report   []byte
}

type RemoveTokenFromWhitelistInput struct {
	Token common.Address
}

type SetExpectedAuthorInput struct {
	Author common.Address
}

type SetExpectedWorkflowIdInput struct {
	Id [32]byte
}

type SetExpectedWorkflowNameInput struct {
	Name string
}

type SetForwarderAddressInput struct {
	Forwarder common.Address
}

type SupportsInterfaceInput struct {
	InterfaceId [4]byte
}

type TransferOwnershipInput struct {
	NewOwner common.Address
}

type WhitelistTokenInput struct {
	Token common.Address
}

type WhitelistedTokensInput struct {
	Arg0 common.Address
}

// Contract Method Outputs
type CampaignsOutput struct {
	Advertiser       common.Address
	Influencer       common.Address
	Token            common.Address
	Amount           *big.Int
	ContentText      string
	MinViews         *big.Int
	CampaignDuration uint64
	Deadline         *big.Int
	State            uint8
}

// Errors
type InvalidAuthor struct {
	Received common.Address
	Expected common.Address
}

type InvalidSender struct {
	Sender   common.Address
	Expected common.Address
}

type InvalidWorkflowId struct {
	Received [32]byte
	Expected [32]byte
}

type InvalidWorkflowName struct {
	Received [10]byte
	Expected [10]byte
}

type OwnableInvalidOwner struct {
	Owner common.Address
}

type OwnableUnauthorizedAccount struct {
	Account common.Address
}

type SafeERC20FailedOperation struct {
	Token common.Address
}

// Events
// The <Event>Topics struct should be used as a filter (for log triggers).
// Note: It is only possible to filter on indexed fields.
// Indexed (string and bytes) fields will be of type common.Hash.
// They need to he (crypto.Keccak256) hashed and passed in.
// Indexed (tuple/slice/array) fields can be passed in as is, the Encode<Event>Topics function will handle the hashing.
//
// The <Event>Decoded struct will be the result of calling decode (Adapt) on the log trigger result.
// Indexed dynamic type fields will be of type common.Hash.

type CampaignDepositedTopics struct {
	CampaignId *big.Int
	Advertiser common.Address
	Influencer common.Address
}

type CampaignDepositedDecoded struct {
	CampaignId *big.Int
	Advertiser common.Address
	Influencer common.Address
	Token      common.Address
	Amount     *big.Int
}

type CampaignExpiredTopics struct {
	CampaignId *big.Int
	Advertiser common.Address
}

type CampaignExpiredDecoded struct {
	CampaignId *big.Int
	Advertiser common.Address
	Token      common.Address
	Amount     *big.Int
}

type DeliveryActionCalledTopics struct {
	CampaignId *big.Int
}

type DeliveryActionCalledDecoded struct {
	CampaignId *big.Int
}

type FundsWithdrawnTopics struct {
	CampaignId *big.Int
	Influencer common.Address
}

type FundsWithdrawnDecoded struct {
	CampaignId *big.Int
	Influencer common.Address
	Token      common.Address
	Amount     *big.Int
}

type OwnershipTransferredTopics struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

type OwnershipTransferredDecoded struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

type TokenRemovedFromWhitelistTopics struct {
	Token common.Address
}

type TokenRemovedFromWhitelistDecoded struct {
	Token common.Address
}

type TokenWhitelistedTopics struct {
	Token common.Address
}

type TokenWhitelistedDecoded struct {
	Token common.Address
}

// Main Binding Type for Escrow
type Escrow struct {
	Address common.Address
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   EscrowCodec
}

type EscrowCodec interface {
	EncodeNATIVEETHADDRESSMethodCall() ([]byte, error)
	DecodeNATIVEETHADDRESSMethodOutput(data []byte) (common.Address, error)
	EncodeCampaignCounterMethodCall() ([]byte, error)
	DecodeCampaignCounterMethodOutput(data []byte) (*big.Int, error)
	EncodeCampaignsMethodCall(in CampaignsInput) ([]byte, error)
	DecodeCampaignsMethodOutput(data []byte) (CampaignsOutput, error)
	EncodeDeliveryActionMethodCall(in DeliveryActionInput) ([]byte, error)
	EncodeDepositMethodCall(in DepositInput) ([]byte, error)
	DecodeDepositMethodOutput(data []byte) (*big.Int, error)
	EncodeExpectedAuthorMethodCall() ([]byte, error)
	DecodeExpectedAuthorMethodOutput(data []byte) (common.Address, error)
	EncodeExpectedWorkflowIdMethodCall() ([]byte, error)
	DecodeExpectedWorkflowIdMethodOutput(data []byte) ([32]byte, error)
	EncodeExpectedWorkflowNameMethodCall() ([]byte, error)
	DecodeExpectedWorkflowNameMethodOutput(data []byte) ([10]byte, error)
	EncodeForwarderAddressMethodCall() ([]byte, error)
	DecodeForwarderAddressMethodOutput(data []byte) (common.Address, error)
	EncodeGetCampaignMethodCall(in GetCampaignInput) ([]byte, error)
	DecodeGetCampaignMethodOutput(data []byte) (AdEscrowCampaign, error)
	EncodeIsExpiredMethodCall(in IsExpiredInput) ([]byte, error)
	DecodeIsExpiredMethodOutput(data []byte) (bool, error)
	EncodeOnReportMethodCall(in OnReportInput) ([]byte, error)
	EncodeOwnerMethodCall() ([]byte, error)
	DecodeOwnerMethodOutput(data []byte) (common.Address, error)
	EncodeRemoveTokenFromWhitelistMethodCall(in RemoveTokenFromWhitelistInput) ([]byte, error)
	EncodeRenounceOwnershipMethodCall() ([]byte, error)
	EncodeSetExpectedAuthorMethodCall(in SetExpectedAuthorInput) ([]byte, error)
	EncodeSetExpectedWorkflowIdMethodCall(in SetExpectedWorkflowIdInput) ([]byte, error)
	EncodeSetExpectedWorkflowNameMethodCall(in SetExpectedWorkflowNameInput) ([]byte, error)
	EncodeSetForwarderAddressMethodCall(in SetForwarderAddressInput) ([]byte, error)
	EncodeSupportsInterfaceMethodCall(in SupportsInterfaceInput) ([]byte, error)
	DecodeSupportsInterfaceMethodOutput(data []byte) (bool, error)
	EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error)
	EncodeWhitelistTokenMethodCall(in WhitelistTokenInput) ([]byte, error)
	EncodeWhitelistedTokensMethodCall(in WhitelistedTokensInput) ([]byte, error)
	DecodeWhitelistedTokensMethodOutput(data []byte) (bool, error)
	EncodeAdEscrowCampaignStruct(in AdEscrowCampaign) ([]byte, error)
	CampaignDepositedLogHash() []byte
	EncodeCampaignDepositedTopics(evt abi.Event, values []CampaignDepositedTopics) ([]*evm.TopicValues, error)
	DecodeCampaignDeposited(log *evm.Log) (*CampaignDepositedDecoded, error)
	CampaignExpiredLogHash() []byte
	EncodeCampaignExpiredTopics(evt abi.Event, values []CampaignExpiredTopics) ([]*evm.TopicValues, error)
	DecodeCampaignExpired(log *evm.Log) (*CampaignExpiredDecoded, error)
	DeliveryActionCalledLogHash() []byte
	EncodeDeliveryActionCalledTopics(evt abi.Event, values []DeliveryActionCalledTopics) ([]*evm.TopicValues, error)
	DecodeDeliveryActionCalled(log *evm.Log) (*DeliveryActionCalledDecoded, error)
	FundsWithdrawnLogHash() []byte
	EncodeFundsWithdrawnTopics(evt abi.Event, values []FundsWithdrawnTopics) ([]*evm.TopicValues, error)
	DecodeFundsWithdrawn(log *evm.Log) (*FundsWithdrawnDecoded, error)
	OwnershipTransferredLogHash() []byte
	EncodeOwnershipTransferredTopics(evt abi.Event, values []OwnershipTransferredTopics) ([]*evm.TopicValues, error)
	DecodeOwnershipTransferred(log *evm.Log) (*OwnershipTransferredDecoded, error)
	TokenRemovedFromWhitelistLogHash() []byte
	EncodeTokenRemovedFromWhitelistTopics(evt abi.Event, values []TokenRemovedFromWhitelistTopics) ([]*evm.TopicValues, error)
	DecodeTokenRemovedFromWhitelist(log *evm.Log) (*TokenRemovedFromWhitelistDecoded, error)
	TokenWhitelistedLogHash() []byte
	EncodeTokenWhitelistedTopics(evt abi.Event, values []TokenWhitelistedTopics) ([]*evm.TopicValues, error)
	DecodeTokenWhitelisted(log *evm.Log) (*TokenWhitelistedDecoded, error)
}

func NewEscrow(
	client *evm.Client,
	address common.Address,
	options *bindings.ContractInitOptions,
) (*Escrow, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewCodec()
	if err != nil {
		return nil, err
	}
	return &Escrow{
		Address: address,
		Options: options,
		ABI:     &parsed,
		client:  client,
		Codec:   codec,
	}, nil
}

type Codec struct {
	abi *abi.ABI
}

func NewCodec() (EscrowCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &Codec{abi: &parsed}, nil
}

func (c *Codec) EncodeNATIVEETHADDRESSMethodCall() ([]byte, error) {
	return c.abi.Pack("NATIVE_ETH_ADDRESS")
}

func (c *Codec) DecodeNATIVEETHADDRESSMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["NATIVE_ETH_ADDRESS"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeCampaignCounterMethodCall() ([]byte, error) {
	return c.abi.Pack("campaignCounter")
}

func (c *Codec) DecodeCampaignCounterMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["campaignCounter"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeCampaignsMethodCall(in CampaignsInput) ([]byte, error) {
	return c.abi.Pack("campaigns", in.Arg0)
}

func (c *Codec) DecodeCampaignsMethodOutput(data []byte) (CampaignsOutput, error) {
	vals, err := c.abi.Methods["campaigns"].Outputs.Unpack(data)
	if err != nil {
		return CampaignsOutput{}, err
	}
	if len(vals) != 9 {
		return CampaignsOutput{}, fmt.Errorf("expected 9 values, got %d", len(vals))
	}
	jsonData0, err := json.Marshal(vals[0])
	if err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to marshal ABI result 0: %w", err)
	}

	var result0 common.Address
	if err := json.Unmarshal(jsonData0, &result0); err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}
	jsonData1, err := json.Marshal(vals[1])
	if err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to marshal ABI result 1: %w", err)
	}

	var result1 common.Address
	if err := json.Unmarshal(jsonData1, &result1); err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}
	jsonData2, err := json.Marshal(vals[2])
	if err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to marshal ABI result 2: %w", err)
	}

	var result2 common.Address
	if err := json.Unmarshal(jsonData2, &result2); err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}
	jsonData3, err := json.Marshal(vals[3])
	if err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to marshal ABI result 3: %w", err)
	}

	var result3 *big.Int
	if err := json.Unmarshal(jsonData3, &result3); err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData4, err := json.Marshal(vals[4])
	if err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to marshal ABI result 4: %w", err)
	}

	var result4 string
	if err := json.Unmarshal(jsonData4, &result4); err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to unmarshal to string: %w", err)
	}
	jsonData5, err := json.Marshal(vals[5])
	if err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to marshal ABI result 5: %w", err)
	}

	var result5 *big.Int
	if err := json.Unmarshal(jsonData5, &result5); err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData6, err := json.Marshal(vals[6])
	if err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to marshal ABI result 6: %w", err)
	}

	var result6 uint64
	if err := json.Unmarshal(jsonData6, &result6); err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to unmarshal to uint64: %w", err)
	}
	jsonData7, err := json.Marshal(vals[7])
	if err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to marshal ABI result 7: %w", err)
	}

	var result7 *big.Int
	if err := json.Unmarshal(jsonData7, &result7); err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData8, err := json.Marshal(vals[8])
	if err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to marshal ABI result 8: %w", err)
	}

	var result8 uint8
	if err := json.Unmarshal(jsonData8, &result8); err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to unmarshal to uint8: %w", err)
	}

	return CampaignsOutput{
		Advertiser:       result0,
		Influencer:       result1,
		Token:            result2,
		Amount:           result3,
		ContentText:      result4,
		MinViews:         result5,
		CampaignDuration: result6,
		Deadline:         result7,
		State:            result8,
	}, nil
}

func (c *Codec) EncodeDeliveryActionMethodCall(in DeliveryActionInput) ([]byte, error) {
	return c.abi.Pack("deliveryAction", in.CampaignId)
}

func (c *Codec) EncodeDepositMethodCall(in DepositInput) ([]byte, error) {
	return c.abi.Pack("deposit", in.Token, in.Influencer, in.Amount, in.ContentText, in.MinViews, in.ExpiryDeadline, in.CampaignDuration)
}

func (c *Codec) DecodeDepositMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["deposit"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeExpectedAuthorMethodCall() ([]byte, error) {
	return c.abi.Pack("expectedAuthor")
}

func (c *Codec) DecodeExpectedAuthorMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["expectedAuthor"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeExpectedWorkflowIdMethodCall() ([]byte, error) {
	return c.abi.Pack("expectedWorkflowId")
}

func (c *Codec) DecodeExpectedWorkflowIdMethodOutput(data []byte) ([32]byte, error) {
	vals, err := c.abi.Methods["expectedWorkflowId"].Outputs.Unpack(data)
	if err != nil {
		return *new([32]byte), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new([32]byte), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result [32]byte
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new([32]byte), fmt.Errorf("failed to unmarshal to [32]byte: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeExpectedWorkflowNameMethodCall() ([]byte, error) {
	return c.abi.Pack("expectedWorkflowName")
}

func (c *Codec) DecodeExpectedWorkflowNameMethodOutput(data []byte) ([10]byte, error) {
	vals, err := c.abi.Methods["expectedWorkflowName"].Outputs.Unpack(data)
	if err != nil {
		return *new([10]byte), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new([10]byte), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result [10]byte
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new([10]byte), fmt.Errorf("failed to unmarshal to [10]byte: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeForwarderAddressMethodCall() ([]byte, error) {
	return c.abi.Pack("forwarderAddress")
}

func (c *Codec) DecodeForwarderAddressMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["forwarderAddress"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetCampaignMethodCall(in GetCampaignInput) ([]byte, error) {
	return c.abi.Pack("getCampaign", in.CampaignId)
}

func (c *Codec) DecodeGetCampaignMethodOutput(data []byte) (AdEscrowCampaign, error) {
	vals, err := c.abi.Methods["getCampaign"].Outputs.Unpack(data)
	if err != nil {
		return *new(AdEscrowCampaign), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(AdEscrowCampaign), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result AdEscrowCampaign
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(AdEscrowCampaign), fmt.Errorf("failed to unmarshal to AdEscrowCampaign: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeIsExpiredMethodCall(in IsExpiredInput) ([]byte, error) {
	return c.abi.Pack("isExpired", in.CampaignId)
}

func (c *Codec) DecodeIsExpiredMethodOutput(data []byte) (bool, error) {
	vals, err := c.abi.Methods["isExpired"].Outputs.Unpack(data)
	if err != nil {
		return *new(bool), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(bool), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result bool
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(bool), fmt.Errorf("failed to unmarshal to bool: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeOnReportMethodCall(in OnReportInput) ([]byte, error) {
	return c.abi.Pack("onReport", in.Metadata, in.Report)
}

func (c *Codec) EncodeOwnerMethodCall() ([]byte, error) {
	return c.abi.Pack("owner")
}

func (c *Codec) DecodeOwnerMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["owner"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeRemoveTokenFromWhitelistMethodCall(in RemoveTokenFromWhitelistInput) ([]byte, error) {
	return c.abi.Pack("removeTokenFromWhitelist", in.Token)
}

func (c *Codec) EncodeRenounceOwnershipMethodCall() ([]byte, error) {
	return c.abi.Pack("renounceOwnership")
}

func (c *Codec) EncodeSetExpectedAuthorMethodCall(in SetExpectedAuthorInput) ([]byte, error) {
	return c.abi.Pack("setExpectedAuthor", in.Author)
}

func (c *Codec) EncodeSetExpectedWorkflowIdMethodCall(in SetExpectedWorkflowIdInput) ([]byte, error) {
	return c.abi.Pack("setExpectedWorkflowId", in.Id)
}

func (c *Codec) EncodeSetExpectedWorkflowNameMethodCall(in SetExpectedWorkflowNameInput) ([]byte, error) {
	return c.abi.Pack("setExpectedWorkflowName", in.Name)
}

func (c *Codec) EncodeSetForwarderAddressMethodCall(in SetForwarderAddressInput) ([]byte, error) {
	return c.abi.Pack("setForwarderAddress", in.Forwarder)
}

func (c *Codec) EncodeSupportsInterfaceMethodCall(in SupportsInterfaceInput) ([]byte, error) {
	return c.abi.Pack("supportsInterface", in.InterfaceId)
}

func (c *Codec) DecodeSupportsInterfaceMethodOutput(data []byte) (bool, error) {
	vals, err := c.abi.Methods["supportsInterface"].Outputs.Unpack(data)
	if err != nil {
		return *new(bool), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(bool), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result bool
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(bool), fmt.Errorf("failed to unmarshal to bool: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error) {
	return c.abi.Pack("transferOwnership", in.NewOwner)
}

func (c *Codec) EncodeWhitelistTokenMethodCall(in WhitelistTokenInput) ([]byte, error) {
	return c.abi.Pack("whitelistToken", in.Token)
}

func (c *Codec) EncodeWhitelistedTokensMethodCall(in WhitelistedTokensInput) ([]byte, error) {
	return c.abi.Pack("whitelistedTokens", in.Arg0)
}

func (c *Codec) DecodeWhitelistedTokensMethodOutput(data []byte) (bool, error) {
	vals, err := c.abi.Methods["whitelistedTokens"].Outputs.Unpack(data)
	if err != nil {
		return *new(bool), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(bool), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result bool
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(bool), fmt.Errorf("failed to unmarshal to bool: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeAdEscrowCampaignStruct(in AdEscrowCampaign) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "advertiser", Type: "address"},
			{Name: "influencer", Type: "address"},
			{Name: "token", Type: "address"},
			{Name: "amount", Type: "uint256"},
			{Name: "contentText", Type: "string"},
			{Name: "minViews", Type: "uint256"},
			{Name: "campaignDuration", Type: "uint64"},
			{Name: "deadline", Type: "uint256"},
			{Name: "state", Type: "uint8"},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for AdEscrowCampaign: %w", err)
	}
	args := abi.Arguments{
		{Name: "adEscrowCampaign", Type: tupleType},
	}

	return args.Pack(in)
}

func (c *Codec) CampaignDepositedLogHash() []byte {
	return c.abi.Events["CampaignDeposited"].ID.Bytes()
}

func (c *Codec) EncodeCampaignDepositedTopics(
	evt abi.Event,
	values []CampaignDepositedTopics,
) ([]*evm.TopicValues, error) {
	var campaignIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.CampaignId).IsZero() {
			campaignIdRule = append(campaignIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.CampaignId)
		if err != nil {
			return nil, err
		}
		campaignIdRule = append(campaignIdRule, fieldVal)
	}
	var advertiserRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Advertiser).IsZero() {
			advertiserRule = append(advertiserRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.Advertiser)
		if err != nil {
			return nil, err
		}
		advertiserRule = append(advertiserRule, fieldVal)
	}
	var influencerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Influencer).IsZero() {
			influencerRule = append(influencerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[2], v.Influencer)
		if err != nil {
			return nil, err
		}
		influencerRule = append(influencerRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		campaignIdRule,
		advertiserRule,
		influencerRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeCampaignDeposited decodes a log into a CampaignDeposited struct.
func (c *Codec) DecodeCampaignDeposited(log *evm.Log) (*CampaignDepositedDecoded, error) {
	event := new(CampaignDepositedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "CampaignDeposited", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["CampaignDeposited"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) CampaignExpiredLogHash() []byte {
	return c.abi.Events["CampaignExpired"].ID.Bytes()
}

func (c *Codec) EncodeCampaignExpiredTopics(
	evt abi.Event,
	values []CampaignExpiredTopics,
) ([]*evm.TopicValues, error) {
	var campaignIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.CampaignId).IsZero() {
			campaignIdRule = append(campaignIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.CampaignId)
		if err != nil {
			return nil, err
		}
		campaignIdRule = append(campaignIdRule, fieldVal)
	}
	var advertiserRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Advertiser).IsZero() {
			advertiserRule = append(advertiserRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.Advertiser)
		if err != nil {
			return nil, err
		}
		advertiserRule = append(advertiserRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		campaignIdRule,
		advertiserRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeCampaignExpired decodes a log into a CampaignExpired struct.
func (c *Codec) DecodeCampaignExpired(log *evm.Log) (*CampaignExpiredDecoded, error) {
	event := new(CampaignExpiredDecoded)
	if err := c.abi.UnpackIntoInterface(event, "CampaignExpired", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["CampaignExpired"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) DeliveryActionCalledLogHash() []byte {
	return c.abi.Events["DeliveryActionCalled"].ID.Bytes()
}

func (c *Codec) EncodeDeliveryActionCalledTopics(
	evt abi.Event,
	values []DeliveryActionCalledTopics,
) ([]*evm.TopicValues, error) {
	var campaignIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.CampaignId).IsZero() {
			campaignIdRule = append(campaignIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.CampaignId)
		if err != nil {
			return nil, err
		}
		campaignIdRule = append(campaignIdRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		campaignIdRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeDeliveryActionCalled decodes a log into a DeliveryActionCalled struct.
func (c *Codec) DecodeDeliveryActionCalled(log *evm.Log) (*DeliveryActionCalledDecoded, error) {
	event := new(DeliveryActionCalledDecoded)
	if err := c.abi.UnpackIntoInterface(event, "DeliveryActionCalled", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["DeliveryActionCalled"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) FundsWithdrawnLogHash() []byte {
	return c.abi.Events["FundsWithdrawn"].ID.Bytes()
}

func (c *Codec) EncodeFundsWithdrawnTopics(
	evt abi.Event,
	values []FundsWithdrawnTopics,
) ([]*evm.TopicValues, error) {
	var campaignIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.CampaignId).IsZero() {
			campaignIdRule = append(campaignIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.CampaignId)
		if err != nil {
			return nil, err
		}
		campaignIdRule = append(campaignIdRule, fieldVal)
	}
	var influencerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Influencer).IsZero() {
			influencerRule = append(influencerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.Influencer)
		if err != nil {
			return nil, err
		}
		influencerRule = append(influencerRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		campaignIdRule,
		influencerRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeFundsWithdrawn decodes a log into a FundsWithdrawn struct.
func (c *Codec) DecodeFundsWithdrawn(log *evm.Log) (*FundsWithdrawnDecoded, error) {
	event := new(FundsWithdrawnDecoded)
	if err := c.abi.UnpackIntoInterface(event, "FundsWithdrawn", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["FundsWithdrawn"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) OwnershipTransferredLogHash() []byte {
	return c.abi.Events["OwnershipTransferred"].ID.Bytes()
}

func (c *Codec) EncodeOwnershipTransferredTopics(
	evt abi.Event,
	values []OwnershipTransferredTopics,
) ([]*evm.TopicValues, error) {
	var previousOwnerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.PreviousOwner).IsZero() {
			previousOwnerRule = append(previousOwnerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.PreviousOwner)
		if err != nil {
			return nil, err
		}
		previousOwnerRule = append(previousOwnerRule, fieldVal)
	}
	var newOwnerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.NewOwner).IsZero() {
			newOwnerRule = append(newOwnerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.NewOwner)
		if err != nil {
			return nil, err
		}
		newOwnerRule = append(newOwnerRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		previousOwnerRule,
		newOwnerRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeOwnershipTransferred decodes a log into a OwnershipTransferred struct.
func (c *Codec) DecodeOwnershipTransferred(log *evm.Log) (*OwnershipTransferredDecoded, error) {
	event := new(OwnershipTransferredDecoded)
	if err := c.abi.UnpackIntoInterface(event, "OwnershipTransferred", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["OwnershipTransferred"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) TokenRemovedFromWhitelistLogHash() []byte {
	return c.abi.Events["TokenRemovedFromWhitelist"].ID.Bytes()
}

func (c *Codec) EncodeTokenRemovedFromWhitelistTopics(
	evt abi.Event,
	values []TokenRemovedFromWhitelistTopics,
) ([]*evm.TopicValues, error) {
	var tokenRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Token).IsZero() {
			tokenRule = append(tokenRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Token)
		if err != nil {
			return nil, err
		}
		tokenRule = append(tokenRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		tokenRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeTokenRemovedFromWhitelist decodes a log into a TokenRemovedFromWhitelist struct.
func (c *Codec) DecodeTokenRemovedFromWhitelist(log *evm.Log) (*TokenRemovedFromWhitelistDecoded, error) {
	event := new(TokenRemovedFromWhitelistDecoded)
	if err := c.abi.UnpackIntoInterface(event, "TokenRemovedFromWhitelist", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["TokenRemovedFromWhitelist"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) TokenWhitelistedLogHash() []byte {
	return c.abi.Events["TokenWhitelisted"].ID.Bytes()
}

func (c *Codec) EncodeTokenWhitelistedTopics(
	evt abi.Event,
	values []TokenWhitelistedTopics,
) ([]*evm.TopicValues, error) {
	var tokenRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Token).IsZero() {
			tokenRule = append(tokenRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Token)
		if err != nil {
			return nil, err
		}
		tokenRule = append(tokenRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		tokenRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeTokenWhitelisted decodes a log into a TokenWhitelisted struct.
func (c *Codec) DecodeTokenWhitelisted(log *evm.Log) (*TokenWhitelistedDecoded, error) {
	event := new(TokenWhitelistedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "TokenWhitelisted", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["TokenWhitelisted"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c Escrow) NATIVEETHADDRESS(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeNATIVEETHADDRESSMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeNATIVEETHADDRESSMethodOutput(response.Data)
	})

}

func (c Escrow) CampaignCounter(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeCampaignCounterMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*big.Int](*new(*big.Int), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (*big.Int, error) {
		return c.Codec.DecodeCampaignCounterMethodOutput(response.Data)
	})

}

func (c Escrow) Campaigns(
	runtime cre.Runtime,
	args CampaignsInput,
	blockNumber *big.Int,
) cre.Promise[CampaignsOutput] {
	calldata, err := c.Codec.EncodeCampaignsMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[CampaignsOutput](CampaignsOutput{}, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (CampaignsOutput, error) {
		return c.Codec.DecodeCampaignsMethodOutput(response.Data)
	})

}

func (c Escrow) ExpectedAuthor(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeExpectedAuthorMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeExpectedAuthorMethodOutput(response.Data)
	})

}

func (c Escrow) ExpectedWorkflowId(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[[32]byte] {
	calldata, err := c.Codec.EncodeExpectedWorkflowIdMethodCall()
	if err != nil {
		return cre.PromiseFromResult[[32]byte](*new([32]byte), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) ([32]byte, error) {
		return c.Codec.DecodeExpectedWorkflowIdMethodOutput(response.Data)
	})

}

func (c Escrow) ExpectedWorkflowName(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[[10]byte] {
	calldata, err := c.Codec.EncodeExpectedWorkflowNameMethodCall()
	if err != nil {
		return cre.PromiseFromResult[[10]byte](*new([10]byte), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) ([10]byte, error) {
		return c.Codec.DecodeExpectedWorkflowNameMethodOutput(response.Data)
	})

}

func (c Escrow) ForwarderAddress(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeForwarderAddressMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeForwarderAddressMethodOutput(response.Data)
	})

}

func (c Escrow) GetCampaign(
	runtime cre.Runtime,
	args GetCampaignInput,
	blockNumber *big.Int,
) cre.Promise[AdEscrowCampaign] {
	calldata, err := c.Codec.EncodeGetCampaignMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[AdEscrowCampaign](*new(AdEscrowCampaign), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (AdEscrowCampaign, error) {
		return c.Codec.DecodeGetCampaignMethodOutput(response.Data)
	})

}

func (c Escrow) IsExpired(
	runtime cre.Runtime,
	args IsExpiredInput,
	blockNumber *big.Int,
) cre.Promise[bool] {
	calldata, err := c.Codec.EncodeIsExpiredMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[bool](*new(bool), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (bool, error) {
		return c.Codec.DecodeIsExpiredMethodOutput(response.Data)
	})

}

func (c Escrow) Owner(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeOwnerMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeOwnerMethodOutput(response.Data)
	})

}

func (c Escrow) WhitelistedTokens(
	runtime cre.Runtime,
	args WhitelistedTokensInput,
	blockNumber *big.Int,
) cre.Promise[bool] {
	calldata, err := c.Codec.EncodeWhitelistedTokensMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[bool](*new(bool), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (bool, error) {
		return c.Codec.DecodeWhitelistedTokensMethodOutput(response.Data)
	})

}

func (c Escrow) WriteReportFromAdEscrowCampaign(
	runtime cre.Runtime,
	input AdEscrowCampaign,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeAdEscrowCampaignStruct(input)
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *cre.Report) cre.Promise[*evm.WriteReportReply] {
		return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
			Receiver:  c.Address.Bytes(),
			Report:    report,
			GasConfig: gasConfig,
		})
	})
}

func (c Escrow) WriteReport(
	runtime cre.Runtime,
	report *cre.Report,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver:  c.Address.Bytes(),
		Report:    report,
		GasConfig: gasConfig,
	})
}

// DecodeInvalidAuthorError decodes a InvalidAuthor error from revert data.
func (c *Escrow) DecodeInvalidAuthorError(data []byte) (*InvalidAuthor, error) {
	args := c.ABI.Errors["InvalidAuthor"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 2 {
		return nil, fmt.Errorf("expected 2 values, got %d", len(values))
	}

	received, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for received in InvalidAuthor error")
	}

	expected, ok1 := values[1].(common.Address)
	if !ok1 {
		return nil, fmt.Errorf("unexpected type for expected in InvalidAuthor error")
	}

	return &InvalidAuthor{
		Received: received,
		Expected: expected,
	}, nil
}

// Error implements the error interface for InvalidAuthor.
func (e *InvalidAuthor) Error() string {
	return fmt.Sprintf("InvalidAuthor error: received=%v; expected=%v;", e.Received, e.Expected)
}

// DecodeInvalidSenderError decodes a InvalidSender error from revert data.
func (c *Escrow) DecodeInvalidSenderError(data []byte) (*InvalidSender, error) {
	args := c.ABI.Errors["InvalidSender"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 2 {
		return nil, fmt.Errorf("expected 2 values, got %d", len(values))
	}

	sender, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for sender in InvalidSender error")
	}

	expected, ok1 := values[1].(common.Address)
	if !ok1 {
		return nil, fmt.Errorf("unexpected type for expected in InvalidSender error")
	}

	return &InvalidSender{
		Sender:   sender,
		Expected: expected,
	}, nil
}

// Error implements the error interface for InvalidSender.
func (e *InvalidSender) Error() string {
	return fmt.Sprintf("InvalidSender error: sender=%v; expected=%v;", e.Sender, e.Expected)
}

// DecodeInvalidWorkflowIdError decodes a InvalidWorkflowId error from revert data.
func (c *Escrow) DecodeInvalidWorkflowIdError(data []byte) (*InvalidWorkflowId, error) {
	args := c.ABI.Errors["InvalidWorkflowId"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 2 {
		return nil, fmt.Errorf("expected 2 values, got %d", len(values))
	}

	received, ok0 := values[0].([32]byte)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for received in InvalidWorkflowId error")
	}

	expected, ok1 := values[1].([32]byte)
	if !ok1 {
		return nil, fmt.Errorf("unexpected type for expected in InvalidWorkflowId error")
	}

	return &InvalidWorkflowId{
		Received: received,
		Expected: expected,
	}, nil
}

// Error implements the error interface for InvalidWorkflowId.
func (e *InvalidWorkflowId) Error() string {
	return fmt.Sprintf("InvalidWorkflowId error: received=%v; expected=%v;", e.Received, e.Expected)
}

// DecodeInvalidWorkflowNameError decodes a InvalidWorkflowName error from revert data.
func (c *Escrow) DecodeInvalidWorkflowNameError(data []byte) (*InvalidWorkflowName, error) {
	args := c.ABI.Errors["InvalidWorkflowName"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 2 {
		return nil, fmt.Errorf("expected 2 values, got %d", len(values))
	}

	received, ok0 := values[0].([10]byte)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for received in InvalidWorkflowName error")
	}

	expected, ok1 := values[1].([10]byte)
	if !ok1 {
		return nil, fmt.Errorf("unexpected type for expected in InvalidWorkflowName error")
	}

	return &InvalidWorkflowName{
		Received: received,
		Expected: expected,
	}, nil
}

// Error implements the error interface for InvalidWorkflowName.
func (e *InvalidWorkflowName) Error() string {
	return fmt.Sprintf("InvalidWorkflowName error: received=%v; expected=%v;", e.Received, e.Expected)
}

// DecodeOwnableInvalidOwnerError decodes a OwnableInvalidOwner error from revert data.
func (c *Escrow) DecodeOwnableInvalidOwnerError(data []byte) (*OwnableInvalidOwner, error) {
	args := c.ABI.Errors["OwnableInvalidOwner"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	owner, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for owner in OwnableInvalidOwner error")
	}

	return &OwnableInvalidOwner{
		Owner: owner,
	}, nil
}

// Error implements the error interface for OwnableInvalidOwner.
func (e *OwnableInvalidOwner) Error() string {
	return fmt.Sprintf("OwnableInvalidOwner error: owner=%v;", e.Owner)
}

// DecodeOwnableUnauthorizedAccountError decodes a OwnableUnauthorizedAccount error from revert data.
func (c *Escrow) DecodeOwnableUnauthorizedAccountError(data []byte) (*OwnableUnauthorizedAccount, error) {
	args := c.ABI.Errors["OwnableUnauthorizedAccount"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	account, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for account in OwnableUnauthorizedAccount error")
	}

	return &OwnableUnauthorizedAccount{
		Account: account,
	}, nil
}

// Error implements the error interface for OwnableUnauthorizedAccount.
func (e *OwnableUnauthorizedAccount) Error() string {
	return fmt.Sprintf("OwnableUnauthorizedAccount error: account=%v;", e.Account)
}

// DecodeSafeERC20FailedOperationError decodes a SafeERC20FailedOperation error from revert data.
func (c *Escrow) DecodeSafeERC20FailedOperationError(data []byte) (*SafeERC20FailedOperation, error) {
	args := c.ABI.Errors["SafeERC20FailedOperation"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	token, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for token in SafeERC20FailedOperation error")
	}

	return &SafeERC20FailedOperation{
		Token: token,
	}, nil
}

// Error implements the error interface for SafeERC20FailedOperation.
func (e *SafeERC20FailedOperation) Error() string {
	return fmt.Sprintf("SafeERC20FailedOperation error: token=%v;", e.Token)
}

func (c *Escrow) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	case common.Bytes2Hex(c.ABI.Errors["InvalidAuthor"].ID.Bytes()[:4]):
		return c.DecodeInvalidAuthorError(data)
	case common.Bytes2Hex(c.ABI.Errors["InvalidSender"].ID.Bytes()[:4]):
		return c.DecodeInvalidSenderError(data)
	case common.Bytes2Hex(c.ABI.Errors["InvalidWorkflowId"].ID.Bytes()[:4]):
		return c.DecodeInvalidWorkflowIdError(data)
	case common.Bytes2Hex(c.ABI.Errors["InvalidWorkflowName"].ID.Bytes()[:4]):
		return c.DecodeInvalidWorkflowNameError(data)
	case common.Bytes2Hex(c.ABI.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]):
		return c.DecodeOwnableInvalidOwnerError(data)
	case common.Bytes2Hex(c.ABI.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]):
		return c.DecodeOwnableUnauthorizedAccountError(data)
	case common.Bytes2Hex(c.ABI.Errors["SafeERC20FailedOperation"].ID.Bytes()[:4]):
		return c.DecodeSafeERC20FailedOperationError(data)
	default:
		return nil, errors.New("unknown error selector")
	}
}

// CampaignDepositedTrigger wraps the raw log trigger and provides decoded CampaignDepositedDecoded data
type CampaignDepositedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Escrow // Keep reference for decoding
}

// Adapt method that decodes the log into CampaignDeposited data
func (t *CampaignDepositedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[CampaignDepositedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeCampaignDeposited(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode CampaignDeposited log: %w", err)
	}

	return &bindings.DecodedLog[CampaignDepositedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Escrow) LogTriggerCampaignDepositedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []CampaignDepositedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[CampaignDepositedDecoded]], error) {
	event := c.ABI.Events["CampaignDeposited"]
	topics, err := c.Codec.EncodeCampaignDepositedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for CampaignDeposited: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &CampaignDepositedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Escrow) FilterLogsCampaignDeposited(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.CampaignDepositedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// CampaignExpiredTrigger wraps the raw log trigger and provides decoded CampaignExpiredDecoded data
type CampaignExpiredTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Escrow // Keep reference for decoding
}

// Adapt method that decodes the log into CampaignExpired data
func (t *CampaignExpiredTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[CampaignExpiredDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeCampaignExpired(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode CampaignExpired log: %w", err)
	}

	return &bindings.DecodedLog[CampaignExpiredDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Escrow) LogTriggerCampaignExpiredLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []CampaignExpiredTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[CampaignExpiredDecoded]], error) {
	event := c.ABI.Events["CampaignExpired"]
	topics, err := c.Codec.EncodeCampaignExpiredTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for CampaignExpired: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &CampaignExpiredTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Escrow) FilterLogsCampaignExpired(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.CampaignExpiredLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// DeliveryActionCalledTrigger wraps the raw log trigger and provides decoded DeliveryActionCalledDecoded data
type DeliveryActionCalledTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Escrow // Keep reference for decoding
}

// Adapt method that decodes the log into DeliveryActionCalled data
func (t *DeliveryActionCalledTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[DeliveryActionCalledDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeDeliveryActionCalled(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode DeliveryActionCalled log: %w", err)
	}

	return &bindings.DecodedLog[DeliveryActionCalledDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Escrow) LogTriggerDeliveryActionCalledLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []DeliveryActionCalledTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[DeliveryActionCalledDecoded]], error) {
	event := c.ABI.Events["DeliveryActionCalled"]
	topics, err := c.Codec.EncodeDeliveryActionCalledTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for DeliveryActionCalled: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &DeliveryActionCalledTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Escrow) FilterLogsDeliveryActionCalled(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.DeliveryActionCalledLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// FundsWithdrawnTrigger wraps the raw log trigger and provides decoded FundsWithdrawnDecoded data
type FundsWithdrawnTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Escrow // Keep reference for decoding
}

// Adapt method that decodes the log into FundsWithdrawn data
func (t *FundsWithdrawnTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[FundsWithdrawnDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeFundsWithdrawn(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode FundsWithdrawn log: %w", err)
	}

	return &bindings.DecodedLog[FundsWithdrawnDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Escrow) LogTriggerFundsWithdrawnLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []FundsWithdrawnTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[FundsWithdrawnDecoded]], error) {
	event := c.ABI.Events["FundsWithdrawn"]
	topics, err := c.Codec.EncodeFundsWithdrawnTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for FundsWithdrawn: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &FundsWithdrawnTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Escrow) FilterLogsFundsWithdrawn(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.FundsWithdrawnLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// OwnershipTransferredTrigger wraps the raw log trigger and provides decoded OwnershipTransferredDecoded data
type OwnershipTransferredTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Escrow // Keep reference for decoding
}

// Adapt method that decodes the log into OwnershipTransferred data
func (t *OwnershipTransferredTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[OwnershipTransferredDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeOwnershipTransferred(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode OwnershipTransferred log: %w", err)
	}

	return &bindings.DecodedLog[OwnershipTransferredDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Escrow) LogTriggerOwnershipTransferredLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []OwnershipTransferredTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[OwnershipTransferredDecoded]], error) {
	event := c.ABI.Events["OwnershipTransferred"]
	topics, err := c.Codec.EncodeOwnershipTransferredTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for OwnershipTransferred: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &OwnershipTransferredTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Escrow) FilterLogsOwnershipTransferred(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.OwnershipTransferredLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// TokenRemovedFromWhitelistTrigger wraps the raw log trigger and provides decoded TokenRemovedFromWhitelistDecoded data
type TokenRemovedFromWhitelistTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Escrow // Keep reference for decoding
}

// Adapt method that decodes the log into TokenRemovedFromWhitelist data
func (t *TokenRemovedFromWhitelistTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[TokenRemovedFromWhitelistDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeTokenRemovedFromWhitelist(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode TokenRemovedFromWhitelist log: %w", err)
	}

	return &bindings.DecodedLog[TokenRemovedFromWhitelistDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Escrow) LogTriggerTokenRemovedFromWhitelistLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []TokenRemovedFromWhitelistTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[TokenRemovedFromWhitelistDecoded]], error) {
	event := c.ABI.Events["TokenRemovedFromWhitelist"]
	topics, err := c.Codec.EncodeTokenRemovedFromWhitelistTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for TokenRemovedFromWhitelist: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &TokenRemovedFromWhitelistTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Escrow) FilterLogsTokenRemovedFromWhitelist(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.TokenRemovedFromWhitelistLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// TokenWhitelistedTrigger wraps the raw log trigger and provides decoded TokenWhitelistedDecoded data
type TokenWhitelistedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Escrow // Keep reference for decoding
}

// Adapt method that decodes the log into TokenWhitelisted data
func (t *TokenWhitelistedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[TokenWhitelistedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeTokenWhitelisted(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode TokenWhitelisted log: %w", err)
	}

	return &bindings.DecodedLog[TokenWhitelistedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Escrow) LogTriggerTokenWhitelistedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []TokenWhitelistedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[TokenWhitelistedDecoded]], error) {
	event := c.ABI.Events["TokenWhitelisted"]
	topics, err := c.Codec.EncodeTokenWhitelistedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for TokenWhitelisted: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &TokenWhitelistedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Escrow) FilterLogsTokenWhitelisted(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.TokenWhitelistedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}
