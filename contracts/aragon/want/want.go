// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AragonWrappedANTTokenContract

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

// AragonWrappedANTTokenContractMetaData contains all meta data concerning the AragonWrappedANTTokenContract contract.
var AragonWrappedANTTokenContractMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_script\",\"type\":\"bytes\"}],\"name\":\"getEVMScriptExecutor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecoveryVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"allowRecoverability\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitializationBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_depositedToken\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"transferToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"bytes32\"},{\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"canPerform\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEVMScriptRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depositedToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPetrified\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"entity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"entity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"script\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ScriptResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoverToVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
}

// AragonWrappedANTTokenContractABI is the input ABI used to generate the binding from.
// Deprecated: Use AragonWrappedANTTokenContractMetaData.ABI instead.
var AragonWrappedANTTokenContractABI = AragonWrappedANTTokenContractMetaData.ABI

// AragonWrappedANTTokenContract is an auto generated Go binding around an Ethereum contract.
type AragonWrappedANTTokenContract struct {
	AragonWrappedANTTokenContractCaller     // Read-only binding to the contract
	AragonWrappedANTTokenContractTransactor // Write-only binding to the contract
	AragonWrappedANTTokenContractFilterer   // Log filterer for contract events
}

// AragonWrappedANTTokenContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AragonWrappedANTTokenContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AragonWrappedANTTokenContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AragonWrappedANTTokenContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AragonWrappedANTTokenContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AragonWrappedANTTokenContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AragonWrappedANTTokenContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AragonWrappedANTTokenContractSession struct {
	Contract     *AragonWrappedANTTokenContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AragonWrappedANTTokenContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AragonWrappedANTTokenContractCallerSession struct {
	Contract *AragonWrappedANTTokenContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// AragonWrappedANTTokenContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AragonWrappedANTTokenContractTransactorSession struct {
	Contract     *AragonWrappedANTTokenContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// AragonWrappedANTTokenContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AragonWrappedANTTokenContractRaw struct {
	Contract *AragonWrappedANTTokenContract // Generic contract binding to access the raw methods on
}

// AragonWrappedANTTokenContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AragonWrappedANTTokenContractCallerRaw struct {
	Contract *AragonWrappedANTTokenContractCaller // Generic read-only contract binding to access the raw methods on
}

// AragonWrappedANTTokenContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AragonWrappedANTTokenContractTransactorRaw struct {
	Contract *AragonWrappedANTTokenContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAragonWrappedANTTokenContract creates a new instance of AragonWrappedANTTokenContract, bound to a specific deployed contract.
func NewAragonWrappedANTTokenContract(address common.Address, backend bind.ContractBackend) (*AragonWrappedANTTokenContract, error) {
	contract, err := bindAragonWrappedANTTokenContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AragonWrappedANTTokenContract{AragonWrappedANTTokenContractCaller: AragonWrappedANTTokenContractCaller{contract: contract}, AragonWrappedANTTokenContractTransactor: AragonWrappedANTTokenContractTransactor{contract: contract}, AragonWrappedANTTokenContractFilterer: AragonWrappedANTTokenContractFilterer{contract: contract}}, nil
}

// NewAragonWrappedANTTokenContractCaller creates a new read-only instance of AragonWrappedANTTokenContract, bound to a specific deployed contract.
func NewAragonWrappedANTTokenContractCaller(address common.Address, caller bind.ContractCaller) (*AragonWrappedANTTokenContractCaller, error) {
	contract, err := bindAragonWrappedANTTokenContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AragonWrappedANTTokenContractCaller{contract: contract}, nil
}

// NewAragonWrappedANTTokenContractTransactor creates a new write-only instance of AragonWrappedANTTokenContract, bound to a specific deployed contract.
func NewAragonWrappedANTTokenContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AragonWrappedANTTokenContractTransactor, error) {
	contract, err := bindAragonWrappedANTTokenContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AragonWrappedANTTokenContractTransactor{contract: contract}, nil
}

// NewAragonWrappedANTTokenContractFilterer creates a new log filterer instance of AragonWrappedANTTokenContract, bound to a specific deployed contract.
func NewAragonWrappedANTTokenContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AragonWrappedANTTokenContractFilterer, error) {
	contract, err := bindAragonWrappedANTTokenContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AragonWrappedANTTokenContractFilterer{contract: contract}, nil
}

// bindAragonWrappedANTTokenContract binds a generic wrapper to an already deployed contract.
func bindAragonWrappedANTTokenContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AragonWrappedANTTokenContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AragonWrappedANTTokenContract.Contract.AragonWrappedANTTokenContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.AragonWrappedANTTokenContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.AragonWrappedANTTokenContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AragonWrappedANTTokenContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.contract.Transact(opts, method, params...)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address _token) view returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) AllowRecoverability(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "allowRecoverability", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address _token) view returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) AllowRecoverability(_token common.Address) (bool, error) {
	return _AragonWrappedANTTokenContract.Contract.AllowRecoverability(&_AragonWrappedANTTokenContract.CallOpts, _token)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address _token) view returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) AllowRecoverability(_token common.Address) (bool, error) {
	return _AragonWrappedANTTokenContract.Contract.AllowRecoverability(&_AragonWrappedANTTokenContract.CallOpts, _token)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _AragonWrappedANTTokenContract.Contract.Allowance(&_AragonWrappedANTTokenContract.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _AragonWrappedANTTokenContract.Contract.Allowance(&_AragonWrappedANTTokenContract.CallOpts, arg0, arg1)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "appId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) AppId() ([32]byte, error) {
	return _AragonWrappedANTTokenContract.Contract.AppId(&_AragonWrappedANTTokenContract.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) AppId() ([32]byte, error) {
	return _AragonWrappedANTTokenContract.Contract.AppId(&_AragonWrappedANTTokenContract.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _AragonWrappedANTTokenContract.Contract.BalanceOf(&_AragonWrappedANTTokenContract.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _AragonWrappedANTTokenContract.Contract.BalanceOf(&_AragonWrappedANTTokenContract.CallOpts, _owner)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) BalanceOfAt(opts *bind.CallOpts, _owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "balanceOfAt", _owner, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) BalanceOfAt(_owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _AragonWrappedANTTokenContract.Contract.BalanceOfAt(&_AragonWrappedANTTokenContract.CallOpts, _owner, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) BalanceOfAt(_owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _AragonWrappedANTTokenContract.Contract.BalanceOfAt(&_AragonWrappedANTTokenContract.CallOpts, _owner, _blockNumber)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) CanPerform(opts *bind.CallOpts, _sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "canPerform", _sender, _role, _params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _AragonWrappedANTTokenContract.Contract.CanPerform(&_AragonWrappedANTTokenContract.CallOpts, _sender, _role, _params)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _AragonWrappedANTTokenContract.Contract.CanPerform(&_AragonWrappedANTTokenContract.CallOpts, _sender, _role, _params)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) Decimals() (uint8, error) {
	return _AragonWrappedANTTokenContract.Contract.Decimals(&_AragonWrappedANTTokenContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) Decimals() (uint8, error) {
	return _AragonWrappedANTTokenContract.Contract.Decimals(&_AragonWrappedANTTokenContract.CallOpts)
}

// DepositedToken is a free data retrieval call binding the contract method 0xdad9b086.
//
// Solidity: function depositedToken() view returns(address)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) DepositedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "depositedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositedToken is a free data retrieval call binding the contract method 0xdad9b086.
//
// Solidity: function depositedToken() view returns(address)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) DepositedToken() (common.Address, error) {
	return _AragonWrappedANTTokenContract.Contract.DepositedToken(&_AragonWrappedANTTokenContract.CallOpts)
}

// DepositedToken is a free data retrieval call binding the contract method 0xdad9b086.
//
// Solidity: function depositedToken() view returns(address)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) DepositedToken() (common.Address, error) {
	return _AragonWrappedANTTokenContract.Contract.DepositedToken(&_AragonWrappedANTTokenContract.CallOpts)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) GetEVMScriptExecutor(opts *bind.CallOpts, _script []byte) (common.Address, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "getEVMScriptExecutor", _script)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _AragonWrappedANTTokenContract.Contract.GetEVMScriptExecutor(&_AragonWrappedANTTokenContract.CallOpts, _script)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _AragonWrappedANTTokenContract.Contract.GetEVMScriptExecutor(&_AragonWrappedANTTokenContract.CallOpts, _script)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) GetEVMScriptRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "getEVMScriptRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) GetEVMScriptRegistry() (common.Address, error) {
	return _AragonWrappedANTTokenContract.Contract.GetEVMScriptRegistry(&_AragonWrappedANTTokenContract.CallOpts)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) GetEVMScriptRegistry() (common.Address, error) {
	return _AragonWrappedANTTokenContract.Contract.GetEVMScriptRegistry(&_AragonWrappedANTTokenContract.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) GetInitializationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "getInitializationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) GetInitializationBlock() (*big.Int, error) {
	return _AragonWrappedANTTokenContract.Contract.GetInitializationBlock(&_AragonWrappedANTTokenContract.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) GetInitializationBlock() (*big.Int, error) {
	return _AragonWrappedANTTokenContract.Contract.GetInitializationBlock(&_AragonWrappedANTTokenContract.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) GetRecoveryVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "getRecoveryVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) GetRecoveryVault() (common.Address, error) {
	return _AragonWrappedANTTokenContract.Contract.GetRecoveryVault(&_AragonWrappedANTTokenContract.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) GetRecoveryVault() (common.Address, error) {
	return _AragonWrappedANTTokenContract.Contract.GetRecoveryVault(&_AragonWrappedANTTokenContract.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) HasInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "hasInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) HasInitialized() (bool, error) {
	return _AragonWrappedANTTokenContract.Contract.HasInitialized(&_AragonWrappedANTTokenContract.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) HasInitialized() (bool, error) {
	return _AragonWrappedANTTokenContract.Contract.HasInitialized(&_AragonWrappedANTTokenContract.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) IsPetrified(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "isPetrified")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) IsPetrified() (bool, error) {
	return _AragonWrappedANTTokenContract.Contract.IsPetrified(&_AragonWrappedANTTokenContract.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) IsPetrified() (bool, error) {
	return _AragonWrappedANTTokenContract.Contract.IsPetrified(&_AragonWrappedANTTokenContract.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "kernel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) Kernel() (common.Address, error) {
	return _AragonWrappedANTTokenContract.Contract.Kernel(&_AragonWrappedANTTokenContract.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) Kernel() (common.Address, error) {
	return _AragonWrappedANTTokenContract.Contract.Kernel(&_AragonWrappedANTTokenContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) Name() (string, error) {
	return _AragonWrappedANTTokenContract.Contract.Name(&_AragonWrappedANTTokenContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) Name() (string, error) {
	return _AragonWrappedANTTokenContract.Contract.Name(&_AragonWrappedANTTokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) Symbol() (string, error) {
	return _AragonWrappedANTTokenContract.Contract.Symbol(&_AragonWrappedANTTokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) Symbol() (string, error) {
	return _AragonWrappedANTTokenContract.Contract.Symbol(&_AragonWrappedANTTokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) TotalSupply() (*big.Int, error) {
	return _AragonWrappedANTTokenContract.Contract.TotalSupply(&_AragonWrappedANTTokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) TotalSupply() (*big.Int, error) {
	return _AragonWrappedANTTokenContract.Contract.TotalSupply(&_AragonWrappedANTTokenContract.CallOpts)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _blockNumber) view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCaller) TotalSupplyAt(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AragonWrappedANTTokenContract.contract.Call(opts, &out, "totalSupplyAt", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _blockNumber) view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) TotalSupplyAt(_blockNumber *big.Int) (*big.Int, error) {
	return _AragonWrappedANTTokenContract.Contract.TotalSupplyAt(&_AragonWrappedANTTokenContract.CallOpts, _blockNumber)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _blockNumber) view returns(uint256)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractCallerSession) TotalSupplyAt(_blockNumber *big.Int) (*big.Int, error) {
	return _AragonWrappedANTTokenContract.Contract.TotalSupplyAt(&_AragonWrappedANTTokenContract.CallOpts, _blockNumber)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.Approve(&_AragonWrappedANTTokenContract.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.Approve(&_AragonWrappedANTTokenContract.TransactOpts, arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.Deposit(&_AragonWrappedANTTokenContract.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.Deposit(&_AragonWrappedANTTokenContract.TransactOpts, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(address _depositedToken, string _name, string _symbol) returns()
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactor) Initialize(opts *bind.TransactOpts, _depositedToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.contract.Transact(opts, "initialize", _depositedToken, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(address _depositedToken, string _name, string _symbol) returns()
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) Initialize(_depositedToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.Initialize(&_AragonWrappedANTTokenContract.TransactOpts, _depositedToken, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(address _depositedToken, string _name, string _symbol) returns()
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactorSession) Initialize(_depositedToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.Initialize(&_AragonWrappedANTTokenContract.TransactOpts, _depositedToken, _name, _symbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactor) Transfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.contract.Transact(opts, "transfer", arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.Transfer(&_AragonWrappedANTTokenContract.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactorSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.Transfer(&_AragonWrappedANTTokenContract.TransactOpts, arg0, arg1)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactor) TransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.TransferFrom(&_AragonWrappedANTTokenContract.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactorSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.TransferFrom(&_AragonWrappedANTTokenContract.TransactOpts, arg0, arg1, arg2)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactor) TransferToVault(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.contract.Transact(opts, "transferToVault", _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.TransferToVault(&_AragonWrappedANTTokenContract.TransactOpts, _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactorSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.TransferToVault(&_AragonWrappedANTTokenContract.TransactOpts, _token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.Withdraw(&_AragonWrappedANTTokenContract.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _AragonWrappedANTTokenContract.Contract.Withdraw(&_AragonWrappedANTTokenContract.TransactOpts, _amount)
}

// AragonWrappedANTTokenContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AragonWrappedANTTokenContract contract.
type AragonWrappedANTTokenContractApprovalIterator struct {
	Event *AragonWrappedANTTokenContractApproval // Event containing the contract specifics and raw log

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
func (it *AragonWrappedANTTokenContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonWrappedANTTokenContractApproval)
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
		it.Event = new(AragonWrappedANTTokenContractApproval)
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
func (it *AragonWrappedANTTokenContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonWrappedANTTokenContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonWrappedANTTokenContractApproval represents a Approval event raised by the AragonWrappedANTTokenContract contract.
type AragonWrappedANTTokenContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AragonWrappedANTTokenContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AragonWrappedANTTokenContract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AragonWrappedANTTokenContractApprovalIterator{contract: _AragonWrappedANTTokenContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AragonWrappedANTTokenContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AragonWrappedANTTokenContract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonWrappedANTTokenContractApproval)
				if err := _AragonWrappedANTTokenContract.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) ParseApproval(log types.Log) (*AragonWrappedANTTokenContractApproval, error) {
	event := new(AragonWrappedANTTokenContractApproval)
	if err := _AragonWrappedANTTokenContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AragonWrappedANTTokenContractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the AragonWrappedANTTokenContract contract.
type AragonWrappedANTTokenContractDepositIterator struct {
	Event *AragonWrappedANTTokenContractDeposit // Event containing the contract specifics and raw log

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
func (it *AragonWrappedANTTokenContractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonWrappedANTTokenContractDeposit)
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
		it.Event = new(AragonWrappedANTTokenContractDeposit)
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
func (it *AragonWrappedANTTokenContractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonWrappedANTTokenContractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonWrappedANTTokenContractDeposit represents a Deposit event raised by the AragonWrappedANTTokenContract contract.
type AragonWrappedANTTokenContractDeposit struct {
	Entity common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed entity, uint256 amount)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) FilterDeposit(opts *bind.FilterOpts, entity []common.Address) (*AragonWrappedANTTokenContractDepositIterator, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _AragonWrappedANTTokenContract.contract.FilterLogs(opts, "Deposit", entityRule)
	if err != nil {
		return nil, err
	}
	return &AragonWrappedANTTokenContractDepositIterator{contract: _AragonWrappedANTTokenContract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed entity, uint256 amount)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *AragonWrappedANTTokenContractDeposit, entity []common.Address) (event.Subscription, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _AragonWrappedANTTokenContract.contract.WatchLogs(opts, "Deposit", entityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonWrappedANTTokenContractDeposit)
				if err := _AragonWrappedANTTokenContract.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed entity, uint256 amount)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) ParseDeposit(log types.Log) (*AragonWrappedANTTokenContractDeposit, error) {
	event := new(AragonWrappedANTTokenContractDeposit)
	if err := _AragonWrappedANTTokenContract.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AragonWrappedANTTokenContractRecoverToVaultIterator is returned from FilterRecoverToVault and is used to iterate over the raw logs and unpacked data for RecoverToVault events raised by the AragonWrappedANTTokenContract contract.
type AragonWrappedANTTokenContractRecoverToVaultIterator struct {
	Event *AragonWrappedANTTokenContractRecoverToVault // Event containing the contract specifics and raw log

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
func (it *AragonWrappedANTTokenContractRecoverToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonWrappedANTTokenContractRecoverToVault)
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
		it.Event = new(AragonWrappedANTTokenContractRecoverToVault)
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
func (it *AragonWrappedANTTokenContractRecoverToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonWrappedANTTokenContractRecoverToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonWrappedANTTokenContractRecoverToVault represents a RecoverToVault event raised by the AragonWrappedANTTokenContract contract.
type AragonWrappedANTTokenContractRecoverToVault struct {
	Vault  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverToVault is a free log retrieval operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) FilterRecoverToVault(opts *bind.FilterOpts, vault []common.Address, token []common.Address) (*AragonWrappedANTTokenContractRecoverToVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _AragonWrappedANTTokenContract.contract.FilterLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &AragonWrappedANTTokenContractRecoverToVaultIterator{contract: _AragonWrappedANTTokenContract.contract, event: "RecoverToVault", logs: logs, sub: sub}, nil
}

// WatchRecoverToVault is a free log subscription operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) WatchRecoverToVault(opts *bind.WatchOpts, sink chan<- *AragonWrappedANTTokenContractRecoverToVault, vault []common.Address, token []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _AragonWrappedANTTokenContract.contract.WatchLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonWrappedANTTokenContractRecoverToVault)
				if err := _AragonWrappedANTTokenContract.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
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

// ParseRecoverToVault is a log parse operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) ParseRecoverToVault(log types.Log) (*AragonWrappedANTTokenContractRecoverToVault, error) {
	event := new(AragonWrappedANTTokenContractRecoverToVault)
	if err := _AragonWrappedANTTokenContract.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AragonWrappedANTTokenContractScriptResultIterator is returned from FilterScriptResult and is used to iterate over the raw logs and unpacked data for ScriptResult events raised by the AragonWrappedANTTokenContract contract.
type AragonWrappedANTTokenContractScriptResultIterator struct {
	Event *AragonWrappedANTTokenContractScriptResult // Event containing the contract specifics and raw log

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
func (it *AragonWrappedANTTokenContractScriptResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonWrappedANTTokenContractScriptResult)
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
		it.Event = new(AragonWrappedANTTokenContractScriptResult)
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
func (it *AragonWrappedANTTokenContractScriptResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonWrappedANTTokenContractScriptResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonWrappedANTTokenContractScriptResult represents a ScriptResult event raised by the AragonWrappedANTTokenContract contract.
type AragonWrappedANTTokenContractScriptResult struct {
	Executor   common.Address
	Script     []byte
	Input      []byte
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScriptResult is a free log retrieval operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) FilterScriptResult(opts *bind.FilterOpts, executor []common.Address) (*AragonWrappedANTTokenContractScriptResultIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _AragonWrappedANTTokenContract.contract.FilterLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return &AragonWrappedANTTokenContractScriptResultIterator{contract: _AragonWrappedANTTokenContract.contract, event: "ScriptResult", logs: logs, sub: sub}, nil
}

// WatchScriptResult is a free log subscription operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) WatchScriptResult(opts *bind.WatchOpts, sink chan<- *AragonWrappedANTTokenContractScriptResult, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _AragonWrappedANTTokenContract.contract.WatchLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonWrappedANTTokenContractScriptResult)
				if err := _AragonWrappedANTTokenContract.contract.UnpackLog(event, "ScriptResult", log); err != nil {
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

// ParseScriptResult is a log parse operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) ParseScriptResult(log types.Log) (*AragonWrappedANTTokenContractScriptResult, error) {
	event := new(AragonWrappedANTTokenContractScriptResult)
	if err := _AragonWrappedANTTokenContract.contract.UnpackLog(event, "ScriptResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AragonWrappedANTTokenContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AragonWrappedANTTokenContract contract.
type AragonWrappedANTTokenContractTransferIterator struct {
	Event *AragonWrappedANTTokenContractTransfer // Event containing the contract specifics and raw log

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
func (it *AragonWrappedANTTokenContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonWrappedANTTokenContractTransfer)
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
		it.Event = new(AragonWrappedANTTokenContractTransfer)
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
func (it *AragonWrappedANTTokenContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonWrappedANTTokenContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonWrappedANTTokenContractTransfer represents a Transfer event raised by the AragonWrappedANTTokenContract contract.
type AragonWrappedANTTokenContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AragonWrappedANTTokenContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AragonWrappedANTTokenContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AragonWrappedANTTokenContractTransferIterator{contract: _AragonWrappedANTTokenContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AragonWrappedANTTokenContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AragonWrappedANTTokenContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonWrappedANTTokenContractTransfer)
				if err := _AragonWrappedANTTokenContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) ParseTransfer(log types.Log) (*AragonWrappedANTTokenContractTransfer, error) {
	event := new(AragonWrappedANTTokenContractTransfer)
	if err := _AragonWrappedANTTokenContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AragonWrappedANTTokenContractWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the AragonWrappedANTTokenContract contract.
type AragonWrappedANTTokenContractWithdrawalIterator struct {
	Event *AragonWrappedANTTokenContractWithdrawal // Event containing the contract specifics and raw log

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
func (it *AragonWrappedANTTokenContractWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonWrappedANTTokenContractWithdrawal)
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
		it.Event = new(AragonWrappedANTTokenContractWithdrawal)
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
func (it *AragonWrappedANTTokenContractWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonWrappedANTTokenContractWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonWrappedANTTokenContractWithdrawal represents a Withdrawal event raised by the AragonWrappedANTTokenContract contract.
type AragonWrappedANTTokenContractWithdrawal struct {
	Entity common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed entity, uint256 amount)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) FilterWithdrawal(opts *bind.FilterOpts, entity []common.Address) (*AragonWrappedANTTokenContractWithdrawalIterator, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _AragonWrappedANTTokenContract.contract.FilterLogs(opts, "Withdrawal", entityRule)
	if err != nil {
		return nil, err
	}
	return &AragonWrappedANTTokenContractWithdrawalIterator{contract: _AragonWrappedANTTokenContract.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed entity, uint256 amount)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *AragonWrappedANTTokenContractWithdrawal, entity []common.Address) (event.Subscription, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _AragonWrappedANTTokenContract.contract.WatchLogs(opts, "Withdrawal", entityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonWrappedANTTokenContractWithdrawal)
				if err := _AragonWrappedANTTokenContract.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed entity, uint256 amount)
func (_AragonWrappedANTTokenContract *AragonWrappedANTTokenContractFilterer) ParseWithdrawal(log types.Log) (*AragonWrappedANTTokenContractWithdrawal, error) {
	event := new(AragonWrappedANTTokenContractWithdrawal)
	if err := _AragonWrappedANTTokenContract.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
