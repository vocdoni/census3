// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FarcasterIDRegistry

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

// IIdRegistryBulkRegisterData is an auto generated low-level Go binding around an user-defined struct.
type IIdRegistryBulkRegisterData struct {
	Fid      *big.Int
	Custody  common.Address
	Recovery common.Address
}

// IIdRegistryBulkRegisterDefaultRecoveryData is an auto generated low-level Go binding around an user-defined struct.
type IIdRegistryBulkRegisterDefaultRecoveryData struct {
	Fid     *big.Int
	Custody common.Address
}

// FarcasterIDRegistryMetaData contains all meta data concerning the FarcasterIDRegistry contract.
var FarcasterIDRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyMigrated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GatewayFrozen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HasId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HasNoId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGuardian\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyMigrator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionRevoked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"}],\"name\":\"AdminReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"ChangeRecoveryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"idGateway\",\"type\":\"address\"}],\"name\":\"FreezeIdGateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"migratedAt\",\"type\":\"uint256\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Recover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"Remove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCounter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"}],\"name\":\"SetIdCounter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldIdGateway\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIdGateway\",\"type\":\"address\"}],\"name\":\"SetIdGateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldMigrator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newMigrator\",\"type\":\"address\"}],\"name\":\"SetMigrator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHANGE_RECOVERY_ADDRESS_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_AND_CHANGE_RECOVERY_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"addGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"fid\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"custody\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"internalType\":\"structIIdRegistry.BulkRegisterData[]\",\"name\":\"ids\",\"type\":\"tuple[]\"}],\"name\":\"bulkRegisterIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"fid\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"custody\",\"type\":\"address\"}],\"internalType\":\"structIIdRegistry.BulkRegisterDefaultRecoveryData[]\",\"name\":\"ids\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"bulkRegisterIdsWithDefaultRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24[]\",\"name\":\"ids\",\"type\":\"uint24[]\"}],\"name\":\"bulkResetIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"changeRecoveryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"changeRecoveryAddressFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"}],\"name\":\"custodyOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"custody\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparatorV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freezeIdGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatewayFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gracePeriod\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"guardians\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isGuardian\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"structHash\",\"type\":\"bytes32\"}],\"name\":\"hashTypedDataV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"idCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"idGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"idOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMigrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migratedAt\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"recoveryDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recoverySig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"toDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"toSig\",\"type\":\"bytes\"}],\"name\":\"recoverFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"}],\"name\":\"recoveryOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"removeGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_counter\",\"type\":\"uint256\"}],\"name\":\"setIdCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_idGateway\",\"type\":\"address\"}],\"name\":\"setIdGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"transferAndChangeRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"toDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"toSig\",\"type\":\"bytes\"}],\"name\":\"transferAndChangeRecoveryFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fromSig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"toDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"toSig\",\"type\":\"bytes\"}],\"name\":\"transferFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"custodyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verifyFidSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FarcasterIDRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use FarcasterIDRegistryMetaData.ABI instead.
var FarcasterIDRegistryABI = FarcasterIDRegistryMetaData.ABI

// FarcasterIDRegistry is an auto generated Go binding around an Ethereum contract.
type FarcasterIDRegistry struct {
	FarcasterIDRegistryCaller     // Read-only binding to the contract
	FarcasterIDRegistryTransactor // Write-only binding to the contract
	FarcasterIDRegistryFilterer   // Log filterer for contract events
}

// FarcasterIDRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FarcasterIDRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarcasterIDRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FarcasterIDRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarcasterIDRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FarcasterIDRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarcasterIDRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FarcasterIDRegistrySession struct {
	Contract     *FarcasterIDRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FarcasterIDRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FarcasterIDRegistryCallerSession struct {
	Contract *FarcasterIDRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// FarcasterIDRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FarcasterIDRegistryTransactorSession struct {
	Contract     *FarcasterIDRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// FarcasterIDRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FarcasterIDRegistryRaw struct {
	Contract *FarcasterIDRegistry // Generic contract binding to access the raw methods on
}

// FarcasterIDRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FarcasterIDRegistryCallerRaw struct {
	Contract *FarcasterIDRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// FarcasterIDRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FarcasterIDRegistryTransactorRaw struct {
	Contract *FarcasterIDRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFarcasterIDRegistry creates a new instance of FarcasterIDRegistry, bound to a specific deployed contract.
func NewFarcasterIDRegistry(address common.Address, backend bind.ContractBackend) (*FarcasterIDRegistry, error) {
	contract, err := bindFarcasterIDRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistry{FarcasterIDRegistryCaller: FarcasterIDRegistryCaller{contract: contract}, FarcasterIDRegistryTransactor: FarcasterIDRegistryTransactor{contract: contract}, FarcasterIDRegistryFilterer: FarcasterIDRegistryFilterer{contract: contract}}, nil
}

// NewFarcasterIDRegistryCaller creates a new read-only instance of FarcasterIDRegistry, bound to a specific deployed contract.
func NewFarcasterIDRegistryCaller(address common.Address, caller bind.ContractCaller) (*FarcasterIDRegistryCaller, error) {
	contract, err := bindFarcasterIDRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryCaller{contract: contract}, nil
}

// NewFarcasterIDRegistryTransactor creates a new write-only instance of FarcasterIDRegistry, bound to a specific deployed contract.
func NewFarcasterIDRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*FarcasterIDRegistryTransactor, error) {
	contract, err := bindFarcasterIDRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryTransactor{contract: contract}, nil
}

// NewFarcasterIDRegistryFilterer creates a new log filterer instance of FarcasterIDRegistry, bound to a specific deployed contract.
func NewFarcasterIDRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*FarcasterIDRegistryFilterer, error) {
	contract, err := bindFarcasterIDRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryFilterer{contract: contract}, nil
}

// bindFarcasterIDRegistry binds a generic wrapper to an already deployed contract.
func bindFarcasterIDRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FarcasterIDRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FarcasterIDRegistry *FarcasterIDRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FarcasterIDRegistry.Contract.FarcasterIDRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FarcasterIDRegistry *FarcasterIDRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.FarcasterIDRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FarcasterIDRegistry *FarcasterIDRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.FarcasterIDRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FarcasterIDRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.contract.Transact(opts, method, params...)
}

// CHANGERECOVERYADDRESSTYPEHASH is a free data retrieval call binding the contract method 0xd5bac7f3.
//
// Solidity: function CHANGE_RECOVERY_ADDRESS_TYPEHASH() view returns(bytes32)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) CHANGERECOVERYADDRESSTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "CHANGE_RECOVERY_ADDRESS_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CHANGERECOVERYADDRESSTYPEHASH is a free data retrieval call binding the contract method 0xd5bac7f3.
//
// Solidity: function CHANGE_RECOVERY_ADDRESS_TYPEHASH() view returns(bytes32)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) CHANGERECOVERYADDRESSTYPEHASH() ([32]byte, error) {
	return _FarcasterIDRegistry.Contract.CHANGERECOVERYADDRESSTYPEHASH(&_FarcasterIDRegistry.CallOpts)
}

// CHANGERECOVERYADDRESSTYPEHASH is a free data retrieval call binding the contract method 0xd5bac7f3.
//
// Solidity: function CHANGE_RECOVERY_ADDRESS_TYPEHASH() view returns(bytes32)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) CHANGERECOVERYADDRESSTYPEHASH() ([32]byte, error) {
	return _FarcasterIDRegistry.Contract.CHANGERECOVERYADDRESSTYPEHASH(&_FarcasterIDRegistry.CallOpts)
}

// TRANSFERANDCHANGERECOVERYTYPEHASH is a free data retrieval call binding the contract method 0xea2bbb83.
//
// Solidity: function TRANSFER_AND_CHANGE_RECOVERY_TYPEHASH() view returns(bytes32)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) TRANSFERANDCHANGERECOVERYTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "TRANSFER_AND_CHANGE_RECOVERY_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFERANDCHANGERECOVERYTYPEHASH is a free data retrieval call binding the contract method 0xea2bbb83.
//
// Solidity: function TRANSFER_AND_CHANGE_RECOVERY_TYPEHASH() view returns(bytes32)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) TRANSFERANDCHANGERECOVERYTYPEHASH() ([32]byte, error) {
	return _FarcasterIDRegistry.Contract.TRANSFERANDCHANGERECOVERYTYPEHASH(&_FarcasterIDRegistry.CallOpts)
}

// TRANSFERANDCHANGERECOVERYTYPEHASH is a free data retrieval call binding the contract method 0xea2bbb83.
//
// Solidity: function TRANSFER_AND_CHANGE_RECOVERY_TYPEHASH() view returns(bytes32)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) TRANSFERANDCHANGERECOVERYTYPEHASH() ([32]byte, error) {
	return _FarcasterIDRegistry.Contract.TRANSFERANDCHANGERECOVERYTYPEHASH(&_FarcasterIDRegistry.CallOpts)
}

// TRANSFERTYPEHASH is a free data retrieval call binding the contract method 0x00bf26f4.
//
// Solidity: function TRANSFER_TYPEHASH() view returns(bytes32)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) TRANSFERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "TRANSFER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFERTYPEHASH is a free data retrieval call binding the contract method 0x00bf26f4.
//
// Solidity: function TRANSFER_TYPEHASH() view returns(bytes32)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) TRANSFERTYPEHASH() ([32]byte, error) {
	return _FarcasterIDRegistry.Contract.TRANSFERTYPEHASH(&_FarcasterIDRegistry.CallOpts)
}

// TRANSFERTYPEHASH is a free data retrieval call binding the contract method 0x00bf26f4.
//
// Solidity: function TRANSFER_TYPEHASH() view returns(bytes32)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) TRANSFERTYPEHASH() ([32]byte, error) {
	return _FarcasterIDRegistry.Contract.TRANSFERTYPEHASH(&_FarcasterIDRegistry.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) VERSION() (string, error) {
	return _FarcasterIDRegistry.Contract.VERSION(&_FarcasterIDRegistry.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) VERSION() (string, error) {
	return _FarcasterIDRegistry.Contract.VERSION(&_FarcasterIDRegistry.CallOpts)
}

// CustodyOf is a free data retrieval call binding the contract method 0x65269e47.
//
// Solidity: function custodyOf(uint256 fid) view returns(address custody)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) CustodyOf(opts *bind.CallOpts, fid *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "custodyOf", fid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CustodyOf is a free data retrieval call binding the contract method 0x65269e47.
//
// Solidity: function custodyOf(uint256 fid) view returns(address custody)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) CustodyOf(fid *big.Int) (common.Address, error) {
	return _FarcasterIDRegistry.Contract.CustodyOf(&_FarcasterIDRegistry.CallOpts, fid)
}

// CustodyOf is a free data retrieval call binding the contract method 0x65269e47.
//
// Solidity: function custodyOf(uint256 fid) view returns(address custody)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) CustodyOf(fid *big.Int) (common.Address, error) {
	return _FarcasterIDRegistry.Contract.CustodyOf(&_FarcasterIDRegistry.CallOpts, fid)
}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) DomainSeparatorV4(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "domainSeparatorV4")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) DomainSeparatorV4() ([32]byte, error) {
	return _FarcasterIDRegistry.Contract.DomainSeparatorV4(&_FarcasterIDRegistry.CallOpts)
}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) DomainSeparatorV4() ([32]byte, error) {
	return _FarcasterIDRegistry.Contract.DomainSeparatorV4(&_FarcasterIDRegistry.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _FarcasterIDRegistry.Contract.Eip712Domain(&_FarcasterIDRegistry.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _FarcasterIDRegistry.Contract.Eip712Domain(&_FarcasterIDRegistry.CallOpts)
}

// GatewayFrozen is a free data retrieval call binding the contract method 0x95e7549f.
//
// Solidity: function gatewayFrozen() view returns(bool)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) GatewayFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "gatewayFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GatewayFrozen is a free data retrieval call binding the contract method 0x95e7549f.
//
// Solidity: function gatewayFrozen() view returns(bool)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) GatewayFrozen() (bool, error) {
	return _FarcasterIDRegistry.Contract.GatewayFrozen(&_FarcasterIDRegistry.CallOpts)
}

// GatewayFrozen is a free data retrieval call binding the contract method 0x95e7549f.
//
// Solidity: function gatewayFrozen() view returns(bool)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) GatewayFrozen() (bool, error) {
	return _FarcasterIDRegistry.Contract.GatewayFrozen(&_FarcasterIDRegistry.CallOpts)
}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() view returns(uint24)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) GracePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "gracePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() view returns(uint24)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) GracePeriod() (*big.Int, error) {
	return _FarcasterIDRegistry.Contract.GracePeriod(&_FarcasterIDRegistry.CallOpts)
}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() view returns(uint24)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) GracePeriod() (*big.Int, error) {
	return _FarcasterIDRegistry.Contract.GracePeriod(&_FarcasterIDRegistry.CallOpts)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) Guardians(opts *bind.CallOpts, guardian common.Address) (bool, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "guardians", guardian)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) Guardians(guardian common.Address) (bool, error) {
	return _FarcasterIDRegistry.Contract.Guardians(&_FarcasterIDRegistry.CallOpts, guardian)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) Guardians(guardian common.Address) (bool, error) {
	return _FarcasterIDRegistry.Contract.Guardians(&_FarcasterIDRegistry.CallOpts, guardian)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) HashTypedDataV4(opts *bind.CallOpts, structHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "hashTypedDataV4", structHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _FarcasterIDRegistry.Contract.HashTypedDataV4(&_FarcasterIDRegistry.CallOpts, structHash)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _FarcasterIDRegistry.Contract.HashTypedDataV4(&_FarcasterIDRegistry.CallOpts, structHash)
}

// IdCounter is a free data retrieval call binding the contract method 0xeb08ab28.
//
// Solidity: function idCounter() view returns(uint256)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) IdCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "idCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdCounter is a free data retrieval call binding the contract method 0xeb08ab28.
//
// Solidity: function idCounter() view returns(uint256)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) IdCounter() (*big.Int, error) {
	return _FarcasterIDRegistry.Contract.IdCounter(&_FarcasterIDRegistry.CallOpts)
}

// IdCounter is a free data retrieval call binding the contract method 0xeb08ab28.
//
// Solidity: function idCounter() view returns(uint256)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) IdCounter() (*big.Int, error) {
	return _FarcasterIDRegistry.Contract.IdCounter(&_FarcasterIDRegistry.CallOpts)
}

// IdGateway is a free data retrieval call binding the contract method 0x4b57a600.
//
// Solidity: function idGateway() view returns(address)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) IdGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "idGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdGateway is a free data retrieval call binding the contract method 0x4b57a600.
//
// Solidity: function idGateway() view returns(address)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) IdGateway() (common.Address, error) {
	return _FarcasterIDRegistry.Contract.IdGateway(&_FarcasterIDRegistry.CallOpts)
}

// IdGateway is a free data retrieval call binding the contract method 0x4b57a600.
//
// Solidity: function idGateway() view returns(address)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) IdGateway() (common.Address, error) {
	return _FarcasterIDRegistry.Contract.IdGateway(&_FarcasterIDRegistry.CallOpts)
}

// IdOf is a free data retrieval call binding the contract method 0xd94fe832.
//
// Solidity: function idOf(address owner) view returns(uint256 fid)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) IdOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "idOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdOf is a free data retrieval call binding the contract method 0xd94fe832.
//
// Solidity: function idOf(address owner) view returns(uint256 fid)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) IdOf(owner common.Address) (*big.Int, error) {
	return _FarcasterIDRegistry.Contract.IdOf(&_FarcasterIDRegistry.CallOpts, owner)
}

// IdOf is a free data retrieval call binding the contract method 0xd94fe832.
//
// Solidity: function idOf(address owner) view returns(uint256 fid)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) IdOf(owner common.Address) (*big.Int, error) {
	return _FarcasterIDRegistry.Contract.IdOf(&_FarcasterIDRegistry.CallOpts, owner)
}

// IsMigrated is a free data retrieval call binding the contract method 0xb06faf62.
//
// Solidity: function isMigrated() view returns(bool)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) IsMigrated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "isMigrated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMigrated is a free data retrieval call binding the contract method 0xb06faf62.
//
// Solidity: function isMigrated() view returns(bool)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) IsMigrated() (bool, error) {
	return _FarcasterIDRegistry.Contract.IsMigrated(&_FarcasterIDRegistry.CallOpts)
}

// IsMigrated is a free data retrieval call binding the contract method 0xb06faf62.
//
// Solidity: function isMigrated() view returns(bool)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) IsMigrated() (bool, error) {
	return _FarcasterIDRegistry.Contract.IsMigrated(&_FarcasterIDRegistry.CallOpts)
}

// MigratedAt is a free data retrieval call binding the contract method 0x8b21e484.
//
// Solidity: function migratedAt() view returns(uint40)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) MigratedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "migratedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MigratedAt is a free data retrieval call binding the contract method 0x8b21e484.
//
// Solidity: function migratedAt() view returns(uint40)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) MigratedAt() (*big.Int, error) {
	return _FarcasterIDRegistry.Contract.MigratedAt(&_FarcasterIDRegistry.CallOpts)
}

// MigratedAt is a free data retrieval call binding the contract method 0x8b21e484.
//
// Solidity: function migratedAt() view returns(uint40)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) MigratedAt() (*big.Int, error) {
	return _FarcasterIDRegistry.Contract.MigratedAt(&_FarcasterIDRegistry.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) Migrator() (common.Address, error) {
	return _FarcasterIDRegistry.Contract.Migrator(&_FarcasterIDRegistry.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) Migrator() (common.Address, error) {
	return _FarcasterIDRegistry.Contract.Migrator(&_FarcasterIDRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) Name() (string, error) {
	return _FarcasterIDRegistry.Contract.Name(&_FarcasterIDRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) Name() (string, error) {
	return _FarcasterIDRegistry.Contract.Name(&_FarcasterIDRegistry.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) Nonces(owner common.Address) (*big.Int, error) {
	return _FarcasterIDRegistry.Contract.Nonces(&_FarcasterIDRegistry.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _FarcasterIDRegistry.Contract.Nonces(&_FarcasterIDRegistry.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) Owner() (common.Address, error) {
	return _FarcasterIDRegistry.Contract.Owner(&_FarcasterIDRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) Owner() (common.Address, error) {
	return _FarcasterIDRegistry.Contract.Owner(&_FarcasterIDRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) Paused() (bool, error) {
	return _FarcasterIDRegistry.Contract.Paused(&_FarcasterIDRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) Paused() (bool, error) {
	return _FarcasterIDRegistry.Contract.Paused(&_FarcasterIDRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) PendingOwner() (common.Address, error) {
	return _FarcasterIDRegistry.Contract.PendingOwner(&_FarcasterIDRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) PendingOwner() (common.Address, error) {
	return _FarcasterIDRegistry.Contract.PendingOwner(&_FarcasterIDRegistry.CallOpts)
}

// RecoveryOf is a free data retrieval call binding the contract method 0xfa1a1b25.
//
// Solidity: function recoveryOf(uint256 fid) view returns(address recovery)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) RecoveryOf(opts *bind.CallOpts, fid *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "recoveryOf", fid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoveryOf is a free data retrieval call binding the contract method 0xfa1a1b25.
//
// Solidity: function recoveryOf(uint256 fid) view returns(address recovery)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) RecoveryOf(fid *big.Int) (common.Address, error) {
	return _FarcasterIDRegistry.Contract.RecoveryOf(&_FarcasterIDRegistry.CallOpts, fid)
}

// RecoveryOf is a free data retrieval call binding the contract method 0xfa1a1b25.
//
// Solidity: function recoveryOf(uint256 fid) view returns(address recovery)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) RecoveryOf(fid *big.Int) (common.Address, error) {
	return _FarcasterIDRegistry.Contract.RecoveryOf(&_FarcasterIDRegistry.CallOpts, fid)
}

// VerifyFidSignature is a free data retrieval call binding the contract method 0x32faac70.
//
// Solidity: function verifyFidSignature(address custodyAddress, uint256 fid, bytes32 digest, bytes sig) view returns(bool isValid)
func (_FarcasterIDRegistry *FarcasterIDRegistryCaller) VerifyFidSignature(opts *bind.CallOpts, custodyAddress common.Address, fid *big.Int, digest [32]byte, sig []byte) (bool, error) {
	var out []interface{}
	err := _FarcasterIDRegistry.contract.Call(opts, &out, "verifyFidSignature", custodyAddress, fid, digest, sig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyFidSignature is a free data retrieval call binding the contract method 0x32faac70.
//
// Solidity: function verifyFidSignature(address custodyAddress, uint256 fid, bytes32 digest, bytes sig) view returns(bool isValid)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) VerifyFidSignature(custodyAddress common.Address, fid *big.Int, digest [32]byte, sig []byte) (bool, error) {
	return _FarcasterIDRegistry.Contract.VerifyFidSignature(&_FarcasterIDRegistry.CallOpts, custodyAddress, fid, digest, sig)
}

// VerifyFidSignature is a free data retrieval call binding the contract method 0x32faac70.
//
// Solidity: function verifyFidSignature(address custodyAddress, uint256 fid, bytes32 digest, bytes sig) view returns(bool isValid)
func (_FarcasterIDRegistry *FarcasterIDRegistryCallerSession) VerifyFidSignature(custodyAddress common.Address, fid *big.Int, digest [32]byte, sig []byte) (bool, error) {
	return _FarcasterIDRegistry.Contract.VerifyFidSignature(&_FarcasterIDRegistry.CallOpts, custodyAddress, fid, digest, sig)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.AcceptOwnership(&_FarcasterIDRegistry.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.AcceptOwnership(&_FarcasterIDRegistry.TransactOpts)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) AddGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "addGuardian", guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.AddGuardian(&_FarcasterIDRegistry.TransactOpts, guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.AddGuardian(&_FarcasterIDRegistry.TransactOpts, guardian)
}

// BulkRegisterIds is a paid mutator transaction binding the contract method 0x55c5b358.
//
// Solidity: function bulkRegisterIds((uint24,address,address)[] ids) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) BulkRegisterIds(opts *bind.TransactOpts, ids []IIdRegistryBulkRegisterData) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "bulkRegisterIds", ids)
}

// BulkRegisterIds is a paid mutator transaction binding the contract method 0x55c5b358.
//
// Solidity: function bulkRegisterIds((uint24,address,address)[] ids) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) BulkRegisterIds(ids []IIdRegistryBulkRegisterData) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.BulkRegisterIds(&_FarcasterIDRegistry.TransactOpts, ids)
}

// BulkRegisterIds is a paid mutator transaction binding the contract method 0x55c5b358.
//
// Solidity: function bulkRegisterIds((uint24,address,address)[] ids) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) BulkRegisterIds(ids []IIdRegistryBulkRegisterData) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.BulkRegisterIds(&_FarcasterIDRegistry.TransactOpts, ids)
}

// BulkRegisterIdsWithDefaultRecovery is a paid mutator transaction binding the contract method 0x8d8043e2.
//
// Solidity: function bulkRegisterIdsWithDefaultRecovery((uint24,address)[] ids, address recovery) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) BulkRegisterIdsWithDefaultRecovery(opts *bind.TransactOpts, ids []IIdRegistryBulkRegisterDefaultRecoveryData, recovery common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "bulkRegisterIdsWithDefaultRecovery", ids, recovery)
}

// BulkRegisterIdsWithDefaultRecovery is a paid mutator transaction binding the contract method 0x8d8043e2.
//
// Solidity: function bulkRegisterIdsWithDefaultRecovery((uint24,address)[] ids, address recovery) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) BulkRegisterIdsWithDefaultRecovery(ids []IIdRegistryBulkRegisterDefaultRecoveryData, recovery common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.BulkRegisterIdsWithDefaultRecovery(&_FarcasterIDRegistry.TransactOpts, ids, recovery)
}

// BulkRegisterIdsWithDefaultRecovery is a paid mutator transaction binding the contract method 0x8d8043e2.
//
// Solidity: function bulkRegisterIdsWithDefaultRecovery((uint24,address)[] ids, address recovery) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) BulkRegisterIdsWithDefaultRecovery(ids []IIdRegistryBulkRegisterDefaultRecoveryData, recovery common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.BulkRegisterIdsWithDefaultRecovery(&_FarcasterIDRegistry.TransactOpts, ids, recovery)
}

// BulkResetIds is a paid mutator transaction binding the contract method 0xff126441.
//
// Solidity: function bulkResetIds(uint24[] ids) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) BulkResetIds(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "bulkResetIds", ids)
}

// BulkResetIds is a paid mutator transaction binding the contract method 0xff126441.
//
// Solidity: function bulkResetIds(uint24[] ids) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) BulkResetIds(ids []*big.Int) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.BulkResetIds(&_FarcasterIDRegistry.TransactOpts, ids)
}

// BulkResetIds is a paid mutator transaction binding the contract method 0xff126441.
//
// Solidity: function bulkResetIds(uint24[] ids) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) BulkResetIds(ids []*big.Int) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.BulkResetIds(&_FarcasterIDRegistry.TransactOpts, ids)
}

// ChangeRecoveryAddress is a paid mutator transaction binding the contract method 0xf1f0b224.
//
// Solidity: function changeRecoveryAddress(address recovery) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) ChangeRecoveryAddress(opts *bind.TransactOpts, recovery common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "changeRecoveryAddress", recovery)
}

// ChangeRecoveryAddress is a paid mutator transaction binding the contract method 0xf1f0b224.
//
// Solidity: function changeRecoveryAddress(address recovery) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) ChangeRecoveryAddress(recovery common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.ChangeRecoveryAddress(&_FarcasterIDRegistry.TransactOpts, recovery)
}

// ChangeRecoveryAddress is a paid mutator transaction binding the contract method 0xf1f0b224.
//
// Solidity: function changeRecoveryAddress(address recovery) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) ChangeRecoveryAddress(recovery common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.ChangeRecoveryAddress(&_FarcasterIDRegistry.TransactOpts, recovery)
}

// ChangeRecoveryAddressFor is a paid mutator transaction binding the contract method 0x9cbef8dc.
//
// Solidity: function changeRecoveryAddressFor(address owner, address recovery, uint256 deadline, bytes sig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) ChangeRecoveryAddressFor(opts *bind.TransactOpts, owner common.Address, recovery common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "changeRecoveryAddressFor", owner, recovery, deadline, sig)
}

// ChangeRecoveryAddressFor is a paid mutator transaction binding the contract method 0x9cbef8dc.
//
// Solidity: function changeRecoveryAddressFor(address owner, address recovery, uint256 deadline, bytes sig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) ChangeRecoveryAddressFor(owner common.Address, recovery common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.ChangeRecoveryAddressFor(&_FarcasterIDRegistry.TransactOpts, owner, recovery, deadline, sig)
}

// ChangeRecoveryAddressFor is a paid mutator transaction binding the contract method 0x9cbef8dc.
//
// Solidity: function changeRecoveryAddressFor(address owner, address recovery, uint256 deadline, bytes sig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) ChangeRecoveryAddressFor(owner common.Address, recovery common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.ChangeRecoveryAddressFor(&_FarcasterIDRegistry.TransactOpts, owner, recovery, deadline, sig)
}

// FreezeIdGateway is a paid mutator transaction binding the contract method 0xddd76649.
//
// Solidity: function freezeIdGateway() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) FreezeIdGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "freezeIdGateway")
}

// FreezeIdGateway is a paid mutator transaction binding the contract method 0xddd76649.
//
// Solidity: function freezeIdGateway() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) FreezeIdGateway() (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.FreezeIdGateway(&_FarcasterIDRegistry.TransactOpts)
}

// FreezeIdGateway is a paid mutator transaction binding the contract method 0xddd76649.
//
// Solidity: function freezeIdGateway() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) FreezeIdGateway() (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.FreezeIdGateway(&_FarcasterIDRegistry.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) Migrate() (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.Migrate(&_FarcasterIDRegistry.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) Migrate() (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.Migrate(&_FarcasterIDRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) Pause() (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.Pause(&_FarcasterIDRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.Pause(&_FarcasterIDRegistry.TransactOpts)
}

// Recover is a paid mutator transaction binding the contract method 0x2a42ede3.
//
// Solidity: function recover(address from, address to, uint256 deadline, bytes sig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) Recover(opts *bind.TransactOpts, from common.Address, to common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "recover", from, to, deadline, sig)
}

// Recover is a paid mutator transaction binding the contract method 0x2a42ede3.
//
// Solidity: function recover(address from, address to, uint256 deadline, bytes sig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) Recover(from common.Address, to common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.Recover(&_FarcasterIDRegistry.TransactOpts, from, to, deadline, sig)
}

// Recover is a paid mutator transaction binding the contract method 0x2a42ede3.
//
// Solidity: function recover(address from, address to, uint256 deadline, bytes sig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) Recover(from common.Address, to common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.Recover(&_FarcasterIDRegistry.TransactOpts, from, to, deadline, sig)
}

// RecoverFor is a paid mutator transaction binding the contract method 0xba656434.
//
// Solidity: function recoverFor(address from, address to, uint256 recoveryDeadline, bytes recoverySig, uint256 toDeadline, bytes toSig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) RecoverFor(opts *bind.TransactOpts, from common.Address, to common.Address, recoveryDeadline *big.Int, recoverySig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "recoverFor", from, to, recoveryDeadline, recoverySig, toDeadline, toSig)
}

// RecoverFor is a paid mutator transaction binding the contract method 0xba656434.
//
// Solidity: function recoverFor(address from, address to, uint256 recoveryDeadline, bytes recoverySig, uint256 toDeadline, bytes toSig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) RecoverFor(from common.Address, to common.Address, recoveryDeadline *big.Int, recoverySig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.RecoverFor(&_FarcasterIDRegistry.TransactOpts, from, to, recoveryDeadline, recoverySig, toDeadline, toSig)
}

// RecoverFor is a paid mutator transaction binding the contract method 0xba656434.
//
// Solidity: function recoverFor(address from, address to, uint256 recoveryDeadline, bytes recoverySig, uint256 toDeadline, bytes toSig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) RecoverFor(from common.Address, to common.Address, recoveryDeadline *big.Int, recoverySig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.RecoverFor(&_FarcasterIDRegistry.TransactOpts, from, to, recoveryDeadline, recoverySig, toDeadline, toSig)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address to, address recovery) returns(uint256 fid)
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) Register(opts *bind.TransactOpts, to common.Address, recovery common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "register", to, recovery)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address to, address recovery) returns(uint256 fid)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) Register(to common.Address, recovery common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.Register(&_FarcasterIDRegistry.TransactOpts, to, recovery)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address to, address recovery) returns(uint256 fid)
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) Register(to common.Address, recovery common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.Register(&_FarcasterIDRegistry.TransactOpts, to, recovery)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) RemoveGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "removeGuardian", guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.RemoveGuardian(&_FarcasterIDRegistry.TransactOpts, guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.RemoveGuardian(&_FarcasterIDRegistry.TransactOpts, guardian)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.RenounceOwnership(&_FarcasterIDRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.RenounceOwnership(&_FarcasterIDRegistry.TransactOpts)
}

// SetIdCounter is a paid mutator transaction binding the contract method 0xa5ed6a6a.
//
// Solidity: function setIdCounter(uint256 _counter) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) SetIdCounter(opts *bind.TransactOpts, _counter *big.Int) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "setIdCounter", _counter)
}

// SetIdCounter is a paid mutator transaction binding the contract method 0xa5ed6a6a.
//
// Solidity: function setIdCounter(uint256 _counter) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) SetIdCounter(_counter *big.Int) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.SetIdCounter(&_FarcasterIDRegistry.TransactOpts, _counter)
}

// SetIdCounter is a paid mutator transaction binding the contract method 0xa5ed6a6a.
//
// Solidity: function setIdCounter(uint256 _counter) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) SetIdCounter(_counter *big.Int) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.SetIdCounter(&_FarcasterIDRegistry.TransactOpts, _counter)
}

// SetIdGateway is a paid mutator transaction binding the contract method 0x033e2cb3.
//
// Solidity: function setIdGateway(address _idGateway) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) SetIdGateway(opts *bind.TransactOpts, _idGateway common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "setIdGateway", _idGateway)
}

// SetIdGateway is a paid mutator transaction binding the contract method 0x033e2cb3.
//
// Solidity: function setIdGateway(address _idGateway) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) SetIdGateway(_idGateway common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.SetIdGateway(&_FarcasterIDRegistry.TransactOpts, _idGateway)
}

// SetIdGateway is a paid mutator transaction binding the contract method 0x033e2cb3.
//
// Solidity: function setIdGateway(address _idGateway) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) SetIdGateway(_idGateway common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.SetIdGateway(&_FarcasterIDRegistry.TransactOpts, _idGateway)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.SetMigrator(&_FarcasterIDRegistry.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.SetMigrator(&_FarcasterIDRegistry.TransactOpts, _migrator)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address to, uint256 deadline, bytes sig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) Transfer(opts *bind.TransactOpts, to common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "transfer", to, deadline, sig)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address to, uint256 deadline, bytes sig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) Transfer(to common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.Transfer(&_FarcasterIDRegistry.TransactOpts, to, deadline, sig)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address to, uint256 deadline, bytes sig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) Transfer(to common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.Transfer(&_FarcasterIDRegistry.TransactOpts, to, deadline, sig)
}

// TransferAndChangeRecovery is a paid mutator transaction binding the contract method 0x3ab8465d.
//
// Solidity: function transferAndChangeRecovery(address to, address recovery, uint256 deadline, bytes sig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) TransferAndChangeRecovery(opts *bind.TransactOpts, to common.Address, recovery common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "transferAndChangeRecovery", to, recovery, deadline, sig)
}

// TransferAndChangeRecovery is a paid mutator transaction binding the contract method 0x3ab8465d.
//
// Solidity: function transferAndChangeRecovery(address to, address recovery, uint256 deadline, bytes sig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) TransferAndChangeRecovery(to common.Address, recovery common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.TransferAndChangeRecovery(&_FarcasterIDRegistry.TransactOpts, to, recovery, deadline, sig)
}

// TransferAndChangeRecovery is a paid mutator transaction binding the contract method 0x3ab8465d.
//
// Solidity: function transferAndChangeRecovery(address to, address recovery, uint256 deadline, bytes sig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) TransferAndChangeRecovery(to common.Address, recovery common.Address, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.TransferAndChangeRecovery(&_FarcasterIDRegistry.TransactOpts, to, recovery, deadline, sig)
}

// TransferAndChangeRecoveryFor is a paid mutator transaction binding the contract method 0x4c5cbb34.
//
// Solidity: function transferAndChangeRecoveryFor(address from, address to, address recovery, uint256 fromDeadline, bytes fromSig, uint256 toDeadline, bytes toSig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) TransferAndChangeRecoveryFor(opts *bind.TransactOpts, from common.Address, to common.Address, recovery common.Address, fromDeadline *big.Int, fromSig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "transferAndChangeRecoveryFor", from, to, recovery, fromDeadline, fromSig, toDeadline, toSig)
}

// TransferAndChangeRecoveryFor is a paid mutator transaction binding the contract method 0x4c5cbb34.
//
// Solidity: function transferAndChangeRecoveryFor(address from, address to, address recovery, uint256 fromDeadline, bytes fromSig, uint256 toDeadline, bytes toSig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) TransferAndChangeRecoveryFor(from common.Address, to common.Address, recovery common.Address, fromDeadline *big.Int, fromSig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.TransferAndChangeRecoveryFor(&_FarcasterIDRegistry.TransactOpts, from, to, recovery, fromDeadline, fromSig, toDeadline, toSig)
}

// TransferAndChangeRecoveryFor is a paid mutator transaction binding the contract method 0x4c5cbb34.
//
// Solidity: function transferAndChangeRecoveryFor(address from, address to, address recovery, uint256 fromDeadline, bytes fromSig, uint256 toDeadline, bytes toSig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) TransferAndChangeRecoveryFor(from common.Address, to common.Address, recovery common.Address, fromDeadline *big.Int, fromSig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.TransferAndChangeRecoveryFor(&_FarcasterIDRegistry.TransactOpts, from, to, recovery, fromDeadline, fromSig, toDeadline, toSig)
}

// TransferFor is a paid mutator transaction binding the contract method 0x16f72842.
//
// Solidity: function transferFor(address from, address to, uint256 fromDeadline, bytes fromSig, uint256 toDeadline, bytes toSig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) TransferFor(opts *bind.TransactOpts, from common.Address, to common.Address, fromDeadline *big.Int, fromSig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "transferFor", from, to, fromDeadline, fromSig, toDeadline, toSig)
}

// TransferFor is a paid mutator transaction binding the contract method 0x16f72842.
//
// Solidity: function transferFor(address from, address to, uint256 fromDeadline, bytes fromSig, uint256 toDeadline, bytes toSig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) TransferFor(from common.Address, to common.Address, fromDeadline *big.Int, fromSig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.TransferFor(&_FarcasterIDRegistry.TransactOpts, from, to, fromDeadline, fromSig, toDeadline, toSig)
}

// TransferFor is a paid mutator transaction binding the contract method 0x16f72842.
//
// Solidity: function transferFor(address from, address to, uint256 fromDeadline, bytes fromSig, uint256 toDeadline, bytes toSig) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) TransferFor(from common.Address, to common.Address, fromDeadline *big.Int, fromSig []byte, toDeadline *big.Int, toSig []byte) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.TransferFor(&_FarcasterIDRegistry.TransactOpts, from, to, fromDeadline, fromSig, toDeadline, toSig)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.TransferOwnership(&_FarcasterIDRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.TransferOwnership(&_FarcasterIDRegistry.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) Unpause() (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.Unpause(&_FarcasterIDRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.Unpause(&_FarcasterIDRegistry.TransactOpts)
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactor) UseNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterIDRegistry.contract.Transact(opts, "useNonce")
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_FarcasterIDRegistry *FarcasterIDRegistrySession) UseNonce() (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.UseNonce(&_FarcasterIDRegistry.TransactOpts)
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_FarcasterIDRegistry *FarcasterIDRegistryTransactorSession) UseNonce() (*types.Transaction, error) {
	return _FarcasterIDRegistry.Contract.UseNonce(&_FarcasterIDRegistry.TransactOpts)
}

// FarcasterIDRegistryAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryAddIterator struct {
	Event *FarcasterIDRegistryAdd // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistryAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistryAdd)
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
		it.Event = new(FarcasterIDRegistryAdd)
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
func (it *FarcasterIDRegistryAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistryAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistryAdd represents a Add event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryAdd struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterAdd(opts *bind.FilterOpts, guardian []common.Address) (*FarcasterIDRegistryAddIterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "Add", guardianRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryAddIterator{contract: _FarcasterIDRegistry.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistryAdd, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "Add", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistryAdd)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Add", log); err != nil {
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

// ParseAdd is a log parse operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseAdd(log types.Log) (*FarcasterIDRegistryAdd, error) {
	event := new(FarcasterIDRegistryAdd)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistryAdminResetIterator is returned from FilterAdminReset and is used to iterate over the raw logs and unpacked data for AdminReset events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryAdminResetIterator struct {
	Event *FarcasterIDRegistryAdminReset // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistryAdminResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistryAdminReset)
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
		it.Event = new(FarcasterIDRegistryAdminReset)
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
func (it *FarcasterIDRegistryAdminResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistryAdminResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistryAdminReset represents a AdminReset event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryAdminReset struct {
	Fid *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAdminReset is a free log retrieval operation binding the contract event 0x8b4b4c6da5b89da518fb865149e01ad2863b48861a8b952e11645f663959fa70.
//
// Solidity: event AdminReset(uint256 indexed fid)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterAdminReset(opts *bind.FilterOpts, fid []*big.Int) (*FarcasterIDRegistryAdminResetIterator, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "AdminReset", fidRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryAdminResetIterator{contract: _FarcasterIDRegistry.contract, event: "AdminReset", logs: logs, sub: sub}, nil
}

// WatchAdminReset is a free log subscription operation binding the contract event 0x8b4b4c6da5b89da518fb865149e01ad2863b48861a8b952e11645f663959fa70.
//
// Solidity: event AdminReset(uint256 indexed fid)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchAdminReset(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistryAdminReset, fid []*big.Int) (event.Subscription, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "AdminReset", fidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistryAdminReset)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "AdminReset", log); err != nil {
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

// ParseAdminReset is a log parse operation binding the contract event 0x8b4b4c6da5b89da518fb865149e01ad2863b48861a8b952e11645f663959fa70.
//
// Solidity: event AdminReset(uint256 indexed fid)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseAdminReset(log types.Log) (*FarcasterIDRegistryAdminReset, error) {
	event := new(FarcasterIDRegistryAdminReset)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "AdminReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistryChangeRecoveryAddressIterator is returned from FilterChangeRecoveryAddress and is used to iterate over the raw logs and unpacked data for ChangeRecoveryAddress events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryChangeRecoveryAddressIterator struct {
	Event *FarcasterIDRegistryChangeRecoveryAddress // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistryChangeRecoveryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistryChangeRecoveryAddress)
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
		it.Event = new(FarcasterIDRegistryChangeRecoveryAddress)
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
func (it *FarcasterIDRegistryChangeRecoveryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistryChangeRecoveryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistryChangeRecoveryAddress represents a ChangeRecoveryAddress event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryChangeRecoveryAddress struct {
	Id       *big.Int
	Recovery common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChangeRecoveryAddress is a free log retrieval operation binding the contract event 0x8e700b803af43e14651431cd73c9fe7d11b131ad797576a70b893ce5766f65c3.
//
// Solidity: event ChangeRecoveryAddress(uint256 indexed id, address indexed recovery)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterChangeRecoveryAddress(opts *bind.FilterOpts, id []*big.Int, recovery []common.Address) (*FarcasterIDRegistryChangeRecoveryAddressIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var recoveryRule []interface{}
	for _, recoveryItem := range recovery {
		recoveryRule = append(recoveryRule, recoveryItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "ChangeRecoveryAddress", idRule, recoveryRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryChangeRecoveryAddressIterator{contract: _FarcasterIDRegistry.contract, event: "ChangeRecoveryAddress", logs: logs, sub: sub}, nil
}

// WatchChangeRecoveryAddress is a free log subscription operation binding the contract event 0x8e700b803af43e14651431cd73c9fe7d11b131ad797576a70b893ce5766f65c3.
//
// Solidity: event ChangeRecoveryAddress(uint256 indexed id, address indexed recovery)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchChangeRecoveryAddress(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistryChangeRecoveryAddress, id []*big.Int, recovery []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var recoveryRule []interface{}
	for _, recoveryItem := range recovery {
		recoveryRule = append(recoveryRule, recoveryItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "ChangeRecoveryAddress", idRule, recoveryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistryChangeRecoveryAddress)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "ChangeRecoveryAddress", log); err != nil {
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

// ParseChangeRecoveryAddress is a log parse operation binding the contract event 0x8e700b803af43e14651431cd73c9fe7d11b131ad797576a70b893ce5766f65c3.
//
// Solidity: event ChangeRecoveryAddress(uint256 indexed id, address indexed recovery)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseChangeRecoveryAddress(log types.Log) (*FarcasterIDRegistryChangeRecoveryAddress, error) {
	event := new(FarcasterIDRegistryChangeRecoveryAddress)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "ChangeRecoveryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistryEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryEIP712DomainChangedIterator struct {
	Event *FarcasterIDRegistryEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistryEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistryEIP712DomainChanged)
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
		it.Event = new(FarcasterIDRegistryEIP712DomainChanged)
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
func (it *FarcasterIDRegistryEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistryEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistryEIP712DomainChanged represents a EIP712DomainChanged event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*FarcasterIDRegistryEIP712DomainChangedIterator, error) {

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryEIP712DomainChangedIterator{contract: _FarcasterIDRegistry.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistryEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistryEIP712DomainChanged)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseEIP712DomainChanged(log types.Log) (*FarcasterIDRegistryEIP712DomainChanged, error) {
	event := new(FarcasterIDRegistryEIP712DomainChanged)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistryFreezeIdGatewayIterator is returned from FilterFreezeIdGateway and is used to iterate over the raw logs and unpacked data for FreezeIdGateway events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryFreezeIdGatewayIterator struct {
	Event *FarcasterIDRegistryFreezeIdGateway // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistryFreezeIdGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistryFreezeIdGateway)
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
		it.Event = new(FarcasterIDRegistryFreezeIdGateway)
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
func (it *FarcasterIDRegistryFreezeIdGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistryFreezeIdGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistryFreezeIdGateway represents a FreezeIdGateway event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryFreezeIdGateway struct {
	IdGateway common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreezeIdGateway is a free log retrieval operation binding the contract event 0x1f54688ee839cb2e57222a4f7482fd67a532a36666748891a7634428b2e8a153.
//
// Solidity: event FreezeIdGateway(address idGateway)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterFreezeIdGateway(opts *bind.FilterOpts) (*FarcasterIDRegistryFreezeIdGatewayIterator, error) {

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "FreezeIdGateway")
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryFreezeIdGatewayIterator{contract: _FarcasterIDRegistry.contract, event: "FreezeIdGateway", logs: logs, sub: sub}, nil
}

// WatchFreezeIdGateway is a free log subscription operation binding the contract event 0x1f54688ee839cb2e57222a4f7482fd67a532a36666748891a7634428b2e8a153.
//
// Solidity: event FreezeIdGateway(address idGateway)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchFreezeIdGateway(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistryFreezeIdGateway) (event.Subscription, error) {

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "FreezeIdGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistryFreezeIdGateway)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "FreezeIdGateway", log); err != nil {
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

// ParseFreezeIdGateway is a log parse operation binding the contract event 0x1f54688ee839cb2e57222a4f7482fd67a532a36666748891a7634428b2e8a153.
//
// Solidity: event FreezeIdGateway(address idGateway)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseFreezeIdGateway(log types.Log) (*FarcasterIDRegistryFreezeIdGateway, error) {
	event := new(FarcasterIDRegistryFreezeIdGateway)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "FreezeIdGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistryMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryMigratedIterator struct {
	Event *FarcasterIDRegistryMigrated // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistryMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistryMigrated)
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
		it.Event = new(FarcasterIDRegistryMigrated)
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
func (it *FarcasterIDRegistryMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistryMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistryMigrated represents a Migrated event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryMigrated struct {
	MigratedAt *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xe4a25c0c2cbe89d6ad8b64c61a7dbdd20d1f781f6023f1ab94ebb7fe0aef6ab8.
//
// Solidity: event Migrated(uint256 indexed migratedAt)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterMigrated(opts *bind.FilterOpts, migratedAt []*big.Int) (*FarcasterIDRegistryMigratedIterator, error) {

	var migratedAtRule []interface{}
	for _, migratedAtItem := range migratedAt {
		migratedAtRule = append(migratedAtRule, migratedAtItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "Migrated", migratedAtRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryMigratedIterator{contract: _FarcasterIDRegistry.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xe4a25c0c2cbe89d6ad8b64c61a7dbdd20d1f781f6023f1ab94ebb7fe0aef6ab8.
//
// Solidity: event Migrated(uint256 indexed migratedAt)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistryMigrated, migratedAt []*big.Int) (event.Subscription, error) {

	var migratedAtRule []interface{}
	for _, migratedAtItem := range migratedAt {
		migratedAtRule = append(migratedAtRule, migratedAtItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "Migrated", migratedAtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistryMigrated)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Migrated", log); err != nil {
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

// ParseMigrated is a log parse operation binding the contract event 0xe4a25c0c2cbe89d6ad8b64c61a7dbdd20d1f781f6023f1ab94ebb7fe0aef6ab8.
//
// Solidity: event Migrated(uint256 indexed migratedAt)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseMigrated(log types.Log) (*FarcasterIDRegistryMigrated, error) {
	event := new(FarcasterIDRegistryMigrated)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Migrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistryOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryOwnershipTransferStartedIterator struct {
	Event *FarcasterIDRegistryOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistryOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistryOwnershipTransferStarted)
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
		it.Event = new(FarcasterIDRegistryOwnershipTransferStarted)
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
func (it *FarcasterIDRegistryOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistryOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistryOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FarcasterIDRegistryOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryOwnershipTransferStartedIterator{contract: _FarcasterIDRegistry.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistryOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistryOwnershipTransferStarted)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseOwnershipTransferStarted(log types.Log) (*FarcasterIDRegistryOwnershipTransferStarted, error) {
	event := new(FarcasterIDRegistryOwnershipTransferStarted)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryOwnershipTransferredIterator struct {
	Event *FarcasterIDRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistryOwnershipTransferred)
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
		it.Event = new(FarcasterIDRegistryOwnershipTransferred)
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
func (it *FarcasterIDRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FarcasterIDRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryOwnershipTransferredIterator{contract: _FarcasterIDRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistryOwnershipTransferred)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*FarcasterIDRegistryOwnershipTransferred, error) {
	event := new(FarcasterIDRegistryOwnershipTransferred)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryPausedIterator struct {
	Event *FarcasterIDRegistryPaused // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistryPaused)
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
		it.Event = new(FarcasterIDRegistryPaused)
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
func (it *FarcasterIDRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistryPaused represents a Paused event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*FarcasterIDRegistryPausedIterator, error) {

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryPausedIterator{contract: _FarcasterIDRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistryPaused)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParsePaused(log types.Log) (*FarcasterIDRegistryPaused, error) {
	event := new(FarcasterIDRegistryPaused)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistryRecoverIterator is returned from FilterRecover and is used to iterate over the raw logs and unpacked data for Recover events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryRecoverIterator struct {
	Event *FarcasterIDRegistryRecover // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistryRecoverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistryRecover)
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
		it.Event = new(FarcasterIDRegistryRecover)
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
func (it *FarcasterIDRegistryRecoverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistryRecoverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistryRecover represents a Recover event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryRecover struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRecover is a free log retrieval operation binding the contract event 0xf6891c84a6c6af32a6d052172a8acc4c631b1d5057ffa2bc1da268b6938ea2da.
//
// Solidity: event Recover(address indexed from, address indexed to, uint256 indexed id)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterRecover(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*FarcasterIDRegistryRecoverIterator, error) {

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

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "Recover", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryRecoverIterator{contract: _FarcasterIDRegistry.contract, event: "Recover", logs: logs, sub: sub}, nil
}

// WatchRecover is a free log subscription operation binding the contract event 0xf6891c84a6c6af32a6d052172a8acc4c631b1d5057ffa2bc1da268b6938ea2da.
//
// Solidity: event Recover(address indexed from, address indexed to, uint256 indexed id)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchRecover(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistryRecover, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "Recover", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistryRecover)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Recover", log); err != nil {
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

// ParseRecover is a log parse operation binding the contract event 0xf6891c84a6c6af32a6d052172a8acc4c631b1d5057ffa2bc1da268b6938ea2da.
//
// Solidity: event Recover(address indexed from, address indexed to, uint256 indexed id)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseRecover(log types.Log) (*FarcasterIDRegistryRecover, error) {
	event := new(FarcasterIDRegistryRecover)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Recover", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistryRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryRegisterIterator struct {
	Event *FarcasterIDRegistryRegister // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistryRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistryRegister)
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
		it.Event = new(FarcasterIDRegistryRegister)
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
func (it *FarcasterIDRegistryRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistryRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistryRegister represents a Register event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryRegister struct {
	To       common.Address
	Id       *big.Int
	Recovery common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0xf2e19a901b0748d8b08e428d0468896a039ac751ec4fec49b44b7b9c28097e45.
//
// Solidity: event Register(address indexed to, uint256 indexed id, address recovery)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterRegister(opts *bind.FilterOpts, to []common.Address, id []*big.Int) (*FarcasterIDRegistryRegisterIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "Register", toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryRegisterIterator{contract: _FarcasterIDRegistry.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0xf2e19a901b0748d8b08e428d0468896a039ac751ec4fec49b44b7b9c28097e45.
//
// Solidity: event Register(address indexed to, uint256 indexed id, address recovery)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistryRegister, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "Register", toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistryRegister)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Register", log); err != nil {
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

// ParseRegister is a log parse operation binding the contract event 0xf2e19a901b0748d8b08e428d0468896a039ac751ec4fec49b44b7b9c28097e45.
//
// Solidity: event Register(address indexed to, uint256 indexed id, address recovery)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseRegister(log types.Log) (*FarcasterIDRegistryRegister, error) {
	event := new(FarcasterIDRegistryRegister)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Register", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistryRemoveIterator is returned from FilterRemove and is used to iterate over the raw logs and unpacked data for Remove events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryRemoveIterator struct {
	Event *FarcasterIDRegistryRemove // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistryRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistryRemove)
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
		it.Event = new(FarcasterIDRegistryRemove)
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
func (it *FarcasterIDRegistryRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistryRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistryRemove represents a Remove event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryRemove struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemove is a free log retrieval operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterRemove(opts *bind.FilterOpts, guardian []common.Address) (*FarcasterIDRegistryRemoveIterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "Remove", guardianRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryRemoveIterator{contract: _FarcasterIDRegistry.contract, event: "Remove", logs: logs, sub: sub}, nil
}

// WatchRemove is a free log subscription operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchRemove(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistryRemove, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "Remove", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistryRemove)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Remove", log); err != nil {
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

// ParseRemove is a log parse operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseRemove(log types.Log) (*FarcasterIDRegistryRemove, error) {
	event := new(FarcasterIDRegistryRemove)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Remove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistrySetIdCounterIterator is returned from FilterSetIdCounter and is used to iterate over the raw logs and unpacked data for SetIdCounter events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistrySetIdCounterIterator struct {
	Event *FarcasterIDRegistrySetIdCounter // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistrySetIdCounterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistrySetIdCounter)
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
		it.Event = new(FarcasterIDRegistrySetIdCounter)
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
func (it *FarcasterIDRegistrySetIdCounterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistrySetIdCounterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistrySetIdCounter represents a SetIdCounter event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistrySetIdCounter struct {
	OldCounter *big.Int
	NewCounter *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetIdCounter is a free log retrieval operation binding the contract event 0x562044dce594b5c0ac495e6cf3717dbef4dcc96bf978ff452457bfccd68a4eed.
//
// Solidity: event SetIdCounter(uint256 oldCounter, uint256 newCounter)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterSetIdCounter(opts *bind.FilterOpts) (*FarcasterIDRegistrySetIdCounterIterator, error) {

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "SetIdCounter")
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistrySetIdCounterIterator{contract: _FarcasterIDRegistry.contract, event: "SetIdCounter", logs: logs, sub: sub}, nil
}

// WatchSetIdCounter is a free log subscription operation binding the contract event 0x562044dce594b5c0ac495e6cf3717dbef4dcc96bf978ff452457bfccd68a4eed.
//
// Solidity: event SetIdCounter(uint256 oldCounter, uint256 newCounter)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchSetIdCounter(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistrySetIdCounter) (event.Subscription, error) {

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "SetIdCounter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistrySetIdCounter)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "SetIdCounter", log); err != nil {
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

// ParseSetIdCounter is a log parse operation binding the contract event 0x562044dce594b5c0ac495e6cf3717dbef4dcc96bf978ff452457bfccd68a4eed.
//
// Solidity: event SetIdCounter(uint256 oldCounter, uint256 newCounter)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseSetIdCounter(log types.Log) (*FarcasterIDRegistrySetIdCounter, error) {
	event := new(FarcasterIDRegistrySetIdCounter)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "SetIdCounter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistrySetIdGatewayIterator is returned from FilterSetIdGateway and is used to iterate over the raw logs and unpacked data for SetIdGateway events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistrySetIdGatewayIterator struct {
	Event *FarcasterIDRegistrySetIdGateway // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistrySetIdGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistrySetIdGateway)
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
		it.Event = new(FarcasterIDRegistrySetIdGateway)
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
func (it *FarcasterIDRegistrySetIdGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistrySetIdGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistrySetIdGateway represents a SetIdGateway event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistrySetIdGateway struct {
	OldIdGateway common.Address
	NewIdGateway common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetIdGateway is a free log retrieval operation binding the contract event 0x306b123921c19a8629c68977f4dfea9ef9d5a6dedfafcd0d4a70ac6c9b763ac2.
//
// Solidity: event SetIdGateway(address oldIdGateway, address newIdGateway)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterSetIdGateway(opts *bind.FilterOpts) (*FarcasterIDRegistrySetIdGatewayIterator, error) {

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "SetIdGateway")
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistrySetIdGatewayIterator{contract: _FarcasterIDRegistry.contract, event: "SetIdGateway", logs: logs, sub: sub}, nil
}

// WatchSetIdGateway is a free log subscription operation binding the contract event 0x306b123921c19a8629c68977f4dfea9ef9d5a6dedfafcd0d4a70ac6c9b763ac2.
//
// Solidity: event SetIdGateway(address oldIdGateway, address newIdGateway)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchSetIdGateway(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistrySetIdGateway) (event.Subscription, error) {

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "SetIdGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistrySetIdGateway)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "SetIdGateway", log); err != nil {
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

// ParseSetIdGateway is a log parse operation binding the contract event 0x306b123921c19a8629c68977f4dfea9ef9d5a6dedfafcd0d4a70ac6c9b763ac2.
//
// Solidity: event SetIdGateway(address oldIdGateway, address newIdGateway)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseSetIdGateway(log types.Log) (*FarcasterIDRegistrySetIdGateway, error) {
	event := new(FarcasterIDRegistrySetIdGateway)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "SetIdGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistrySetMigratorIterator is returned from FilterSetMigrator and is used to iterate over the raw logs and unpacked data for SetMigrator events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistrySetMigratorIterator struct {
	Event *FarcasterIDRegistrySetMigrator // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistrySetMigratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistrySetMigrator)
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
		it.Event = new(FarcasterIDRegistrySetMigrator)
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
func (it *FarcasterIDRegistrySetMigratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistrySetMigratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistrySetMigrator represents a SetMigrator event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistrySetMigrator struct {
	OldMigrator common.Address
	NewMigrator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetMigrator is a free log retrieval operation binding the contract event 0xd8ad954fe808212ab9ed7139873e40807dff7995fe36e3d6cdeb8fa00fcebf10.
//
// Solidity: event SetMigrator(address oldMigrator, address newMigrator)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterSetMigrator(opts *bind.FilterOpts) (*FarcasterIDRegistrySetMigratorIterator, error) {

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "SetMigrator")
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistrySetMigratorIterator{contract: _FarcasterIDRegistry.contract, event: "SetMigrator", logs: logs, sub: sub}, nil
}

// WatchSetMigrator is a free log subscription operation binding the contract event 0xd8ad954fe808212ab9ed7139873e40807dff7995fe36e3d6cdeb8fa00fcebf10.
//
// Solidity: event SetMigrator(address oldMigrator, address newMigrator)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchSetMigrator(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistrySetMigrator) (event.Subscription, error) {

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "SetMigrator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistrySetMigrator)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "SetMigrator", log); err != nil {
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

// ParseSetMigrator is a log parse operation binding the contract event 0xd8ad954fe808212ab9ed7139873e40807dff7995fe36e3d6cdeb8fa00fcebf10.
//
// Solidity: event SetMigrator(address oldMigrator, address newMigrator)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseSetMigrator(log types.Log) (*FarcasterIDRegistrySetMigrator, error) {
	event := new(FarcasterIDRegistrySetMigrator)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "SetMigrator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryTransferIterator struct {
	Event *FarcasterIDRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistryTransfer)
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
		it.Event = new(FarcasterIDRegistryTransfer)
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
func (it *FarcasterIDRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistryTransfer represents a Transfer event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryTransfer struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*FarcasterIDRegistryTransferIterator, error) {

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

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryTransferIterator{contract: _FarcasterIDRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistryTransfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistryTransfer)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseTransfer(log types.Log) (*FarcasterIDRegistryTransfer, error) {
	event := new(FarcasterIDRegistryTransfer)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterIDRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryUnpausedIterator struct {
	Event *FarcasterIDRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *FarcasterIDRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterIDRegistryUnpaused)
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
		it.Event = new(FarcasterIDRegistryUnpaused)
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
func (it *FarcasterIDRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterIDRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterIDRegistryUnpaused represents a Unpaused event raised by the FarcasterIDRegistry contract.
type FarcasterIDRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*FarcasterIDRegistryUnpausedIterator, error) {

	logs, sub, err := _FarcasterIDRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &FarcasterIDRegistryUnpausedIterator{contract: _FarcasterIDRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *FarcasterIDRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _FarcasterIDRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterIDRegistryUnpaused)
				if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_FarcasterIDRegistry *FarcasterIDRegistryFilterer) ParseUnpaused(log types.Log) (*FarcasterIDRegistryUnpaused, error) {
	event := new(FarcasterIDRegistryUnpaused)
	if err := _FarcasterIDRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
