// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nation3Passportcontracts

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

// Nation3PassportcontractsMetaData contains all meta data concerning the Nation3Passportcontracts contract.
var Nation3PassportcontractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerIsNotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFrom\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMinted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSafeRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetIsZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousController\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"ControlTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeControl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renderer\",\"outputs\":[{\"internalType\":\"contractRenderer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"safeMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRenderer\",\"name\":\"_renderer\",\"type\":\"address\"}],\"name\":\"setRenderer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"signerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"timestampOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"transferControl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Nation3PassportcontractsABI is the input ABI used to generate the binding from.
// Deprecated: Use Nation3PassportcontractsMetaData.ABI instead.
var Nation3PassportcontractsABI = Nation3PassportcontractsMetaData.ABI

// Nation3Passportcontracts is an auto generated Go binding around an Ethereum contract.
type Nation3Passportcontracts struct {
	Nation3PassportcontractsCaller     // Read-only binding to the contract
	Nation3PassportcontractsTransactor // Write-only binding to the contract
	Nation3PassportcontractsFilterer   // Log filterer for contract events
}

// Nation3PassportcontractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type Nation3PassportcontractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nation3PassportcontractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Nation3PassportcontractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nation3PassportcontractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Nation3PassportcontractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nation3PassportcontractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Nation3PassportcontractsSession struct {
	Contract     *Nation3Passportcontracts // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Nation3PassportcontractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Nation3PassportcontractsCallerSession struct {
	Contract *Nation3PassportcontractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Nation3PassportcontractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Nation3PassportcontractsTransactorSession struct {
	Contract     *Nation3PassportcontractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Nation3PassportcontractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type Nation3PassportcontractsRaw struct {
	Contract *Nation3Passportcontracts // Generic contract binding to access the raw methods on
}

// Nation3PassportcontractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Nation3PassportcontractsCallerRaw struct {
	Contract *Nation3PassportcontractsCaller // Generic read-only contract binding to access the raw methods on
}

// Nation3PassportcontractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Nation3PassportcontractsTransactorRaw struct {
	Contract *Nation3PassportcontractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNation3Passportcontracts creates a new instance of Nation3Passportcontracts, bound to a specific deployed contract.
func NewNation3Passportcontracts(address common.Address, backend bind.ContractBackend) (*Nation3Passportcontracts, error) {
	contract, err := bindNation3Passportcontracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nation3Passportcontracts{Nation3PassportcontractsCaller: Nation3PassportcontractsCaller{contract: contract}, Nation3PassportcontractsTransactor: Nation3PassportcontractsTransactor{contract: contract}, Nation3PassportcontractsFilterer: Nation3PassportcontractsFilterer{contract: contract}}, nil
}

// NewNation3PassportcontractsCaller creates a new read-only instance of Nation3Passportcontracts, bound to a specific deployed contract.
func NewNation3PassportcontractsCaller(address common.Address, caller bind.ContractCaller) (*Nation3PassportcontractsCaller, error) {
	contract, err := bindNation3Passportcontracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportcontractsCaller{contract: contract}, nil
}

// NewNation3PassportcontractsTransactor creates a new write-only instance of Nation3Passportcontracts, bound to a specific deployed contract.
func NewNation3PassportcontractsTransactor(address common.Address, transactor bind.ContractTransactor) (*Nation3PassportcontractsTransactor, error) {
	contract, err := bindNation3Passportcontracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportcontractsTransactor{contract: contract}, nil
}

// NewNation3PassportcontractsFilterer creates a new log filterer instance of Nation3Passportcontracts, bound to a specific deployed contract.
func NewNation3PassportcontractsFilterer(address common.Address, filterer bind.ContractFilterer) (*Nation3PassportcontractsFilterer, error) {
	contract, err := bindNation3Passportcontracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportcontractsFilterer{contract: contract}, nil
}

// bindNation3Passportcontracts binds a generic wrapper to an already deployed contract.
func bindNation3Passportcontracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Nation3PassportcontractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nation3Passportcontracts *Nation3PassportcontractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nation3Passportcontracts.Contract.Nation3PassportcontractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nation3Passportcontracts *Nation3PassportcontractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.Nation3PassportcontractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nation3Passportcontracts *Nation3PassportcontractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.Nation3PassportcontractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nation3Passportcontracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nation3Passportcontracts *Nation3PassportcontractsCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nation3Passportcontracts.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Nation3Passportcontracts.Contract.BalanceOf(&_Nation3Passportcontracts.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Nation3Passportcontracts.Contract.BalanceOf(&_Nation3Passportcontracts.CallOpts, owner)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Nation3Passportcontracts *Nation3PassportcontractsCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3Passportcontracts.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) Controller() (common.Address, error) {
	return _Nation3Passportcontracts.Contract.Controller(&_Nation3Passportcontracts.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerSession) Controller() (common.Address, error) {
	return _Nation3Passportcontracts.Contract.Controller(&_Nation3Passportcontracts.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Nation3Passportcontracts *Nation3PassportcontractsCaller) GetApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nation3Passportcontracts.contract.Call(opts, &out, "getApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Nation3Passportcontracts.Contract.GetApproved(&_Nation3Passportcontracts.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Nation3Passportcontracts.Contract.GetApproved(&_Nation3Passportcontracts.CallOpts, arg0)
}

// GetNextId is a free data retrieval call binding the contract method 0xbc968326.
//
// Solidity: function getNextId() view returns(uint256)
func (_Nation3Passportcontracts *Nation3PassportcontractsCaller) GetNextId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nation3Passportcontracts.contract.Call(opts, &out, "getNextId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextId is a free data retrieval call binding the contract method 0xbc968326.
//
// Solidity: function getNextId() view returns(uint256)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) GetNextId() (*big.Int, error) {
	return _Nation3Passportcontracts.Contract.GetNextId(&_Nation3Passportcontracts.CallOpts)
}

// GetNextId is a free data retrieval call binding the contract method 0xbc968326.
//
// Solidity: function getNextId() view returns(uint256)
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerSession) GetNextId() (*big.Int, error) {
	return _Nation3Passportcontracts.Contract.GetNextId(&_Nation3Passportcontracts.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Nation3Passportcontracts *Nation3PassportcontractsCaller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Nation3Passportcontracts.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Nation3Passportcontracts.Contract.IsApprovedForAll(&_Nation3Passportcontracts.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Nation3Passportcontracts.Contract.IsApprovedForAll(&_Nation3Passportcontracts.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nation3Passportcontracts *Nation3PassportcontractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nation3Passportcontracts.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) Name() (string, error) {
	return _Nation3Passportcontracts.Contract.Name(&_Nation3Passportcontracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerSession) Name() (string, error) {
	return _Nation3Passportcontracts.Contract.Name(&_Nation3Passportcontracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nation3Passportcontracts *Nation3PassportcontractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3Passportcontracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) Owner() (common.Address, error) {
	return _Nation3Passportcontracts.Contract.Owner(&_Nation3Passportcontracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerSession) Owner() (common.Address, error) {
	return _Nation3Passportcontracts.Contract.Owner(&_Nation3Passportcontracts.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Nation3Passportcontracts *Nation3PassportcontractsCaller) OwnerOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nation3Passportcontracts.contract.Call(opts, &out, "ownerOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Nation3Passportcontracts.Contract.OwnerOf(&_Nation3Passportcontracts.CallOpts, id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Nation3Passportcontracts.Contract.OwnerOf(&_Nation3Passportcontracts.CallOpts, id)
}

// Renderer is a free data retrieval call binding the contract method 0x8ada6b0f.
//
// Solidity: function renderer() view returns(address)
func (_Nation3Passportcontracts *Nation3PassportcontractsCaller) Renderer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3Passportcontracts.contract.Call(opts, &out, "renderer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Renderer is a free data retrieval call binding the contract method 0x8ada6b0f.
//
// Solidity: function renderer() view returns(address)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) Renderer() (common.Address, error) {
	return _Nation3Passportcontracts.Contract.Renderer(&_Nation3Passportcontracts.CallOpts)
}

// Renderer is a free data retrieval call binding the contract method 0x8ada6b0f.
//
// Solidity: function renderer() view returns(address)
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerSession) Renderer() (common.Address, error) {
	return _Nation3Passportcontracts.Contract.Renderer(&_Nation3Passportcontracts.CallOpts)
}

// SignerOf is a free data retrieval call binding the contract method 0x5161fdf5.
//
// Solidity: function signerOf(uint256 id) view returns(address)
func (_Nation3Passportcontracts *Nation3PassportcontractsCaller) SignerOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nation3Passportcontracts.contract.Call(opts, &out, "signerOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerOf is a free data retrieval call binding the contract method 0x5161fdf5.
//
// Solidity: function signerOf(uint256 id) view returns(address)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) SignerOf(id *big.Int) (common.Address, error) {
	return _Nation3Passportcontracts.Contract.SignerOf(&_Nation3Passportcontracts.CallOpts, id)
}

// SignerOf is a free data retrieval call binding the contract method 0x5161fdf5.
//
// Solidity: function signerOf(uint256 id) view returns(address)
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerSession) SignerOf(id *big.Int) (common.Address, error) {
	return _Nation3Passportcontracts.Contract.SignerOf(&_Nation3Passportcontracts.CallOpts, id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nation3Passportcontracts *Nation3PassportcontractsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Nation3Passportcontracts.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Nation3Passportcontracts.Contract.SupportsInterface(&_Nation3Passportcontracts.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Nation3Passportcontracts.Contract.SupportsInterface(&_Nation3Passportcontracts.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nation3Passportcontracts *Nation3PassportcontractsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nation3Passportcontracts.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) Symbol() (string, error) {
	return _Nation3Passportcontracts.Contract.Symbol(&_Nation3Passportcontracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerSession) Symbol() (string, error) {
	return _Nation3Passportcontracts.Contract.Symbol(&_Nation3Passportcontracts.CallOpts)
}

// TimestampOf is a free data retrieval call binding the contract method 0x2d9c77e1.
//
// Solidity: function timestampOf(uint256 id) view returns(uint256)
func (_Nation3Passportcontracts *Nation3PassportcontractsCaller) TimestampOf(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nation3Passportcontracts.contract.Call(opts, &out, "timestampOf", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimestampOf is a free data retrieval call binding the contract method 0x2d9c77e1.
//
// Solidity: function timestampOf(uint256 id) view returns(uint256)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) TimestampOf(id *big.Int) (*big.Int, error) {
	return _Nation3Passportcontracts.Contract.TimestampOf(&_Nation3Passportcontracts.CallOpts, id)
}

// TimestampOf is a free data retrieval call binding the contract method 0x2d9c77e1.
//
// Solidity: function timestampOf(uint256 id) view returns(uint256)
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerSession) TimestampOf(id *big.Int) (*big.Int, error) {
	return _Nation3Passportcontracts.Contract.TimestampOf(&_Nation3Passportcontracts.CallOpts, id)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_Nation3Passportcontracts *Nation3PassportcontractsCaller) TokenURI(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _Nation3Passportcontracts.contract.Call(opts, &out, "tokenURI", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) TokenURI(id *big.Int) (string, error) {
	return _Nation3Passportcontracts.Contract.TokenURI(&_Nation3Passportcontracts.CallOpts, id)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerSession) TokenURI(id *big.Int) (string, error) {
	return _Nation3Passportcontracts.Contract.TokenURI(&_Nation3Passportcontracts.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nation3Passportcontracts *Nation3PassportcontractsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nation3Passportcontracts.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) TotalSupply() (*big.Int, error) {
	return _Nation3Passportcontracts.Contract.TotalSupply(&_Nation3Passportcontracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nation3Passportcontracts *Nation3PassportcontractsCallerSession) TotalSupply() (*big.Int, error) {
	return _Nation3Passportcontracts.Contract.TotalSupply(&_Nation3Passportcontracts.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3Passportcontracts.contract.Transact(opts, "approve", spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.Approve(&_Nation3Passportcontracts.TransactOpts, spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.Approve(&_Nation3Passportcontracts.TransactOpts, spender, id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 id) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactor) Burn(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Nation3Passportcontracts.contract.Transact(opts, "burn", id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 id) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) Burn(id *big.Int) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.Burn(&_Nation3Passportcontracts.TransactOpts, id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 id) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorSession) Burn(id *big.Int) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.Burn(&_Nation3Passportcontracts.TransactOpts, id)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 tokenId)
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 tokenId)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.Mint(&_Nation3Passportcontracts.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 tokenId)
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.Mint(&_Nation3Passportcontracts.TransactOpts, to)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x056097ac.
//
// Solidity: function recoverTokens(address token, address to) returns(uint256 amount)
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactor) RecoverTokens(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.contract.Transact(opts, "recoverTokens", token, to)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x056097ac.
//
// Solidity: function recoverTokens(address token, address to) returns(uint256 amount)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) RecoverTokens(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.RecoverTokens(&_Nation3Passportcontracts.TransactOpts, token, to)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x056097ac.
//
// Solidity: function recoverTokens(address token, address to) returns(uint256 amount)
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorSession) RecoverTokens(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.RecoverTokens(&_Nation3Passportcontracts.TransactOpts, token, to)
}

// RemoveControl is a paid mutator transaction binding the contract method 0x7bee684b.
//
// Solidity: function removeControl() returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactor) RemoveControl(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3Passportcontracts.contract.Transact(opts, "removeControl")
}

// RemoveControl is a paid mutator transaction binding the contract method 0x7bee684b.
//
// Solidity: function removeControl() returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) RemoveControl() (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.RemoveControl(&_Nation3Passportcontracts.TransactOpts)
}

// RemoveControl is a paid mutator transaction binding the contract method 0x7bee684b.
//
// Solidity: function removeControl() returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorSession) RemoveControl() (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.RemoveControl(&_Nation3Passportcontracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3Passportcontracts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.RenounceOwnership(&_Nation3Passportcontracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.RenounceOwnership(&_Nation3Passportcontracts.TransactOpts)
}

// SafeMint is a paid mutator transaction binding the contract method 0x40d097c3.
//
// Solidity: function safeMint(address to) returns(uint256 tokenId)
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactor) SafeMint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.contract.Transact(opts, "safeMint", to)
}

// SafeMint is a paid mutator transaction binding the contract method 0x40d097c3.
//
// Solidity: function safeMint(address to) returns(uint256 tokenId)
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) SafeMint(to common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.SafeMint(&_Nation3Passportcontracts.TransactOpts, to)
}

// SafeMint is a paid mutator transaction binding the contract method 0x40d097c3.
//
// Solidity: function safeMint(address to) returns(uint256 tokenId)
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorSession) SafeMint(to common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.SafeMint(&_Nation3Passportcontracts.TransactOpts, to)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3Passportcontracts.contract.Transact(opts, "safeTransferFrom", from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.SafeTransferFrom(&_Nation3Passportcontracts.TransactOpts, from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.SafeTransferFrom(&_Nation3Passportcontracts.TransactOpts, from, to, id)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Nation3Passportcontracts.contract.Transact(opts, "safeTransferFrom0", from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.SafeTransferFrom0(&_Nation3Passportcontracts.TransactOpts, from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.SafeTransferFrom0(&_Nation3Passportcontracts.TransactOpts, from, to, id, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nation3Passportcontracts.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.SetApprovalForAll(&_Nation3Passportcontracts.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.SetApprovalForAll(&_Nation3Passportcontracts.TransactOpts, operator, approved)
}

// SetRenderer is a paid mutator transaction binding the contract method 0x56d3163d.
//
// Solidity: function setRenderer(address _renderer) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactor) SetRenderer(opts *bind.TransactOpts, _renderer common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.contract.Transact(opts, "setRenderer", _renderer)
}

// SetRenderer is a paid mutator transaction binding the contract method 0x56d3163d.
//
// Solidity: function setRenderer(address _renderer) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) SetRenderer(_renderer common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.SetRenderer(&_Nation3Passportcontracts.TransactOpts, _renderer)
}

// SetRenderer is a paid mutator transaction binding the contract method 0x56d3163d.
//
// Solidity: function setRenderer(address _renderer) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorSession) SetRenderer(_renderer common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.SetRenderer(&_Nation3Passportcontracts.TransactOpts, _renderer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x85267907.
//
// Solidity: function setSigner(uint256 id, address signer) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactor) SetSigner(opts *bind.TransactOpts, id *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.contract.Transact(opts, "setSigner", id, signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x85267907.
//
// Solidity: function setSigner(uint256 id, address signer) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) SetSigner(id *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.SetSigner(&_Nation3Passportcontracts.TransactOpts, id, signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x85267907.
//
// Solidity: function setSigner(uint256 id, address signer) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorSession) SetSigner(id *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.SetSigner(&_Nation3Passportcontracts.TransactOpts, id, signer)
}

// TransferControl is a paid mutator transaction binding the contract method 0x6d16fa41.
//
// Solidity: function transferControl(address newController) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactor) TransferControl(opts *bind.TransactOpts, newController common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.contract.Transact(opts, "transferControl", newController)
}

// TransferControl is a paid mutator transaction binding the contract method 0x6d16fa41.
//
// Solidity: function transferControl(address newController) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) TransferControl(newController common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.TransferControl(&_Nation3Passportcontracts.TransactOpts, newController)
}

// TransferControl is a paid mutator transaction binding the contract method 0x6d16fa41.
//
// Solidity: function transferControl(address newController) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorSession) TransferControl(newController common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.TransferControl(&_Nation3Passportcontracts.TransactOpts, newController)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3Passportcontracts.contract.Transact(opts, "transferFrom", from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.TransferFrom(&_Nation3Passportcontracts.TransactOpts, from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.TransferFrom(&_Nation3Passportcontracts.TransactOpts, from, to, id)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.TransferOwnership(&_Nation3Passportcontracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nation3Passportcontracts *Nation3PassportcontractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nation3Passportcontracts.Contract.TransferOwnership(&_Nation3Passportcontracts.TransactOpts, newOwner)
}

// Nation3PassportcontractsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Nation3Passportcontracts contract.
type Nation3PassportcontractsApprovalIterator struct {
	Event *Nation3PassportcontractsApproval // Event containing the contract specifics and raw log

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
func (it *Nation3PassportcontractsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3PassportcontractsApproval)
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
		it.Event = new(Nation3PassportcontractsApproval)
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
func (it *Nation3PassportcontractsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3PassportcontractsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3PassportcontractsApproval represents a Approval event raised by the Nation3Passportcontracts contract.
type Nation3PassportcontractsApproval struct {
	Owner   common.Address
	Spender common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Nation3Passportcontracts *Nation3PassportcontractsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address, id []*big.Int) (*Nation3PassportcontractsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Nation3Passportcontracts.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportcontractsApprovalIterator{contract: _Nation3Passportcontracts.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Nation3Passportcontracts *Nation3PassportcontractsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Nation3PassportcontractsApproval, owner []common.Address, spender []common.Address, id []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Nation3Passportcontracts.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3PassportcontractsApproval)
				if err := _Nation3Passportcontracts.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Nation3Passportcontracts *Nation3PassportcontractsFilterer) ParseApproval(log types.Log) (*Nation3PassportcontractsApproval, error) {
	event := new(Nation3PassportcontractsApproval)
	if err := _Nation3Passportcontracts.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3PassportcontractsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Nation3Passportcontracts contract.
type Nation3PassportcontractsApprovalForAllIterator struct {
	Event *Nation3PassportcontractsApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Nation3PassportcontractsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3PassportcontractsApprovalForAll)
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
		it.Event = new(Nation3PassportcontractsApprovalForAll)
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
func (it *Nation3PassportcontractsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3PassportcontractsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3PassportcontractsApprovalForAll represents a ApprovalForAll event raised by the Nation3Passportcontracts contract.
type Nation3PassportcontractsApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Nation3Passportcontracts *Nation3PassportcontractsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Nation3PassportcontractsApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Nation3Passportcontracts.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportcontractsApprovalForAllIterator{contract: _Nation3Passportcontracts.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Nation3Passportcontracts *Nation3PassportcontractsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Nation3PassportcontractsApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Nation3Passportcontracts.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3PassportcontractsApprovalForAll)
				if err := _Nation3Passportcontracts.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Nation3Passportcontracts *Nation3PassportcontractsFilterer) ParseApprovalForAll(log types.Log) (*Nation3PassportcontractsApprovalForAll, error) {
	event := new(Nation3PassportcontractsApprovalForAll)
	if err := _Nation3Passportcontracts.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3PassportcontractsControlTransferredIterator is returned from FilterControlTransferred and is used to iterate over the raw logs and unpacked data for ControlTransferred events raised by the Nation3Passportcontracts contract.
type Nation3PassportcontractsControlTransferredIterator struct {
	Event *Nation3PassportcontractsControlTransferred // Event containing the contract specifics and raw log

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
func (it *Nation3PassportcontractsControlTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3PassportcontractsControlTransferred)
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
		it.Event = new(Nation3PassportcontractsControlTransferred)
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
func (it *Nation3PassportcontractsControlTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3PassportcontractsControlTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3PassportcontractsControlTransferred represents a ControlTransferred event raised by the Nation3Passportcontracts contract.
type Nation3PassportcontractsControlTransferred struct {
	PreviousController common.Address
	NewController      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterControlTransferred is a free log retrieval operation binding the contract event 0xa06677f7b64342b4bcbde423684dbdb5356acfe41ad0285b6ecbe6dc4bf427f2.
//
// Solidity: event ControlTransferred(address indexed previousController, address indexed newController)
func (_Nation3Passportcontracts *Nation3PassportcontractsFilterer) FilterControlTransferred(opts *bind.FilterOpts, previousController []common.Address, newController []common.Address) (*Nation3PassportcontractsControlTransferredIterator, error) {

	var previousControllerRule []interface{}
	for _, previousControllerItem := range previousController {
		previousControllerRule = append(previousControllerRule, previousControllerItem)
	}
	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _Nation3Passportcontracts.contract.FilterLogs(opts, "ControlTransferred", previousControllerRule, newControllerRule)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportcontractsControlTransferredIterator{contract: _Nation3Passportcontracts.contract, event: "ControlTransferred", logs: logs, sub: sub}, nil
}

// WatchControlTransferred is a free log subscription operation binding the contract event 0xa06677f7b64342b4bcbde423684dbdb5356acfe41ad0285b6ecbe6dc4bf427f2.
//
// Solidity: event ControlTransferred(address indexed previousController, address indexed newController)
func (_Nation3Passportcontracts *Nation3PassportcontractsFilterer) WatchControlTransferred(opts *bind.WatchOpts, sink chan<- *Nation3PassportcontractsControlTransferred, previousController []common.Address, newController []common.Address) (event.Subscription, error) {

	var previousControllerRule []interface{}
	for _, previousControllerItem := range previousController {
		previousControllerRule = append(previousControllerRule, previousControllerItem)
	}
	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _Nation3Passportcontracts.contract.WatchLogs(opts, "ControlTransferred", previousControllerRule, newControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3PassportcontractsControlTransferred)
				if err := _Nation3Passportcontracts.contract.UnpackLog(event, "ControlTransferred", log); err != nil {
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
func (_Nation3Passportcontracts *Nation3PassportcontractsFilterer) ParseControlTransferred(log types.Log) (*Nation3PassportcontractsControlTransferred, error) {
	event := new(Nation3PassportcontractsControlTransferred)
	if err := _Nation3Passportcontracts.contract.UnpackLog(event, "ControlTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3PassportcontractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Nation3Passportcontracts contract.
type Nation3PassportcontractsOwnershipTransferredIterator struct {
	Event *Nation3PassportcontractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Nation3PassportcontractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3PassportcontractsOwnershipTransferred)
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
		it.Event = new(Nation3PassportcontractsOwnershipTransferred)
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
func (it *Nation3PassportcontractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3PassportcontractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3PassportcontractsOwnershipTransferred represents a OwnershipTransferred event raised by the Nation3Passportcontracts contract.
type Nation3PassportcontractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nation3Passportcontracts *Nation3PassportcontractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Nation3PassportcontractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nation3Passportcontracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportcontractsOwnershipTransferredIterator{contract: _Nation3Passportcontracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nation3Passportcontracts *Nation3PassportcontractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Nation3PassportcontractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nation3Passportcontracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3PassportcontractsOwnershipTransferred)
				if err := _Nation3Passportcontracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Nation3Passportcontracts *Nation3PassportcontractsFilterer) ParseOwnershipTransferred(log types.Log) (*Nation3PassportcontractsOwnershipTransferred, error) {
	event := new(Nation3PassportcontractsOwnershipTransferred)
	if err := _Nation3Passportcontracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3PassportcontractsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Nation3Passportcontracts contract.
type Nation3PassportcontractsTransferIterator struct {
	Event *Nation3PassportcontractsTransfer // Event containing the contract specifics and raw log

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
func (it *Nation3PassportcontractsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3PassportcontractsTransfer)
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
		it.Event = new(Nation3PassportcontractsTransfer)
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
func (it *Nation3PassportcontractsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3PassportcontractsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3PassportcontractsTransfer represents a Transfer event raised by the Nation3Passportcontracts contract.
type Nation3PassportcontractsTransfer struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Nation3Passportcontracts *Nation3PassportcontractsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*Nation3PassportcontractsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Nation3Passportcontracts.contract.FilterLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportcontractsTransferIterator{contract: _Nation3Passportcontracts.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Nation3Passportcontracts *Nation3PassportcontractsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Nation3PassportcontractsTransfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Nation3Passportcontracts.contract.WatchLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3PassportcontractsTransfer)
				if err := _Nation3Passportcontracts.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Nation3Passportcontracts *Nation3PassportcontractsFilterer) ParseTransfer(log types.Log) (*Nation3PassportcontractsTransfer, error) {
	event := new(Nation3PassportcontractsTransfer)
	if err := _Nation3Passportcontracts.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
