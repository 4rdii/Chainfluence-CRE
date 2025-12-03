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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_authorizedWithdrawer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"advertiser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"influencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CampaignDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"}],\"name\":\"DeliveryActionCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"influencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"authorizedWithdrawer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"campaigns\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"advertiser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"influencer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentText\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minViews\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"withdrawn\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createCampaignId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"}],\"name\":\"deliveryAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"influencer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contentText\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minViews\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"}],\"name\":\"getCampaign\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"advertiser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"influencer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentText\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minViews\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"withdrawn\",\"type\":\"bool\"}],\"internalType\":\"structAdEscrow.Campaign\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextCampaignId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAuthorizedWithdrawer\",\"type\":\"address\"}],\"name\":\"setAuthorizedWithdrawer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualViews\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Structs
type AdEscrowCampaign struct {
	Advertiser  common.Address
	Influencer  common.Address
	Amount      *big.Int
	ContentText string
	MinViews    *big.Int
	Deadline    *big.Int
	Fulfilled   bool
	Withdrawn   bool
}

// Contract Method Inputs
type CampaignsInput struct {
	Arg0 *big.Int
}

type DeliveryActionInput struct {
	CampaignId *big.Int
}

type DepositInput struct {
	CampaignId  *big.Int
	Influencer  common.Address
	ContentText string
	MinViews    *big.Int
	Deadline    *big.Int
}

type GetCampaignInput struct {
	CampaignId *big.Int
}

type SetAuthorizedWithdrawerInput struct {
	NewAuthorizedWithdrawer common.Address
}

type WithdrawInput struct {
	CampaignId  *big.Int
	ActualViews *big.Int
}

// Contract Method Outputs
type CampaignsOutput struct {
	Advertiser  common.Address
	Influencer  common.Address
	Amount      *big.Int
	ContentText string
	MinViews    *big.Int
	Deadline    *big.Int
	Fulfilled   bool
	Withdrawn   bool
}

// Errors

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
	Amount     *big.Int
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
	EncodeAuthorizedWithdrawerMethodCall() ([]byte, error)
	DecodeAuthorizedWithdrawerMethodOutput(data []byte) (common.Address, error)
	EncodeCampaignsMethodCall(in CampaignsInput) ([]byte, error)
	DecodeCampaignsMethodOutput(data []byte) (CampaignsOutput, error)
	EncodeCreateCampaignIdMethodCall() ([]byte, error)
	DecodeCreateCampaignIdMethodOutput(data []byte) (*big.Int, error)
	EncodeDeliveryActionMethodCall(in DeliveryActionInput) ([]byte, error)
	EncodeDepositMethodCall(in DepositInput) ([]byte, error)
	EncodeGetCampaignMethodCall(in GetCampaignInput) ([]byte, error)
	DecodeGetCampaignMethodOutput(data []byte) (AdEscrowCampaign, error)
	EncodeGetNextCampaignIdMethodCall() ([]byte, error)
	DecodeGetNextCampaignIdMethodOutput(data []byte) (*big.Int, error)
	EncodeOwnerMethodCall() ([]byte, error)
	DecodeOwnerMethodOutput(data []byte) (common.Address, error)
	EncodeSetAuthorizedWithdrawerMethodCall(in SetAuthorizedWithdrawerInput) ([]byte, error)
	EncodeWithdrawMethodCall(in WithdrawInput) ([]byte, error)
	EncodeAdEscrowCampaignStruct(in AdEscrowCampaign) ([]byte, error)
	CampaignDepositedLogHash() []byte
	EncodeCampaignDepositedTopics(evt abi.Event, values []CampaignDepositedTopics) ([]*evm.TopicValues, error)
	DecodeCampaignDeposited(log *evm.Log) (*CampaignDepositedDecoded, error)
	DeliveryActionCalledLogHash() []byte
	EncodeDeliveryActionCalledTopics(evt abi.Event, values []DeliveryActionCalledTopics) ([]*evm.TopicValues, error)
	DecodeDeliveryActionCalled(log *evm.Log) (*DeliveryActionCalledDecoded, error)
	FundsWithdrawnLogHash() []byte
	EncodeFundsWithdrawnTopics(evt abi.Event, values []FundsWithdrawnTopics) ([]*evm.TopicValues, error)
	DecodeFundsWithdrawn(log *evm.Log) (*FundsWithdrawnDecoded, error)
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

func (c *Codec) EncodeAuthorizedWithdrawerMethodCall() ([]byte, error) {
	return c.abi.Pack("authorizedWithdrawer")
}

func (c *Codec) DecodeAuthorizedWithdrawerMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["authorizedWithdrawer"].Outputs.Unpack(data)
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

func (c *Codec) EncodeCampaignsMethodCall(in CampaignsInput) ([]byte, error) {
	return c.abi.Pack("campaigns", in.Arg0)
}

func (c *Codec) DecodeCampaignsMethodOutput(data []byte) (CampaignsOutput, error) {
	vals, err := c.abi.Methods["campaigns"].Outputs.Unpack(data)
	if err != nil {
		return CampaignsOutput{}, err
	}
	if len(vals) != 8 {
		return CampaignsOutput{}, fmt.Errorf("expected 8 values, got %d", len(vals))
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

	var result2 *big.Int
	if err := json.Unmarshal(jsonData2, &result2); err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData3, err := json.Marshal(vals[3])
	if err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to marshal ABI result 3: %w", err)
	}

	var result3 string
	if err := json.Unmarshal(jsonData3, &result3); err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to unmarshal to string: %w", err)
	}
	jsonData4, err := json.Marshal(vals[4])
	if err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to marshal ABI result 4: %w", err)
	}

	var result4 *big.Int
	if err := json.Unmarshal(jsonData4, &result4); err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
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

	var result6 bool
	if err := json.Unmarshal(jsonData6, &result6); err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to unmarshal to bool: %w", err)
	}
	jsonData7, err := json.Marshal(vals[7])
	if err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to marshal ABI result 7: %w", err)
	}

	var result7 bool
	if err := json.Unmarshal(jsonData7, &result7); err != nil {
		return CampaignsOutput{}, fmt.Errorf("failed to unmarshal to bool: %w", err)
	}

	return CampaignsOutput{
		Advertiser:  result0,
		Influencer:  result1,
		Amount:      result2,
		ContentText: result3,
		MinViews:    result4,
		Deadline:    result5,
		Fulfilled:   result6,
		Withdrawn:   result7,
	}, nil
}

func (c *Codec) EncodeCreateCampaignIdMethodCall() ([]byte, error) {
	return c.abi.Pack("createCampaignId")
}

func (c *Codec) DecodeCreateCampaignIdMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["createCampaignId"].Outputs.Unpack(data)
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

func (c *Codec) EncodeDeliveryActionMethodCall(in DeliveryActionInput) ([]byte, error) {
	return c.abi.Pack("deliveryAction", in.CampaignId)
}

func (c *Codec) EncodeDepositMethodCall(in DepositInput) ([]byte, error) {
	return c.abi.Pack("deposit", in.CampaignId, in.Influencer, in.ContentText, in.MinViews, in.Deadline)
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

func (c *Codec) EncodeGetNextCampaignIdMethodCall() ([]byte, error) {
	return c.abi.Pack("getNextCampaignId")
}

func (c *Codec) DecodeGetNextCampaignIdMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["getNextCampaignId"].Outputs.Unpack(data)
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

func (c *Codec) EncodeSetAuthorizedWithdrawerMethodCall(in SetAuthorizedWithdrawerInput) ([]byte, error) {
	return c.abi.Pack("setAuthorizedWithdrawer", in.NewAuthorizedWithdrawer)
}

func (c *Codec) EncodeWithdrawMethodCall(in WithdrawInput) ([]byte, error) {
	return c.abi.Pack("withdraw", in.CampaignId, in.ActualViews)
}

func (c *Codec) EncodeAdEscrowCampaignStruct(in AdEscrowCampaign) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "advertiser", Type: "address"},
			{Name: "influencer", Type: "address"},
			{Name: "amount", Type: "uint256"},
			{Name: "contentText", Type: "string"},
			{Name: "minViews", Type: "uint256"},
			{Name: "deadline", Type: "uint256"},
			{Name: "fulfilled", Type: "bool"},
			{Name: "withdrawn", Type: "bool"},
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

func (c Escrow) AuthorizedWithdrawer(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeAuthorizedWithdrawerMethodCall()
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
		return c.Codec.DecodeAuthorizedWithdrawerMethodOutput(response.Data)
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

func (c Escrow) GetNextCampaignId(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeGetNextCampaignIdMethodCall()
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
		return c.Codec.DecodeGetNextCampaignIdMethodOutput(response.Data)
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

func (c *Escrow) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
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

func (c *Escrow) FilterLogsCampaignDeposited(runtime cre.Runtime, options *bindings.FilterOptions) cre.Promise[*evm.FilterLogsReply] {
	if options == nil {
		options = &bindings.FilterOptions{
			ToBlock: options.ToBlock,
		}
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
	})
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

func (c *Escrow) FilterLogsDeliveryActionCalled(runtime cre.Runtime, options *bindings.FilterOptions) cre.Promise[*evm.FilterLogsReply] {
	if options == nil {
		options = &bindings.FilterOptions{
			ToBlock: options.ToBlock,
		}
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
	})
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

func (c *Escrow) FilterLogsFundsWithdrawn(runtime cre.Runtime, options *bindings.FilterOptions) cre.Promise[*evm.FilterLogsReply] {
	if options == nil {
		options = &bindings.FilterOptions{
			ToBlock: options.ToBlock,
		}
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
	})
}
