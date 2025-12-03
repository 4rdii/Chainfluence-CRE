// Code generated — DO NOT EDIT.

//go:build !wasip1

package storage

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

// StorageMock is a mock implementation of Storage for testing.
type StorageMock struct {
	Get   func() (*big.Int, error)
	Value func() (*big.Int, error)
}

// NewStorageMock creates a new StorageMock for testing.
func NewStorageMock(address common.Address, clientMock *evmmock.ClientCapability) *StorageMock {
	mock := &StorageMock{}

	codec, err := NewCodec()
	if err != nil {
		panic("failed to create codec for mock: " + err.Error())
	}

	abi := codec.(*Codec).abi
	_ = abi

	funcMap := map[string]func([]byte) ([]byte, error){
		string(abi.Methods["get"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Get == nil {
				return nil, errors.New("get method not mocked")
			}
			result, err := mock.Get()
			if err != nil {
				return nil, err
			}
			return abi.Methods["get"].Outputs.Pack(result)
		},
		string(abi.Methods["value"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Value == nil {
				return nil, errors.New("value method not mocked")
			}
			result, err := mock.Value()
			if err != nil {
				return nil, err
			}
			return abi.Methods["value"].Outputs.Pack(result)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}
