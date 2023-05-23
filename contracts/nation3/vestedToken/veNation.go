// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Nation3VestedTokenContract

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

// Nation3VestedTokenContractMetaData contains all meta data concerning the Nation3VestedTokenContract contract.
var Nation3VestedTokenContractMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"CommitOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"DepositsLockedChange\",\"inputs\":[{\"type\":\"bool\",\"name\":\"status\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"locktime\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"type\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"ts\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"ts\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Supply\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"prevSupply\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token_addr\"},{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"string\",\"name\":\"_version\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37597},{\"name\":\"set_depositsLocked\",\"outputs\":[],\"inputs\":[{\"type\":\"bool\",\"name\":\"new_status\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37624},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38527},{\"name\":\"commit_smart_wallet_checker\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36337},{\"name\":\"apply_smart_wallet_checker\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37125},{\"name\":\"get_last_user_slope\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2599},{\"name\":\"user_point_history__ts\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"},{\"type\":\"uint256\",\"name\":\"_idx\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1702},{\"name\":\"locked__end\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1623},{\"name\":\"checkpoint\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37052402},{\"name\":\"deposit_for\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74280011},{\"name\":\"create_lock\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"},{\"type\":\"uint256\",\"name\":\"_unlock_time\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74281615},{\"name\":\"increase_amount\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74280980},{\"name\":\"increase_unlock_time\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_unlock_time\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74281728},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37224478},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"uint256\",\"name\":\"_t\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"balanceOfAt\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"uint256\",\"name\":\"_block\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":514393},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"t\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"totalSupplyAt\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_block\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":812650},{\"name\":\"changeController\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_newController\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36937},{\"name\":\"token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1871},{\"name\":\"supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1901},{\"name\":\"locked\",\"outputs\":[{\"type\":\"int128\",\"name\":\"amount\"},{\"type\":\"uint256\",\"name\":\"end\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3389},{\"name\":\"epoch\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1961},{\"name\":\"point_history\",\"outputs\":[{\"type\":\"int128\",\"name\":\"bias\"},{\"type\":\"int128\",\"name\":\"slope\"},{\"type\":\"uint256\",\"name\":\"ts\"},{\"type\":\"uint256\",\"name\":\"blk\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":5580},{\"name\":\"user_point_history\",\"outputs\":[{\"type\":\"int128\",\"name\":\"bias\"},{\"type\":\"int128\",\"name\":\"slope\"},{\"type\":\"uint256\",\"name\":\"ts\"},{\"type\":\"uint256\",\"name\":\"blk\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"uint256\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":6109},{\"name\":\"user_point_epoch\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2205},{\"name\":\"slope_changes\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2196},{\"name\":\"controller\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2111},{\"name\":\"transfersEnabled\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2141},{\"name\":\"depositsLocked\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2171},{\"name\":\"name\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":8603},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":7656},{\"name\":\"version\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":7686},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2291},{\"name\":\"future_smart_wallet_checker\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2321},{\"name\":\"smart_wallet_checker\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2351},{\"name\":\"admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2381},{\"name\":\"future_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2411}]",
}

// Nation3VestedTokenContractABI is the input ABI used to generate the binding from.
// Deprecated: Use Nation3VestedTokenContractMetaData.ABI instead.
var Nation3VestedTokenContractABI = Nation3VestedTokenContractMetaData.ABI

// Nation3VestedTokenContract is an auto generated Go binding around an Ethereum contract.
type Nation3VestedTokenContract struct {
	Nation3VestedTokenContractCaller     // Read-only binding to the contract
	Nation3VestedTokenContractTransactor // Write-only binding to the contract
	Nation3VestedTokenContractFilterer   // Log filterer for contract events
}

// Nation3VestedTokenContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type Nation3VestedTokenContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nation3VestedTokenContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Nation3VestedTokenContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nation3VestedTokenContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Nation3VestedTokenContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nation3VestedTokenContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Nation3VestedTokenContractSession struct {
	Contract     *Nation3VestedTokenContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Nation3VestedTokenContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Nation3VestedTokenContractCallerSession struct {
	Contract *Nation3VestedTokenContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// Nation3VestedTokenContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Nation3VestedTokenContractTransactorSession struct {
	Contract     *Nation3VestedTokenContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// Nation3VestedTokenContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type Nation3VestedTokenContractRaw struct {
	Contract *Nation3VestedTokenContract // Generic contract binding to access the raw methods on
}

// Nation3VestedTokenContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Nation3VestedTokenContractCallerRaw struct {
	Contract *Nation3VestedTokenContractCaller // Generic read-only contract binding to access the raw methods on
}

// Nation3VestedTokenContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Nation3VestedTokenContractTransactorRaw struct {
	Contract *Nation3VestedTokenContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNation3VestedTokenContract creates a new instance of Nation3VestedTokenContract, bound to a specific deployed contract.
func NewNation3VestedTokenContract(address common.Address, backend bind.ContractBackend) (*Nation3VestedTokenContract, error) {
	contract, err := bindNation3VestedTokenContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokenContract{Nation3VestedTokenContractCaller: Nation3VestedTokenContractCaller{contract: contract}, Nation3VestedTokenContractTransactor: Nation3VestedTokenContractTransactor{contract: contract}, Nation3VestedTokenContractFilterer: Nation3VestedTokenContractFilterer{contract: contract}}, nil
}

// NewNation3VestedTokenContractCaller creates a new read-only instance of Nation3VestedTokenContract, bound to a specific deployed contract.
func NewNation3VestedTokenContractCaller(address common.Address, caller bind.ContractCaller) (*Nation3VestedTokenContractCaller, error) {
	contract, err := bindNation3VestedTokenContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokenContractCaller{contract: contract}, nil
}

// NewNation3VestedTokenContractTransactor creates a new write-only instance of Nation3VestedTokenContract, bound to a specific deployed contract.
func NewNation3VestedTokenContractTransactor(address common.Address, transactor bind.ContractTransactor) (*Nation3VestedTokenContractTransactor, error) {
	contract, err := bindNation3VestedTokenContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokenContractTransactor{contract: contract}, nil
}

// NewNation3VestedTokenContractFilterer creates a new log filterer instance of Nation3VestedTokenContract, bound to a specific deployed contract.
func NewNation3VestedTokenContractFilterer(address common.Address, filterer bind.ContractFilterer) (*Nation3VestedTokenContractFilterer, error) {
	contract, err := bindNation3VestedTokenContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokenContractFilterer{contract: contract}, nil
}

// bindNation3VestedTokenContract binds a generic wrapper to an already deployed contract.
func bindNation3VestedTokenContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Nation3VestedTokenContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nation3VestedTokenContract *Nation3VestedTokenContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nation3VestedTokenContract.Contract.Nation3VestedTokenContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nation3VestedTokenContract *Nation3VestedTokenContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.Nation3VestedTokenContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nation3VestedTokenContract *Nation3VestedTokenContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.Nation3VestedTokenContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nation3VestedTokenContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "admin")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) Admin() (common.Address, error) {
	return _Nation3VestedTokenContract.Contract.Admin(&_Nation3VestedTokenContract.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) Admin() (common.Address, error) {
	return _Nation3VestedTokenContract.Contract.Admin(&_Nation3VestedTokenContract.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "balanceOf", addr)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.BalanceOf(&_Nation3VestedTokenContract.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.BalanceOf(&_Nation3VestedTokenContract.CallOpts, addr)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) BalanceOf0(opts *bind.CallOpts, addr common.Address, _t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "balanceOf0", addr, _t)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) BalanceOf0(addr common.Address, _t *big.Int) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.BalanceOf0(&_Nation3VestedTokenContract.CallOpts, addr, _t)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) BalanceOf0(addr common.Address, _t *big.Int) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.BalanceOf0(&_Nation3VestedTokenContract.CallOpts, addr, _t)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) BalanceOfAt(opts *bind.CallOpts, addr common.Address, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "balanceOfAt", addr, _block)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) BalanceOfAt(addr common.Address, _block *big.Int) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.BalanceOfAt(&_Nation3VestedTokenContract.CallOpts, addr, _block)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) BalanceOfAt(addr common.Address, _block *big.Int) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.BalanceOfAt(&_Nation3VestedTokenContract.CallOpts, addr, _block)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "controller")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) Controller() (common.Address, error) {
	return _Nation3VestedTokenContract.Contract.Controller(&_Nation3VestedTokenContract.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) Controller() (common.Address, error) {
	return _Nation3VestedTokenContract.Contract.Controller(&_Nation3VestedTokenContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "decimals")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) Decimals() (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.Decimals(&_Nation3VestedTokenContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) Decimals() (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.Decimals(&_Nation3VestedTokenContract.CallOpts)
}

// DepositsLocked is a free data retrieval call binding the contract method 0xf60ab774.
//
// Solidity: function depositsLocked() view returns(bool)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) DepositsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "depositsLocked")
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// DepositsLocked is a free data retrieval call binding the contract method 0xf60ab774.
//
// Solidity: function depositsLocked() view returns(bool)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) DepositsLocked() (bool, error) {
	return _Nation3VestedTokenContract.Contract.DepositsLocked(&_Nation3VestedTokenContract.CallOpts)
}

// DepositsLocked is a free data retrieval call binding the contract method 0xf60ab774.
//
// Solidity: function depositsLocked() view returns(bool)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) DepositsLocked() (bool, error) {
	return _Nation3VestedTokenContract.Contract.DepositsLocked(&_Nation3VestedTokenContract.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "epoch")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) Epoch() (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.Epoch(&_Nation3VestedTokenContract.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) Epoch() (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.Epoch(&_Nation3VestedTokenContract.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "future_admin")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) FutureAdmin() (common.Address, error) {
	return _Nation3VestedTokenContract.Contract.FutureAdmin(&_Nation3VestedTokenContract.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) FutureAdmin() (common.Address, error) {
	return _Nation3VestedTokenContract.Contract.FutureAdmin(&_Nation3VestedTokenContract.CallOpts)
}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) FutureSmartWalletChecker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "future_smart_wallet_checker")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) FutureSmartWalletChecker() (common.Address, error) {
	return _Nation3VestedTokenContract.Contract.FutureSmartWalletChecker(&_Nation3VestedTokenContract.CallOpts)
}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) FutureSmartWalletChecker() (common.Address, error) {
	return _Nation3VestedTokenContract.Contract.FutureSmartWalletChecker(&_Nation3VestedTokenContract.CallOpts)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) GetLastUserSlope(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "get_last_user_slope", addr)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) GetLastUserSlope(addr common.Address) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.GetLastUserSlope(&_Nation3VestedTokenContract.CallOpts, addr)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) GetLastUserSlope(addr common.Address) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.GetLastUserSlope(&_Nation3VestedTokenContract.CallOpts, addr)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns(int128 amount, uint256 end)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) Locked(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error,
) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "locked", arg0)

	outstruct := new(struct {
		Amount *big.Int
		End    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns(int128 amount, uint256 end)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) Locked(arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error,
) {
	return _Nation3VestedTokenContract.Contract.Locked(&_Nation3VestedTokenContract.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns(int128 amount, uint256 end)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) Locked(arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error,
) {
	return _Nation3VestedTokenContract.Contract.Locked(&_Nation3VestedTokenContract.CallOpts, arg0)
}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) LockedEnd(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "locked__end", _addr)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) LockedEnd(_addr common.Address) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.LockedEnd(&_Nation3VestedTokenContract.CallOpts, _addr)
}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) LockedEnd(_addr common.Address) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.LockedEnd(&_Nation3VestedTokenContract.CallOpts, _addr)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "name")
	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) Name() (string, error) {
	return _Nation3VestedTokenContract.Contract.Name(&_Nation3VestedTokenContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) Name() (string, error) {
	return _Nation3VestedTokenContract.Contract.Name(&_Nation3VestedTokenContract.CallOpts)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) PointHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error,
) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "point_history", arg0)

	outstruct := new(struct {
		Bias  *big.Int
		Slope *big.Int
		Ts    *big.Int
		Blk   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bias = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ts = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Blk = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) PointHistory(arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error,
) {
	return _Nation3VestedTokenContract.Contract.PointHistory(&_Nation3VestedTokenContract.CallOpts, arg0)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) PointHistory(arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error,
) {
	return _Nation3VestedTokenContract.Contract.PointHistory(&_Nation3VestedTokenContract.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) SlopeChanges(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "slope_changes", arg0)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.SlopeChanges(&_Nation3VestedTokenContract.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.SlopeChanges(&_Nation3VestedTokenContract.CallOpts, arg0)
}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) SmartWalletChecker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "smart_wallet_checker")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) SmartWalletChecker() (common.Address, error) {
	return _Nation3VestedTokenContract.Contract.SmartWalletChecker(&_Nation3VestedTokenContract.CallOpts)
}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) SmartWalletChecker() (common.Address, error) {
	return _Nation3VestedTokenContract.Contract.SmartWalletChecker(&_Nation3VestedTokenContract.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "supply")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) Supply() (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.Supply(&_Nation3VestedTokenContract.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) Supply() (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.Supply(&_Nation3VestedTokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "symbol")
	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) Symbol() (string, error) {
	return _Nation3VestedTokenContract.Contract.Symbol(&_Nation3VestedTokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) Symbol() (string, error) {
	return _Nation3VestedTokenContract.Contract.Symbol(&_Nation3VestedTokenContract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "token")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) Token() (common.Address, error) {
	return _Nation3VestedTokenContract.Contract.Token(&_Nation3VestedTokenContract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) Token() (common.Address, error) {
	return _Nation3VestedTokenContract.Contract.Token(&_Nation3VestedTokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "totalSupply")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) TotalSupply() (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.TotalSupply(&_Nation3VestedTokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.TotalSupply(&_Nation3VestedTokenContract.CallOpts)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) TotalSupply0(opts *bind.CallOpts, t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "totalSupply0", t)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) TotalSupply0(t *big.Int) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.TotalSupply0(&_Nation3VestedTokenContract.CallOpts, t)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) TotalSupply0(t *big.Int) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.TotalSupply0(&_Nation3VestedTokenContract.CallOpts, t)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) TotalSupplyAt(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "totalSupplyAt", _block)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.TotalSupplyAt(&_Nation3VestedTokenContract.CallOpts, _block)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.TotalSupplyAt(&_Nation3VestedTokenContract.CallOpts, _block)
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) TransfersEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "transfersEnabled")
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) TransfersEnabled() (bool, error) {
	return _Nation3VestedTokenContract.Contract.TransfersEnabled(&_Nation3VestedTokenContract.CallOpts)
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) TransfersEnabled() (bool, error) {
	return _Nation3VestedTokenContract.Contract.TransfersEnabled(&_Nation3VestedTokenContract.CallOpts)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) UserPointEpoch(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "user_point_epoch", arg0)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) UserPointEpoch(arg0 common.Address) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.UserPointEpoch(&_Nation3VestedTokenContract.CallOpts, arg0)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) UserPointEpoch(arg0 common.Address) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.UserPointEpoch(&_Nation3VestedTokenContract.CallOpts, arg0)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) UserPointHistory(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error,
) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "user_point_history", arg0, arg1)

	outstruct := new(struct {
		Bias  *big.Int
		Slope *big.Int
		Ts    *big.Int
		Blk   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bias = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ts = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Blk = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err
}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) UserPointHistory(arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error,
) {
	return _Nation3VestedTokenContract.Contract.UserPointHistory(&_Nation3VestedTokenContract.CallOpts, arg0, arg1)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) UserPointHistory(arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error,
) {
	return _Nation3VestedTokenContract.Contract.UserPointHistory(&_Nation3VestedTokenContract.CallOpts, arg0, arg1)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) UserPointHistoryTs(opts *bind.CallOpts, _addr common.Address, _idx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "user_point_history__ts", _addr, _idx)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) UserPointHistoryTs(_addr common.Address, _idx *big.Int) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.UserPointHistoryTs(&_Nation3VestedTokenContract.CallOpts, _addr, _idx)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) UserPointHistoryTs(_addr common.Address, _idx *big.Int) (*big.Int, error) {
	return _Nation3VestedTokenContract.Contract.UserPointHistoryTs(&_Nation3VestedTokenContract.CallOpts, _addr, _idx)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nation3VestedTokenContract.contract.Call(opts, &out, "version")
	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) Version() (string, error) {
	return _Nation3VestedTokenContract.Contract.Version(&_Nation3VestedTokenContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractCallerSession) Version() (string, error) {
	return _Nation3VestedTokenContract.Contract.Version(&_Nation3VestedTokenContract.CallOpts)
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactor) ApplySmartWalletChecker(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.contract.Transact(opts, "apply_smart_wallet_checker")
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) ApplySmartWalletChecker() (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.ApplySmartWalletChecker(&_Nation3VestedTokenContract.TransactOpts)
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactorSession) ApplySmartWalletChecker() (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.ApplySmartWalletChecker(&_Nation3VestedTokenContract.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.ApplyTransferOwnership(&_Nation3VestedTokenContract.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.ApplyTransferOwnership(&_Nation3VestedTokenContract.TransactOpts)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactor) ChangeController(opts *bind.TransactOpts, _newController common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.contract.Transact(opts, "changeController", _newController)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) ChangeController(_newController common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.ChangeController(&_Nation3VestedTokenContract.TransactOpts, _newController)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactorSession) ChangeController(_newController common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.ChangeController(&_Nation3VestedTokenContract.TransactOpts, _newController)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactor) Checkpoint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.contract.Transact(opts, "checkpoint")
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) Checkpoint() (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.Checkpoint(&_Nation3VestedTokenContract.TransactOpts)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactorSession) Checkpoint() (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.Checkpoint(&_Nation3VestedTokenContract.TransactOpts)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactor) CommitSmartWalletChecker(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.contract.Transact(opts, "commit_smart_wallet_checker", addr)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) CommitSmartWalletChecker(addr common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.CommitSmartWalletChecker(&_Nation3VestedTokenContract.TransactOpts, addr)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactorSession) CommitSmartWalletChecker(addr common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.CommitSmartWalletChecker(&_Nation3VestedTokenContract.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactor) CommitTransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.contract.Transact(opts, "commit_transfer_ownership", addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.CommitTransferOwnership(&_Nation3VestedTokenContract.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactorSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.CommitTransferOwnership(&_Nation3VestedTokenContract.TransactOpts, addr)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactor) CreateLock(opts *bind.TransactOpts, _value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.contract.Transact(opts, "create_lock", _value, _unlock_time)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) CreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.CreateLock(&_Nation3VestedTokenContract.TransactOpts, _value, _unlock_time)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactorSession) CreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.CreateLock(&_Nation3VestedTokenContract.TransactOpts, _value, _unlock_time)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactor) DepositFor(opts *bind.TransactOpts, _addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.contract.Transact(opts, "deposit_for", _addr, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) DepositFor(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.DepositFor(&_Nation3VestedTokenContract.TransactOpts, _addr, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactorSession) DepositFor(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.DepositFor(&_Nation3VestedTokenContract.TransactOpts, _addr, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactor) IncreaseAmount(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.contract.Transact(opts, "increase_amount", _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) IncreaseAmount(_value *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.IncreaseAmount(&_Nation3VestedTokenContract.TransactOpts, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactorSession) IncreaseAmount(_value *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.IncreaseAmount(&_Nation3VestedTokenContract.TransactOpts, _value)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactor) IncreaseUnlockTime(opts *bind.TransactOpts, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.contract.Transact(opts, "increase_unlock_time", _unlock_time)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) IncreaseUnlockTime(_unlock_time *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.IncreaseUnlockTime(&_Nation3VestedTokenContract.TransactOpts, _unlock_time)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactorSession) IncreaseUnlockTime(_unlock_time *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.IncreaseUnlockTime(&_Nation3VestedTokenContract.TransactOpts, _unlock_time)
}

// SetDepositsLocked is a paid mutator transaction binding the contract method 0x4e782813.
//
// Solidity: function set_depositsLocked(bool new_status) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactor) SetDepositsLocked(opts *bind.TransactOpts, new_status bool) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.contract.Transact(opts, "set_depositsLocked", new_status)
}

// SetDepositsLocked is a paid mutator transaction binding the contract method 0x4e782813.
//
// Solidity: function set_depositsLocked(bool new_status) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) SetDepositsLocked(new_status bool) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.SetDepositsLocked(&_Nation3VestedTokenContract.TransactOpts, new_status)
}

// SetDepositsLocked is a paid mutator transaction binding the contract method 0x4e782813.
//
// Solidity: function set_depositsLocked(bool new_status) returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactorSession) SetDepositsLocked(new_status bool) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.SetDepositsLocked(&_Nation3VestedTokenContract.TransactOpts, new_status)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3VestedTokenContract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractSession) Withdraw() (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.Withdraw(&_Nation3VestedTokenContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Nation3VestedTokenContract *Nation3VestedTokenContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Nation3VestedTokenContract.Contract.Withdraw(&_Nation3VestedTokenContract.TransactOpts)
}

// Nation3VestedTokenContractApplyOwnershipIterator is returned from FilterApplyOwnership and is used to iterate over the raw logs and unpacked data for ApplyOwnership events raised by the Nation3VestedTokenContract contract.
type Nation3VestedTokenContractApplyOwnershipIterator struct {
	Event *Nation3VestedTokenContractApplyOwnership // Event containing the contract specifics and raw log

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
func (it *Nation3VestedTokenContractApplyOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3VestedTokenContractApplyOwnership)
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
		it.Event = new(Nation3VestedTokenContractApplyOwnership)
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
func (it *Nation3VestedTokenContractApplyOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3VestedTokenContractApplyOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3VestedTokenContractApplyOwnership represents a ApplyOwnership event raised by the Nation3VestedTokenContract contract.
type Nation3VestedTokenContractApplyOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterApplyOwnership is a free log retrieval operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) FilterApplyOwnership(opts *bind.FilterOpts) (*Nation3VestedTokenContractApplyOwnershipIterator, error) {
	logs, sub, err := _Nation3VestedTokenContract.contract.FilterLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokenContractApplyOwnershipIterator{contract: _Nation3VestedTokenContract.contract, event: "ApplyOwnership", logs: logs, sub: sub}, nil
}

// WatchApplyOwnership is a free log subscription operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) WatchApplyOwnership(opts *bind.WatchOpts, sink chan<- *Nation3VestedTokenContractApplyOwnership) (event.Subscription, error) {
	logs, sub, err := _Nation3VestedTokenContract.contract.WatchLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3VestedTokenContractApplyOwnership)
				if err := _Nation3VestedTokenContract.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
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

// ParseApplyOwnership is a log parse operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) ParseApplyOwnership(log types.Log) (*Nation3VestedTokenContractApplyOwnership, error) {
	event := new(Nation3VestedTokenContractApplyOwnership)
	if err := _Nation3VestedTokenContract.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3VestedTokenContractCommitOwnershipIterator is returned from FilterCommitOwnership and is used to iterate over the raw logs and unpacked data for CommitOwnership events raised by the Nation3VestedTokenContract contract.
type Nation3VestedTokenContractCommitOwnershipIterator struct {
	Event *Nation3VestedTokenContractCommitOwnership // Event containing the contract specifics and raw log

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
func (it *Nation3VestedTokenContractCommitOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3VestedTokenContractCommitOwnership)
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
		it.Event = new(Nation3VestedTokenContractCommitOwnership)
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
func (it *Nation3VestedTokenContractCommitOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3VestedTokenContractCommitOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3VestedTokenContractCommitOwnership represents a CommitOwnership event raised by the Nation3VestedTokenContract contract.
type Nation3VestedTokenContractCommitOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCommitOwnership is a free log retrieval operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) FilterCommitOwnership(opts *bind.FilterOpts) (*Nation3VestedTokenContractCommitOwnershipIterator, error) {
	logs, sub, err := _Nation3VestedTokenContract.contract.FilterLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokenContractCommitOwnershipIterator{contract: _Nation3VestedTokenContract.contract, event: "CommitOwnership", logs: logs, sub: sub}, nil
}

// WatchCommitOwnership is a free log subscription operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) WatchCommitOwnership(opts *bind.WatchOpts, sink chan<- *Nation3VestedTokenContractCommitOwnership) (event.Subscription, error) {
	logs, sub, err := _Nation3VestedTokenContract.contract.WatchLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3VestedTokenContractCommitOwnership)
				if err := _Nation3VestedTokenContract.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
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

// ParseCommitOwnership is a log parse operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) ParseCommitOwnership(log types.Log) (*Nation3VestedTokenContractCommitOwnership, error) {
	event := new(Nation3VestedTokenContractCommitOwnership)
	if err := _Nation3VestedTokenContract.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3VestedTokenContractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Nation3VestedTokenContract contract.
type Nation3VestedTokenContractDepositIterator struct {
	Event *Nation3VestedTokenContractDeposit // Event containing the contract specifics and raw log

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
func (it *Nation3VestedTokenContractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3VestedTokenContractDeposit)
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
		it.Event = new(Nation3VestedTokenContractDeposit)
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
func (it *Nation3VestedTokenContractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3VestedTokenContractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3VestedTokenContractDeposit represents a Deposit event raised by the Nation3VestedTokenContract contract.
type Nation3VestedTokenContractDeposit struct {
	Provider common.Address
	Value    *big.Int
	Locktime *big.Int
	Arg3     *big.Int
	Ts       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) FilterDeposit(opts *bind.FilterOpts, provider []common.Address, locktime []*big.Int) (*Nation3VestedTokenContractDepositIterator, error) {
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _Nation3VestedTokenContract.contract.FilterLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokenContractDepositIterator{contract: _Nation3VestedTokenContract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *Nation3VestedTokenContractDeposit, provider []common.Address, locktime []*big.Int) (event.Subscription, error) {
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _Nation3VestedTokenContract.contract.WatchLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3VestedTokenContractDeposit)
				if err := _Nation3VestedTokenContract.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) ParseDeposit(log types.Log) (*Nation3VestedTokenContractDeposit, error) {
	event := new(Nation3VestedTokenContractDeposit)
	if err := _Nation3VestedTokenContract.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3VestedTokenContractDepositsLockedChangeIterator is returned from FilterDepositsLockedChange and is used to iterate over the raw logs and unpacked data for DepositsLockedChange events raised by the Nation3VestedTokenContract contract.
type Nation3VestedTokenContractDepositsLockedChangeIterator struct {
	Event *Nation3VestedTokenContractDepositsLockedChange // Event containing the contract specifics and raw log

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
func (it *Nation3VestedTokenContractDepositsLockedChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3VestedTokenContractDepositsLockedChange)
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
		it.Event = new(Nation3VestedTokenContractDepositsLockedChange)
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
func (it *Nation3VestedTokenContractDepositsLockedChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3VestedTokenContractDepositsLockedChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3VestedTokenContractDepositsLockedChange represents a DepositsLockedChange event raised by the Nation3VestedTokenContract contract.
type Nation3VestedTokenContractDepositsLockedChange struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositsLockedChange is a free log retrieval operation binding the contract event 0x3a6f72b4ed02517aec1d580529f912e860631ecb6fec4712d400294932aaffba.
//
// Solidity: event DepositsLockedChange(bool status)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) FilterDepositsLockedChange(opts *bind.FilterOpts) (*Nation3VestedTokenContractDepositsLockedChangeIterator, error) {
	logs, sub, err := _Nation3VestedTokenContract.contract.FilterLogs(opts, "DepositsLockedChange")
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokenContractDepositsLockedChangeIterator{contract: _Nation3VestedTokenContract.contract, event: "DepositsLockedChange", logs: logs, sub: sub}, nil
}

// WatchDepositsLockedChange is a free log subscription operation binding the contract event 0x3a6f72b4ed02517aec1d580529f912e860631ecb6fec4712d400294932aaffba.
//
// Solidity: event DepositsLockedChange(bool status)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) WatchDepositsLockedChange(opts *bind.WatchOpts, sink chan<- *Nation3VestedTokenContractDepositsLockedChange) (event.Subscription, error) {
	logs, sub, err := _Nation3VestedTokenContract.contract.WatchLogs(opts, "DepositsLockedChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3VestedTokenContractDepositsLockedChange)
				if err := _Nation3VestedTokenContract.contract.UnpackLog(event, "DepositsLockedChange", log); err != nil {
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

// ParseDepositsLockedChange is a log parse operation binding the contract event 0x3a6f72b4ed02517aec1d580529f912e860631ecb6fec4712d400294932aaffba.
//
// Solidity: event DepositsLockedChange(bool status)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) ParseDepositsLockedChange(log types.Log) (*Nation3VestedTokenContractDepositsLockedChange, error) {
	event := new(Nation3VestedTokenContractDepositsLockedChange)
	if err := _Nation3VestedTokenContract.contract.UnpackLog(event, "DepositsLockedChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3VestedTokenContractSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the Nation3VestedTokenContract contract.
type Nation3VestedTokenContractSupplyIterator struct {
	Event *Nation3VestedTokenContractSupply // Event containing the contract specifics and raw log

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
func (it *Nation3VestedTokenContractSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3VestedTokenContractSupply)
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
		it.Event = new(Nation3VestedTokenContractSupply)
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
func (it *Nation3VestedTokenContractSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3VestedTokenContractSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3VestedTokenContractSupply represents a Supply event raised by the Nation3VestedTokenContract contract.
type Nation3VestedTokenContractSupply struct {
	PrevSupply *big.Int
	Supply     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) FilterSupply(opts *bind.FilterOpts) (*Nation3VestedTokenContractSupplyIterator, error) {
	logs, sub, err := _Nation3VestedTokenContract.contract.FilterLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokenContractSupplyIterator{contract: _Nation3VestedTokenContract.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *Nation3VestedTokenContractSupply) (event.Subscription, error) {
	logs, sub, err := _Nation3VestedTokenContract.contract.WatchLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3VestedTokenContractSupply)
				if err := _Nation3VestedTokenContract.contract.UnpackLog(event, "Supply", log); err != nil {
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

// ParseSupply is a log parse operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) ParseSupply(log types.Log) (*Nation3VestedTokenContractSupply, error) {
	event := new(Nation3VestedTokenContractSupply)
	if err := _Nation3VestedTokenContract.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3VestedTokenContractWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Nation3VestedTokenContract contract.
type Nation3VestedTokenContractWithdrawIterator struct {
	Event *Nation3VestedTokenContractWithdraw // Event containing the contract specifics and raw log

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
func (it *Nation3VestedTokenContractWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3VestedTokenContractWithdraw)
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
		it.Event = new(Nation3VestedTokenContractWithdraw)
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
func (it *Nation3VestedTokenContractWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3VestedTokenContractWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3VestedTokenContractWithdraw represents a Withdraw event raised by the Nation3VestedTokenContract contract.
type Nation3VestedTokenContractWithdraw struct {
	Provider common.Address
	Value    *big.Int
	Ts       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) FilterWithdraw(opts *bind.FilterOpts, provider []common.Address) (*Nation3VestedTokenContractWithdrawIterator, error) {
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Nation3VestedTokenContract.contract.FilterLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokenContractWithdrawIterator{contract: _Nation3VestedTokenContract.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Nation3VestedTokenContractWithdraw, provider []common.Address) (event.Subscription, error) {
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Nation3VestedTokenContract.contract.WatchLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3VestedTokenContractWithdraw)
				if err := _Nation3VestedTokenContract.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_Nation3VestedTokenContract *Nation3VestedTokenContractFilterer) ParseWithdraw(log types.Log) (*Nation3VestedTokenContractWithdraw, error) {
	event := new(Nation3VestedTokenContractWithdraw)
	if err := _Nation3VestedTokenContract.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
