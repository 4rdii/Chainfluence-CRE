// Code generated — DO NOT EDIT.

package escrow_v2

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

var EscrowV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_forwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_whitelistedTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_platformFeeBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_feeRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"MAX_FEE_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NATIVE_ETH_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDeal\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tweetUrl\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDeal\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimExpired\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deals\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"advertiser\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"influencer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"minViews\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"campaignDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumAdEscrowV2.CampaignState\"},{\"name\":\"platformFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"influencerAccepted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"channelName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tweetUrl\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"internalType\":\"structAdEscrowV2.DepositParams\",\"components\":[{\"name\":\"dealId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"influencer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"minViews\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"campaignDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"channelName\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"expectedAuthor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"expectedWorkflowId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"expectedWorkflowName\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeRecipient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forwarderAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeal\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structAdEscrowV2.Campaign\",\"components\":[{\"name\":\"advertiser\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"influencer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"minViews\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"campaignDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumAdEscrowV2.CampaignState\"},{\"name\":\"platformFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"influencerAccepted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"channelName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tweetUrl\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isExpired\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onReport\",\"inputs\":[{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"platformFeeBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseDispute\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeTokenFromWhitelist\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolveDispute\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"releaseToInfluencer\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedAuthor\",\"inputs\":[{\"name\":\"_author\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedWorkflowId\",\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedWorkflowName\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeRecipient\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setForwarderAddress\",\"inputs\":[{\"name\":\"_forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPlatformFeeBps\",\"inputs\":[{\"name\":\"_feeBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"whitelistToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"whitelistedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DealAccepted\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"influencer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tweetUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DealCancelled\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"advertiser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DealCreated\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"advertiser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"influencer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"contentHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"channelName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeRaised\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"raisedBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeResolved\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"releasedToInfluencer\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundsRefunded\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"advertiser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundsReleased\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"influencer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"influencerAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenRemovedFromWhitelist\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenWhitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyAccepted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CampaignExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CampaignNotExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DealAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAuthor\",\"inputs\":[{\"name\":\"received\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidState\",\"inputs\":[{\"name\":\"current\",\"type\":\"uint8\",\"internalType\":\"enumAdEscrowV2.CampaignState\"},{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"enumAdEscrowV2.CampaignState\"}]},{\"type\":\"error\",\"name\":\"InvalidWorkflowId\",\"inputs\":[{\"name\":\"received\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidWorkflowName\",\"inputs\":[{\"name\":\"received\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"},{\"name\":\"expected\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"}]},{\"type\":\"error\",\"name\":\"NotAdvertiser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInfluencer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenNotWhitelisted\",\"inputs\":[]}]",
}

// Structs
type AdEscrowV2Campaign struct {
	Advertiser         common.Address
	Influencer         common.Address
	Token              common.Address
	Amount             *big.Int
	ContentHash        [32]byte
	MinViews           *big.Int
	CampaignDuration   uint64
	Deadline           *big.Int
	State              uint8
	PlatformFee        *big.Int
	InfluencerAccepted bool
	ChannelName        string
	TweetUrl           string
}

type AdEscrowV2DepositParams struct {
	DealId           *big.Int
	Token            common.Address
	Influencer       common.Address
	Amount           *big.Int
	ContentHash      [32]byte
	MinViews         *big.Int
	ExpiryDeadline   *big.Int
	CampaignDuration uint64
	ChannelName      string
}

// Contract Method Inputs
type AcceptDealInput struct {
	DealId   *big.Int
	TweetUrl string
}

type CancelDealInput struct {
	DealId *big.Int
}

type ClaimExpiredInput struct {
	DealId *big.Int
}

type DealsInput struct {
	Arg0 *big.Int
}

type DepositInput struct {
	P AdEscrowV2DepositParams
}

type GetDealInput struct {
	DealId *big.Int
}

type IsExpiredInput struct {
	DealId *big.Int
}

type OnReportInput struct {
	Metadata []byte
	Report   []byte
}

type RaiseDisputeInput struct {
	DealId *big.Int
}

type RemoveTokenFromWhitelistInput struct {
	Token common.Address
}

type ResolveDisputeInput struct {
	DealId              *big.Int
	ReleaseToInfluencer bool
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

type SetFeeRecipientInput struct {
	Recipient common.Address
}

type SetForwarderAddressInput struct {
	Forwarder common.Address
}

type SetPlatformFeeBpsInput struct {
	FeeBps *big.Int
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
type DealsOutput struct {
	Advertiser         common.Address
	Influencer         common.Address
	Token              common.Address
	Amount             *big.Int
	ContentHash        [32]byte
	MinViews           *big.Int
	CampaignDuration   uint64
	Deadline           *big.Int
	State              uint8
	PlatformFee        *big.Int
	InfluencerAccepted bool
	ChannelName        string
	TweetUrl           string
}

// Errors
type AlreadyAccepted struct {
}

type CampaignExpired struct {
}

type CampaignNotExpired struct {
}

type DealAlreadyExists struct {
}

type FeeTooHigh struct {
}

type InvalidAmount struct {
}

type InvalidAuthor struct {
	Received common.Address
	Expected common.Address
}

type InvalidSender struct {
	Sender   common.Address
	Expected common.Address
}

type InvalidState struct {
	Current  uint8
	Expected uint8
}

type InvalidWorkflowId struct {
	Received [32]byte
	Expected [32]byte
}

type InvalidWorkflowName struct {
	Received [10]byte
	Expected [10]byte
}

type NotAdvertiser struct {
}

type NotInfluencer struct {
}

type NotParticipant struct {
}

type OwnableInvalidOwner struct {
	Owner common.Address
}

type OwnableUnauthorizedAccount struct {
	Account common.Address
}

type ReentrancyGuardReentrantCall struct {
}

type SafeERC20FailedOperation struct {
	Token common.Address
}

type TokenNotWhitelisted struct {
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

type DealAcceptedTopics struct {
	DealId     *big.Int
	Influencer common.Address
}

type DealAcceptedDecoded struct {
	DealId     *big.Int
	Influencer common.Address
	TweetUrl   string
}

type DealCancelledTopics struct {
	DealId     *big.Int
	Advertiser common.Address
}

type DealCancelledDecoded struct {
	DealId     *big.Int
	Advertiser common.Address
}

type DealCreatedTopics struct {
	DealId     *big.Int
	Advertiser common.Address
	Influencer common.Address
}

type DealCreatedDecoded struct {
	DealId      *big.Int
	Advertiser  common.Address
	Influencer  common.Address
	Token       common.Address
	Amount      *big.Int
	ContentHash [32]byte
	ChannelName string
}

type DisputeRaisedTopics struct {
	DealId   *big.Int
	RaisedBy common.Address
}

type DisputeRaisedDecoded struct {
	DealId   *big.Int
	RaisedBy common.Address
}

type DisputeResolvedTopics struct {
	DealId *big.Int
}

type DisputeResolvedDecoded struct {
	DealId               *big.Int
	ReleasedToInfluencer bool
}

type FundsRefundedTopics struct {
	DealId     *big.Int
	Advertiser common.Address
}

type FundsRefundedDecoded struct {
	DealId     *big.Int
	Advertiser common.Address
	Token      common.Address
	Amount     *big.Int
}

type FundsReleasedTopics struct {
	DealId     *big.Int
	Influencer common.Address
}

type FundsReleasedDecoded struct {
	DealId           *big.Int
	Influencer       common.Address
	Token            common.Address
	InfluencerAmount *big.Int
	FeeAmount        *big.Int
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

// Main Binding Type for EscrowV2
type EscrowV2 struct {
	Address common.Address
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   EscrowV2Codec
}

type EscrowV2Codec interface {
	EncodeMAXFEEBPSMethodCall() ([]byte, error)
	DecodeMAXFEEBPSMethodOutput(data []byte) (*big.Int, error)
	EncodeNATIVEETHADDRESSMethodCall() ([]byte, error)
	DecodeNATIVEETHADDRESSMethodOutput(data []byte) (common.Address, error)
	EncodeAcceptDealMethodCall(in AcceptDealInput) ([]byte, error)
	EncodeCancelDealMethodCall(in CancelDealInput) ([]byte, error)
	EncodeClaimExpiredMethodCall(in ClaimExpiredInput) ([]byte, error)
	EncodeDealsMethodCall(in DealsInput) ([]byte, error)
	DecodeDealsMethodOutput(data []byte) (DealsOutput, error)
	EncodeDepositMethodCall(in DepositInput) ([]byte, error)
	EncodeExpectedAuthorMethodCall() ([]byte, error)
	DecodeExpectedAuthorMethodOutput(data []byte) (common.Address, error)
	EncodeExpectedWorkflowIdMethodCall() ([]byte, error)
	DecodeExpectedWorkflowIdMethodOutput(data []byte) ([32]byte, error)
	EncodeExpectedWorkflowNameMethodCall() ([]byte, error)
	DecodeExpectedWorkflowNameMethodOutput(data []byte) ([10]byte, error)
	EncodeFeeRecipientMethodCall() ([]byte, error)
	DecodeFeeRecipientMethodOutput(data []byte) (common.Address, error)
	EncodeForwarderAddressMethodCall() ([]byte, error)
	DecodeForwarderAddressMethodOutput(data []byte) (common.Address, error)
	EncodeGetDealMethodCall(in GetDealInput) ([]byte, error)
	DecodeGetDealMethodOutput(data []byte) (AdEscrowV2Campaign, error)
	EncodeIsExpiredMethodCall(in IsExpiredInput) ([]byte, error)
	DecodeIsExpiredMethodOutput(data []byte) (bool, error)
	EncodeOnReportMethodCall(in OnReportInput) ([]byte, error)
	EncodeOwnerMethodCall() ([]byte, error)
	DecodeOwnerMethodOutput(data []byte) (common.Address, error)
	EncodePlatformFeeBpsMethodCall() ([]byte, error)
	DecodePlatformFeeBpsMethodOutput(data []byte) (*big.Int, error)
	EncodeRaiseDisputeMethodCall(in RaiseDisputeInput) ([]byte, error)
	EncodeRemoveTokenFromWhitelistMethodCall(in RemoveTokenFromWhitelistInput) ([]byte, error)
	EncodeRenounceOwnershipMethodCall() ([]byte, error)
	EncodeResolveDisputeMethodCall(in ResolveDisputeInput) ([]byte, error)
	EncodeSetExpectedAuthorMethodCall(in SetExpectedAuthorInput) ([]byte, error)
	EncodeSetExpectedWorkflowIdMethodCall(in SetExpectedWorkflowIdInput) ([]byte, error)
	EncodeSetExpectedWorkflowNameMethodCall(in SetExpectedWorkflowNameInput) ([]byte, error)
	EncodeSetFeeRecipientMethodCall(in SetFeeRecipientInput) ([]byte, error)
	EncodeSetForwarderAddressMethodCall(in SetForwarderAddressInput) ([]byte, error)
	EncodeSetPlatformFeeBpsMethodCall(in SetPlatformFeeBpsInput) ([]byte, error)
	EncodeSupportsInterfaceMethodCall(in SupportsInterfaceInput) ([]byte, error)
	DecodeSupportsInterfaceMethodOutput(data []byte) (bool, error)
	EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error)
	EncodeWhitelistTokenMethodCall(in WhitelistTokenInput) ([]byte, error)
	EncodeWhitelistedTokensMethodCall(in WhitelistedTokensInput) ([]byte, error)
	DecodeWhitelistedTokensMethodOutput(data []byte) (bool, error)
	EncodeAdEscrowV2CampaignStruct(in AdEscrowV2Campaign) ([]byte, error)
	EncodeAdEscrowV2DepositParamsStruct(in AdEscrowV2DepositParams) ([]byte, error)
	DealAcceptedLogHash() []byte
	EncodeDealAcceptedTopics(evt abi.Event, values []DealAcceptedTopics) ([]*evm.TopicValues, error)
	DecodeDealAccepted(log *evm.Log) (*DealAcceptedDecoded, error)
	DealCancelledLogHash() []byte
	EncodeDealCancelledTopics(evt abi.Event, values []DealCancelledTopics) ([]*evm.TopicValues, error)
	DecodeDealCancelled(log *evm.Log) (*DealCancelledDecoded, error)
	DealCreatedLogHash() []byte
	EncodeDealCreatedTopics(evt abi.Event, values []DealCreatedTopics) ([]*evm.TopicValues, error)
	DecodeDealCreated(log *evm.Log) (*DealCreatedDecoded, error)
	DisputeRaisedLogHash() []byte
	EncodeDisputeRaisedTopics(evt abi.Event, values []DisputeRaisedTopics) ([]*evm.TopicValues, error)
	DecodeDisputeRaised(log *evm.Log) (*DisputeRaisedDecoded, error)
	DisputeResolvedLogHash() []byte
	EncodeDisputeResolvedTopics(evt abi.Event, values []DisputeResolvedTopics) ([]*evm.TopicValues, error)
	DecodeDisputeResolved(log *evm.Log) (*DisputeResolvedDecoded, error)
	FundsRefundedLogHash() []byte
	EncodeFundsRefundedTopics(evt abi.Event, values []FundsRefundedTopics) ([]*evm.TopicValues, error)
	DecodeFundsRefunded(log *evm.Log) (*FundsRefundedDecoded, error)
	FundsReleasedLogHash() []byte
	EncodeFundsReleasedTopics(evt abi.Event, values []FundsReleasedTopics) ([]*evm.TopicValues, error)
	DecodeFundsReleased(log *evm.Log) (*FundsReleasedDecoded, error)
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

func NewEscrowV2(
	client *evm.Client,
	address common.Address,
	options *bindings.ContractInitOptions,
) (*EscrowV2, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowV2MetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewCodec()
	if err != nil {
		return nil, err
	}
	return &EscrowV2{
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

func NewCodec() (EscrowV2Codec, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowV2MetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &Codec{abi: &parsed}, nil
}

func (c *Codec) EncodeMAXFEEBPSMethodCall() ([]byte, error) {
	return c.abi.Pack("MAX_FEE_BPS")
}

func (c *Codec) DecodeMAXFEEBPSMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["MAX_FEE_BPS"].Outputs.Unpack(data)
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

func (c *Codec) EncodeAcceptDealMethodCall(in AcceptDealInput) ([]byte, error) {
	return c.abi.Pack("acceptDeal", in.DealId, in.TweetUrl)
}

func (c *Codec) EncodeCancelDealMethodCall(in CancelDealInput) ([]byte, error) {
	return c.abi.Pack("cancelDeal", in.DealId)
}

func (c *Codec) EncodeClaimExpiredMethodCall(in ClaimExpiredInput) ([]byte, error) {
	return c.abi.Pack("claimExpired", in.DealId)
}

func (c *Codec) EncodeDealsMethodCall(in DealsInput) ([]byte, error) {
	return c.abi.Pack("deals", in.Arg0)
}

func (c *Codec) DecodeDealsMethodOutput(data []byte) (DealsOutput, error) {
	vals, err := c.abi.Methods["deals"].Outputs.Unpack(data)
	if err != nil {
		return DealsOutput{}, err
	}
	if len(vals) != 13 {
		return DealsOutput{}, fmt.Errorf("expected 13 values, got %d", len(vals))
	}
	jsonData0, err := json.Marshal(vals[0])
	if err != nil {
		return DealsOutput{}, fmt.Errorf("failed to marshal ABI result 0: %w", err)
	}

	var result0 common.Address
	if err := json.Unmarshal(jsonData0, &result0); err != nil {
		return DealsOutput{}, fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}
	jsonData1, err := json.Marshal(vals[1])
	if err != nil {
		return DealsOutput{}, fmt.Errorf("failed to marshal ABI result 1: %w", err)
	}

	var result1 common.Address
	if err := json.Unmarshal(jsonData1, &result1); err != nil {
		return DealsOutput{}, fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}
	jsonData2, err := json.Marshal(vals[2])
	if err != nil {
		return DealsOutput{}, fmt.Errorf("failed to marshal ABI result 2: %w", err)
	}

	var result2 common.Address
	if err := json.Unmarshal(jsonData2, &result2); err != nil {
		return DealsOutput{}, fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}
	jsonData3, err := json.Marshal(vals[3])
	if err != nil {
		return DealsOutput{}, fmt.Errorf("failed to marshal ABI result 3: %w", err)
	}

	var result3 *big.Int
	if err := json.Unmarshal(jsonData3, &result3); err != nil {
		return DealsOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData4, err := json.Marshal(vals[4])
	if err != nil {
		return DealsOutput{}, fmt.Errorf("failed to marshal ABI result 4: %w", err)
	}

	var result4 [32]byte
	if err := json.Unmarshal(jsonData4, &result4); err != nil {
		return DealsOutput{}, fmt.Errorf("failed to unmarshal to [32]byte: %w", err)
	}
	jsonData5, err := json.Marshal(vals[5])
	if err != nil {
		return DealsOutput{}, fmt.Errorf("failed to marshal ABI result 5: %w", err)
	}

	var result5 *big.Int
	if err := json.Unmarshal(jsonData5, &result5); err != nil {
		return DealsOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData6, err := json.Marshal(vals[6])
	if err != nil {
		return DealsOutput{}, fmt.Errorf("failed to marshal ABI result 6: %w", err)
	}

	var result6 uint64
	if err := json.Unmarshal(jsonData6, &result6); err != nil {
		return DealsOutput{}, fmt.Errorf("failed to unmarshal to uint64: %w", err)
	}
	jsonData7, err := json.Marshal(vals[7])
	if err != nil {
		return DealsOutput{}, fmt.Errorf("failed to marshal ABI result 7: %w", err)
	}

	var result7 *big.Int
	if err := json.Unmarshal(jsonData7, &result7); err != nil {
		return DealsOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData8, err := json.Marshal(vals[8])
	if err != nil {
		return DealsOutput{}, fmt.Errorf("failed to marshal ABI result 8: %w", err)
	}

	var result8 uint8
	if err := json.Unmarshal(jsonData8, &result8); err != nil {
		return DealsOutput{}, fmt.Errorf("failed to unmarshal to uint8: %w", err)
	}
	jsonData9, err := json.Marshal(vals[9])
	if err != nil {
		return DealsOutput{}, fmt.Errorf("failed to marshal ABI result 9: %w", err)
	}

	var result9 *big.Int
	if err := json.Unmarshal(jsonData9, &result9); err != nil {
		return DealsOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData10, err := json.Marshal(vals[10])
	if err != nil {
		return DealsOutput{}, fmt.Errorf("failed to marshal ABI result 10: %w", err)
	}

	var result10 bool
	if err := json.Unmarshal(jsonData10, &result10); err != nil {
		return DealsOutput{}, fmt.Errorf("failed to unmarshal to bool: %w", err)
	}
	jsonData11, err := json.Marshal(vals[11])
	if err != nil {
		return DealsOutput{}, fmt.Errorf("failed to marshal ABI result 11: %w", err)
	}

	var result11 string
	if err := json.Unmarshal(jsonData11, &result11); err != nil {
		return DealsOutput{}, fmt.Errorf("failed to unmarshal to string: %w", err)
	}
	jsonData12, err := json.Marshal(vals[12])
	if err != nil {
		return DealsOutput{}, fmt.Errorf("failed to marshal ABI result 12: %w", err)
	}

	var result12 string
	if err := json.Unmarshal(jsonData12, &result12); err != nil {
		return DealsOutput{}, fmt.Errorf("failed to unmarshal to string: %w", err)
	}

	return DealsOutput{
		Advertiser:         result0,
		Influencer:         result1,
		Token:              result2,
		Amount:             result3,
		ContentHash:        result4,
		MinViews:           result5,
		CampaignDuration:   result6,
		Deadline:           result7,
		State:              result8,
		PlatformFee:        result9,
		InfluencerAccepted: result10,
		ChannelName:        result11,
		TweetUrl:           result12,
	}, nil
}

func (c *Codec) EncodeDepositMethodCall(in DepositInput) ([]byte, error) {
	return c.abi.Pack("deposit", in.P)
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

func (c *Codec) EncodeFeeRecipientMethodCall() ([]byte, error) {
	return c.abi.Pack("feeRecipient")
}

func (c *Codec) DecodeFeeRecipientMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["feeRecipient"].Outputs.Unpack(data)
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

func (c *Codec) EncodeGetDealMethodCall(in GetDealInput) ([]byte, error) {
	return c.abi.Pack("getDeal", in.DealId)
}

func (c *Codec) DecodeGetDealMethodOutput(data []byte) (AdEscrowV2Campaign, error) {
	vals, err := c.abi.Methods["getDeal"].Outputs.Unpack(data)
	if err != nil {
		return *new(AdEscrowV2Campaign), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(AdEscrowV2Campaign), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result AdEscrowV2Campaign
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(AdEscrowV2Campaign), fmt.Errorf("failed to unmarshal to AdEscrowV2Campaign: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeIsExpiredMethodCall(in IsExpiredInput) ([]byte, error) {
	return c.abi.Pack("isExpired", in.DealId)
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

func (c *Codec) EncodePlatformFeeBpsMethodCall() ([]byte, error) {
	return c.abi.Pack("platformFeeBps")
}

func (c *Codec) DecodePlatformFeeBpsMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["platformFeeBps"].Outputs.Unpack(data)
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

func (c *Codec) EncodeRaiseDisputeMethodCall(in RaiseDisputeInput) ([]byte, error) {
	return c.abi.Pack("raiseDispute", in.DealId)
}

func (c *Codec) EncodeRemoveTokenFromWhitelistMethodCall(in RemoveTokenFromWhitelistInput) ([]byte, error) {
	return c.abi.Pack("removeTokenFromWhitelist", in.Token)
}

func (c *Codec) EncodeRenounceOwnershipMethodCall() ([]byte, error) {
	return c.abi.Pack("renounceOwnership")
}

func (c *Codec) EncodeResolveDisputeMethodCall(in ResolveDisputeInput) ([]byte, error) {
	return c.abi.Pack("resolveDispute", in.DealId, in.ReleaseToInfluencer)
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

func (c *Codec) EncodeSetFeeRecipientMethodCall(in SetFeeRecipientInput) ([]byte, error) {
	return c.abi.Pack("setFeeRecipient", in.Recipient)
}

func (c *Codec) EncodeSetForwarderAddressMethodCall(in SetForwarderAddressInput) ([]byte, error) {
	return c.abi.Pack("setForwarderAddress", in.Forwarder)
}

func (c *Codec) EncodeSetPlatformFeeBpsMethodCall(in SetPlatformFeeBpsInput) ([]byte, error) {
	return c.abi.Pack("setPlatformFeeBps", in.FeeBps)
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

func (c *Codec) EncodeAdEscrowV2CampaignStruct(in AdEscrowV2Campaign) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "advertiser", Type: "address"},
			{Name: "influencer", Type: "address"},
			{Name: "token", Type: "address"},
			{Name: "amount", Type: "uint256"},
			{Name: "contentHash", Type: "bytes32"},
			{Name: "minViews", Type: "uint256"},
			{Name: "campaignDuration", Type: "uint64"},
			{Name: "deadline", Type: "uint256"},
			{Name: "state", Type: "uint8"},
			{Name: "platformFee", Type: "uint256"},
			{Name: "influencerAccepted", Type: "bool"},
			{Name: "channelName", Type: "string"},
			{Name: "tweetUrl", Type: "string"},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for AdEscrowV2Campaign: %w", err)
	}
	args := abi.Arguments{
		{Name: "adEscrowV2Campaign", Type: tupleType},
	}

	return args.Pack(in)
}
func (c *Codec) EncodeAdEscrowV2DepositParamsStruct(in AdEscrowV2DepositParams) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "dealId", Type: "uint256"},
			{Name: "token", Type: "address"},
			{Name: "influencer", Type: "address"},
			{Name: "amount", Type: "uint256"},
			{Name: "contentHash", Type: "bytes32"},
			{Name: "minViews", Type: "uint256"},
			{Name: "expiryDeadline", Type: "uint256"},
			{Name: "campaignDuration", Type: "uint64"},
			{Name: "channelName", Type: "string"},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for AdEscrowV2DepositParams: %w", err)
	}
	args := abi.Arguments{
		{Name: "adEscrowV2DepositParams", Type: tupleType},
	}

	return args.Pack(in)
}

func (c *Codec) DealAcceptedLogHash() []byte {
	return c.abi.Events["DealAccepted"].ID.Bytes()
}

func (c *Codec) EncodeDealAcceptedTopics(
	evt abi.Event,
	values []DealAcceptedTopics,
) ([]*evm.TopicValues, error) {
	var dealIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.DealId).IsZero() {
			dealIdRule = append(dealIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.DealId)
		if err != nil {
			return nil, err
		}
		dealIdRule = append(dealIdRule, fieldVal)
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
		dealIdRule,
		influencerRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeDealAccepted decodes a log into a DealAccepted struct.
func (c *Codec) DecodeDealAccepted(log *evm.Log) (*DealAcceptedDecoded, error) {
	event := new(DealAcceptedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "DealAccepted", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["DealAccepted"].Inputs {
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

func (c *Codec) DealCancelledLogHash() []byte {
	return c.abi.Events["DealCancelled"].ID.Bytes()
}

func (c *Codec) EncodeDealCancelledTopics(
	evt abi.Event,
	values []DealCancelledTopics,
) ([]*evm.TopicValues, error) {
	var dealIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.DealId).IsZero() {
			dealIdRule = append(dealIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.DealId)
		if err != nil {
			return nil, err
		}
		dealIdRule = append(dealIdRule, fieldVal)
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
		dealIdRule,
		advertiserRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeDealCancelled decodes a log into a DealCancelled struct.
func (c *Codec) DecodeDealCancelled(log *evm.Log) (*DealCancelledDecoded, error) {
	event := new(DealCancelledDecoded)
	if err := c.abi.UnpackIntoInterface(event, "DealCancelled", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["DealCancelled"].Inputs {
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

func (c *Codec) DealCreatedLogHash() []byte {
	return c.abi.Events["DealCreated"].ID.Bytes()
}

func (c *Codec) EncodeDealCreatedTopics(
	evt abi.Event,
	values []DealCreatedTopics,
) ([]*evm.TopicValues, error) {
	var dealIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.DealId).IsZero() {
			dealIdRule = append(dealIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.DealId)
		if err != nil {
			return nil, err
		}
		dealIdRule = append(dealIdRule, fieldVal)
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
		dealIdRule,
		advertiserRule,
		influencerRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeDealCreated decodes a log into a DealCreated struct.
func (c *Codec) DecodeDealCreated(log *evm.Log) (*DealCreatedDecoded, error) {
	event := new(DealCreatedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "DealCreated", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["DealCreated"].Inputs {
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

func (c *Codec) DisputeRaisedLogHash() []byte {
	return c.abi.Events["DisputeRaised"].ID.Bytes()
}

func (c *Codec) EncodeDisputeRaisedTopics(
	evt abi.Event,
	values []DisputeRaisedTopics,
) ([]*evm.TopicValues, error) {
	var dealIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.DealId).IsZero() {
			dealIdRule = append(dealIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.DealId)
		if err != nil {
			return nil, err
		}
		dealIdRule = append(dealIdRule, fieldVal)
	}
	var raisedByRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.RaisedBy).IsZero() {
			raisedByRule = append(raisedByRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.RaisedBy)
		if err != nil {
			return nil, err
		}
		raisedByRule = append(raisedByRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		dealIdRule,
		raisedByRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeDisputeRaised decodes a log into a DisputeRaised struct.
func (c *Codec) DecodeDisputeRaised(log *evm.Log) (*DisputeRaisedDecoded, error) {
	event := new(DisputeRaisedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "DisputeRaised", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["DisputeRaised"].Inputs {
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

func (c *Codec) DisputeResolvedLogHash() []byte {
	return c.abi.Events["DisputeResolved"].ID.Bytes()
}

func (c *Codec) EncodeDisputeResolvedTopics(
	evt abi.Event,
	values []DisputeResolvedTopics,
) ([]*evm.TopicValues, error) {
	var dealIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.DealId).IsZero() {
			dealIdRule = append(dealIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.DealId)
		if err != nil {
			return nil, err
		}
		dealIdRule = append(dealIdRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		dealIdRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeDisputeResolved decodes a log into a DisputeResolved struct.
func (c *Codec) DecodeDisputeResolved(log *evm.Log) (*DisputeResolvedDecoded, error) {
	event := new(DisputeResolvedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "DisputeResolved", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["DisputeResolved"].Inputs {
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

func (c *Codec) FundsRefundedLogHash() []byte {
	return c.abi.Events["FundsRefunded"].ID.Bytes()
}

func (c *Codec) EncodeFundsRefundedTopics(
	evt abi.Event,
	values []FundsRefundedTopics,
) ([]*evm.TopicValues, error) {
	var dealIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.DealId).IsZero() {
			dealIdRule = append(dealIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.DealId)
		if err != nil {
			return nil, err
		}
		dealIdRule = append(dealIdRule, fieldVal)
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
		dealIdRule,
		advertiserRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeFundsRefunded decodes a log into a FundsRefunded struct.
func (c *Codec) DecodeFundsRefunded(log *evm.Log) (*FundsRefundedDecoded, error) {
	event := new(FundsRefundedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "FundsRefunded", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["FundsRefunded"].Inputs {
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

func (c *Codec) FundsReleasedLogHash() []byte {
	return c.abi.Events["FundsReleased"].ID.Bytes()
}

func (c *Codec) EncodeFundsReleasedTopics(
	evt abi.Event,
	values []FundsReleasedTopics,
) ([]*evm.TopicValues, error) {
	var dealIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.DealId).IsZero() {
			dealIdRule = append(dealIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.DealId)
		if err != nil {
			return nil, err
		}
		dealIdRule = append(dealIdRule, fieldVal)
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
		dealIdRule,
		influencerRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeFundsReleased decodes a log into a FundsReleased struct.
func (c *Codec) DecodeFundsReleased(log *evm.Log) (*FundsReleasedDecoded, error) {
	event := new(FundsReleasedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "FundsReleased", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["FundsReleased"].Inputs {
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

func (c EscrowV2) MAXFEEBPS(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeMAXFEEBPSMethodCall()
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
		return c.Codec.DecodeMAXFEEBPSMethodOutput(response.Data)
	})

}

func (c EscrowV2) NATIVEETHADDRESS(
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

func (c EscrowV2) Deals(
	runtime cre.Runtime,
	args DealsInput,
	blockNumber *big.Int,
) cre.Promise[DealsOutput] {
	calldata, err := c.Codec.EncodeDealsMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[DealsOutput](DealsOutput{}, err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (DealsOutput, error) {
		return c.Codec.DecodeDealsMethodOutput(response.Data)
	})

}

func (c EscrowV2) ExpectedAuthor(
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

func (c EscrowV2) ExpectedWorkflowId(
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

func (c EscrowV2) ExpectedWorkflowName(
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

func (c EscrowV2) FeeRecipient(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeFeeRecipientMethodCall()
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
		return c.Codec.DecodeFeeRecipientMethodOutput(response.Data)
	})

}

func (c EscrowV2) ForwarderAddress(
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

func (c EscrowV2) GetDeal(
	runtime cre.Runtime,
	args GetDealInput,
	blockNumber *big.Int,
) cre.Promise[AdEscrowV2Campaign] {
	calldata, err := c.Codec.EncodeGetDealMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[AdEscrowV2Campaign](*new(AdEscrowV2Campaign), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (AdEscrowV2Campaign, error) {
		return c.Codec.DecodeGetDealMethodOutput(response.Data)
	})

}

func (c EscrowV2) IsExpired(
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

func (c EscrowV2) Owner(
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

func (c EscrowV2) PlatformFeeBps(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodePlatformFeeBpsMethodCall()
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
		return c.Codec.DecodePlatformFeeBpsMethodOutput(response.Data)
	})

}

func (c EscrowV2) WhitelistedTokens(
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

func (c EscrowV2) WriteReportFromAdEscrowV2Campaign(
	runtime cre.Runtime,
	input AdEscrowV2Campaign,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeAdEscrowV2CampaignStruct(input)
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

func (c EscrowV2) WriteReportFromAdEscrowV2DepositParams(
	runtime cre.Runtime,
	input AdEscrowV2DepositParams,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeAdEscrowV2DepositParamsStruct(input)
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

func (c EscrowV2) WriteReport(
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

// DecodeAlreadyAcceptedError decodes a AlreadyAccepted error from revert data.
func (c *EscrowV2) DecodeAlreadyAcceptedError(data []byte) (*AlreadyAccepted, error) {
	args := c.ABI.Errors["AlreadyAccepted"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &AlreadyAccepted{}, nil
}

// Error implements the error interface for AlreadyAccepted.
func (e *AlreadyAccepted) Error() string {
	return fmt.Sprintf("AlreadyAccepted error:")
}

// DecodeCampaignExpiredError decodes a CampaignExpired error from revert data.
func (c *EscrowV2) DecodeCampaignExpiredError(data []byte) (*CampaignExpired, error) {
	args := c.ABI.Errors["CampaignExpired"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &CampaignExpired{}, nil
}

// Error implements the error interface for CampaignExpired.
func (e *CampaignExpired) Error() string {
	return fmt.Sprintf("CampaignExpired error:")
}

// DecodeCampaignNotExpiredError decodes a CampaignNotExpired error from revert data.
func (c *EscrowV2) DecodeCampaignNotExpiredError(data []byte) (*CampaignNotExpired, error) {
	args := c.ABI.Errors["CampaignNotExpired"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &CampaignNotExpired{}, nil
}

// Error implements the error interface for CampaignNotExpired.
func (e *CampaignNotExpired) Error() string {
	return fmt.Sprintf("CampaignNotExpired error:")
}

// DecodeDealAlreadyExistsError decodes a DealAlreadyExists error from revert data.
func (c *EscrowV2) DecodeDealAlreadyExistsError(data []byte) (*DealAlreadyExists, error) {
	args := c.ABI.Errors["DealAlreadyExists"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &DealAlreadyExists{}, nil
}

// Error implements the error interface for DealAlreadyExists.
func (e *DealAlreadyExists) Error() string {
	return fmt.Sprintf("DealAlreadyExists error:")
}

// DecodeFeeTooHighError decodes a FeeTooHigh error from revert data.
func (c *EscrowV2) DecodeFeeTooHighError(data []byte) (*FeeTooHigh, error) {
	args := c.ABI.Errors["FeeTooHigh"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &FeeTooHigh{}, nil
}

// Error implements the error interface for FeeTooHigh.
func (e *FeeTooHigh) Error() string {
	return fmt.Sprintf("FeeTooHigh error:")
}

// DecodeInvalidAmountError decodes a InvalidAmount error from revert data.
func (c *EscrowV2) DecodeInvalidAmountError(data []byte) (*InvalidAmount, error) {
	args := c.ABI.Errors["InvalidAmount"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &InvalidAmount{}, nil
}

// Error implements the error interface for InvalidAmount.
func (e *InvalidAmount) Error() string {
	return fmt.Sprintf("InvalidAmount error:")
}

// DecodeInvalidAuthorError decodes a InvalidAuthor error from revert data.
func (c *EscrowV2) DecodeInvalidAuthorError(data []byte) (*InvalidAuthor, error) {
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
func (c *EscrowV2) DecodeInvalidSenderError(data []byte) (*InvalidSender, error) {
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

// DecodeInvalidStateError decodes a InvalidState error from revert data.
func (c *EscrowV2) DecodeInvalidStateError(data []byte) (*InvalidState, error) {
	args := c.ABI.Errors["InvalidState"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 2 {
		return nil, fmt.Errorf("expected 2 values, got %d", len(values))
	}

	current, ok0 := values[0].(uint8)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for current in InvalidState error")
	}

	expected, ok1 := values[1].(uint8)
	if !ok1 {
		return nil, fmt.Errorf("unexpected type for expected in InvalidState error")
	}

	return &InvalidState{
		Current:  current,
		Expected: expected,
	}, nil
}

// Error implements the error interface for InvalidState.
func (e *InvalidState) Error() string {
	return fmt.Sprintf("InvalidState error: current=%v; expected=%v;", e.Current, e.Expected)
}

// DecodeInvalidWorkflowIdError decodes a InvalidWorkflowId error from revert data.
func (c *EscrowV2) DecodeInvalidWorkflowIdError(data []byte) (*InvalidWorkflowId, error) {
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
func (c *EscrowV2) DecodeInvalidWorkflowNameError(data []byte) (*InvalidWorkflowName, error) {
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

// DecodeNotAdvertiserError decodes a NotAdvertiser error from revert data.
func (c *EscrowV2) DecodeNotAdvertiserError(data []byte) (*NotAdvertiser, error) {
	args := c.ABI.Errors["NotAdvertiser"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &NotAdvertiser{}, nil
}

// Error implements the error interface for NotAdvertiser.
func (e *NotAdvertiser) Error() string {
	return fmt.Sprintf("NotAdvertiser error:")
}

// DecodeNotInfluencerError decodes a NotInfluencer error from revert data.
func (c *EscrowV2) DecodeNotInfluencerError(data []byte) (*NotInfluencer, error) {
	args := c.ABI.Errors["NotInfluencer"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &NotInfluencer{}, nil
}

// Error implements the error interface for NotInfluencer.
func (e *NotInfluencer) Error() string {
	return fmt.Sprintf("NotInfluencer error:")
}

// DecodeNotParticipantError decodes a NotParticipant error from revert data.
func (c *EscrowV2) DecodeNotParticipantError(data []byte) (*NotParticipant, error) {
	args := c.ABI.Errors["NotParticipant"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &NotParticipant{}, nil
}

// Error implements the error interface for NotParticipant.
func (e *NotParticipant) Error() string {
	return fmt.Sprintf("NotParticipant error:")
}

// DecodeOwnableInvalidOwnerError decodes a OwnableInvalidOwner error from revert data.
func (c *EscrowV2) DecodeOwnableInvalidOwnerError(data []byte) (*OwnableInvalidOwner, error) {
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
func (c *EscrowV2) DecodeOwnableUnauthorizedAccountError(data []byte) (*OwnableUnauthorizedAccount, error) {
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

// DecodeReentrancyGuardReentrantCallError decodes a ReentrancyGuardReentrantCall error from revert data.
func (c *EscrowV2) DecodeReentrancyGuardReentrantCallError(data []byte) (*ReentrancyGuardReentrantCall, error) {
	args := c.ABI.Errors["ReentrancyGuardReentrantCall"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &ReentrancyGuardReentrantCall{}, nil
}

// Error implements the error interface for ReentrancyGuardReentrantCall.
func (e *ReentrancyGuardReentrantCall) Error() string {
	return fmt.Sprintf("ReentrancyGuardReentrantCall error:")
}

// DecodeSafeERC20FailedOperationError decodes a SafeERC20FailedOperation error from revert data.
func (c *EscrowV2) DecodeSafeERC20FailedOperationError(data []byte) (*SafeERC20FailedOperation, error) {
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

// DecodeTokenNotWhitelistedError decodes a TokenNotWhitelisted error from revert data.
func (c *EscrowV2) DecodeTokenNotWhitelistedError(data []byte) (*TokenNotWhitelisted, error) {
	args := c.ABI.Errors["TokenNotWhitelisted"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &TokenNotWhitelisted{}, nil
}

// Error implements the error interface for TokenNotWhitelisted.
func (e *TokenNotWhitelisted) Error() string {
	return fmt.Sprintf("TokenNotWhitelisted error:")
}

func (c *EscrowV2) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	case common.Bytes2Hex(c.ABI.Errors["AlreadyAccepted"].ID.Bytes()[:4]):
		return c.DecodeAlreadyAcceptedError(data)
	case common.Bytes2Hex(c.ABI.Errors["CampaignExpired"].ID.Bytes()[:4]):
		return c.DecodeCampaignExpiredError(data)
	case common.Bytes2Hex(c.ABI.Errors["CampaignNotExpired"].ID.Bytes()[:4]):
		return c.DecodeCampaignNotExpiredError(data)
	case common.Bytes2Hex(c.ABI.Errors["DealAlreadyExists"].ID.Bytes()[:4]):
		return c.DecodeDealAlreadyExistsError(data)
	case common.Bytes2Hex(c.ABI.Errors["FeeTooHigh"].ID.Bytes()[:4]):
		return c.DecodeFeeTooHighError(data)
	case common.Bytes2Hex(c.ABI.Errors["InvalidAmount"].ID.Bytes()[:4]):
		return c.DecodeInvalidAmountError(data)
	case common.Bytes2Hex(c.ABI.Errors["InvalidAuthor"].ID.Bytes()[:4]):
		return c.DecodeInvalidAuthorError(data)
	case common.Bytes2Hex(c.ABI.Errors["InvalidSender"].ID.Bytes()[:4]):
		return c.DecodeInvalidSenderError(data)
	case common.Bytes2Hex(c.ABI.Errors["InvalidState"].ID.Bytes()[:4]):
		return c.DecodeInvalidStateError(data)
	case common.Bytes2Hex(c.ABI.Errors["InvalidWorkflowId"].ID.Bytes()[:4]):
		return c.DecodeInvalidWorkflowIdError(data)
	case common.Bytes2Hex(c.ABI.Errors["InvalidWorkflowName"].ID.Bytes()[:4]):
		return c.DecodeInvalidWorkflowNameError(data)
	case common.Bytes2Hex(c.ABI.Errors["NotAdvertiser"].ID.Bytes()[:4]):
		return c.DecodeNotAdvertiserError(data)
	case common.Bytes2Hex(c.ABI.Errors["NotInfluencer"].ID.Bytes()[:4]):
		return c.DecodeNotInfluencerError(data)
	case common.Bytes2Hex(c.ABI.Errors["NotParticipant"].ID.Bytes()[:4]):
		return c.DecodeNotParticipantError(data)
	case common.Bytes2Hex(c.ABI.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]):
		return c.DecodeOwnableInvalidOwnerError(data)
	case common.Bytes2Hex(c.ABI.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]):
		return c.DecodeOwnableUnauthorizedAccountError(data)
	case common.Bytes2Hex(c.ABI.Errors["ReentrancyGuardReentrantCall"].ID.Bytes()[:4]):
		return c.DecodeReentrancyGuardReentrantCallError(data)
	case common.Bytes2Hex(c.ABI.Errors["SafeERC20FailedOperation"].ID.Bytes()[:4]):
		return c.DecodeSafeERC20FailedOperationError(data)
	case common.Bytes2Hex(c.ABI.Errors["TokenNotWhitelisted"].ID.Bytes()[:4]):
		return c.DecodeTokenNotWhitelistedError(data)
	default:
		return nil, errors.New("unknown error selector")
	}
}

// DealAcceptedTrigger wraps the raw log trigger and provides decoded DealAcceptedDecoded data
type DealAcceptedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]           // Embed the raw trigger
	contract                        *EscrowV2 // Keep reference for decoding
}

// Adapt method that decodes the log into DealAccepted data
func (t *DealAcceptedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[DealAcceptedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeDealAccepted(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode DealAccepted log: %w", err)
	}

	return &bindings.DecodedLog[DealAcceptedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *EscrowV2) LogTriggerDealAcceptedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []DealAcceptedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[DealAcceptedDecoded]], error) {
	event := c.ABI.Events["DealAccepted"]
	topics, err := c.Codec.EncodeDealAcceptedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for DealAccepted: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &DealAcceptedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *EscrowV2) FilterLogsDealAccepted(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.DealAcceptedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// DealCancelledTrigger wraps the raw log trigger and provides decoded DealCancelledDecoded data
type DealCancelledTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]           // Embed the raw trigger
	contract                        *EscrowV2 // Keep reference for decoding
}

// Adapt method that decodes the log into DealCancelled data
func (t *DealCancelledTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[DealCancelledDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeDealCancelled(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode DealCancelled log: %w", err)
	}

	return &bindings.DecodedLog[DealCancelledDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *EscrowV2) LogTriggerDealCancelledLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []DealCancelledTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[DealCancelledDecoded]], error) {
	event := c.ABI.Events["DealCancelled"]
	topics, err := c.Codec.EncodeDealCancelledTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for DealCancelled: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &DealCancelledTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *EscrowV2) FilterLogsDealCancelled(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.DealCancelledLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// DealCreatedTrigger wraps the raw log trigger and provides decoded DealCreatedDecoded data
type DealCreatedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]           // Embed the raw trigger
	contract                        *EscrowV2 // Keep reference for decoding
}

// Adapt method that decodes the log into DealCreated data
func (t *DealCreatedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[DealCreatedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeDealCreated(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode DealCreated log: %w", err)
	}

	return &bindings.DecodedLog[DealCreatedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *EscrowV2) LogTriggerDealCreatedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []DealCreatedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[DealCreatedDecoded]], error) {
	event := c.ABI.Events["DealCreated"]
	topics, err := c.Codec.EncodeDealCreatedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for DealCreated: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &DealCreatedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *EscrowV2) FilterLogsDealCreated(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.DealCreatedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// DisputeRaisedTrigger wraps the raw log trigger and provides decoded DisputeRaisedDecoded data
type DisputeRaisedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]           // Embed the raw trigger
	contract                        *EscrowV2 // Keep reference for decoding
}

// Adapt method that decodes the log into DisputeRaised data
func (t *DisputeRaisedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[DisputeRaisedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeDisputeRaised(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode DisputeRaised log: %w", err)
	}

	return &bindings.DecodedLog[DisputeRaisedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *EscrowV2) LogTriggerDisputeRaisedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []DisputeRaisedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[DisputeRaisedDecoded]], error) {
	event := c.ABI.Events["DisputeRaised"]
	topics, err := c.Codec.EncodeDisputeRaisedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for DisputeRaised: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &DisputeRaisedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *EscrowV2) FilterLogsDisputeRaised(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.DisputeRaisedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// DisputeResolvedTrigger wraps the raw log trigger and provides decoded DisputeResolvedDecoded data
type DisputeResolvedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]           // Embed the raw trigger
	contract                        *EscrowV2 // Keep reference for decoding
}

// Adapt method that decodes the log into DisputeResolved data
func (t *DisputeResolvedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[DisputeResolvedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeDisputeResolved(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode DisputeResolved log: %w", err)
	}

	return &bindings.DecodedLog[DisputeResolvedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *EscrowV2) LogTriggerDisputeResolvedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []DisputeResolvedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[DisputeResolvedDecoded]], error) {
	event := c.ABI.Events["DisputeResolved"]
	topics, err := c.Codec.EncodeDisputeResolvedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for DisputeResolved: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &DisputeResolvedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *EscrowV2) FilterLogsDisputeResolved(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.DisputeResolvedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// FundsRefundedTrigger wraps the raw log trigger and provides decoded FundsRefundedDecoded data
type FundsRefundedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]           // Embed the raw trigger
	contract                        *EscrowV2 // Keep reference for decoding
}

// Adapt method that decodes the log into FundsRefunded data
func (t *FundsRefundedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[FundsRefundedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeFundsRefunded(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode FundsRefunded log: %w", err)
	}

	return &bindings.DecodedLog[FundsRefundedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *EscrowV2) LogTriggerFundsRefundedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []FundsRefundedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[FundsRefundedDecoded]], error) {
	event := c.ABI.Events["FundsRefunded"]
	topics, err := c.Codec.EncodeFundsRefundedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for FundsRefunded: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &FundsRefundedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *EscrowV2) FilterLogsFundsRefunded(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.FundsRefundedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// FundsReleasedTrigger wraps the raw log trigger and provides decoded FundsReleasedDecoded data
type FundsReleasedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]           // Embed the raw trigger
	contract                        *EscrowV2 // Keep reference for decoding
}

// Adapt method that decodes the log into FundsReleased data
func (t *FundsReleasedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[FundsReleasedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeFundsReleased(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode FundsReleased log: %w", err)
	}

	return &bindings.DecodedLog[FundsReleasedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *EscrowV2) LogTriggerFundsReleasedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []FundsReleasedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[FundsReleasedDecoded]], error) {
	event := c.ABI.Events["FundsReleased"]
	topics, err := c.Codec.EncodeFundsReleasedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for FundsReleased: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &FundsReleasedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *EscrowV2) FilterLogsFundsReleased(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.FundsReleasedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// OwnershipTransferredTrigger wraps the raw log trigger and provides decoded OwnershipTransferredDecoded data
type OwnershipTransferredTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]           // Embed the raw trigger
	contract                        *EscrowV2 // Keep reference for decoding
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

func (c *EscrowV2) LogTriggerOwnershipTransferredLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []OwnershipTransferredTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[OwnershipTransferredDecoded]], error) {
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

func (c *EscrowV2) FilterLogsOwnershipTransferred(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
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
	cre.Trigger[*evm.Log, *evm.Log]           // Embed the raw trigger
	contract                        *EscrowV2 // Keep reference for decoding
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

func (c *EscrowV2) LogTriggerTokenRemovedFromWhitelistLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []TokenRemovedFromWhitelistTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[TokenRemovedFromWhitelistDecoded]], error) {
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

func (c *EscrowV2) FilterLogsTokenRemovedFromWhitelist(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
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
	cre.Trigger[*evm.Log, *evm.Log]           // Embed the raw trigger
	contract                        *EscrowV2 // Keep reference for decoding
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

func (c *EscrowV2) LogTriggerTokenWhitelistedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []TokenWhitelistedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[TokenWhitelistedDecoded]], error) {
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

func (c *EscrowV2) FilterLogsTokenWhitelisted(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
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
