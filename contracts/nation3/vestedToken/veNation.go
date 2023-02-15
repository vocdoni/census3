// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nation3VestedTokencontracts

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

// Nation3VestedTokencontractsMetaData contains all meta data concerning the Nation3VestedTokencontracts contract.
var Nation3VestedTokencontractsMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"CommitOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"DepositsLockedChange\",\"inputs\":[{\"type\":\"bool\",\"name\":\"status\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"locktime\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"type\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"ts\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"ts\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Supply\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"prevSupply\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token_addr\"},{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"string\",\"name\":\"_version\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37597},{\"name\":\"set_depositsLocked\",\"outputs\":[],\"inputs\":[{\"type\":\"bool\",\"name\":\"new_status\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37624},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38527},{\"name\":\"commit_smart_wallet_checker\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36337},{\"name\":\"apply_smart_wallet_checker\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37125},{\"name\":\"get_last_user_slope\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2599},{\"name\":\"user_point_history__ts\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"},{\"type\":\"uint256\",\"name\":\"_idx\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1702},{\"name\":\"locked__end\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1623},{\"name\":\"checkpoint\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37052402},{\"name\":\"deposit_for\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74280011},{\"name\":\"create_lock\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"},{\"type\":\"uint256\",\"name\":\"_unlock_time\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74281615},{\"name\":\"increase_amount\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74280980},{\"name\":\"increase_unlock_time\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_unlock_time\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74281728},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37224478},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"uint256\",\"name\":\"_t\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"balanceOfAt\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"uint256\",\"name\":\"_block\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":514393},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"t\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"totalSupplyAt\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_block\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":812650},{\"name\":\"changeController\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_newController\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36937},{\"name\":\"token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1871},{\"name\":\"supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1901},{\"name\":\"locked\",\"outputs\":[{\"type\":\"int128\",\"name\":\"amount\"},{\"type\":\"uint256\",\"name\":\"end\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3389},{\"name\":\"epoch\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1961},{\"name\":\"point_history\",\"outputs\":[{\"type\":\"int128\",\"name\":\"bias\"},{\"type\":\"int128\",\"name\":\"slope\"},{\"type\":\"uint256\",\"name\":\"ts\"},{\"type\":\"uint256\",\"name\":\"blk\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":5580},{\"name\":\"user_point_history\",\"outputs\":[{\"type\":\"int128\",\"name\":\"bias\"},{\"type\":\"int128\",\"name\":\"slope\"},{\"type\":\"uint256\",\"name\":\"ts\"},{\"type\":\"uint256\",\"name\":\"blk\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"uint256\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":6109},{\"name\":\"user_point_epoch\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2205},{\"name\":\"slope_changes\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2196},{\"name\":\"controller\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2111},{\"name\":\"transfersEnabled\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2141},{\"name\":\"depositsLocked\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2171},{\"name\":\"name\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":8603},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":7656},{\"name\":\"version\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":7686},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2291},{\"name\":\"future_smart_wallet_checker\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2321},{\"name\":\"smart_wallet_checker\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2351},{\"name\":\"admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2381},{\"name\":\"future_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2411}]",
}

// Nation3VestedTokencontractsABI is the input ABI used to generate the binding from.
// Deprecated: Use Nation3VestedTokencontractsMetaData.ABI instead.
var Nation3VestedTokencontractsABI = Nation3VestedTokencontractsMetaData.ABI

// Nation3VestedTokencontracts is an auto generated Go binding around an Ethereum contract.
type Nation3VestedTokencontracts struct {
	Nation3VestedTokencontractsCaller     // Read-only binding to the contract
	Nation3VestedTokencontractsTransactor // Write-only binding to the contract
	Nation3VestedTokencontractsFilterer   // Log filterer for contract events
}

// Nation3VestedTokencontractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type Nation3VestedTokencontractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nation3VestedTokencontractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Nation3VestedTokencontractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nation3VestedTokencontractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Nation3VestedTokencontractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nation3VestedTokencontractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Nation3VestedTokencontractsSession struct {
	Contract     *Nation3VestedTokencontracts // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Nation3VestedTokencontractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Nation3VestedTokencontractsCallerSession struct {
	Contract *Nation3VestedTokencontractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Nation3VestedTokencontractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Nation3VestedTokencontractsTransactorSession struct {
	Contract     *Nation3VestedTokencontractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Nation3VestedTokencontractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type Nation3VestedTokencontractsRaw struct {
	Contract *Nation3VestedTokencontracts // Generic contract binding to access the raw methods on
}

// Nation3VestedTokencontractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Nation3VestedTokencontractsCallerRaw struct {
	Contract *Nation3VestedTokencontractsCaller // Generic read-only contract binding to access the raw methods on
}

// Nation3VestedTokencontractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Nation3VestedTokencontractsTransactorRaw struct {
	Contract *Nation3VestedTokencontractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNation3VestedTokencontracts creates a new instance of Nation3VestedTokencontracts, bound to a specific deployed contract.
func NewNation3VestedTokencontracts(address common.Address, backend bind.ContractBackend) (*Nation3VestedTokencontracts, error) {
	contract, err := bindNation3VestedTokencontracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokencontracts{Nation3VestedTokencontractsCaller: Nation3VestedTokencontractsCaller{contract: contract}, Nation3VestedTokencontractsTransactor: Nation3VestedTokencontractsTransactor{contract: contract}, Nation3VestedTokencontractsFilterer: Nation3VestedTokencontractsFilterer{contract: contract}}, nil
}

// NewNation3VestedTokencontractsCaller creates a new read-only instance of Nation3VestedTokencontracts, bound to a specific deployed contract.
func NewNation3VestedTokencontractsCaller(address common.Address, caller bind.ContractCaller) (*Nation3VestedTokencontractsCaller, error) {
	contract, err := bindNation3VestedTokencontracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokencontractsCaller{contract: contract}, nil
}

// NewNation3VestedTokencontractsTransactor creates a new write-only instance of Nation3VestedTokencontracts, bound to a specific deployed contract.
func NewNation3VestedTokencontractsTransactor(address common.Address, transactor bind.ContractTransactor) (*Nation3VestedTokencontractsTransactor, error) {
	contract, err := bindNation3VestedTokencontracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokencontractsTransactor{contract: contract}, nil
}

// NewNation3VestedTokencontractsFilterer creates a new log filterer instance of Nation3VestedTokencontracts, bound to a specific deployed contract.
func NewNation3VestedTokencontractsFilterer(address common.Address, filterer bind.ContractFilterer) (*Nation3VestedTokencontractsFilterer, error) {
	contract, err := bindNation3VestedTokencontracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokencontractsFilterer{contract: contract}, nil
}

// bindNation3VestedTokencontracts binds a generic wrapper to an already deployed contract.
func bindNation3VestedTokencontracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Nation3VestedTokencontractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nation3VestedTokencontracts.Contract.Nation3VestedTokencontractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.Nation3VestedTokencontractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.Nation3VestedTokencontractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nation3VestedTokencontracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) Admin() (common.Address, error) {
	return _Nation3VestedTokencontracts.Contract.Admin(&_Nation3VestedTokencontracts.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) Admin() (common.Address, error) {
	return _Nation3VestedTokencontracts.Contract.Admin(&_Nation3VestedTokencontracts.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "balanceOf", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.BalanceOf(&_Nation3VestedTokencontracts.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.BalanceOf(&_Nation3VestedTokencontracts.CallOpts, addr)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) BalanceOf0(opts *bind.CallOpts, addr common.Address, _t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "balanceOf0", addr, _t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) BalanceOf0(addr common.Address, _t *big.Int) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.BalanceOf0(&_Nation3VestedTokencontracts.CallOpts, addr, _t)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) BalanceOf0(addr common.Address, _t *big.Int) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.BalanceOf0(&_Nation3VestedTokencontracts.CallOpts, addr, _t)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) BalanceOfAt(opts *bind.CallOpts, addr common.Address, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "balanceOfAt", addr, _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) BalanceOfAt(addr common.Address, _block *big.Int) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.BalanceOfAt(&_Nation3VestedTokencontracts.CallOpts, addr, _block)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) BalanceOfAt(addr common.Address, _block *big.Int) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.BalanceOfAt(&_Nation3VestedTokencontracts.CallOpts, addr, _block)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) Controller() (common.Address, error) {
	return _Nation3VestedTokencontracts.Contract.Controller(&_Nation3VestedTokencontracts.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) Controller() (common.Address, error) {
	return _Nation3VestedTokencontracts.Contract.Controller(&_Nation3VestedTokencontracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) Decimals() (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.Decimals(&_Nation3VestedTokencontracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) Decimals() (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.Decimals(&_Nation3VestedTokencontracts.CallOpts)
}

// DepositsLocked is a free data retrieval call binding the contract method 0xf60ab774.
//
// Solidity: function depositsLocked() view returns(bool)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) DepositsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "depositsLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositsLocked is a free data retrieval call binding the contract method 0xf60ab774.
//
// Solidity: function depositsLocked() view returns(bool)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) DepositsLocked() (bool, error) {
	return _Nation3VestedTokencontracts.Contract.DepositsLocked(&_Nation3VestedTokencontracts.CallOpts)
}

// DepositsLocked is a free data retrieval call binding the contract method 0xf60ab774.
//
// Solidity: function depositsLocked() view returns(bool)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) DepositsLocked() (bool, error) {
	return _Nation3VestedTokencontracts.Contract.DepositsLocked(&_Nation3VestedTokencontracts.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) Epoch() (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.Epoch(&_Nation3VestedTokencontracts.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) Epoch() (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.Epoch(&_Nation3VestedTokencontracts.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) FutureAdmin() (common.Address, error) {
	return _Nation3VestedTokencontracts.Contract.FutureAdmin(&_Nation3VestedTokencontracts.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) FutureAdmin() (common.Address, error) {
	return _Nation3VestedTokencontracts.Contract.FutureAdmin(&_Nation3VestedTokencontracts.CallOpts)
}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) FutureSmartWalletChecker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "future_smart_wallet_checker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) FutureSmartWalletChecker() (common.Address, error) {
	return _Nation3VestedTokencontracts.Contract.FutureSmartWalletChecker(&_Nation3VestedTokencontracts.CallOpts)
}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) FutureSmartWalletChecker() (common.Address, error) {
	return _Nation3VestedTokencontracts.Contract.FutureSmartWalletChecker(&_Nation3VestedTokencontracts.CallOpts)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) GetLastUserSlope(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "get_last_user_slope", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) GetLastUserSlope(addr common.Address) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.GetLastUserSlope(&_Nation3VestedTokencontracts.CallOpts, addr)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) GetLastUserSlope(addr common.Address) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.GetLastUserSlope(&_Nation3VestedTokencontracts.CallOpts, addr)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns(int128 amount, uint256 end)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) Locked(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "locked", arg0)

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
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) Locked(arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	return _Nation3VestedTokencontracts.Contract.Locked(&_Nation3VestedTokencontracts.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns(int128 amount, uint256 end)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) Locked(arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	return _Nation3VestedTokencontracts.Contract.Locked(&_Nation3VestedTokencontracts.CallOpts, arg0)
}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) LockedEnd(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "locked__end", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) LockedEnd(_addr common.Address) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.LockedEnd(&_Nation3VestedTokencontracts.CallOpts, _addr)
}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) LockedEnd(_addr common.Address) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.LockedEnd(&_Nation3VestedTokencontracts.CallOpts, _addr)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) Name() (string, error) {
	return _Nation3VestedTokencontracts.Contract.Name(&_Nation3VestedTokencontracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) Name() (string, error) {
	return _Nation3VestedTokencontracts.Contract.Name(&_Nation3VestedTokencontracts.CallOpts)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) PointHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "point_history", arg0)

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
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) PointHistory(arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Nation3VestedTokencontracts.Contract.PointHistory(&_Nation3VestedTokencontracts.CallOpts, arg0)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) PointHistory(arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Nation3VestedTokencontracts.Contract.PointHistory(&_Nation3VestedTokencontracts.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) SlopeChanges(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "slope_changes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.SlopeChanges(&_Nation3VestedTokencontracts.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.SlopeChanges(&_Nation3VestedTokencontracts.CallOpts, arg0)
}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) SmartWalletChecker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "smart_wallet_checker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) SmartWalletChecker() (common.Address, error) {
	return _Nation3VestedTokencontracts.Contract.SmartWalletChecker(&_Nation3VestedTokencontracts.CallOpts)
}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) SmartWalletChecker() (common.Address, error) {
	return _Nation3VestedTokencontracts.Contract.SmartWalletChecker(&_Nation3VestedTokencontracts.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) Supply() (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.Supply(&_Nation3VestedTokencontracts.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) Supply() (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.Supply(&_Nation3VestedTokencontracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) Symbol() (string, error) {
	return _Nation3VestedTokencontracts.Contract.Symbol(&_Nation3VestedTokencontracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) Symbol() (string, error) {
	return _Nation3VestedTokencontracts.Contract.Symbol(&_Nation3VestedTokencontracts.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) Token() (common.Address, error) {
	return _Nation3VestedTokencontracts.Contract.Token(&_Nation3VestedTokencontracts.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) Token() (common.Address, error) {
	return _Nation3VestedTokencontracts.Contract.Token(&_Nation3VestedTokencontracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) TotalSupply() (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.TotalSupply(&_Nation3VestedTokencontracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) TotalSupply() (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.TotalSupply(&_Nation3VestedTokencontracts.CallOpts)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) TotalSupply0(opts *bind.CallOpts, t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "totalSupply0", t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) TotalSupply0(t *big.Int) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.TotalSupply0(&_Nation3VestedTokencontracts.CallOpts, t)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) TotalSupply0(t *big.Int) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.TotalSupply0(&_Nation3VestedTokencontracts.CallOpts, t)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) TotalSupplyAt(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "totalSupplyAt", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.TotalSupplyAt(&_Nation3VestedTokencontracts.CallOpts, _block)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.TotalSupplyAt(&_Nation3VestedTokencontracts.CallOpts, _block)
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) TransfersEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "transfersEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) TransfersEnabled() (bool, error) {
	return _Nation3VestedTokencontracts.Contract.TransfersEnabled(&_Nation3VestedTokencontracts.CallOpts)
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) TransfersEnabled() (bool, error) {
	return _Nation3VestedTokencontracts.Contract.TransfersEnabled(&_Nation3VestedTokencontracts.CallOpts)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) UserPointEpoch(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "user_point_epoch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) UserPointEpoch(arg0 common.Address) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.UserPointEpoch(&_Nation3VestedTokencontracts.CallOpts, arg0)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) UserPointEpoch(arg0 common.Address) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.UserPointEpoch(&_Nation3VestedTokencontracts.CallOpts, arg0)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) UserPointHistory(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "user_point_history", arg0, arg1)

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
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) UserPointHistory(arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Nation3VestedTokencontracts.Contract.UserPointHistory(&_Nation3VestedTokencontracts.CallOpts, arg0, arg1)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) UserPointHistory(arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Nation3VestedTokencontracts.Contract.UserPointHistory(&_Nation3VestedTokencontracts.CallOpts, arg0, arg1)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) UserPointHistoryTs(opts *bind.CallOpts, _addr common.Address, _idx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "user_point_history__ts", _addr, _idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) UserPointHistoryTs(_addr common.Address, _idx *big.Int) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.UserPointHistoryTs(&_Nation3VestedTokencontracts.CallOpts, _addr, _idx)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) UserPointHistoryTs(_addr common.Address, _idx *big.Int) (*big.Int, error) {
	return _Nation3VestedTokencontracts.Contract.UserPointHistoryTs(&_Nation3VestedTokencontracts.CallOpts, _addr, _idx)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nation3VestedTokencontracts.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) Version() (string, error) {
	return _Nation3VestedTokencontracts.Contract.Version(&_Nation3VestedTokencontracts.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsCallerSession) Version() (string, error) {
	return _Nation3VestedTokencontracts.Contract.Version(&_Nation3VestedTokencontracts.CallOpts)
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactor) ApplySmartWalletChecker(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.contract.Transact(opts, "apply_smart_wallet_checker")
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) ApplySmartWalletChecker() (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.ApplySmartWalletChecker(&_Nation3VestedTokencontracts.TransactOpts)
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactorSession) ApplySmartWalletChecker() (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.ApplySmartWalletChecker(&_Nation3VestedTokencontracts.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.ApplyTransferOwnership(&_Nation3VestedTokencontracts.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.ApplyTransferOwnership(&_Nation3VestedTokencontracts.TransactOpts)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactor) ChangeController(opts *bind.TransactOpts, _newController common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.contract.Transact(opts, "changeController", _newController)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) ChangeController(_newController common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.ChangeController(&_Nation3VestedTokencontracts.TransactOpts, _newController)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactorSession) ChangeController(_newController common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.ChangeController(&_Nation3VestedTokencontracts.TransactOpts, _newController)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactor) Checkpoint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.contract.Transact(opts, "checkpoint")
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) Checkpoint() (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.Checkpoint(&_Nation3VestedTokencontracts.TransactOpts)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactorSession) Checkpoint() (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.Checkpoint(&_Nation3VestedTokencontracts.TransactOpts)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactor) CommitSmartWalletChecker(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.contract.Transact(opts, "commit_smart_wallet_checker", addr)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) CommitSmartWalletChecker(addr common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.CommitSmartWalletChecker(&_Nation3VestedTokencontracts.TransactOpts, addr)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactorSession) CommitSmartWalletChecker(addr common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.CommitSmartWalletChecker(&_Nation3VestedTokencontracts.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactor) CommitTransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.contract.Transact(opts, "commit_transfer_ownership", addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.CommitTransferOwnership(&_Nation3VestedTokencontracts.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactorSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.CommitTransferOwnership(&_Nation3VestedTokencontracts.TransactOpts, addr)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactor) CreateLock(opts *bind.TransactOpts, _value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.contract.Transact(opts, "create_lock", _value, _unlock_time)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) CreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.CreateLock(&_Nation3VestedTokencontracts.TransactOpts, _value, _unlock_time)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactorSession) CreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.CreateLock(&_Nation3VestedTokencontracts.TransactOpts, _value, _unlock_time)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactor) DepositFor(opts *bind.TransactOpts, _addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.contract.Transact(opts, "deposit_for", _addr, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) DepositFor(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.DepositFor(&_Nation3VestedTokencontracts.TransactOpts, _addr, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactorSession) DepositFor(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.DepositFor(&_Nation3VestedTokencontracts.TransactOpts, _addr, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactor) IncreaseAmount(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.contract.Transact(opts, "increase_amount", _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) IncreaseAmount(_value *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.IncreaseAmount(&_Nation3VestedTokencontracts.TransactOpts, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactorSession) IncreaseAmount(_value *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.IncreaseAmount(&_Nation3VestedTokencontracts.TransactOpts, _value)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactor) IncreaseUnlockTime(opts *bind.TransactOpts, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.contract.Transact(opts, "increase_unlock_time", _unlock_time)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) IncreaseUnlockTime(_unlock_time *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.IncreaseUnlockTime(&_Nation3VestedTokencontracts.TransactOpts, _unlock_time)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactorSession) IncreaseUnlockTime(_unlock_time *big.Int) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.IncreaseUnlockTime(&_Nation3VestedTokencontracts.TransactOpts, _unlock_time)
}

// SetDepositsLocked is a paid mutator transaction binding the contract method 0x4e782813.
//
// Solidity: function set_depositsLocked(bool new_status) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactor) SetDepositsLocked(opts *bind.TransactOpts, new_status bool) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.contract.Transact(opts, "set_depositsLocked", new_status)
}

// SetDepositsLocked is a paid mutator transaction binding the contract method 0x4e782813.
//
// Solidity: function set_depositsLocked(bool new_status) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) SetDepositsLocked(new_status bool) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.SetDepositsLocked(&_Nation3VestedTokencontracts.TransactOpts, new_status)
}

// SetDepositsLocked is a paid mutator transaction binding the contract method 0x4e782813.
//
// Solidity: function set_depositsLocked(bool new_status) returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactorSession) SetDepositsLocked(new_status bool) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.SetDepositsLocked(&_Nation3VestedTokencontracts.TransactOpts, new_status)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsSession) Withdraw() (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.Withdraw(&_Nation3VestedTokencontracts.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Nation3VestedTokencontracts.Contract.Withdraw(&_Nation3VestedTokencontracts.TransactOpts)
}

// Nation3VestedTokencontractsApplyOwnershipIterator is returned from FilterApplyOwnership and is used to iterate over the raw logs and unpacked data for ApplyOwnership events raised by the Nation3VestedTokencontracts contract.
type Nation3VestedTokencontractsApplyOwnershipIterator struct {
	Event *Nation3VestedTokencontractsApplyOwnership // Event containing the contract specifics and raw log

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
func (it *Nation3VestedTokencontractsApplyOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3VestedTokencontractsApplyOwnership)
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
		it.Event = new(Nation3VestedTokencontractsApplyOwnership)
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
func (it *Nation3VestedTokencontractsApplyOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3VestedTokencontractsApplyOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3VestedTokencontractsApplyOwnership represents a ApplyOwnership event raised by the Nation3VestedTokencontracts contract.
type Nation3VestedTokencontractsApplyOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterApplyOwnership is a free log retrieval operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) FilterApplyOwnership(opts *bind.FilterOpts) (*Nation3VestedTokencontractsApplyOwnershipIterator, error) {

	logs, sub, err := _Nation3VestedTokencontracts.contract.FilterLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokencontractsApplyOwnershipIterator{contract: _Nation3VestedTokencontracts.contract, event: "ApplyOwnership", logs: logs, sub: sub}, nil
}

// WatchApplyOwnership is a free log subscription operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) WatchApplyOwnership(opts *bind.WatchOpts, sink chan<- *Nation3VestedTokencontractsApplyOwnership) (event.Subscription, error) {

	logs, sub, err := _Nation3VestedTokencontracts.contract.WatchLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3VestedTokencontractsApplyOwnership)
				if err := _Nation3VestedTokencontracts.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
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
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) ParseApplyOwnership(log types.Log) (*Nation3VestedTokencontractsApplyOwnership, error) {
	event := new(Nation3VestedTokencontractsApplyOwnership)
	if err := _Nation3VestedTokencontracts.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3VestedTokencontractsCommitOwnershipIterator is returned from FilterCommitOwnership and is used to iterate over the raw logs and unpacked data for CommitOwnership events raised by the Nation3VestedTokencontracts contract.
type Nation3VestedTokencontractsCommitOwnershipIterator struct {
	Event *Nation3VestedTokencontractsCommitOwnership // Event containing the contract specifics and raw log

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
func (it *Nation3VestedTokencontractsCommitOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3VestedTokencontractsCommitOwnership)
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
		it.Event = new(Nation3VestedTokencontractsCommitOwnership)
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
func (it *Nation3VestedTokencontractsCommitOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3VestedTokencontractsCommitOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3VestedTokencontractsCommitOwnership represents a CommitOwnership event raised by the Nation3VestedTokencontracts contract.
type Nation3VestedTokencontractsCommitOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCommitOwnership is a free log retrieval operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) FilterCommitOwnership(opts *bind.FilterOpts) (*Nation3VestedTokencontractsCommitOwnershipIterator, error) {

	logs, sub, err := _Nation3VestedTokencontracts.contract.FilterLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokencontractsCommitOwnershipIterator{contract: _Nation3VestedTokencontracts.contract, event: "CommitOwnership", logs: logs, sub: sub}, nil
}

// WatchCommitOwnership is a free log subscription operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) WatchCommitOwnership(opts *bind.WatchOpts, sink chan<- *Nation3VestedTokencontractsCommitOwnership) (event.Subscription, error) {

	logs, sub, err := _Nation3VestedTokencontracts.contract.WatchLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3VestedTokencontractsCommitOwnership)
				if err := _Nation3VestedTokencontracts.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
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
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) ParseCommitOwnership(log types.Log) (*Nation3VestedTokencontractsCommitOwnership, error) {
	event := new(Nation3VestedTokencontractsCommitOwnership)
	if err := _Nation3VestedTokencontracts.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3VestedTokencontractsDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Nation3VestedTokencontracts contract.
type Nation3VestedTokencontractsDepositIterator struct {
	Event *Nation3VestedTokencontractsDeposit // Event containing the contract specifics and raw log

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
func (it *Nation3VestedTokencontractsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3VestedTokencontractsDeposit)
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
		it.Event = new(Nation3VestedTokencontractsDeposit)
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
func (it *Nation3VestedTokencontractsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3VestedTokencontractsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3VestedTokencontractsDeposit represents a Deposit event raised by the Nation3VestedTokencontracts contract.
type Nation3VestedTokencontractsDeposit struct {
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
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) FilterDeposit(opts *bind.FilterOpts, provider []common.Address, locktime []*big.Int) (*Nation3VestedTokencontractsDepositIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _Nation3VestedTokencontracts.contract.FilterLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokencontractsDepositIterator{contract: _Nation3VestedTokencontracts.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *Nation3VestedTokencontractsDeposit, provider []common.Address, locktime []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _Nation3VestedTokencontracts.contract.WatchLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3VestedTokencontractsDeposit)
				if err := _Nation3VestedTokencontracts.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) ParseDeposit(log types.Log) (*Nation3VestedTokencontractsDeposit, error) {
	event := new(Nation3VestedTokencontractsDeposit)
	if err := _Nation3VestedTokencontracts.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3VestedTokencontractsDepositsLockedChangeIterator is returned from FilterDepositsLockedChange and is used to iterate over the raw logs and unpacked data for DepositsLockedChange events raised by the Nation3VestedTokencontracts contract.
type Nation3VestedTokencontractsDepositsLockedChangeIterator struct {
	Event *Nation3VestedTokencontractsDepositsLockedChange // Event containing the contract specifics and raw log

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
func (it *Nation3VestedTokencontractsDepositsLockedChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3VestedTokencontractsDepositsLockedChange)
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
		it.Event = new(Nation3VestedTokencontractsDepositsLockedChange)
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
func (it *Nation3VestedTokencontractsDepositsLockedChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3VestedTokencontractsDepositsLockedChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3VestedTokencontractsDepositsLockedChange represents a DepositsLockedChange event raised by the Nation3VestedTokencontracts contract.
type Nation3VestedTokencontractsDepositsLockedChange struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositsLockedChange is a free log retrieval operation binding the contract event 0x3a6f72b4ed02517aec1d580529f912e860631ecb6fec4712d400294932aaffba.
//
// Solidity: event DepositsLockedChange(bool status)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) FilterDepositsLockedChange(opts *bind.FilterOpts) (*Nation3VestedTokencontractsDepositsLockedChangeIterator, error) {

	logs, sub, err := _Nation3VestedTokencontracts.contract.FilterLogs(opts, "DepositsLockedChange")
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokencontractsDepositsLockedChangeIterator{contract: _Nation3VestedTokencontracts.contract, event: "DepositsLockedChange", logs: logs, sub: sub}, nil
}

// WatchDepositsLockedChange is a free log subscription operation binding the contract event 0x3a6f72b4ed02517aec1d580529f912e860631ecb6fec4712d400294932aaffba.
//
// Solidity: event DepositsLockedChange(bool status)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) WatchDepositsLockedChange(opts *bind.WatchOpts, sink chan<- *Nation3VestedTokencontractsDepositsLockedChange) (event.Subscription, error) {

	logs, sub, err := _Nation3VestedTokencontracts.contract.WatchLogs(opts, "DepositsLockedChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3VestedTokencontractsDepositsLockedChange)
				if err := _Nation3VestedTokencontracts.contract.UnpackLog(event, "DepositsLockedChange", log); err != nil {
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
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) ParseDepositsLockedChange(log types.Log) (*Nation3VestedTokencontractsDepositsLockedChange, error) {
	event := new(Nation3VestedTokencontractsDepositsLockedChange)
	if err := _Nation3VestedTokencontracts.contract.UnpackLog(event, "DepositsLockedChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3VestedTokencontractsSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the Nation3VestedTokencontracts contract.
type Nation3VestedTokencontractsSupplyIterator struct {
	Event *Nation3VestedTokencontractsSupply // Event containing the contract specifics and raw log

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
func (it *Nation3VestedTokencontractsSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3VestedTokencontractsSupply)
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
		it.Event = new(Nation3VestedTokencontractsSupply)
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
func (it *Nation3VestedTokencontractsSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3VestedTokencontractsSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3VestedTokencontractsSupply represents a Supply event raised by the Nation3VestedTokencontracts contract.
type Nation3VestedTokencontractsSupply struct {
	PrevSupply *big.Int
	Supply     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) FilterSupply(opts *bind.FilterOpts) (*Nation3VestedTokencontractsSupplyIterator, error) {

	logs, sub, err := _Nation3VestedTokencontracts.contract.FilterLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokencontractsSupplyIterator{contract: _Nation3VestedTokencontracts.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *Nation3VestedTokencontractsSupply) (event.Subscription, error) {

	logs, sub, err := _Nation3VestedTokencontracts.contract.WatchLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3VestedTokencontractsSupply)
				if err := _Nation3VestedTokencontracts.contract.UnpackLog(event, "Supply", log); err != nil {
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
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) ParseSupply(log types.Log) (*Nation3VestedTokencontractsSupply, error) {
	event := new(Nation3VestedTokencontractsSupply)
	if err := _Nation3VestedTokencontracts.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Nation3VestedTokencontractsWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Nation3VestedTokencontracts contract.
type Nation3VestedTokencontractsWithdrawIterator struct {
	Event *Nation3VestedTokencontractsWithdraw // Event containing the contract specifics and raw log

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
func (it *Nation3VestedTokencontractsWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Nation3VestedTokencontractsWithdraw)
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
		it.Event = new(Nation3VestedTokencontractsWithdraw)
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
func (it *Nation3VestedTokencontractsWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Nation3VestedTokencontractsWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Nation3VestedTokencontractsWithdraw represents a Withdraw event raised by the Nation3VestedTokencontracts contract.
type Nation3VestedTokencontractsWithdraw struct {
	Provider common.Address
	Value    *big.Int
	Ts       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) FilterWithdraw(opts *bind.FilterOpts, provider []common.Address) (*Nation3VestedTokencontractsWithdrawIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Nation3VestedTokencontracts.contract.FilterLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return &Nation3VestedTokencontractsWithdrawIterator{contract: _Nation3VestedTokencontracts.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Nation3VestedTokencontractsWithdraw, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Nation3VestedTokencontracts.contract.WatchLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Nation3VestedTokencontractsWithdraw)
				if err := _Nation3VestedTokencontracts.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Nation3VestedTokencontracts *Nation3VestedTokencontractsFilterer) ParseWithdraw(log types.Log) (*Nation3VestedTokencontractsWithdraw, error) {
	event := new(Nation3VestedTokencontractsWithdraw)
	if err := _Nation3VestedTokencontracts.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
