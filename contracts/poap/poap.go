// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package POAPContract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// POAPContractMetaData contains all meta data concerning the POAPContract contract.
var POAPContractMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"eventId\",\"type\":\"uint256\"}],\"name\":\"renounceEventMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenEvent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"eventId\",\"type\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeEventMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newLastId\",\"type\":\"uint256\"}],\"name\":\"setLastId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"eventId\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"address[]\"}],\"name\":\"mintEventToManyUsers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"eventId\",\"type\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isEventMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freezeDuration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"unfreeze\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenDetailsOfOwnerByIndex\",\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"eventId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"setFreezeDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"__name\",\"type\":\"string\"},{\"name\":\"__symbol\",\"type\":\"string\"},{\"name\":\"__baseURI\",\"type\":\"string\"},{\"name\":\"admins\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getFreezeTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"eventId\",\"type\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addEventMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isFrozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"eventId\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"freeze\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"eventIds\",\"type\":\"uint256[]\"},{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintUserToManyEvents\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"eventId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"EventToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Frozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Unfrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"eventId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"EventMinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"eventId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"EventMinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]",
}

// POAPContractABI is the input ABI used to generate the binding from.
// Deprecated: Use POAPContractMetaData.ABI instead.
var POAPContractABI = POAPContractMetaData.ABI

// POAPContract is an auto generated Go binding around an Ethereum contract.
type POAPContract struct {
	POAPContractCaller     // Read-only binding to the contract
	POAPContractTransactor // Write-only binding to the contract
	POAPContractFilterer   // Log filterer for contract events
}

// POAPContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type POAPContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// POAPContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type POAPContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// POAPContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type POAPContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// POAPContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type POAPContractSession struct {
	Contract     *POAPContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// POAPContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type POAPContractCallerSession struct {
	Contract *POAPContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// POAPContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type POAPContractTransactorSession struct {
	Contract     *POAPContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// POAPContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type POAPContractRaw struct {
	Contract *POAPContract // Generic contract binding to access the raw methods on
}

// POAPContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type POAPContractCallerRaw struct {
	Contract *POAPContractCaller // Generic read-only contract binding to access the raw methods on
}

// POAPContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type POAPContractTransactorRaw struct {
	Contract *POAPContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPOAPContract creates a new instance of POAPContract, bound to a specific deployed contract.
func NewPOAPContract(address common.Address, backend bind.ContractBackend) (*POAPContract, error) {
	contract, err := bindPOAPContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &POAPContract{POAPContractCaller: POAPContractCaller{contract: contract}, POAPContractTransactor: POAPContractTransactor{contract: contract}, POAPContractFilterer: POAPContractFilterer{contract: contract}}, nil
}

// NewPOAPContractCaller creates a new read-only instance of POAPContract, bound to a specific deployed contract.
func NewPOAPContractCaller(address common.Address, caller bind.ContractCaller) (*POAPContractCaller, error) {
	contract, err := bindPOAPContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &POAPContractCaller{contract: contract}, nil
}

// NewPOAPContractTransactor creates a new write-only instance of POAPContract, bound to a specific deployed contract.
func NewPOAPContractTransactor(address common.Address, transactor bind.ContractTransactor) (*POAPContractTransactor, error) {
	contract, err := bindPOAPContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &POAPContractTransactor{contract: contract}, nil
}

// NewPOAPContractFilterer creates a new log filterer instance of POAPContract, bound to a specific deployed contract.
func NewPOAPContractFilterer(address common.Address, filterer bind.ContractFilterer) (*POAPContractFilterer, error) {
	contract, err := bindPOAPContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &POAPContractFilterer{contract: contract}, nil
}

// bindPOAPContract binds a generic wrapper to an already deployed contract.
func bindPOAPContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := POAPContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_POAPContract *POAPContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _POAPContract.Contract.POAPContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_POAPContract *POAPContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _POAPContract.Contract.POAPContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_POAPContract *POAPContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _POAPContract.Contract.POAPContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_POAPContract *POAPContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _POAPContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_POAPContract *POAPContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _POAPContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_POAPContract *POAPContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _POAPContract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_POAPContract *POAPContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_POAPContract *POAPContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _POAPContract.Contract.BalanceOf(&_POAPContract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_POAPContract *POAPContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _POAPContract.Contract.BalanceOf(&_POAPContract.CallOpts, owner)
}

// FreezeDuration is a free data retrieval call binding the contract method 0x440991bd.
//
// Solidity: function freezeDuration() view returns(uint256)
func (_POAPContract *POAPContractCaller) FreezeDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "freezeDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreezeDuration is a free data retrieval call binding the contract method 0x440991bd.
//
// Solidity: function freezeDuration() view returns(uint256)
func (_POAPContract *POAPContractSession) FreezeDuration() (*big.Int, error) {
	return _POAPContract.Contract.FreezeDuration(&_POAPContract.CallOpts)
}

// FreezeDuration is a free data retrieval call binding the contract method 0x440991bd.
//
// Solidity: function freezeDuration() view returns(uint256)
func (_POAPContract *POAPContractCallerSession) FreezeDuration() (*big.Int, error) {
	return _POAPContract.Contract.FreezeDuration(&_POAPContract.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_POAPContract *POAPContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_POAPContract *POAPContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _POAPContract.Contract.GetApproved(&_POAPContract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_POAPContract *POAPContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _POAPContract.Contract.GetApproved(&_POAPContract.CallOpts, tokenId)
}

// GetFreezeTime is a free data retrieval call binding the contract method 0x90fdd897.
//
// Solidity: function getFreezeTime(uint256 tokenId) view returns(uint256)
func (_POAPContract *POAPContractCaller) GetFreezeTime(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "getFreezeTime", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFreezeTime is a free data retrieval call binding the contract method 0x90fdd897.
//
// Solidity: function getFreezeTime(uint256 tokenId) view returns(uint256)
func (_POAPContract *POAPContractSession) GetFreezeTime(tokenId *big.Int) (*big.Int, error) {
	return _POAPContract.Contract.GetFreezeTime(&_POAPContract.CallOpts, tokenId)
}

// GetFreezeTime is a free data retrieval call binding the contract method 0x90fdd897.
//
// Solidity: function getFreezeTime(uint256 tokenId) view returns(uint256)
func (_POAPContract *POAPContractCallerSession) GetFreezeTime(tokenId *big.Int) (*big.Int, error) {
	return _POAPContract.Contract.GetFreezeTime(&_POAPContract.CallOpts, tokenId)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_POAPContract *POAPContractCaller) IsAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "isAdmin", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_POAPContract *POAPContractSession) IsAdmin(account common.Address) (bool, error) {
	return _POAPContract.Contract.IsAdmin(&_POAPContract.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_POAPContract *POAPContractCallerSession) IsAdmin(account common.Address) (bool, error) {
	return _POAPContract.Contract.IsAdmin(&_POAPContract.CallOpts, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_POAPContract *POAPContractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_POAPContract *POAPContractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _POAPContract.Contract.IsApprovedForAll(&_POAPContract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_POAPContract *POAPContractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _POAPContract.Contract.IsApprovedForAll(&_POAPContract.CallOpts, owner, operator)
}

// IsEventMinter is a free data retrieval call binding the contract method 0x28db38b4.
//
// Solidity: function isEventMinter(uint256 eventId, address account) view returns(bool)
func (_POAPContract *POAPContractCaller) IsEventMinter(opts *bind.CallOpts, eventId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "isEventMinter", eventId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEventMinter is a free data retrieval call binding the contract method 0x28db38b4.
//
// Solidity: function isEventMinter(uint256 eventId, address account) view returns(bool)
func (_POAPContract *POAPContractSession) IsEventMinter(eventId *big.Int, account common.Address) (bool, error) {
	return _POAPContract.Contract.IsEventMinter(&_POAPContract.CallOpts, eventId, account)
}

// IsEventMinter is a free data retrieval call binding the contract method 0x28db38b4.
//
// Solidity: function isEventMinter(uint256 eventId, address account) view returns(bool)
func (_POAPContract *POAPContractCallerSession) IsEventMinter(eventId *big.Int, account common.Address) (bool, error) {
	return _POAPContract.Contract.IsEventMinter(&_POAPContract.CallOpts, eventId, account)
}

// IsFrozen is a free data retrieval call binding the contract method 0xa0894799.
//
// Solidity: function isFrozen(uint256 tokenId) view returns(bool)
func (_POAPContract *POAPContractCaller) IsFrozen(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "isFrozen", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0xa0894799.
//
// Solidity: function isFrozen(uint256 tokenId) view returns(bool)
func (_POAPContract *POAPContractSession) IsFrozen(tokenId *big.Int) (bool, error) {
	return _POAPContract.Contract.IsFrozen(&_POAPContract.CallOpts, tokenId)
}

// IsFrozen is a free data retrieval call binding the contract method 0xa0894799.
//
// Solidity: function isFrozen(uint256 tokenId) view returns(bool)
func (_POAPContract *POAPContractCallerSession) IsFrozen(tokenId *big.Int) (bool, error) {
	return _POAPContract.Contract.IsFrozen(&_POAPContract.CallOpts, tokenId)
}

// LastId is a free data retrieval call binding the contract method 0xc1292cc3.
//
// Solidity: function lastId() view returns(uint256)
func (_POAPContract *POAPContractCaller) LastId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "lastId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastId is a free data retrieval call binding the contract method 0xc1292cc3.
//
// Solidity: function lastId() view returns(uint256)
func (_POAPContract *POAPContractSession) LastId() (*big.Int, error) {
	return _POAPContract.Contract.LastId(&_POAPContract.CallOpts)
}

// LastId is a free data retrieval call binding the contract method 0xc1292cc3.
//
// Solidity: function lastId() view returns(uint256)
func (_POAPContract *POAPContractCallerSession) LastId() (*big.Int, error) {
	return _POAPContract.Contract.LastId(&_POAPContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_POAPContract *POAPContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_POAPContract *POAPContractSession) Name() (string, error) {
	return _POAPContract.Contract.Name(&_POAPContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_POAPContract *POAPContractCallerSession) Name() (string, error) {
	return _POAPContract.Contract.Name(&_POAPContract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_POAPContract *POAPContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_POAPContract *POAPContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _POAPContract.Contract.OwnerOf(&_POAPContract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_POAPContract *POAPContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _POAPContract.Contract.OwnerOf(&_POAPContract.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_POAPContract *POAPContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_POAPContract *POAPContractSession) Paused() (bool, error) {
	return _POAPContract.Contract.Paused(&_POAPContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_POAPContract *POAPContractCallerSession) Paused() (bool, error) {
	return _POAPContract.Contract.Paused(&_POAPContract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_POAPContract *POAPContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_POAPContract *POAPContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _POAPContract.Contract.SupportsInterface(&_POAPContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_POAPContract *POAPContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _POAPContract.Contract.SupportsInterface(&_POAPContract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_POAPContract *POAPContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_POAPContract *POAPContractSession) Symbol() (string, error) {
	return _POAPContract.Contract.Symbol(&_POAPContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_POAPContract *POAPContractCallerSession) Symbol() (string, error) {
	return _POAPContract.Contract.Symbol(&_POAPContract.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_POAPContract *POAPContractCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_POAPContract *POAPContractSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _POAPContract.Contract.TokenByIndex(&_POAPContract.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_POAPContract *POAPContractCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _POAPContract.Contract.TokenByIndex(&_POAPContract.CallOpts, index)
}

// TokenDetailsOfOwnerByIndex is a free data retrieval call binding the contract method 0x67e971ce.
//
// Solidity: function tokenDetailsOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId, uint256 eventId)
func (_POAPContract *POAPContractCaller) TokenDetailsOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (struct {
	TokenId *big.Int
	EventId *big.Int
}, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "tokenDetailsOfOwnerByIndex", owner, index)

	outstruct := new(struct {
		TokenId *big.Int
		EventId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EventId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenDetailsOfOwnerByIndex is a free data retrieval call binding the contract method 0x67e971ce.
//
// Solidity: function tokenDetailsOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId, uint256 eventId)
func (_POAPContract *POAPContractSession) TokenDetailsOfOwnerByIndex(owner common.Address, index *big.Int) (struct {
	TokenId *big.Int
	EventId *big.Int
}, error) {
	return _POAPContract.Contract.TokenDetailsOfOwnerByIndex(&_POAPContract.CallOpts, owner, index)
}

// TokenDetailsOfOwnerByIndex is a free data retrieval call binding the contract method 0x67e971ce.
//
// Solidity: function tokenDetailsOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId, uint256 eventId)
func (_POAPContract *POAPContractCallerSession) TokenDetailsOfOwnerByIndex(owner common.Address, index *big.Int) (struct {
	TokenId *big.Int
	EventId *big.Int
}, error) {
	return _POAPContract.Contract.TokenDetailsOfOwnerByIndex(&_POAPContract.CallOpts, owner, index)
}

// TokenEvent is a free data retrieval call binding the contract method 0x127a5298.
//
// Solidity: function tokenEvent(uint256 tokenId) view returns(uint256)
func (_POAPContract *POAPContractCaller) TokenEvent(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "tokenEvent", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenEvent is a free data retrieval call binding the contract method 0x127a5298.
//
// Solidity: function tokenEvent(uint256 tokenId) view returns(uint256)
func (_POAPContract *POAPContractSession) TokenEvent(tokenId *big.Int) (*big.Int, error) {
	return _POAPContract.Contract.TokenEvent(&_POAPContract.CallOpts, tokenId)
}

// TokenEvent is a free data retrieval call binding the contract method 0x127a5298.
//
// Solidity: function tokenEvent(uint256 tokenId) view returns(uint256)
func (_POAPContract *POAPContractCallerSession) TokenEvent(tokenId *big.Int) (*big.Int, error) {
	return _POAPContract.Contract.TokenEvent(&_POAPContract.CallOpts, tokenId)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_POAPContract *POAPContractCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_POAPContract *POAPContractSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _POAPContract.Contract.TokenOfOwnerByIndex(&_POAPContract.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_POAPContract *POAPContractCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _POAPContract.Contract.TokenOfOwnerByIndex(&_POAPContract.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_POAPContract *POAPContractCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_POAPContract *POAPContractSession) TokenURI(tokenId *big.Int) (string, error) {
	return _POAPContract.Contract.TokenURI(&_POAPContract.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_POAPContract *POAPContractCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _POAPContract.Contract.TokenURI(&_POAPContract.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_POAPContract *POAPContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _POAPContract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_POAPContract *POAPContractSession) TotalSupply() (*big.Int, error) {
	return _POAPContract.Contract.TotalSupply(&_POAPContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_POAPContract *POAPContractCallerSession) TotalSupply() (*big.Int, error) {
	return _POAPContract.Contract.TotalSupply(&_POAPContract.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_POAPContract *POAPContractTransactor) AddAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "addAdmin", account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_POAPContract *POAPContractSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.AddAdmin(&_POAPContract.TransactOpts, account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_POAPContract *POAPContractTransactorSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.AddAdmin(&_POAPContract.TransactOpts, account)
}

// AddEventMinter is a paid mutator transaction binding the contract method 0x9cd3cad6.
//
// Solidity: function addEventMinter(uint256 eventId, address account) returns()
func (_POAPContract *POAPContractTransactor) AddEventMinter(opts *bind.TransactOpts, eventId *big.Int, account common.Address) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "addEventMinter", eventId, account)
}

// AddEventMinter is a paid mutator transaction binding the contract method 0x9cd3cad6.
//
// Solidity: function addEventMinter(uint256 eventId, address account) returns()
func (_POAPContract *POAPContractSession) AddEventMinter(eventId *big.Int, account common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.AddEventMinter(&_POAPContract.TransactOpts, eventId, account)
}

// AddEventMinter is a paid mutator transaction binding the contract method 0x9cd3cad6.
//
// Solidity: function addEventMinter(uint256 eventId, address account) returns()
func (_POAPContract *POAPContractTransactorSession) AddEventMinter(eventId *big.Int, account common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.AddEventMinter(&_POAPContract.TransactOpts, eventId, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_POAPContract *POAPContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_POAPContract *POAPContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.Approve(&_POAPContract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_POAPContract *POAPContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.Approve(&_POAPContract.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_POAPContract *POAPContractTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_POAPContract *POAPContractSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.Burn(&_POAPContract.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_POAPContract *POAPContractTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.Burn(&_POAPContract.TransactOpts, tokenId)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 tokenId) returns()
func (_POAPContract *POAPContractTransactor) Freeze(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "freeze", tokenId)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 tokenId) returns()
func (_POAPContract *POAPContractSession) Freeze(tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.Freeze(&_POAPContract.TransactOpts, tokenId)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 tokenId) returns()
func (_POAPContract *POAPContractTransactorSession) Freeze(tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.Freeze(&_POAPContract.TransactOpts, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_POAPContract *POAPContractTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_POAPContract *POAPContractSession) Initialize() (*types.Transaction, error) {
	return _POAPContract.Contract.Initialize(&_POAPContract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_POAPContract *POAPContractTransactorSession) Initialize() (*types.Transaction, error) {
	return _POAPContract.Contract.Initialize(&_POAPContract.TransactOpts)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x8d232094.
//
// Solidity: function initialize(string __name, string __symbol, string __baseURI, address[] admins) returns()
func (_POAPContract *POAPContractTransactor) Initialize0(opts *bind.TransactOpts, __name string, __symbol string, __baseURI string, admins []common.Address) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "initialize0", __name, __symbol, __baseURI, admins)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x8d232094.
//
// Solidity: function initialize(string __name, string __symbol, string __baseURI, address[] admins) returns()
func (_POAPContract *POAPContractSession) Initialize0(__name string, __symbol string, __baseURI string, admins []common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.Initialize0(&_POAPContract.TransactOpts, __name, __symbol, __baseURI, admins)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x8d232094.
//
// Solidity: function initialize(string __name, string __symbol, string __baseURI, address[] admins) returns()
func (_POAPContract *POAPContractTransactorSession) Initialize0(__name string, __symbol string, __baseURI string, admins []common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.Initialize0(&_POAPContract.TransactOpts, __name, __symbol, __baseURI, admins)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_POAPContract *POAPContractTransactor) Initialize1(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "initialize1", sender)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_POAPContract *POAPContractSession) Initialize1(sender common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.Initialize1(&_POAPContract.TransactOpts, sender)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_POAPContract *POAPContractTransactorSession) Initialize1(sender common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.Initialize1(&_POAPContract.TransactOpts, sender)
}

// MintEventToManyUsers is a paid mutator transaction binding the contract method 0x278d9c41.
//
// Solidity: function mintEventToManyUsers(uint256 eventId, address[] to) returns(bool)
func (_POAPContract *POAPContractTransactor) MintEventToManyUsers(opts *bind.TransactOpts, eventId *big.Int, to []common.Address) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "mintEventToManyUsers", eventId, to)
}

// MintEventToManyUsers is a paid mutator transaction binding the contract method 0x278d9c41.
//
// Solidity: function mintEventToManyUsers(uint256 eventId, address[] to) returns(bool)
func (_POAPContract *POAPContractSession) MintEventToManyUsers(eventId *big.Int, to []common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.MintEventToManyUsers(&_POAPContract.TransactOpts, eventId, to)
}

// MintEventToManyUsers is a paid mutator transaction binding the contract method 0x278d9c41.
//
// Solidity: function mintEventToManyUsers(uint256 eventId, address[] to) returns(bool)
func (_POAPContract *POAPContractTransactorSession) MintEventToManyUsers(eventId *big.Int, to []common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.MintEventToManyUsers(&_POAPContract.TransactOpts, eventId, to)
}

// MintToken is a paid mutator transaction binding the contract method 0xa140ae23.
//
// Solidity: function mintToken(uint256 eventId, address to) returns(bool)
func (_POAPContract *POAPContractTransactor) MintToken(opts *bind.TransactOpts, eventId *big.Int, to common.Address) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "mintToken", eventId, to)
}

// MintToken is a paid mutator transaction binding the contract method 0xa140ae23.
//
// Solidity: function mintToken(uint256 eventId, address to) returns(bool)
func (_POAPContract *POAPContractSession) MintToken(eventId *big.Int, to common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.MintToken(&_POAPContract.TransactOpts, eventId, to)
}

// MintToken is a paid mutator transaction binding the contract method 0xa140ae23.
//
// Solidity: function mintToken(uint256 eventId, address to) returns(bool)
func (_POAPContract *POAPContractTransactorSession) MintToken(eventId *big.Int, to common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.MintToken(&_POAPContract.TransactOpts, eventId, to)
}

// MintUserToManyEvents is a paid mutator transaction binding the contract method 0xf980f3dc.
//
// Solidity: function mintUserToManyEvents(uint256[] eventIds, address to) returns(bool)
func (_POAPContract *POAPContractTransactor) MintUserToManyEvents(opts *bind.TransactOpts, eventIds []*big.Int, to common.Address) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "mintUserToManyEvents", eventIds, to)
}

// MintUserToManyEvents is a paid mutator transaction binding the contract method 0xf980f3dc.
//
// Solidity: function mintUserToManyEvents(uint256[] eventIds, address to) returns(bool)
func (_POAPContract *POAPContractSession) MintUserToManyEvents(eventIds []*big.Int, to common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.MintUserToManyEvents(&_POAPContract.TransactOpts, eventIds, to)
}

// MintUserToManyEvents is a paid mutator transaction binding the contract method 0xf980f3dc.
//
// Solidity: function mintUserToManyEvents(uint256[] eventIds, address to) returns(bool)
func (_POAPContract *POAPContractTransactorSession) MintUserToManyEvents(eventIds []*big.Int, to common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.MintUserToManyEvents(&_POAPContract.TransactOpts, eventIds, to)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_POAPContract *POAPContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_POAPContract *POAPContractSession) Pause() (*types.Transaction, error) {
	return _POAPContract.Contract.Pause(&_POAPContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_POAPContract *POAPContractTransactorSession) Pause() (*types.Transaction, error) {
	return _POAPContract.Contract.Pause(&_POAPContract.TransactOpts)
}

// RemoveEventMinter is a paid mutator transaction binding the contract method 0x166c4b05.
//
// Solidity: function removeEventMinter(uint256 eventId, address account) returns()
func (_POAPContract *POAPContractTransactor) RemoveEventMinter(opts *bind.TransactOpts, eventId *big.Int, account common.Address) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "removeEventMinter", eventId, account)
}

// RemoveEventMinter is a paid mutator transaction binding the contract method 0x166c4b05.
//
// Solidity: function removeEventMinter(uint256 eventId, address account) returns()
func (_POAPContract *POAPContractSession) RemoveEventMinter(eventId *big.Int, account common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.RemoveEventMinter(&_POAPContract.TransactOpts, eventId, account)
}

// RemoveEventMinter is a paid mutator transaction binding the contract method 0x166c4b05.
//
// Solidity: function removeEventMinter(uint256 eventId, address account) returns()
func (_POAPContract *POAPContractTransactorSession) RemoveEventMinter(eventId *big.Int, account common.Address) (*types.Transaction, error) {
	return _POAPContract.Contract.RemoveEventMinter(&_POAPContract.TransactOpts, eventId, account)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_POAPContract *POAPContractTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_POAPContract *POAPContractSession) RenounceAdmin() (*types.Transaction, error) {
	return _POAPContract.Contract.RenounceAdmin(&_POAPContract.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_POAPContract *POAPContractTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _POAPContract.Contract.RenounceAdmin(&_POAPContract.TransactOpts)
}

// RenounceEventMinter is a paid mutator transaction binding the contract method 0x02c37ddc.
//
// Solidity: function renounceEventMinter(uint256 eventId) returns()
func (_POAPContract *POAPContractTransactor) RenounceEventMinter(opts *bind.TransactOpts, eventId *big.Int) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "renounceEventMinter", eventId)
}

// RenounceEventMinter is a paid mutator transaction binding the contract method 0x02c37ddc.
//
// Solidity: function renounceEventMinter(uint256 eventId) returns()
func (_POAPContract *POAPContractSession) RenounceEventMinter(eventId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.RenounceEventMinter(&_POAPContract.TransactOpts, eventId)
}

// RenounceEventMinter is a paid mutator transaction binding the contract method 0x02c37ddc.
//
// Solidity: function renounceEventMinter(uint256 eventId) returns()
func (_POAPContract *POAPContractTransactorSession) RenounceEventMinter(eventId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.RenounceEventMinter(&_POAPContract.TransactOpts, eventId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_POAPContract *POAPContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_POAPContract *POAPContractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.SafeTransferFrom(&_POAPContract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_POAPContract *POAPContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.SafeTransferFrom(&_POAPContract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_POAPContract *POAPContractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_POAPContract *POAPContractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _POAPContract.Contract.SafeTransferFrom0(&_POAPContract.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_POAPContract *POAPContractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _POAPContract.Contract.SafeTransferFrom0(&_POAPContract.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_POAPContract *POAPContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_POAPContract *POAPContractSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _POAPContract.Contract.SetApprovalForAll(&_POAPContract.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_POAPContract *POAPContractTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _POAPContract.Contract.SetApprovalForAll(&_POAPContract.TransactOpts, to, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_POAPContract *POAPContractTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_POAPContract *POAPContractSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _POAPContract.Contract.SetBaseURI(&_POAPContract.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_POAPContract *POAPContractTransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _POAPContract.Contract.SetBaseURI(&_POAPContract.TransactOpts, baseURI)
}

// SetFreezeDuration is a paid mutator transaction binding the contract method 0x6ca2aa95.
//
// Solidity: function setFreezeDuration(uint256 time) returns()
func (_POAPContract *POAPContractTransactor) SetFreezeDuration(opts *bind.TransactOpts, time *big.Int) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "setFreezeDuration", time)
}

// SetFreezeDuration is a paid mutator transaction binding the contract method 0x6ca2aa95.
//
// Solidity: function setFreezeDuration(uint256 time) returns()
func (_POAPContract *POAPContractSession) SetFreezeDuration(time *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.SetFreezeDuration(&_POAPContract.TransactOpts, time)
}

// SetFreezeDuration is a paid mutator transaction binding the contract method 0x6ca2aa95.
//
// Solidity: function setFreezeDuration(uint256 time) returns()
func (_POAPContract *POAPContractTransactorSession) SetFreezeDuration(time *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.SetFreezeDuration(&_POAPContract.TransactOpts, time)
}

// SetLastId is a paid mutator transaction binding the contract method 0x1a27e85f.
//
// Solidity: function setLastId(uint256 newLastId) returns()
func (_POAPContract *POAPContractTransactor) SetLastId(opts *bind.TransactOpts, newLastId *big.Int) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "setLastId", newLastId)
}

// SetLastId is a paid mutator transaction binding the contract method 0x1a27e85f.
//
// Solidity: function setLastId(uint256 newLastId) returns()
func (_POAPContract *POAPContractSession) SetLastId(newLastId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.SetLastId(&_POAPContract.TransactOpts, newLastId)
}

// SetLastId is a paid mutator transaction binding the contract method 0x1a27e85f.
//
// Solidity: function setLastId(uint256 newLastId) returns()
func (_POAPContract *POAPContractTransactorSession) SetLastId(newLastId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.SetLastId(&_POAPContract.TransactOpts, newLastId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_POAPContract *POAPContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_POAPContract *POAPContractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.TransferFrom(&_POAPContract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_POAPContract *POAPContractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.TransferFrom(&_POAPContract.TransactOpts, from, to, tokenId)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 tokenId) returns()
func (_POAPContract *POAPContractTransactor) Unfreeze(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "unfreeze", tokenId)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 tokenId) returns()
func (_POAPContract *POAPContractSession) Unfreeze(tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.Unfreeze(&_POAPContract.TransactOpts, tokenId)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 tokenId) returns()
func (_POAPContract *POAPContractTransactorSession) Unfreeze(tokenId *big.Int) (*types.Transaction, error) {
	return _POAPContract.Contract.Unfreeze(&_POAPContract.TransactOpts, tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_POAPContract *POAPContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _POAPContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_POAPContract *POAPContractSession) Unpause() (*types.Transaction, error) {
	return _POAPContract.Contract.Unpause(&_POAPContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_POAPContract *POAPContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _POAPContract.Contract.Unpause(&_POAPContract.TransactOpts)
}

// POAPContractAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the POAPContract contract.
type POAPContractAdminAddedIterator struct {
	Event *POAPContractAdminAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POAPContractAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POAPContractAdminAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POAPContractAdminAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POAPContractAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POAPContractAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POAPContractAdminAdded represents a AdminAdded event raised by the POAPContract contract.
type POAPContractAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_POAPContract *POAPContractFilterer) FilterAdminAdded(opts *bind.FilterOpts, account []common.Address) (*POAPContractAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _POAPContract.contract.FilterLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &POAPContractAdminAddedIterator{contract: _POAPContract.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_POAPContract *POAPContractFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *POAPContractAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _POAPContract.contract.WatchLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POAPContractAdminAdded)
				if err := _POAPContract.contract.UnpackLog(event, "AdminAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminAdded is a log parse operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_POAPContract *POAPContractFilterer) ParseAdminAdded(log types.Log) (*POAPContractAdminAdded, error) {
	event := new(POAPContractAdminAdded)
	if err := _POAPContract.contract.UnpackLog(event, "AdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// POAPContractAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the POAPContract contract.
type POAPContractAdminRemovedIterator struct {
	Event *POAPContractAdminRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POAPContractAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POAPContractAdminRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POAPContractAdminRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POAPContractAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POAPContractAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POAPContractAdminRemoved represents a AdminRemoved event raised by the POAPContract contract.
type POAPContractAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_POAPContract *POAPContractFilterer) FilterAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*POAPContractAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _POAPContract.contract.FilterLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &POAPContractAdminRemovedIterator{contract: _POAPContract.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_POAPContract *POAPContractFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *POAPContractAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _POAPContract.contract.WatchLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POAPContractAdminRemoved)
				if err := _POAPContract.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_POAPContract *POAPContractFilterer) ParseAdminRemoved(log types.Log) (*POAPContractAdminRemoved, error) {
	event := new(POAPContractAdminRemoved)
	if err := _POAPContract.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// POAPContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the POAPContract contract.
type POAPContractApprovalIterator struct {
	Event *POAPContractApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POAPContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POAPContractApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POAPContractApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POAPContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POAPContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POAPContractApproval represents a Approval event raised by the POAPContract contract.
type POAPContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_POAPContract *POAPContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*POAPContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _POAPContract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &POAPContractApprovalIterator{contract: _POAPContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_POAPContract *POAPContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *POAPContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _POAPContract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POAPContractApproval)
				if err := _POAPContract.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_POAPContract *POAPContractFilterer) ParseApproval(log types.Log) (*POAPContractApproval, error) {
	event := new(POAPContractApproval)
	if err := _POAPContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// POAPContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the POAPContract contract.
type POAPContractApprovalForAllIterator struct {
	Event *POAPContractApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POAPContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POAPContractApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POAPContractApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POAPContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POAPContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POAPContractApprovalForAll represents a ApprovalForAll event raised by the POAPContract contract.
type POAPContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_POAPContract *POAPContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*POAPContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _POAPContract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &POAPContractApprovalForAllIterator{contract: _POAPContract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_POAPContract *POAPContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *POAPContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _POAPContract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POAPContractApprovalForAll)
				if err := _POAPContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_POAPContract *POAPContractFilterer) ParseApprovalForAll(log types.Log) (*POAPContractApprovalForAll, error) {
	event := new(POAPContractApprovalForAll)
	if err := _POAPContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// POAPContractEventMinterAddedIterator is returned from FilterEventMinterAdded and is used to iterate over the raw logs and unpacked data for EventMinterAdded events raised by the POAPContract contract.
type POAPContractEventMinterAddedIterator struct {
	Event *POAPContractEventMinterAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POAPContractEventMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POAPContractEventMinterAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POAPContractEventMinterAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POAPContractEventMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POAPContractEventMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POAPContractEventMinterAdded represents a EventMinterAdded event raised by the POAPContract contract.
type POAPContractEventMinterAdded struct {
	EventId *big.Int
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEventMinterAdded is a free log retrieval operation binding the contract event 0xe1bd660d9f7c60e6fb12dd6479fdde12d21fc96385dc7b9b022c0b2f319e7391.
//
// Solidity: event EventMinterAdded(uint256 indexed eventId, address indexed account)
func (_POAPContract *POAPContractFilterer) FilterEventMinterAdded(opts *bind.FilterOpts, eventId []*big.Int, account []common.Address) (*POAPContractEventMinterAddedIterator, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _POAPContract.contract.FilterLogs(opts, "EventMinterAdded", eventIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &POAPContractEventMinterAddedIterator{contract: _POAPContract.contract, event: "EventMinterAdded", logs: logs, sub: sub}, nil
}

// WatchEventMinterAdded is a free log subscription operation binding the contract event 0xe1bd660d9f7c60e6fb12dd6479fdde12d21fc96385dc7b9b022c0b2f319e7391.
//
// Solidity: event EventMinterAdded(uint256 indexed eventId, address indexed account)
func (_POAPContract *POAPContractFilterer) WatchEventMinterAdded(opts *bind.WatchOpts, sink chan<- *POAPContractEventMinterAdded, eventId []*big.Int, account []common.Address) (event.Subscription, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _POAPContract.contract.WatchLogs(opts, "EventMinterAdded", eventIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POAPContractEventMinterAdded)
				if err := _POAPContract.contract.UnpackLog(event, "EventMinterAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventMinterAdded is a log parse operation binding the contract event 0xe1bd660d9f7c60e6fb12dd6479fdde12d21fc96385dc7b9b022c0b2f319e7391.
//
// Solidity: event EventMinterAdded(uint256 indexed eventId, address indexed account)
func (_POAPContract *POAPContractFilterer) ParseEventMinterAdded(log types.Log) (*POAPContractEventMinterAdded, error) {
	event := new(POAPContractEventMinterAdded)
	if err := _POAPContract.contract.UnpackLog(event, "EventMinterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// POAPContractEventMinterRemovedIterator is returned from FilterEventMinterRemoved and is used to iterate over the raw logs and unpacked data for EventMinterRemoved events raised by the POAPContract contract.
type POAPContractEventMinterRemovedIterator struct {
	Event *POAPContractEventMinterRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POAPContractEventMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POAPContractEventMinterRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POAPContractEventMinterRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POAPContractEventMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POAPContractEventMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POAPContractEventMinterRemoved represents a EventMinterRemoved event raised by the POAPContract contract.
type POAPContractEventMinterRemoved struct {
	EventId *big.Int
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEventMinterRemoved is a free log retrieval operation binding the contract event 0xb6882c4d609d560f6d57e78e73dd96027f0d9852739b0b922537a6dd3c8e944c.
//
// Solidity: event EventMinterRemoved(uint256 indexed eventId, address indexed account)
func (_POAPContract *POAPContractFilterer) FilterEventMinterRemoved(opts *bind.FilterOpts, eventId []*big.Int, account []common.Address) (*POAPContractEventMinterRemovedIterator, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _POAPContract.contract.FilterLogs(opts, "EventMinterRemoved", eventIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &POAPContractEventMinterRemovedIterator{contract: _POAPContract.contract, event: "EventMinterRemoved", logs: logs, sub: sub}, nil
}

// WatchEventMinterRemoved is a free log subscription operation binding the contract event 0xb6882c4d609d560f6d57e78e73dd96027f0d9852739b0b922537a6dd3c8e944c.
//
// Solidity: event EventMinterRemoved(uint256 indexed eventId, address indexed account)
func (_POAPContract *POAPContractFilterer) WatchEventMinterRemoved(opts *bind.WatchOpts, sink chan<- *POAPContractEventMinterRemoved, eventId []*big.Int, account []common.Address) (event.Subscription, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _POAPContract.contract.WatchLogs(opts, "EventMinterRemoved", eventIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POAPContractEventMinterRemoved)
				if err := _POAPContract.contract.UnpackLog(event, "EventMinterRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventMinterRemoved is a log parse operation binding the contract event 0xb6882c4d609d560f6d57e78e73dd96027f0d9852739b0b922537a6dd3c8e944c.
//
// Solidity: event EventMinterRemoved(uint256 indexed eventId, address indexed account)
func (_POAPContract *POAPContractFilterer) ParseEventMinterRemoved(log types.Log) (*POAPContractEventMinterRemoved, error) {
	event := new(POAPContractEventMinterRemoved)
	if err := _POAPContract.contract.UnpackLog(event, "EventMinterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// POAPContractEventTokenIterator is returned from FilterEventToken and is used to iterate over the raw logs and unpacked data for EventToken events raised by the POAPContract contract.
type POAPContractEventTokenIterator struct {
	Event *POAPContractEventToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POAPContractEventTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POAPContractEventToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POAPContractEventToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POAPContractEventTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POAPContractEventTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POAPContractEventToken represents a EventToken event raised by the POAPContract contract.
type POAPContractEventToken struct {
	EventId *big.Int
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEventToken is a free log retrieval operation binding the contract event 0x4b3711cd7ece062b0828c1b6e08d814a72d4c003383a016c833cbb1b45956e34.
//
// Solidity: event EventToken(uint256 indexed eventId, uint256 tokenId)
func (_POAPContract *POAPContractFilterer) FilterEventToken(opts *bind.FilterOpts, eventId []*big.Int) (*POAPContractEventTokenIterator, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}

	logs, sub, err := _POAPContract.contract.FilterLogs(opts, "EventToken", eventIdRule)
	if err != nil {
		return nil, err
	}
	return &POAPContractEventTokenIterator{contract: _POAPContract.contract, event: "EventToken", logs: logs, sub: sub}, nil
}

// WatchEventToken is a free log subscription operation binding the contract event 0x4b3711cd7ece062b0828c1b6e08d814a72d4c003383a016c833cbb1b45956e34.
//
// Solidity: event EventToken(uint256 indexed eventId, uint256 tokenId)
func (_POAPContract *POAPContractFilterer) WatchEventToken(opts *bind.WatchOpts, sink chan<- *POAPContractEventToken, eventId []*big.Int) (event.Subscription, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}

	logs, sub, err := _POAPContract.contract.WatchLogs(opts, "EventToken", eventIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POAPContractEventToken)
				if err := _POAPContract.contract.UnpackLog(event, "EventToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventToken is a log parse operation binding the contract event 0x4b3711cd7ece062b0828c1b6e08d814a72d4c003383a016c833cbb1b45956e34.
//
// Solidity: event EventToken(uint256 indexed eventId, uint256 tokenId)
func (_POAPContract *POAPContractFilterer) ParseEventToken(log types.Log) (*POAPContractEventToken, error) {
	event := new(POAPContractEventToken)
	if err := _POAPContract.contract.UnpackLog(event, "EventToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// POAPContractFrozenIterator is returned from FilterFrozen and is used to iterate over the raw logs and unpacked data for Frozen events raised by the POAPContract contract.
type POAPContractFrozenIterator struct {
	Event *POAPContractFrozen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POAPContractFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POAPContractFrozen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POAPContractFrozen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POAPContractFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POAPContractFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POAPContractFrozen represents a Frozen event raised by the POAPContract contract.
type POAPContractFrozen struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFrozen is a free log retrieval operation binding the contract event 0x4d69b51fee53c28bd8b61fe008151577ca65160b5248f6225e74d64fd4cf7328.
//
// Solidity: event Frozen(uint256 id)
func (_POAPContract *POAPContractFilterer) FilterFrozen(opts *bind.FilterOpts) (*POAPContractFrozenIterator, error) {

	logs, sub, err := _POAPContract.contract.FilterLogs(opts, "Frozen")
	if err != nil {
		return nil, err
	}
	return &POAPContractFrozenIterator{contract: _POAPContract.contract, event: "Frozen", logs: logs, sub: sub}, nil
}

// WatchFrozen is a free log subscription operation binding the contract event 0x4d69b51fee53c28bd8b61fe008151577ca65160b5248f6225e74d64fd4cf7328.
//
// Solidity: event Frozen(uint256 id)
func (_POAPContract *POAPContractFilterer) WatchFrozen(opts *bind.WatchOpts, sink chan<- *POAPContractFrozen) (event.Subscription, error) {

	logs, sub, err := _POAPContract.contract.WatchLogs(opts, "Frozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POAPContractFrozen)
				if err := _POAPContract.contract.UnpackLog(event, "Frozen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFrozen is a log parse operation binding the contract event 0x4d69b51fee53c28bd8b61fe008151577ca65160b5248f6225e74d64fd4cf7328.
//
// Solidity: event Frozen(uint256 id)
func (_POAPContract *POAPContractFilterer) ParseFrozen(log types.Log) (*POAPContractFrozen, error) {
	event := new(POAPContractFrozen)
	if err := _POAPContract.contract.UnpackLog(event, "Frozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// POAPContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the POAPContract contract.
type POAPContractPausedIterator struct {
	Event *POAPContractPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POAPContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POAPContractPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POAPContractPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POAPContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POAPContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POAPContractPaused represents a Paused event raised by the POAPContract contract.
type POAPContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_POAPContract *POAPContractFilterer) FilterPaused(opts *bind.FilterOpts) (*POAPContractPausedIterator, error) {

	logs, sub, err := _POAPContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &POAPContractPausedIterator{contract: _POAPContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_POAPContract *POAPContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *POAPContractPaused) (event.Subscription, error) {

	logs, sub, err := _POAPContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POAPContractPaused)
				if err := _POAPContract.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_POAPContract *POAPContractFilterer) ParsePaused(log types.Log) (*POAPContractPaused, error) {
	event := new(POAPContractPaused)
	if err := _POAPContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// POAPContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the POAPContract contract.
type POAPContractTransferIterator struct {
	Event *POAPContractTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POAPContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POAPContractTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POAPContractTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POAPContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POAPContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POAPContractTransfer represents a Transfer event raised by the POAPContract contract.
type POAPContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_POAPContract *POAPContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*POAPContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _POAPContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &POAPContractTransferIterator{contract: _POAPContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_POAPContract *POAPContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *POAPContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _POAPContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POAPContractTransfer)
				if err := _POAPContract.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_POAPContract *POAPContractFilterer) ParseTransfer(log types.Log) (*POAPContractTransfer, error) {
	event := new(POAPContractTransfer)
	if err := _POAPContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// POAPContractUnfrozenIterator is returned from FilterUnfrozen and is used to iterate over the raw logs and unpacked data for Unfrozen events raised by the POAPContract contract.
type POAPContractUnfrozenIterator struct {
	Event *POAPContractUnfrozen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POAPContractUnfrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POAPContractUnfrozen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POAPContractUnfrozen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POAPContractUnfrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POAPContractUnfrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POAPContractUnfrozen represents a Unfrozen event raised by the POAPContract contract.
type POAPContractUnfrozen struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnfrozen is a free log retrieval operation binding the contract event 0x083eea12772ab70fba01c8212d02f3bc5dc29b8540dcdc84298e5dfa22731b92.
//
// Solidity: event Unfrozen(uint256 id)
func (_POAPContract *POAPContractFilterer) FilterUnfrozen(opts *bind.FilterOpts) (*POAPContractUnfrozenIterator, error) {

	logs, sub, err := _POAPContract.contract.FilterLogs(opts, "Unfrozen")
	if err != nil {
		return nil, err
	}
	return &POAPContractUnfrozenIterator{contract: _POAPContract.contract, event: "Unfrozen", logs: logs, sub: sub}, nil
}

// WatchUnfrozen is a free log subscription operation binding the contract event 0x083eea12772ab70fba01c8212d02f3bc5dc29b8540dcdc84298e5dfa22731b92.
//
// Solidity: event Unfrozen(uint256 id)
func (_POAPContract *POAPContractFilterer) WatchUnfrozen(opts *bind.WatchOpts, sink chan<- *POAPContractUnfrozen) (event.Subscription, error) {

	logs, sub, err := _POAPContract.contract.WatchLogs(opts, "Unfrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POAPContractUnfrozen)
				if err := _POAPContract.contract.UnpackLog(event, "Unfrozen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnfrozen is a log parse operation binding the contract event 0x083eea12772ab70fba01c8212d02f3bc5dc29b8540dcdc84298e5dfa22731b92.
//
// Solidity: event Unfrozen(uint256 id)
func (_POAPContract *POAPContractFilterer) ParseUnfrozen(log types.Log) (*POAPContractUnfrozen, error) {
	event := new(POAPContractUnfrozen)
	if err := _POAPContract.contract.UnpackLog(event, "Unfrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// POAPContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the POAPContract contract.
type POAPContractUnpausedIterator struct {
	Event *POAPContractUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POAPContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POAPContractUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(POAPContractUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *POAPContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POAPContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POAPContractUnpaused represents a Unpaused event raised by the POAPContract contract.
type POAPContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_POAPContract *POAPContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*POAPContractUnpausedIterator, error) {

	logs, sub, err := _POAPContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &POAPContractUnpausedIterator{contract: _POAPContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_POAPContract *POAPContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *POAPContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _POAPContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POAPContractUnpaused)
				if err := _POAPContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_POAPContract *POAPContractFilterer) ParseUnpaused(log types.Log) (*POAPContractUnpaused, error) {
	event := new(POAPContractUnpaused)
	if err := _POAPContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
