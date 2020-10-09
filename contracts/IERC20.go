// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC20BaseContractABI is the input ABI used to generate the binding from.
const ERC20BaseContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC20BaseContract is an auto generated Go binding around an Ethereum contract.
type ERC20BaseContract struct {
	ERC20BaseContractCaller     // Read-only binding to the contract
	ERC20BaseContractTransactor // Write-only binding to the contract
	ERC20BaseContractFilterer   // Log filterer for contract events
}

// ERC20BaseContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BaseContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BaseContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BaseContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BaseContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BaseContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BaseContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BaseContractSession struct {
	Contract     *ERC20BaseContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC20BaseContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BaseContractCallerSession struct {
	Contract *ERC20BaseContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ERC20BaseContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BaseContractTransactorSession struct {
	Contract     *ERC20BaseContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ERC20BaseContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BaseContractRaw struct {
	Contract *ERC20BaseContract // Generic contract binding to access the raw methods on
}

// ERC20BaseContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BaseContractCallerRaw struct {
	Contract *ERC20BaseContractCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BaseContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BaseContractTransactorRaw struct {
	Contract *ERC20BaseContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20BaseContract creates a new instance of ERC20BaseContract, bound to a specific deployed contract.
func NewERC20BaseContract(address common.Address, backend bind.ContractBackend) (*ERC20BaseContract, error) {
	contract, err := bindERC20BaseContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20BaseContract{ERC20BaseContractCaller: ERC20BaseContractCaller{contract: contract}, ERC20BaseContractTransactor: ERC20BaseContractTransactor{contract: contract}, ERC20BaseContractFilterer: ERC20BaseContractFilterer{contract: contract}}, nil
}

// NewERC20BaseContractCaller creates a new read-only instance of ERC20BaseContract, bound to a specific deployed contract.
func NewERC20BaseContractCaller(address common.Address, caller bind.ContractCaller) (*ERC20BaseContractCaller, error) {
	contract, err := bindERC20BaseContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BaseContractCaller{contract: contract}, nil
}

// NewERC20BaseContractTransactor creates a new write-only instance of ERC20BaseContract, bound to a specific deployed contract.
func NewERC20BaseContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BaseContractTransactor, error) {
	contract, err := bindERC20BaseContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BaseContractTransactor{contract: contract}, nil
}

// NewERC20BaseContractFilterer creates a new log filterer instance of ERC20BaseContract, bound to a specific deployed contract.
func NewERC20BaseContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BaseContractFilterer, error) {
	contract, err := bindERC20BaseContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BaseContractFilterer{contract: contract}, nil
}

// bindERC20BaseContract binds a generic wrapper to an already deployed contract.
func bindERC20BaseContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BaseContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20BaseContract *ERC20BaseContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20BaseContract.Contract.ERC20BaseContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20BaseContract *ERC20BaseContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20BaseContract.Contract.ERC20BaseContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20BaseContract *ERC20BaseContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20BaseContract.Contract.ERC20BaseContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20BaseContract *ERC20BaseContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20BaseContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20BaseContract *ERC20BaseContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20BaseContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20BaseContract *ERC20BaseContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20BaseContract.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20BaseContract *ERC20BaseContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC20BaseContract.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20BaseContract *ERC20BaseContractSession) Decimals() (uint8, error) {
	return _ERC20BaseContract.Contract.Decimals(&_ERC20BaseContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20BaseContract *ERC20BaseContractCallerSession) Decimals() (uint8, error) {
	return _ERC20BaseContract.Contract.Decimals(&_ERC20BaseContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20BaseContract *ERC20BaseContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20BaseContract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20BaseContract *ERC20BaseContractSession) Name() (string, error) {
	return _ERC20BaseContract.Contract.Name(&_ERC20BaseContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20BaseContract *ERC20BaseContractCallerSession) Name() (string, error) {
	return _ERC20BaseContract.Contract.Name(&_ERC20BaseContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20BaseContract *ERC20BaseContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20BaseContract.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20BaseContract *ERC20BaseContractSession) Symbol() (string, error) {
	return _ERC20BaseContract.Contract.Symbol(&_ERC20BaseContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20BaseContract *ERC20BaseContractCallerSession) Symbol() (string, error) {
	return _ERC20BaseContract.Contract.Symbol(&_ERC20BaseContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20BaseContract *ERC20BaseContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20BaseContract.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20BaseContract *ERC20BaseContractSession) TotalSupply() (*big.Int, error) {
	return _ERC20BaseContract.Contract.TotalSupply(&_ERC20BaseContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20BaseContract *ERC20BaseContractCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20BaseContract.Contract.TotalSupply(&_ERC20BaseContract.CallOpts)
}
