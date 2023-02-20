// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC897Contract

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

// ERC897ContractMetaData contains all meta data concerning the ERC897Contract contract.
var ERC897ContractMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"proxyType\",\"outputs\":[{\"name\":\"proxyTypeId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isDepositable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_kernel\",\"type\":\"address\"},{\"name\":\"_appId\",\"type\":\"bytes32\"},{\"name\":\"_initializePayload\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ProxyDeposit\",\"type\":\"event\"}]",
}

// ERC897ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC897ContractMetaData.ABI instead.
var ERC897ContractABI = ERC897ContractMetaData.ABI

// ERC897Contract is an auto generated Go binding around an Ethereum contract.
type ERC897Contract struct {
	ERC897ContractCaller     // Read-only binding to the contract
	ERC897ContractTransactor // Write-only binding to the contract
	ERC897ContractFilterer   // Log filterer for contract events
}

// ERC897ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC897ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC897ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC897ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC897ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC897ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC897ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC897ContractSession struct {
	Contract     *ERC897Contract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC897ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC897ContractCallerSession struct {
	Contract *ERC897ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC897ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC897ContractTransactorSession struct {
	Contract     *ERC897ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC897ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC897ContractRaw struct {
	Contract *ERC897Contract // Generic contract binding to access the raw methods on
}

// ERC897ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC897ContractCallerRaw struct {
	Contract *ERC897ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ERC897ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC897ContractTransactorRaw struct {
	Contract *ERC897ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC897Contract creates a new instance of ERC897Contract, bound to a specific deployed contract.
func NewERC897Contract(address common.Address, backend bind.ContractBackend) (*ERC897Contract, error) {
	contract, err := bindERC897Contract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC897Contract{ERC897ContractCaller: ERC897ContractCaller{contract: contract}, ERC897ContractTransactor: ERC897ContractTransactor{contract: contract}, ERC897ContractFilterer: ERC897ContractFilterer{contract: contract}}, nil
}

// NewERC897ContractCaller creates a new read-only instance of ERC897Contract, bound to a specific deployed contract.
func NewERC897ContractCaller(address common.Address, caller bind.ContractCaller) (*ERC897ContractCaller, error) {
	contract, err := bindERC897Contract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC897ContractCaller{contract: contract}, nil
}

// NewERC897ContractTransactor creates a new write-only instance of ERC897Contract, bound to a specific deployed contract.
func NewERC897ContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC897ContractTransactor, error) {
	contract, err := bindERC897Contract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC897ContractTransactor{contract: contract}, nil
}

// NewERC897ContractFilterer creates a new log filterer instance of ERC897Contract, bound to a specific deployed contract.
func NewERC897ContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC897ContractFilterer, error) {
	contract, err := bindERC897Contract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC897ContractFilterer{contract: contract}, nil
}

// bindERC897Contract binds a generic wrapper to an already deployed contract.
func bindERC897Contract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC897ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC897Contract *ERC897ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC897Contract.Contract.ERC897ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC897Contract *ERC897ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC897Contract.Contract.ERC897ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC897Contract *ERC897ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC897Contract.Contract.ERC897ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC897Contract *ERC897ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC897Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC897Contract *ERC897ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC897Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC897Contract *ERC897ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC897Contract.Contract.contract.Transact(opts, method, params...)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_ERC897Contract *ERC897ContractCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC897Contract.contract.Call(opts, &out, "appId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_ERC897Contract *ERC897ContractSession) AppId() ([32]byte, error) {
	return _ERC897Contract.Contract.AppId(&_ERC897Contract.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_ERC897Contract *ERC897ContractCallerSession) AppId() ([32]byte, error) {
	return _ERC897Contract.Contract.AppId(&_ERC897Contract.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_ERC897Contract *ERC897ContractCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC897Contract.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_ERC897Contract *ERC897ContractSession) Implementation() (common.Address, error) {
	return _ERC897Contract.Contract.Implementation(&_ERC897Contract.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_ERC897Contract *ERC897ContractCallerSession) Implementation() (common.Address, error) {
	return _ERC897Contract.Contract.Implementation(&_ERC897Contract.CallOpts)
}

// IsDepositable is a free data retrieval call binding the contract method 0x48a0c8dd.
//
// Solidity: function isDepositable() view returns(bool)
func (_ERC897Contract *ERC897ContractCaller) IsDepositable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC897Contract.contract.Call(opts, &out, "isDepositable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositable is a free data retrieval call binding the contract method 0x48a0c8dd.
//
// Solidity: function isDepositable() view returns(bool)
func (_ERC897Contract *ERC897ContractSession) IsDepositable() (bool, error) {
	return _ERC897Contract.Contract.IsDepositable(&_ERC897Contract.CallOpts)
}

// IsDepositable is a free data retrieval call binding the contract method 0x48a0c8dd.
//
// Solidity: function isDepositable() view returns(bool)
func (_ERC897Contract *ERC897ContractCallerSession) IsDepositable() (bool, error) {
	return _ERC897Contract.Contract.IsDepositable(&_ERC897Contract.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_ERC897Contract *ERC897ContractCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC897Contract.contract.Call(opts, &out, "kernel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_ERC897Contract *ERC897ContractSession) Kernel() (common.Address, error) {
	return _ERC897Contract.Contract.Kernel(&_ERC897Contract.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_ERC897Contract *ERC897ContractCallerSession) Kernel() (common.Address, error) {
	return _ERC897Contract.Contract.Kernel(&_ERC897Contract.CallOpts)
}

// ProxyType is a free data retrieval call binding the contract method 0x4555d5c9.
//
// Solidity: function proxyType() pure returns(uint256 proxyTypeId)
func (_ERC897Contract *ERC897ContractCaller) ProxyType(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC897Contract.contract.Call(opts, &out, "proxyType")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProxyType is a free data retrieval call binding the contract method 0x4555d5c9.
//
// Solidity: function proxyType() pure returns(uint256 proxyTypeId)
func (_ERC897Contract *ERC897ContractSession) ProxyType() (*big.Int, error) {
	return _ERC897Contract.Contract.ProxyType(&_ERC897Contract.CallOpts)
}

// ProxyType is a free data retrieval call binding the contract method 0x4555d5c9.
//
// Solidity: function proxyType() pure returns(uint256 proxyTypeId)
func (_ERC897Contract *ERC897ContractCallerSession) ProxyType() (*big.Int, error) {
	return _ERC897Contract.Contract.ProxyType(&_ERC897Contract.CallOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ERC897Contract *ERC897ContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ERC897Contract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ERC897Contract *ERC897ContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ERC897Contract.Contract.Fallback(&_ERC897Contract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ERC897Contract *ERC897ContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ERC897Contract.Contract.Fallback(&_ERC897Contract.TransactOpts, calldata)
}

// ERC897ContractProxyDepositIterator is returned from FilterProxyDeposit and is used to iterate over the raw logs and unpacked data for ProxyDeposit events raised by the ERC897Contract contract.
type ERC897ContractProxyDepositIterator struct {
	Event *ERC897ContractProxyDeposit // Event containing the contract specifics and raw log

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
func (it *ERC897ContractProxyDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC897ContractProxyDeposit)
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
		it.Event = new(ERC897ContractProxyDeposit)
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
func (it *ERC897ContractProxyDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC897ContractProxyDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC897ContractProxyDeposit represents a ProxyDeposit event raised by the ERC897Contract contract.
type ERC897ContractProxyDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProxyDeposit is a free log retrieval operation binding the contract event 0x15eeaa57c7bd188c1388020bcadc2c436ec60d647d36ef5b9eb3c742217ddee1.
//
// Solidity: event ProxyDeposit(address sender, uint256 value)
func (_ERC897Contract *ERC897ContractFilterer) FilterProxyDeposit(opts *bind.FilterOpts) (*ERC897ContractProxyDepositIterator, error) {

	logs, sub, err := _ERC897Contract.contract.FilterLogs(opts, "ProxyDeposit")
	if err != nil {
		return nil, err
	}
	return &ERC897ContractProxyDepositIterator{contract: _ERC897Contract.contract, event: "ProxyDeposit", logs: logs, sub: sub}, nil
}

// WatchProxyDeposit is a free log subscription operation binding the contract event 0x15eeaa57c7bd188c1388020bcadc2c436ec60d647d36ef5b9eb3c742217ddee1.
//
// Solidity: event ProxyDeposit(address sender, uint256 value)
func (_ERC897Contract *ERC897ContractFilterer) WatchProxyDeposit(opts *bind.WatchOpts, sink chan<- *ERC897ContractProxyDeposit) (event.Subscription, error) {

	logs, sub, err := _ERC897Contract.contract.WatchLogs(opts, "ProxyDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC897ContractProxyDeposit)
				if err := _ERC897Contract.contract.UnpackLog(event, "ProxyDeposit", log); err != nil {
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

// ParseProxyDeposit is a log parse operation binding the contract event 0x15eeaa57c7bd188c1388020bcadc2c436ec60d647d36ef5b9eb3c742217ddee1.
//
// Solidity: event ProxyDeposit(address sender, uint256 value)
func (_ERC897Contract *ERC897ContractFilterer) ParseProxyDeposit(log types.Log) (*ERC897ContractProxyDeposit, error) {
	event := new(ERC897ContractProxyDeposit)
	if err := _ERC897Contract.contract.UnpackLog(event, "ProxyDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
