// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nation3Tokencontracts

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
)

// Nation3TokencontractsMetaData contains all meta data concerning the Nation3Tokencontracts contract.
var Nation3TokencontractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerIsNotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetIsZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousController\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"ControlTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeControl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"transferControl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Nation3TokencontractsABI is the input ABI used to generate the binding from.
// Deprecated: Use Nation3TokencontractsMetaData.ABI instead.
var Nation3TokencontractsABI = Nation3TokencontractsMetaData.ABI

// Nation3Tokencontracts is an auto generated Go binding around an Ethereum contract.
type Nation3Tokencontracts struct {
	Nation3TokencontractsCaller     // Read-only binding to the contract
	Nation3TokencontractsTransactor // Write-only binding to the contract
	Nation3TokencontractsFilterer   // Log filterer for contract events
}

// Nation3TokencontractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type Nation3TokencontractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nation3TokencontractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Nation3TokencontractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nation3TokencontractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Nation3TokencontractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nation3TokencontractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Nation3TokencontractsSession struct {
	Contract     *Nation3Tokencontracts // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Nation3TokencontractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Nation3TokencontractsCallerSession struct {
	Contract *Nation3TokencontractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Nation3TokencontractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Nation3TokencontractsTransactorSession struct {
	Contract     *Nation3TokencontractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Nation3TokencontractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type Nation3TokencontractsRaw struct {
	Contract *Nation3Tokencontracts // Generic contract binding to access the raw methods on
}

// Nation3TokencontractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Nation3TokencontractsCallerRaw struct {
	Contract *Nation3TokencontractsCaller // Generic read-only contract binding to access the raw methods on
}

// Nation3TokencontractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Nation3TokencontractsTransactorRaw struct {
	Contract *Nation3TokencontractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNation3Tokencontracts creates a new instance of Nation3Tokencontracts, bound to a specific deployed contract.
func NewNation3Tokencontracts(address common.Address, backend bind.ContractBackend) (*Nation3Tokencontracts, error) {
	contract, err := bindNation3Tokencontracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nation3Tokencontracts{Nation3TokencontractsCaller: Nation3TokencontractsCaller{contract: contract}, Nation3TokencontractsTransactor: Nation3TokencontractsTransactor{contract: contract}, Nation3TokencontractsFilterer: Nation3TokencontractsFilterer{contract: contract}}, nil
}

// NewNation3TokencontractsCaller creates a new read-only instance of Nation3Tokencontracts, bound to a specific deployed contract.
func NewNation3TokencontractsCaller(address common.Address, caller bind.ContractCaller) (*Nation3TokencontractsCaller, error) {
	contract, err := bindNation3Tokencontracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Nation3TokencontractsCaller{contract: contract}, nil
}

// NewNation3TokencontractsTransactor creates a new write-only instance of Nation3Tokencontracts, bound to a specific deployed contract.
func NewNation3TokencontractsTransactor(address common.Address, transactor bind.ContractTransactor) (*Nation3TokencontractsTransactor, error) {
	contract, err := bindNation3Tokencontracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Nation3TokencontractsTransactor{contract: contract}, nil
}

// NewNation3TokencontractsFilterer creates a new log filterer instance of Nation3Tokencontracts, bound to a specific deployed contract.
func NewNation3TokencontractsFilterer(address common.Address, filterer bind.ContractFilterer) (*Nation3TokencontractsFilterer, error) {
	contract, err := bindNation3Tokencontracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Nation3TokencontractsFilterer{contract: contract}, nil
}

// bindNation3Tokencontracts binds a generic wrapper to an already deployed contract.
func bindNation3Tokencontracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Nation3TokencontractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nation3Tokencontracts *Nation3TokencontractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nation3Tokencontracts.Contract.Nation3TokencontractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nation3Tokencontracts *Nation3TokencontractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.Nation3TokencontractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nation3Tokencontracts *Nation3TokencontractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.Nation3TokencontractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nation3Tokencontracts *Nation3TokencontractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nation3Tokencontracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nation3Tokencontracts *Nation3TokencontractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nation3Tokencontracts *Nation3TokencontractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Nation3Tokencontracts *Nation3TokencontractsCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nation3Tokencontracts.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Nation3Tokencontracts *Nation3TokencontractsSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Nation3Tokencontracts.Contract.DOMAINSEPARATOR(&_Nation3Tokencontracts.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Nation3Tokencontracts *Nation3TokencontractsCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Nation3Tokencontracts.Contract.DOMAINSEPARATOR(&_Nation3Tokencontracts.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Nation3Tokencontracts *Nation3TokencontractsCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nation3Tokencontracts.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Nation3Tokencontracts *Nation3TokencontractsSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Nation3Tokencontracts.Contract.PERMITTYPEHASH(&_Nation3Tokencontracts.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Nation3Tokencontracts *Nation3TokencontractsCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Nation3Tokencontracts.Contract.PERMITTYPEHASH(&_Nation3Tokencontracts.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Nation3Tokencontracts *Nation3TokencontractsCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nation3Tokencontracts.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Nation3Tokencontracts *Nation3TokencontractsSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Nation3Tokencontracts.Contract.Allowance(&_Nation3Tokencontracts.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Nation3Tokencontracts *Nation3TokencontractsCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Nation3Tokencontracts.Contract.Allowance(&_Nation3Tokencontracts.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Nation3Tokencontracts *Nation3TokencontractsCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nation3Tokencontracts.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Nation3Tokencontracts *Nation3TokencontractsSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Nation3Tokencontracts.Contract.BalanceOf(&_Nation3Tokencontracts.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Nation3Tokencontracts *Nation3TokencontractsCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Nation3Tokencontracts.Contract.BalanceOf(&_Nation3Tokencontracts.CallOpts, arg0)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Nation3Tokencontracts *Nation3TokencontractsCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3Tokencontracts.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Nation3Tokencontracts *Nation3TokencontractsSession) Controller() (common.Address, error) {
	return _Nation3Tokencontracts.Contract.Controller(&_Nation3Tokencontracts.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Nation3Tokencontracts *Nation3TokencontractsCallerSession) Controller() (common.Address, error) {
	return _Nation3Tokencontracts.Contract.Controller(&_Nation3Tokencontracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Nation3Tokencontracts *Nation3TokencontractsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Nation3Tokencontracts.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Nation3Tokencontracts *Nation3TokencontractsSession) Decimals() (uint8, error) {
	return _Nation3Tokencontracts.Contract.Decimals(&_Nation3Tokencontracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Nation3Tokencontracts *Nation3TokencontractsCallerSession) Decimals() (uint8, error) {
	return _Nation3Tokencontracts.Contract.Decimals(&_Nation3Tokencontracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nation3Tokencontracts *Nation3TokencontractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nation3Tokencontracts.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nation3Tokencontracts *Nation3TokencontractsSession) Name() (string, error) {
	return _Nation3Tokencontracts.Contract.Name(&_Nation3Tokencontracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nation3Tokencontracts *Nation3TokencontractsCallerSession) Name() (string, error) {
	return _Nation3Tokencontracts.Contract.Name(&_Nation3Tokencontracts.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Nation3Tokencontracts *Nation3TokencontractsCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nation3Tokencontracts.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Nation3Tokencontracts *Nation3TokencontractsSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Nation3Tokencontracts.Contract.Nonces(&_Nation3Tokencontracts.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Nation3Tokencontracts *Nation3TokencontractsCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Nation3Tokencontracts.Contract.Nonces(&_Nation3Tokencontracts.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nation3Tokencontracts *Nation3TokencontractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3Tokencontracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nation3Tokencontracts *Nation3TokencontractsSession) Owner() (common.Address, error) {
	return _Nation3Tokencontracts.Contract.Owner(&_Nation3Tokencontracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nation3Tokencontracts *Nation3TokencontractsCallerSession) Owner() (common.Address, error) {
	return _Nation3Tokencontracts.Contract.Owner(&_Nation3Tokencontracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nation3Tokencontracts *Nation3TokencontractsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nation3Tokencontracts.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nation3Tokencontracts *Nation3TokencontractsSession) Symbol() (string, error) {
	return _Nation3Tokencontracts.Contract.Symbol(&_Nation3Tokencontracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nation3Tokencontracts *Nation3TokencontractsCallerSession) Symbol() (string, error) {
	return _Nation3Tokencontracts.Contract.Symbol(&_Nation3Tokencontracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nation3Tokencontracts *Nation3TokencontractsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nation3Tokencontracts.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nation3Tokencontracts *Nation3TokencontractsSession) TotalSupply() (*big.Int, error) {
	return _Nation3Tokencontracts.Contract.TotalSupply(&_Nation3Tokencontracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nation3Tokencontracts *Nation3TokencontractsCallerSession) TotalSupply() (*big.Int, error) {
	return _Nation3Tokencontracts.Contract.TotalSupply(&_Nation3Tokencontracts.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Nation3Tokencontracts *Nation3TokencontractsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nation3Tokencontracts.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Nation3Tokencontracts *Nation3TokencontractsSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.Approve(&_Nation3Tokencontracts.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Nation3Tokencontracts *Nation3TokencontractsTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.Approve(&_Nation3Tokencontracts.TransactOpts, spender, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Nation3Tokencontracts *Nation3TokencontractsTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nation3Tokencontracts.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Nation3Tokencontracts *Nation3TokencontractsSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.Mint(&_Nation3Tokencontracts.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Nation3Tokencontracts *Nation3TokencontractsTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.Mint(&_Nation3Tokencontracts.TransactOpts, to, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Nation3Tokencontracts *Nation3TokencontractsTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Nation3Tokencontracts.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Nation3Tokencontracts *Nation3TokencontractsSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.Permit(&_Nation3Tokencontracts.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Nation3Tokencontracts *Nation3TokencontractsTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.Permit(&_Nation3Tokencontracts.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RemoveControl is a paid mutator transaction binding the contract method 0x7bee684b.
//
// Solidity: function removeControl() returns()
func (_Nation3Tokencontracts *Nation3TokencontractsTransactor) RemoveControl(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3Tokencontracts.contract.Transact(opts, "removeControl")
}

// RemoveControl is a paid mutator transaction binding the contract method 0x7bee684b.
//
// Solidity: function removeControl() returns()
func (_Nation3Tokencontracts *Nation3TokencontractsSession) RemoveControl() (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.RemoveControl(&_Nation3Tokencontracts.TransactOpts)
}

// RemoveControl is a paid mutator transaction binding the contract method 0x7bee684b.
//
// Solidity: function removeControl() returns()
func (_Nation3Tokencontracts *Nation3TokencontractsTransactorSession) RemoveControl() (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.RemoveControl(&_Nation3Tokencontracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nation3Tokencontracts *Nation3TokencontractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3Tokencontracts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nation3Tokencontracts *Nation3TokencontractsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.RenounceOwnership(&_Nation3Tokencontracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nation3Tokencontracts *Nation3TokencontractsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.RenounceOwnership(&_Nation3Tokencontracts.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Nation3Tokencontracts *Nation3TokencontractsTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nation3Tokencontracts.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Nation3Tokencontracts *Nation3TokencontractsSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.Transfer(&_Nation3Tokencontracts.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Nation3Tokencontracts *Nation3TokencontractsTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.Transfer(&_Nation3Tokencontracts.TransactOpts, to, amount)
}

// TransferControl is a paid mutator transaction binding the contract method 0x6d16fa41.
//
// Solidity: function transferControl(address newController) returns()
func (_Nation3Tokencontracts *Nation3TokencontractsTransactor) TransferControl(opts *bind.TransactOpts, newController common.Address) (*types.Transaction, error) {
	return _Nation3Tokencontracts.contract.Transact(opts, "transferControl", newController)
}

// TransferControl is a paid mutator transaction binding the contract method 0x6d16fa41.
//
// Solidity: function transferControl(address newController) returns()
func (_Nation3Tokencontracts *Nation3TokencontractsSession) TransferControl(newController common.Address) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.TransferControl(&_Nation3Tokencontracts.TransactOpts, newController)
}

// TransferControl is a paid mutator transaction binding the contract method 0x6d16fa41.
//
// Solidity: function transferControl(address newController) returns()
func (_Nation3Tokencontracts *Nation3TokencontractsTransactorSession) TransferControl(newController common.Address) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.TransferControl(&_Nation3Tokencontracts.TransactOpts, newController)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Nation3Tokencontracts *Nation3TokencontractsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nation3Tokencontracts.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Nation3Tokencontracts *Nation3TokencontractsSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.TransferFrom(&_Nation3Tokencontracts.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Nation3Tokencontracts *Nation3TokencontractsTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.TransferFrom(&_Nation3Tokencontracts.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nation3Tokencontracts *Nation3TokencontractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Nation3Tokencontracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nation3Tokencontracts *Nation3TokencontractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.TransferOwnership(&_Nation3Tokencontracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nation3Tokencontracts *Nation3TokencontractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nation3Tokencontracts.Contract.TransferOwnership(&_Nation3Tokencontracts.TransactOpts, newOwner)
}

// Nation3TokencontractsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Nation3Tokencontracts contract.
type Nation3TokencontractsApprovalIterator struct {
	Event *Nation3TokencontractsApproval // Event containing the contract specifics and raw log

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
func (it *Nation3TokencontractsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3TokencontractsApproval)
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
		it.Event = new(Nation3TokencontractsApproval)
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
func (it *Nation3TokencontractsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3TokencontractsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3TokencontractsApproval represents a Approval event raised by the Nation3Tokencontracts contract.
type Nation3TokencontractsApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Nation3Tokencontracts *Nation3TokencontractsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Nation3TokencontractsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Nation3Tokencontracts.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Nation3TokencontractsApprovalIterator{contract: _Nation3Tokencontracts.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Nation3Tokencontracts *Nation3TokencontractsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Nation3TokencontractsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Nation3Tokencontracts.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3TokencontractsApproval)
				if err := _Nation3Tokencontracts.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Nation3Tokencontracts *Nation3TokencontractsFilterer) ParseApproval(log types.Log) (*Nation3TokencontractsApproval, error) {
	event := new(Nation3TokencontractsApproval)
	if err := _Nation3Tokencontracts.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3TokencontractsControlTransferredIterator is returned from FilterControlTransferred and is used to iterate over the raw logs and unpacked data for ControlTransferred events raised by the Nation3Tokencontracts contract.
type Nation3TokencontractsControlTransferredIterator struct {
	Event *Nation3TokencontractsControlTransferred // Event containing the contract specifics and raw log

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
func (it *Nation3TokencontractsControlTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3TokencontractsControlTransferred)
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
		it.Event = new(Nation3TokencontractsControlTransferred)
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
func (it *Nation3TokencontractsControlTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3TokencontractsControlTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3TokencontractsControlTransferred represents a ControlTransferred event raised by the Nation3Tokencontracts contract.
type Nation3TokencontractsControlTransferred struct {
	PreviousController common.Address
	NewController      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterControlTransferred is a free log retrieval operation binding the contract event 0xa06677f7b64342b4bcbde423684dbdb5356acfe41ad0285b6ecbe6dc4bf427f2.
//
// Solidity: event ControlTransferred(address indexed previousController, address indexed newController)
func (_Nation3Tokencontracts *Nation3TokencontractsFilterer) FilterControlTransferred(opts *bind.FilterOpts, previousController []common.Address, newController []common.Address) (*Nation3TokencontractsControlTransferredIterator, error) {

	var previousControllerRule []interface{}
	for _, previousControllerItem := range previousController {
		previousControllerRule = append(previousControllerRule, previousControllerItem)
	}
	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _Nation3Tokencontracts.contract.FilterLogs(opts, "ControlTransferred", previousControllerRule, newControllerRule)
	if err != nil {
		return nil, err
	}
	return &Nation3TokencontractsControlTransferredIterator{contract: _Nation3Tokencontracts.contract, event: "ControlTransferred", logs: logs, sub: sub}, nil
}

// WatchControlTransferred is a free log subscription operation binding the contract event 0xa06677f7b64342b4bcbde423684dbdb5356acfe41ad0285b6ecbe6dc4bf427f2.
//
// Solidity: event ControlTransferred(address indexed previousController, address indexed newController)
func (_Nation3Tokencontracts *Nation3TokencontractsFilterer) WatchControlTransferred(opts *bind.WatchOpts, sink chan<- *Nation3TokencontractsControlTransferred, previousController []common.Address, newController []common.Address) (event.Subscription, error) {

	var previousControllerRule []interface{}
	for _, previousControllerItem := range previousController {
		previousControllerRule = append(previousControllerRule, previousControllerItem)
	}
	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _Nation3Tokencontracts.contract.WatchLogs(opts, "ControlTransferred", previousControllerRule, newControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3TokencontractsControlTransferred)
				if err := _Nation3Tokencontracts.contract.UnpackLog(event, "ControlTransferred", log); err != nil {
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

// ParseControlTransferred is a log parse operation binding the contract event 0xa06677f7b64342b4bcbde423684dbdb5356acfe41ad0285b6ecbe6dc4bf427f2.
//
// Solidity: event ControlTransferred(address indexed previousController, address indexed newController)
func (_Nation3Tokencontracts *Nation3TokencontractsFilterer) ParseControlTransferred(log types.Log) (*Nation3TokencontractsControlTransferred, error) {
	event := new(Nation3TokencontractsControlTransferred)
	if err := _Nation3Tokencontracts.contract.UnpackLog(event, "ControlTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3TokencontractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Nation3Tokencontracts contract.
type Nation3TokencontractsOwnershipTransferredIterator struct {
	Event *Nation3TokencontractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Nation3TokencontractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3TokencontractsOwnershipTransferred)
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
		it.Event = new(Nation3TokencontractsOwnershipTransferred)
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
func (it *Nation3TokencontractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3TokencontractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3TokencontractsOwnershipTransferred represents a OwnershipTransferred event raised by the Nation3Tokencontracts contract.
type Nation3TokencontractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nation3Tokencontracts *Nation3TokencontractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Nation3TokencontractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nation3Tokencontracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Nation3TokencontractsOwnershipTransferredIterator{contract: _Nation3Tokencontracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nation3Tokencontracts *Nation3TokencontractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Nation3TokencontractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nation3Tokencontracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3TokencontractsOwnershipTransferred)
				if err := _Nation3Tokencontracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nation3Tokencontracts *Nation3TokencontractsFilterer) ParseOwnershipTransferred(log types.Log) (*Nation3TokencontractsOwnershipTransferred, error) {
	event := new(Nation3TokencontractsOwnershipTransferred)
	if err := _Nation3Tokencontracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3TokencontractsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Nation3Tokencontracts contract.
type Nation3TokencontractsTransferIterator struct {
	Event *Nation3TokencontractsTransfer // Event containing the contract specifics and raw log

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
func (it *Nation3TokencontractsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3TokencontractsTransfer)
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
		it.Event = new(Nation3TokencontractsTransfer)
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
func (it *Nation3TokencontractsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3TokencontractsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3TokencontractsTransfer represents a Transfer event raised by the Nation3Tokencontracts contract.
type Nation3TokencontractsTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Nation3Tokencontracts *Nation3TokencontractsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Nation3TokencontractsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Nation3Tokencontracts.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Nation3TokencontractsTransferIterator{contract: _Nation3Tokencontracts.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Nation3Tokencontracts *Nation3TokencontractsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Nation3TokencontractsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Nation3Tokencontracts.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3TokencontractsTransfer)
				if err := _Nation3Tokencontracts.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Nation3Tokencontracts *Nation3TokencontractsFilterer) ParseTransfer(log types.Log) (*Nation3TokencontractsTransfer, error) {
	event := new(Nation3TokencontractsTransfer)
	if err := _Nation3Tokencontracts.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
