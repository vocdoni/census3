// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Nation3PassportContract

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

// Nation3PassportContractMetaData contains all meta data concerning the Nation3PassportContract contract.
var Nation3PassportContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerIsNotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFrom\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMinted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSafeRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetIsZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousController\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"ControlTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeControl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renderer\",\"outputs\":[{\"internalType\":\"contractRenderer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"safeMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRenderer\",\"name\":\"_renderer\",\"type\":\"address\"}],\"name\":\"setRenderer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"signerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"timestampOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"transferControl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Nation3PassportContractABI is the input ABI used to generate the binding from.
// Deprecated: Use Nation3PassportContractMetaData.ABI instead.
var Nation3PassportContractABI = Nation3PassportContractMetaData.ABI

// Nation3PassportContract is an auto generated Go binding around an Ethereum contract.
type Nation3PassportContract struct {
	Nation3PassportContractCaller     // Read-only binding to the contract
	Nation3PassportContractTransactor // Write-only binding to the contract
	Nation3PassportContractFilterer   // Log filterer for contract events
}

// Nation3PassportContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type Nation3PassportContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nation3PassportContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Nation3PassportContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nation3PassportContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Nation3PassportContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nation3PassportContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Nation3PassportContractSession struct {
	Contract     *Nation3PassportContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Nation3PassportContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Nation3PassportContractCallerSession struct {
	Contract *Nation3PassportContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// Nation3PassportContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Nation3PassportContractTransactorSession struct {
	Contract     *Nation3PassportContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// Nation3PassportContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type Nation3PassportContractRaw struct {
	Contract *Nation3PassportContract // Generic contract binding to access the raw methods on
}

// Nation3PassportContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Nation3PassportContractCallerRaw struct {
	Contract *Nation3PassportContractCaller // Generic read-only contract binding to access the raw methods on
}

// Nation3PassportContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Nation3PassportContractTransactorRaw struct {
	Contract *Nation3PassportContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNation3PassportContract creates a new instance of Nation3PassportContract, bound to a specific deployed contract.
func NewNation3PassportContract(address common.Address, backend bind.ContractBackend) (*Nation3PassportContract, error) {
	contract, err := bindNation3PassportContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportContract{Nation3PassportContractCaller: Nation3PassportContractCaller{contract: contract}, Nation3PassportContractTransactor: Nation3PassportContractTransactor{contract: contract}, Nation3PassportContractFilterer: Nation3PassportContractFilterer{contract: contract}}, nil
}

// NewNation3PassportContractCaller creates a new read-only instance of Nation3PassportContract, bound to a specific deployed contract.
func NewNation3PassportContractCaller(address common.Address, caller bind.ContractCaller) (*Nation3PassportContractCaller, error) {
	contract, err := bindNation3PassportContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportContractCaller{contract: contract}, nil
}

// NewNation3PassportContractTransactor creates a new write-only instance of Nation3PassportContract, bound to a specific deployed contract.
func NewNation3PassportContractTransactor(address common.Address, transactor bind.ContractTransactor) (*Nation3PassportContractTransactor, error) {
	contract, err := bindNation3PassportContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportContractTransactor{contract: contract}, nil
}

// NewNation3PassportContractFilterer creates a new log filterer instance of Nation3PassportContract, bound to a specific deployed contract.
func NewNation3PassportContractFilterer(address common.Address, filterer bind.ContractFilterer) (*Nation3PassportContractFilterer, error) {
	contract, err := bindNation3PassportContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportContractFilterer{contract: contract}, nil
}

// bindNation3PassportContract binds a generic wrapper to an already deployed contract.
func bindNation3PassportContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Nation3PassportContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nation3PassportContract *Nation3PassportContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nation3PassportContract.Contract.Nation3PassportContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nation3PassportContract *Nation3PassportContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.Nation3PassportContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nation3PassportContract *Nation3PassportContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.Nation3PassportContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nation3PassportContract *Nation3PassportContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nation3PassportContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nation3PassportContract *Nation3PassportContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nation3PassportContract *Nation3PassportContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nation3PassportContract *Nation3PassportContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nation3PassportContract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nation3PassportContract *Nation3PassportContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Nation3PassportContract.Contract.BalanceOf(&_Nation3PassportContract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nation3PassportContract *Nation3PassportContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Nation3PassportContract.Contract.BalanceOf(&_Nation3PassportContract.CallOpts, owner)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Nation3PassportContract *Nation3PassportContractCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3PassportContract.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Nation3PassportContract *Nation3PassportContractSession) Controller() (common.Address, error) {
	return _Nation3PassportContract.Contract.Controller(&_Nation3PassportContract.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Nation3PassportContract *Nation3PassportContractCallerSession) Controller() (common.Address, error) {
	return _Nation3PassportContract.Contract.Controller(&_Nation3PassportContract.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Nation3PassportContract *Nation3PassportContractCaller) GetApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nation3PassportContract.contract.Call(opts, &out, "getApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Nation3PassportContract *Nation3PassportContractSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Nation3PassportContract.Contract.GetApproved(&_Nation3PassportContract.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Nation3PassportContract *Nation3PassportContractCallerSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Nation3PassportContract.Contract.GetApproved(&_Nation3PassportContract.CallOpts, arg0)
}

// GetNextId is a free data retrieval call binding the contract method 0xbc968326.
//
// Solidity: function getNextId() view returns(uint256)
func (_Nation3PassportContract *Nation3PassportContractCaller) GetNextId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nation3PassportContract.contract.Call(opts, &out, "getNextId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextId is a free data retrieval call binding the contract method 0xbc968326.
//
// Solidity: function getNextId() view returns(uint256)
func (_Nation3PassportContract *Nation3PassportContractSession) GetNextId() (*big.Int, error) {
	return _Nation3PassportContract.Contract.GetNextId(&_Nation3PassportContract.CallOpts)
}

// GetNextId is a free data retrieval call binding the contract method 0xbc968326.
//
// Solidity: function getNextId() view returns(uint256)
func (_Nation3PassportContract *Nation3PassportContractCallerSession) GetNextId() (*big.Int, error) {
	return _Nation3PassportContract.Contract.GetNextId(&_Nation3PassportContract.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Nation3PassportContract *Nation3PassportContractCaller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Nation3PassportContract.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Nation3PassportContract *Nation3PassportContractSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Nation3PassportContract.Contract.IsApprovedForAll(&_Nation3PassportContract.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Nation3PassportContract *Nation3PassportContractCallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Nation3PassportContract.Contract.IsApprovedForAll(&_Nation3PassportContract.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nation3PassportContract *Nation3PassportContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nation3PassportContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nation3PassportContract *Nation3PassportContractSession) Name() (string, error) {
	return _Nation3PassportContract.Contract.Name(&_Nation3PassportContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nation3PassportContract *Nation3PassportContractCallerSession) Name() (string, error) {
	return _Nation3PassportContract.Contract.Name(&_Nation3PassportContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nation3PassportContract *Nation3PassportContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3PassportContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nation3PassportContract *Nation3PassportContractSession) Owner() (common.Address, error) {
	return _Nation3PassportContract.Contract.Owner(&_Nation3PassportContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nation3PassportContract *Nation3PassportContractCallerSession) Owner() (common.Address, error) {
	return _Nation3PassportContract.Contract.Owner(&_Nation3PassportContract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Nation3PassportContract *Nation3PassportContractCaller) OwnerOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nation3PassportContract.contract.Call(opts, &out, "ownerOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Nation3PassportContract *Nation3PassportContractSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Nation3PassportContract.Contract.OwnerOf(&_Nation3PassportContract.CallOpts, id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Nation3PassportContract *Nation3PassportContractCallerSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Nation3PassportContract.Contract.OwnerOf(&_Nation3PassportContract.CallOpts, id)
}

// Renderer is a free data retrieval call binding the contract method 0x8ada6b0f.
//
// Solidity: function renderer() view returns(address)
func (_Nation3PassportContract *Nation3PassportContractCaller) Renderer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3PassportContract.contract.Call(opts, &out, "renderer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Renderer is a free data retrieval call binding the contract method 0x8ada6b0f.
//
// Solidity: function renderer() view returns(address)
func (_Nation3PassportContract *Nation3PassportContractSession) Renderer() (common.Address, error) {
	return _Nation3PassportContract.Contract.Renderer(&_Nation3PassportContract.CallOpts)
}

// Renderer is a free data retrieval call binding the contract method 0x8ada6b0f.
//
// Solidity: function renderer() view returns(address)
func (_Nation3PassportContract *Nation3PassportContractCallerSession) Renderer() (common.Address, error) {
	return _Nation3PassportContract.Contract.Renderer(&_Nation3PassportContract.CallOpts)
}

// SignerOf is a free data retrieval call binding the contract method 0x5161fdf5.
//
// Solidity: function signerOf(uint256 id) view returns(address)
func (_Nation3PassportContract *Nation3PassportContractCaller) SignerOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nation3PassportContract.contract.Call(opts, &out, "signerOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerOf is a free data retrieval call binding the contract method 0x5161fdf5.
//
// Solidity: function signerOf(uint256 id) view returns(address)
func (_Nation3PassportContract *Nation3PassportContractSession) SignerOf(id *big.Int) (common.Address, error) {
	return _Nation3PassportContract.Contract.SignerOf(&_Nation3PassportContract.CallOpts, id)
}

// SignerOf is a free data retrieval call binding the contract method 0x5161fdf5.
//
// Solidity: function signerOf(uint256 id) view returns(address)
func (_Nation3PassportContract *Nation3PassportContractCallerSession) SignerOf(id *big.Int) (common.Address, error) {
	return _Nation3PassportContract.Contract.SignerOf(&_Nation3PassportContract.CallOpts, id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nation3PassportContract *Nation3PassportContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Nation3PassportContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nation3PassportContract *Nation3PassportContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Nation3PassportContract.Contract.SupportsInterface(&_Nation3PassportContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nation3PassportContract *Nation3PassportContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Nation3PassportContract.Contract.SupportsInterface(&_Nation3PassportContract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nation3PassportContract *Nation3PassportContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nation3PassportContract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nation3PassportContract *Nation3PassportContractSession) Symbol() (string, error) {
	return _Nation3PassportContract.Contract.Symbol(&_Nation3PassportContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nation3PassportContract *Nation3PassportContractCallerSession) Symbol() (string, error) {
	return _Nation3PassportContract.Contract.Symbol(&_Nation3PassportContract.CallOpts)
}

// TimestampOf is a free data retrieval call binding the contract method 0x2d9c77e1.
//
// Solidity: function timestampOf(uint256 id) view returns(uint256)
func (_Nation3PassportContract *Nation3PassportContractCaller) TimestampOf(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nation3PassportContract.contract.Call(opts, &out, "timestampOf", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimestampOf is a free data retrieval call binding the contract method 0x2d9c77e1.
//
// Solidity: function timestampOf(uint256 id) view returns(uint256)
func (_Nation3PassportContract *Nation3PassportContractSession) TimestampOf(id *big.Int) (*big.Int, error) {
	return _Nation3PassportContract.Contract.TimestampOf(&_Nation3PassportContract.CallOpts, id)
}

// TimestampOf is a free data retrieval call binding the contract method 0x2d9c77e1.
//
// Solidity: function timestampOf(uint256 id) view returns(uint256)
func (_Nation3PassportContract *Nation3PassportContractCallerSession) TimestampOf(id *big.Int) (*big.Int, error) {
	return _Nation3PassportContract.Contract.TimestampOf(&_Nation3PassportContract.CallOpts, id)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_Nation3PassportContract *Nation3PassportContractCaller) TokenURI(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _Nation3PassportContract.contract.Call(opts, &out, "tokenURI", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_Nation3PassportContract *Nation3PassportContractSession) TokenURI(id *big.Int) (string, error) {
	return _Nation3PassportContract.Contract.TokenURI(&_Nation3PassportContract.CallOpts, id)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_Nation3PassportContract *Nation3PassportContractCallerSession) TokenURI(id *big.Int) (string, error) {
	return _Nation3PassportContract.Contract.TokenURI(&_Nation3PassportContract.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nation3PassportContract *Nation3PassportContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nation3PassportContract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nation3PassportContract *Nation3PassportContractSession) TotalSupply() (*big.Int, error) {
	return _Nation3PassportContract.Contract.TotalSupply(&_Nation3PassportContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nation3PassportContract *Nation3PassportContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Nation3PassportContract.Contract.TotalSupply(&_Nation3PassportContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3PassportContract.contract.Transact(opts, "approve", spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Nation3PassportContract *Nation3PassportContractSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.Approve(&_Nation3PassportContract.TransactOpts, spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactorSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.Approve(&_Nation3PassportContract.TransactOpts, spender, id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 id) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactor) Burn(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Nation3PassportContract.contract.Transact(opts, "burn", id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 id) returns()
func (_Nation3PassportContract *Nation3PassportContractSession) Burn(id *big.Int) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.Burn(&_Nation3PassportContract.TransactOpts, id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 id) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactorSession) Burn(id *big.Int) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.Burn(&_Nation3PassportContract.TransactOpts, id)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 tokenId)
func (_Nation3PassportContract *Nation3PassportContractTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 tokenId)
func (_Nation3PassportContract *Nation3PassportContractSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.Mint(&_Nation3PassportContract.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 tokenId)
func (_Nation3PassportContract *Nation3PassportContractTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.Mint(&_Nation3PassportContract.TransactOpts, to)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x056097ac.
//
// Solidity: function recoverTokens(address token, address to) returns(uint256 amount)
func (_Nation3PassportContract *Nation3PassportContractTransactor) RecoverTokens(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.contract.Transact(opts, "recoverTokens", token, to)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x056097ac.
//
// Solidity: function recoverTokens(address token, address to) returns(uint256 amount)
func (_Nation3PassportContract *Nation3PassportContractSession) RecoverTokens(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.RecoverTokens(&_Nation3PassportContract.TransactOpts, token, to)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x056097ac.
//
// Solidity: function recoverTokens(address token, address to) returns(uint256 amount)
func (_Nation3PassportContract *Nation3PassportContractTransactorSession) RecoverTokens(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.RecoverTokens(&_Nation3PassportContract.TransactOpts, token, to)
}

// RemoveControl is a paid mutator transaction binding the contract method 0x7bee684b.
//
// Solidity: function removeControl() returns()
func (_Nation3PassportContract *Nation3PassportContractTransactor) RemoveControl(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3PassportContract.contract.Transact(opts, "removeControl")
}

// RemoveControl is a paid mutator transaction binding the contract method 0x7bee684b.
//
// Solidity: function removeControl() returns()
func (_Nation3PassportContract *Nation3PassportContractSession) RemoveControl() (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.RemoveControl(&_Nation3PassportContract.TransactOpts)
}

// RemoveControl is a paid mutator transaction binding the contract method 0x7bee684b.
//
// Solidity: function removeControl() returns()
func (_Nation3PassportContract *Nation3PassportContractTransactorSession) RemoveControl() (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.RemoveControl(&_Nation3PassportContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nation3PassportContract *Nation3PassportContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3PassportContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nation3PassportContract *Nation3PassportContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.RenounceOwnership(&_Nation3PassportContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nation3PassportContract *Nation3PassportContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.RenounceOwnership(&_Nation3PassportContract.TransactOpts)
}

// SafeMint is a paid mutator transaction binding the contract method 0x40d097c3.
//
// Solidity: function safeMint(address to) returns(uint256 tokenId)
func (_Nation3PassportContract *Nation3PassportContractTransactor) SafeMint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.contract.Transact(opts, "safeMint", to)
}

// SafeMint is a paid mutator transaction binding the contract method 0x40d097c3.
//
// Solidity: function safeMint(address to) returns(uint256 tokenId)
func (_Nation3PassportContract *Nation3PassportContractSession) SafeMint(to common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.SafeMint(&_Nation3PassportContract.TransactOpts, to)
}

// SafeMint is a paid mutator transaction binding the contract method 0x40d097c3.
//
// Solidity: function safeMint(address to) returns(uint256 tokenId)
func (_Nation3PassportContract *Nation3PassportContractTransactorSession) SafeMint(to common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.SafeMint(&_Nation3PassportContract.TransactOpts, to)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3PassportContract.contract.Transact(opts, "safeTransferFrom", from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Nation3PassportContract *Nation3PassportContractSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.SafeTransferFrom(&_Nation3PassportContract.TransactOpts, from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.SafeTransferFrom(&_Nation3PassportContract.TransactOpts, from, to, id)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Nation3PassportContract.contract.Transact(opts, "safeTransferFrom0", from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Nation3PassportContract *Nation3PassportContractSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.SafeTransferFrom0(&_Nation3PassportContract.TransactOpts, from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.SafeTransferFrom0(&_Nation3PassportContract.TransactOpts, from, to, id, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nation3PassportContract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nation3PassportContract *Nation3PassportContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.SetApprovalForAll(&_Nation3PassportContract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.SetApprovalForAll(&_Nation3PassportContract.TransactOpts, operator, approved)
}

// SetRenderer is a paid mutator transaction binding the contract method 0x56d3163d.
//
// Solidity: function setRenderer(address _renderer) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactor) SetRenderer(opts *bind.TransactOpts, _renderer common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.contract.Transact(opts, "setRenderer", _renderer)
}

// SetRenderer is a paid mutator transaction binding the contract method 0x56d3163d.
//
// Solidity: function setRenderer(address _renderer) returns()
func (_Nation3PassportContract *Nation3PassportContractSession) SetRenderer(_renderer common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.SetRenderer(&_Nation3PassportContract.TransactOpts, _renderer)
}

// SetRenderer is a paid mutator transaction binding the contract method 0x56d3163d.
//
// Solidity: function setRenderer(address _renderer) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactorSession) SetRenderer(_renderer common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.SetRenderer(&_Nation3PassportContract.TransactOpts, _renderer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x85267907.
//
// Solidity: function setSigner(uint256 id, address signer) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactor) SetSigner(opts *bind.TransactOpts, id *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.contract.Transact(opts, "setSigner", id, signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x85267907.
//
// Solidity: function setSigner(uint256 id, address signer) returns()
func (_Nation3PassportContract *Nation3PassportContractSession) SetSigner(id *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.SetSigner(&_Nation3PassportContract.TransactOpts, id, signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x85267907.
//
// Solidity: function setSigner(uint256 id, address signer) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactorSession) SetSigner(id *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.SetSigner(&_Nation3PassportContract.TransactOpts, id, signer)
}

// TransferControl is a paid mutator transaction binding the contract method 0x6d16fa41.
//
// Solidity: function transferControl(address newController) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactor) TransferControl(opts *bind.TransactOpts, newController common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.contract.Transact(opts, "transferControl", newController)
}

// TransferControl is a paid mutator transaction binding the contract method 0x6d16fa41.
//
// Solidity: function transferControl(address newController) returns()
func (_Nation3PassportContract *Nation3PassportContractSession) TransferControl(newController common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.TransferControl(&_Nation3PassportContract.TransactOpts, newController)
}

// TransferControl is a paid mutator transaction binding the contract method 0x6d16fa41.
//
// Solidity: function transferControl(address newController) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactorSession) TransferControl(newController common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.TransferControl(&_Nation3PassportContract.TransactOpts, newController)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3PassportContract.contract.Transact(opts, "transferFrom", from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Nation3PassportContract *Nation3PassportContractSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.TransferFrom(&_Nation3PassportContract.TransactOpts, from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactorSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.TransferFrom(&_Nation3PassportContract.TransactOpts, from, to, id)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nation3PassportContract *Nation3PassportContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.TransferOwnership(&_Nation3PassportContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nation3PassportContract *Nation3PassportContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nation3PassportContract.Contract.TransferOwnership(&_Nation3PassportContract.TransactOpts, newOwner)
}

// Nation3PassportContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Nation3PassportContract contract.
type Nation3PassportContractApprovalIterator struct {
	Event *Nation3PassportContractApproval // Event containing the contract specifics and raw log

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
func (it *Nation3PassportContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3PassportContractApproval)
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
		it.Event = new(Nation3PassportContractApproval)
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
func (it *Nation3PassportContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3PassportContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3PassportContractApproval represents a Approval event raised by the Nation3PassportContract contract.
type Nation3PassportContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Nation3PassportContract *Nation3PassportContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address, id []*big.Int) (*Nation3PassportContractApprovalIterator, error) {

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

	logs, sub, err := _Nation3PassportContract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportContractApprovalIterator{contract: _Nation3PassportContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Nation3PassportContract *Nation3PassportContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Nation3PassportContractApproval, owner []common.Address, spender []common.Address, id []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Nation3PassportContract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3PassportContractApproval)
				if err := _Nation3PassportContract.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Nation3PassportContract *Nation3PassportContractFilterer) ParseApproval(log types.Log) (*Nation3PassportContractApproval, error) {
	event := new(Nation3PassportContractApproval)
	if err := _Nation3PassportContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3PassportContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Nation3PassportContract contract.
type Nation3PassportContractApprovalForAllIterator struct {
	Event *Nation3PassportContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Nation3PassportContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3PassportContractApprovalForAll)
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
		it.Event = new(Nation3PassportContractApprovalForAll)
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
func (it *Nation3PassportContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3PassportContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3PassportContractApprovalForAll represents a ApprovalForAll event raised by the Nation3PassportContract contract.
type Nation3PassportContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Nation3PassportContract *Nation3PassportContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Nation3PassportContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Nation3PassportContract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportContractApprovalForAllIterator{contract: _Nation3PassportContract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Nation3PassportContract *Nation3PassportContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Nation3PassportContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Nation3PassportContract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3PassportContractApprovalForAll)
				if err := _Nation3PassportContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Nation3PassportContract *Nation3PassportContractFilterer) ParseApprovalForAll(log types.Log) (*Nation3PassportContractApprovalForAll, error) {
	event := new(Nation3PassportContractApprovalForAll)
	if err := _Nation3PassportContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3PassportContractControlTransferredIterator is returned from FilterControlTransferred and is used to iterate over the raw logs and unpacked data for ControlTransferred events raised by the Nation3PassportContract contract.
type Nation3PassportContractControlTransferredIterator struct {
	Event *Nation3PassportContractControlTransferred // Event containing the contract specifics and raw log

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
func (it *Nation3PassportContractControlTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3PassportContractControlTransferred)
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
		it.Event = new(Nation3PassportContractControlTransferred)
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
func (it *Nation3PassportContractControlTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3PassportContractControlTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3PassportContractControlTransferred represents a ControlTransferred event raised by the Nation3PassportContract contract.
type Nation3PassportContractControlTransferred struct {
	PreviousController common.Address
	NewController      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterControlTransferred is a free log retrieval operation binding the contract event 0xa06677f7b64342b4bcbde423684dbdb5356acfe41ad0285b6ecbe6dc4bf427f2.
//
// Solidity: event ControlTransferred(address indexed previousController, address indexed newController)
func (_Nation3PassportContract *Nation3PassportContractFilterer) FilterControlTransferred(opts *bind.FilterOpts, previousController []common.Address, newController []common.Address) (*Nation3PassportContractControlTransferredIterator, error) {

	var previousControllerRule []interface{}
	for _, previousControllerItem := range previousController {
		previousControllerRule = append(previousControllerRule, previousControllerItem)
	}
	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _Nation3PassportContract.contract.FilterLogs(opts, "ControlTransferred", previousControllerRule, newControllerRule)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportContractControlTransferredIterator{contract: _Nation3PassportContract.contract, event: "ControlTransferred", logs: logs, sub: sub}, nil
}

// WatchControlTransferred is a free log subscription operation binding the contract event 0xa06677f7b64342b4bcbde423684dbdb5356acfe41ad0285b6ecbe6dc4bf427f2.
//
// Solidity: event ControlTransferred(address indexed previousController, address indexed newController)
func (_Nation3PassportContract *Nation3PassportContractFilterer) WatchControlTransferred(opts *bind.WatchOpts, sink chan<- *Nation3PassportContractControlTransferred, previousController []common.Address, newController []common.Address) (event.Subscription, error) {

	var previousControllerRule []interface{}
	for _, previousControllerItem := range previousController {
		previousControllerRule = append(previousControllerRule, previousControllerItem)
	}
	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _Nation3PassportContract.contract.WatchLogs(opts, "ControlTransferred", previousControllerRule, newControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3PassportContractControlTransferred)
				if err := _Nation3PassportContract.contract.UnpackLog(event, "ControlTransferred", log); err != nil {
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
func (_Nation3PassportContract *Nation3PassportContractFilterer) ParseControlTransferred(log types.Log) (*Nation3PassportContractControlTransferred, error) {
	event := new(Nation3PassportContractControlTransferred)
	if err := _Nation3PassportContract.contract.UnpackLog(event, "ControlTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3PassportContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Nation3PassportContract contract.
type Nation3PassportContractOwnershipTransferredIterator struct {
	Event *Nation3PassportContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Nation3PassportContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3PassportContractOwnershipTransferred)
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
		it.Event = new(Nation3PassportContractOwnershipTransferred)
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
func (it *Nation3PassportContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3PassportContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3PassportContractOwnershipTransferred represents a OwnershipTransferred event raised by the Nation3PassportContract contract.
type Nation3PassportContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nation3PassportContract *Nation3PassportContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Nation3PassportContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nation3PassportContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportContractOwnershipTransferredIterator{contract: _Nation3PassportContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nation3PassportContract *Nation3PassportContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Nation3PassportContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nation3PassportContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3PassportContractOwnershipTransferred)
				if err := _Nation3PassportContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Nation3PassportContract *Nation3PassportContractFilterer) ParseOwnershipTransferred(log types.Log) (*Nation3PassportContractOwnershipTransferred, error) {
	event := new(Nation3PassportContractOwnershipTransferred)
	if err := _Nation3PassportContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3PassportContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Nation3PassportContract contract.
type Nation3PassportContractTransferIterator struct {
	Event *Nation3PassportContractTransfer // Event containing the contract specifics and raw log

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
func (it *Nation3PassportContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3PassportContractTransfer)
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
		it.Event = new(Nation3PassportContractTransfer)
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
func (it *Nation3PassportContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3PassportContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3PassportContractTransfer represents a Transfer event raised by the Nation3PassportContract contract.
type Nation3PassportContractTransfer struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Nation3PassportContract *Nation3PassportContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*Nation3PassportContractTransferIterator, error) {

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

	logs, sub, err := _Nation3PassportContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &Nation3PassportContractTransferIterator{contract: _Nation3PassportContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Nation3PassportContract *Nation3PassportContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Nation3PassportContractTransfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Nation3PassportContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3PassportContractTransfer)
				if err := _Nation3PassportContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Nation3PassportContract *Nation3PassportContractFilterer) ParseTransfer(log types.Log) (*Nation3PassportContractTransfer, error) {
	event := new(Nation3PassportContractTransfer)
	if err := _Nation3PassportContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
