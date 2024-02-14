// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FarcasterKeyRegistry

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

// IKeyRegistryBulkAddData is an auto generated low-level Go binding around an user-defined struct.
type IKeyRegistryBulkAddData struct {
	Fid  *big.Int
	Keys []IKeyRegistryBulkAddKey
}

// IKeyRegistryBulkAddKey is an auto generated low-level Go binding around an user-defined struct.
type IKeyRegistryBulkAddKey struct {
	Key      []byte
	Metadata []byte
}

// IKeyRegistryBulkResetData is an auto generated low-level Go binding around an user-defined struct.
type IKeyRegistryBulkResetData struct {
	Fid  *big.Int
	Keys [][]byte
}

// IKeyRegistryKeyData is an auto generated low-level Go binding around an user-defined struct.
type IKeyRegistryKeyData struct {
	State   uint8
	KeyType uint32
}

// FarcasterKeyRegistryMetaData contains all meta data concerning the FarcasterKeyRegistry contract.
var FarcasterKeyRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_idRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxKeysPerFid\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyMigrated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedsMaximum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GatewayFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxKeys\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMetadata\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMetadataType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGuardian\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyMigrator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionRevoked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"}],\"name\":\"ValidatorNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"keyBytes\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"keyBytes\",\"type\":\"bytes\"}],\"name\":\"AdminReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"keyGateway\",\"type\":\"address\"}],\"name\":\"FreezeKeyGateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"migratedAt\",\"type\":\"uint256\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"keyBytes\",\"type\":\"bytes\"}],\"name\":\"Remove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"Remove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldIdRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIdRegistry\",\"type\":\"address\"}],\"name\":\"SetIdRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldKeyGateway\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newKeyGateway\",\"type\":\"address\"}],\"name\":\"SetKeyGateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMax\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"SetMaxKeysPerFid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldMigrator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newMigrator\",\"type\":\"address\"}],\"name\":\"SetMigrator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValidator\",\"type\":\"address\"}],\"name\":\"SetValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"REMOVE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fidOwner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"addGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structIKeyRegistry.BulkAddKey[]\",\"name\":\"keys\",\"type\":\"tuple[]\"}],\"internalType\":\"structIKeyRegistry.BulkAddData[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"name\":\"bulkAddKeysForMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"keys\",\"type\":\"bytes[]\"}],\"internalType\":\"structIKeyRegistry.BulkResetData[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"name\":\"bulkResetKeysForMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparatorV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freezeKeyGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatewayFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gracePeriod\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"guardians\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isGuardian\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"structHash\",\"type\":\"bytes32\"}],\"name\":\"hashTypedDataV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"idRegistry\",\"outputs\":[{\"internalType\":\"contractIdRegistryLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMigrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"enumIKeyRegistry.KeyState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"keyAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"keyDataOf\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIKeyRegistry.KeyState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"}],\"internalType\":\"structIKeyRegistry.KeyData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keyGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"keys\",\"outputs\":[{\"internalType\":\"enumIKeyRegistry.KeyState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"enumIKeyRegistry.KeyState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"keysOf\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"enumIKeyRegistry.KeyState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"}],\"name\":\"keysOf\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"page\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"nextIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxKeysPerFid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migratedAt\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fidOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"removeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"removeGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_idRegistry\",\"type\":\"address\"}],\"name\":\"setIdRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keyGateway\",\"type\":\"address\"}],\"name\":\"setKeyGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxKeysPerFid\",\"type\":\"uint256\"}],\"name\":\"setMaxKeysPerFid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"},{\"internalType\":\"contractIMetadataValidator\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fid\",\"type\":\"uint256\"},{\"internalType\":\"enumIKeyRegistry.KeyState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"totalKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"keyType\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"metadataType\",\"type\":\"uint8\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"contractIMetadataValidator\",\"name\":\"validator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FarcasterKeyRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use FarcasterKeyRegistryMetaData.ABI instead.
var FarcasterKeyRegistryABI = FarcasterKeyRegistryMetaData.ABI

// FarcasterKeyRegistry is an auto generated Go binding around an Ethereum contract.
type FarcasterKeyRegistry struct {
	FarcasterKeyRegistryCaller     // Read-only binding to the contract
	FarcasterKeyRegistryTransactor // Write-only binding to the contract
	FarcasterKeyRegistryFilterer   // Log filterer for contract events
}

// FarcasterKeyRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FarcasterKeyRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarcasterKeyRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FarcasterKeyRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarcasterKeyRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FarcasterKeyRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarcasterKeyRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FarcasterKeyRegistrySession struct {
	Contract     *FarcasterKeyRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FarcasterKeyRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FarcasterKeyRegistryCallerSession struct {
	Contract *FarcasterKeyRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// FarcasterKeyRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FarcasterKeyRegistryTransactorSession struct {
	Contract     *FarcasterKeyRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// FarcasterKeyRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FarcasterKeyRegistryRaw struct {
	Contract *FarcasterKeyRegistry // Generic contract binding to access the raw methods on
}

// FarcasterKeyRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FarcasterKeyRegistryCallerRaw struct {
	Contract *FarcasterKeyRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// FarcasterKeyRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FarcasterKeyRegistryTransactorRaw struct {
	Contract *FarcasterKeyRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFarcasterKeyRegistry creates a new instance of FarcasterKeyRegistry, bound to a specific deployed contract.
func NewFarcasterKeyRegistry(address common.Address, backend bind.ContractBackend) (*FarcasterKeyRegistry, error) {
	contract, err := bindFarcasterKeyRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistry{FarcasterKeyRegistryCaller: FarcasterKeyRegistryCaller{contract: contract}, FarcasterKeyRegistryTransactor: FarcasterKeyRegistryTransactor{contract: contract}, FarcasterKeyRegistryFilterer: FarcasterKeyRegistryFilterer{contract: contract}}, nil
}

// NewFarcasterKeyRegistryCaller creates a new read-only instance of FarcasterKeyRegistry, bound to a specific deployed contract.
func NewFarcasterKeyRegistryCaller(address common.Address, caller bind.ContractCaller) (*FarcasterKeyRegistryCaller, error) {
	contract, err := bindFarcasterKeyRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistryCaller{contract: contract}, nil
}

// NewFarcasterKeyRegistryTransactor creates a new write-only instance of FarcasterKeyRegistry, bound to a specific deployed contract.
func NewFarcasterKeyRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*FarcasterKeyRegistryTransactor, error) {
	contract, err := bindFarcasterKeyRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistryTransactor{contract: contract}, nil
}

// NewFarcasterKeyRegistryFilterer creates a new log filterer instance of FarcasterKeyRegistry, bound to a specific deployed contract.
func NewFarcasterKeyRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*FarcasterKeyRegistryFilterer, error) {
	contract, err := bindFarcasterKeyRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistryFilterer{contract: contract}, nil
}

// bindFarcasterKeyRegistry binds a generic wrapper to an already deployed contract.
func bindFarcasterKeyRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FarcasterKeyRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FarcasterKeyRegistry *FarcasterKeyRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FarcasterKeyRegistry.Contract.FarcasterKeyRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FarcasterKeyRegistry *FarcasterKeyRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.FarcasterKeyRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FarcasterKeyRegistry *FarcasterKeyRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.FarcasterKeyRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FarcasterKeyRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.contract.Transact(opts, method, params...)
}

// REMOVETYPEHASH is a free data retrieval call binding the contract method 0xb5775561.
//
// Solidity: function REMOVE_TYPEHASH() view returns(bytes32)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) REMOVETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "REMOVE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REMOVETYPEHASH is a free data retrieval call binding the contract method 0xb5775561.
//
// Solidity: function REMOVE_TYPEHASH() view returns(bytes32)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) REMOVETYPEHASH() ([32]byte, error) {
	return _FarcasterKeyRegistry.Contract.REMOVETYPEHASH(&_FarcasterKeyRegistry.CallOpts)
}

// REMOVETYPEHASH is a free data retrieval call binding the contract method 0xb5775561.
//
// Solidity: function REMOVE_TYPEHASH() view returns(bytes32)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) REMOVETYPEHASH() ([32]byte, error) {
	return _FarcasterKeyRegistry.Contract.REMOVETYPEHASH(&_FarcasterKeyRegistry.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) VERSION() (string, error) {
	return _FarcasterKeyRegistry.Contract.VERSION(&_FarcasterKeyRegistry.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) VERSION() (string, error) {
	return _FarcasterKeyRegistry.Contract.VERSION(&_FarcasterKeyRegistry.CallOpts)
}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) DomainSeparatorV4(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "domainSeparatorV4")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) DomainSeparatorV4() ([32]byte, error) {
	return _FarcasterKeyRegistry.Contract.DomainSeparatorV4(&_FarcasterKeyRegistry.CallOpts)
}

// DomainSeparatorV4 is a free data retrieval call binding the contract method 0x78e890ba.
//
// Solidity: function domainSeparatorV4() view returns(bytes32)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) DomainSeparatorV4() ([32]byte, error) {
	return _FarcasterKeyRegistry.Contract.DomainSeparatorV4(&_FarcasterKeyRegistry.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "eip712Domain")

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
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _FarcasterKeyRegistry.Contract.Eip712Domain(&_FarcasterKeyRegistry.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _FarcasterKeyRegistry.Contract.Eip712Domain(&_FarcasterKeyRegistry.CallOpts)
}

// GatewayFrozen is a free data retrieval call binding the contract method 0x95e7549f.
//
// Solidity: function gatewayFrozen() view returns(bool)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) GatewayFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "gatewayFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GatewayFrozen is a free data retrieval call binding the contract method 0x95e7549f.
//
// Solidity: function gatewayFrozen() view returns(bool)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) GatewayFrozen() (bool, error) {
	return _FarcasterKeyRegistry.Contract.GatewayFrozen(&_FarcasterKeyRegistry.CallOpts)
}

// GatewayFrozen is a free data retrieval call binding the contract method 0x95e7549f.
//
// Solidity: function gatewayFrozen() view returns(bool)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) GatewayFrozen() (bool, error) {
	return _FarcasterKeyRegistry.Contract.GatewayFrozen(&_FarcasterKeyRegistry.CallOpts)
}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() view returns(uint24)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) GracePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "gracePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() view returns(uint24)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) GracePeriod() (*big.Int, error) {
	return _FarcasterKeyRegistry.Contract.GracePeriod(&_FarcasterKeyRegistry.CallOpts)
}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() view returns(uint24)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) GracePeriod() (*big.Int, error) {
	return _FarcasterKeyRegistry.Contract.GracePeriod(&_FarcasterKeyRegistry.CallOpts)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) Guardians(opts *bind.CallOpts, guardian common.Address) (bool, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "guardians", guardian)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) Guardians(guardian common.Address) (bool, error) {
	return _FarcasterKeyRegistry.Contract.Guardians(&_FarcasterKeyRegistry.CallOpts, guardian)
}

// Guardians is a free data retrieval call binding the contract method 0x0633b14a.
//
// Solidity: function guardians(address guardian) view returns(bool isGuardian)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) Guardians(guardian common.Address) (bool, error) {
	return _FarcasterKeyRegistry.Contract.Guardians(&_FarcasterKeyRegistry.CallOpts, guardian)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) HashTypedDataV4(opts *bind.CallOpts, structHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "hashTypedDataV4", structHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _FarcasterKeyRegistry.Contract.HashTypedDataV4(&_FarcasterKeyRegistry.CallOpts, structHash)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _FarcasterKeyRegistry.Contract.HashTypedDataV4(&_FarcasterKeyRegistry.CallOpts, structHash)
}

// IdRegistry is a free data retrieval call binding the contract method 0x0aa13b8c.
//
// Solidity: function idRegistry() view returns(address)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) IdRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "idRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdRegistry is a free data retrieval call binding the contract method 0x0aa13b8c.
//
// Solidity: function idRegistry() view returns(address)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) IdRegistry() (common.Address, error) {
	return _FarcasterKeyRegistry.Contract.IdRegistry(&_FarcasterKeyRegistry.CallOpts)
}

// IdRegistry is a free data retrieval call binding the contract method 0x0aa13b8c.
//
// Solidity: function idRegistry() view returns(address)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) IdRegistry() (common.Address, error) {
	return _FarcasterKeyRegistry.Contract.IdRegistry(&_FarcasterKeyRegistry.CallOpts)
}

// IsMigrated is a free data retrieval call binding the contract method 0xb06faf62.
//
// Solidity: function isMigrated() view returns(bool)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) IsMigrated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "isMigrated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMigrated is a free data retrieval call binding the contract method 0xb06faf62.
//
// Solidity: function isMigrated() view returns(bool)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) IsMigrated() (bool, error) {
	return _FarcasterKeyRegistry.Contract.IsMigrated(&_FarcasterKeyRegistry.CallOpts)
}

// IsMigrated is a free data retrieval call binding the contract method 0xb06faf62.
//
// Solidity: function isMigrated() view returns(bool)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) IsMigrated() (bool, error) {
	return _FarcasterKeyRegistry.Contract.IsMigrated(&_FarcasterKeyRegistry.CallOpts)
}

// KeyAt is a free data retrieval call binding the contract method 0x0ea9442c.
//
// Solidity: function keyAt(uint256 fid, uint8 state, uint256 index) view returns(bytes)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) KeyAt(opts *bind.CallOpts, fid *big.Int, state uint8, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "keyAt", fid, state, index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// KeyAt is a free data retrieval call binding the contract method 0x0ea9442c.
//
// Solidity: function keyAt(uint256 fid, uint8 state, uint256 index) view returns(bytes)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) KeyAt(fid *big.Int, state uint8, index *big.Int) ([]byte, error) {
	return _FarcasterKeyRegistry.Contract.KeyAt(&_FarcasterKeyRegistry.CallOpts, fid, state, index)
}

// KeyAt is a free data retrieval call binding the contract method 0x0ea9442c.
//
// Solidity: function keyAt(uint256 fid, uint8 state, uint256 index) view returns(bytes)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) KeyAt(fid *big.Int, state uint8, index *big.Int) ([]byte, error) {
	return _FarcasterKeyRegistry.Contract.KeyAt(&_FarcasterKeyRegistry.CallOpts, fid, state, index)
}

// KeyDataOf is a free data retrieval call binding the contract method 0xac34cc5a.
//
// Solidity: function keyDataOf(uint256 fid, bytes key) view returns((uint8,uint32))
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) KeyDataOf(opts *bind.CallOpts, fid *big.Int, key []byte) (IKeyRegistryKeyData, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "keyDataOf", fid, key)

	if err != nil {
		return *new(IKeyRegistryKeyData), err
	}

	out0 := *abi.ConvertType(out[0], new(IKeyRegistryKeyData)).(*IKeyRegistryKeyData)

	return out0, err

}

// KeyDataOf is a free data retrieval call binding the contract method 0xac34cc5a.
//
// Solidity: function keyDataOf(uint256 fid, bytes key) view returns((uint8,uint32))
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) KeyDataOf(fid *big.Int, key []byte) (IKeyRegistryKeyData, error) {
	return _FarcasterKeyRegistry.Contract.KeyDataOf(&_FarcasterKeyRegistry.CallOpts, fid, key)
}

// KeyDataOf is a free data retrieval call binding the contract method 0xac34cc5a.
//
// Solidity: function keyDataOf(uint256 fid, bytes key) view returns((uint8,uint32))
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) KeyDataOf(fid *big.Int, key []byte) (IKeyRegistryKeyData, error) {
	return _FarcasterKeyRegistry.Contract.KeyDataOf(&_FarcasterKeyRegistry.CallOpts, fid, key)
}

// KeyGateway is a free data retrieval call binding the contract method 0x80334737.
//
// Solidity: function keyGateway() view returns(address)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) KeyGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "keyGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyGateway is a free data retrieval call binding the contract method 0x80334737.
//
// Solidity: function keyGateway() view returns(address)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) KeyGateway() (common.Address, error) {
	return _FarcasterKeyRegistry.Contract.KeyGateway(&_FarcasterKeyRegistry.CallOpts)
}

// KeyGateway is a free data retrieval call binding the contract method 0x80334737.
//
// Solidity: function keyGateway() view returns(address)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) KeyGateway() (common.Address, error) {
	return _FarcasterKeyRegistry.Contract.KeyGateway(&_FarcasterKeyRegistry.CallOpts)
}

// Keys is a free data retrieval call binding the contract method 0x0e24e34c.
//
// Solidity: function keys(uint256 fid, bytes key) view returns(uint8 state, uint32 keyType)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) Keys(opts *bind.CallOpts, fid *big.Int, key []byte) (struct {
	State   uint8
	KeyType uint32
}, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "keys", fid, key)

	outstruct := new(struct {
		State   uint8
		KeyType uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.State = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.KeyType = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// Keys is a free data retrieval call binding the contract method 0x0e24e34c.
//
// Solidity: function keys(uint256 fid, bytes key) view returns(uint8 state, uint32 keyType)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) Keys(fid *big.Int, key []byte) (struct {
	State   uint8
	KeyType uint32
}, error) {
	return _FarcasterKeyRegistry.Contract.Keys(&_FarcasterKeyRegistry.CallOpts, fid, key)
}

// Keys is a free data retrieval call binding the contract method 0x0e24e34c.
//
// Solidity: function keys(uint256 fid, bytes key) view returns(uint8 state, uint32 keyType)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) Keys(fid *big.Int, key []byte) (struct {
	State   uint8
	KeyType uint32
}, error) {
	return _FarcasterKeyRegistry.Contract.Keys(&_FarcasterKeyRegistry.CallOpts, fid, key)
}

// KeysOf is a free data retrieval call binding the contract method 0x1f64222f.
//
// Solidity: function keysOf(uint256 fid, uint8 state) view returns(bytes[])
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) KeysOf(opts *bind.CallOpts, fid *big.Int, state uint8) ([][]byte, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "keysOf", fid, state)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// KeysOf is a free data retrieval call binding the contract method 0x1f64222f.
//
// Solidity: function keysOf(uint256 fid, uint8 state) view returns(bytes[])
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) KeysOf(fid *big.Int, state uint8) ([][]byte, error) {
	return _FarcasterKeyRegistry.Contract.KeysOf(&_FarcasterKeyRegistry.CallOpts, fid, state)
}

// KeysOf is a free data retrieval call binding the contract method 0x1f64222f.
//
// Solidity: function keysOf(uint256 fid, uint8 state) view returns(bytes[])
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) KeysOf(fid *big.Int, state uint8) ([][]byte, error) {
	return _FarcasterKeyRegistry.Contract.KeysOf(&_FarcasterKeyRegistry.CallOpts, fid, state)
}

// KeysOf0 is a free data retrieval call binding the contract method 0xf27995e3.
//
// Solidity: function keysOf(uint256 fid, uint8 state, uint256 startIdx, uint256 batchSize) view returns(bytes[] page, uint256 nextIdx)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) KeysOf0(opts *bind.CallOpts, fid *big.Int, state uint8, startIdx *big.Int, batchSize *big.Int) (struct {
	Page    [][]byte
	NextIdx *big.Int
}, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "keysOf0", fid, state, startIdx, batchSize)

	outstruct := new(struct {
		Page    [][]byte
		NextIdx *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Page = *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	outstruct.NextIdx = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// KeysOf0 is a free data retrieval call binding the contract method 0xf27995e3.
//
// Solidity: function keysOf(uint256 fid, uint8 state, uint256 startIdx, uint256 batchSize) view returns(bytes[] page, uint256 nextIdx)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) KeysOf0(fid *big.Int, state uint8, startIdx *big.Int, batchSize *big.Int) (struct {
	Page    [][]byte
	NextIdx *big.Int
}, error) {
	return _FarcasterKeyRegistry.Contract.KeysOf0(&_FarcasterKeyRegistry.CallOpts, fid, state, startIdx, batchSize)
}

// KeysOf0 is a free data retrieval call binding the contract method 0xf27995e3.
//
// Solidity: function keysOf(uint256 fid, uint8 state, uint256 startIdx, uint256 batchSize) view returns(bytes[] page, uint256 nextIdx)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) KeysOf0(fid *big.Int, state uint8, startIdx *big.Int, batchSize *big.Int) (struct {
	Page    [][]byte
	NextIdx *big.Int
}, error) {
	return _FarcasterKeyRegistry.Contract.KeysOf0(&_FarcasterKeyRegistry.CallOpts, fid, state, startIdx, batchSize)
}

// MaxKeysPerFid is a free data retrieval call binding the contract method 0xe33acf38.
//
// Solidity: function maxKeysPerFid() view returns(uint256)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) MaxKeysPerFid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "maxKeysPerFid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxKeysPerFid is a free data retrieval call binding the contract method 0xe33acf38.
//
// Solidity: function maxKeysPerFid() view returns(uint256)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) MaxKeysPerFid() (*big.Int, error) {
	return _FarcasterKeyRegistry.Contract.MaxKeysPerFid(&_FarcasterKeyRegistry.CallOpts)
}

// MaxKeysPerFid is a free data retrieval call binding the contract method 0xe33acf38.
//
// Solidity: function maxKeysPerFid() view returns(uint256)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) MaxKeysPerFid() (*big.Int, error) {
	return _FarcasterKeyRegistry.Contract.MaxKeysPerFid(&_FarcasterKeyRegistry.CallOpts)
}

// MigratedAt is a free data retrieval call binding the contract method 0x8b21e484.
//
// Solidity: function migratedAt() view returns(uint40)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) MigratedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "migratedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MigratedAt is a free data retrieval call binding the contract method 0x8b21e484.
//
// Solidity: function migratedAt() view returns(uint40)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) MigratedAt() (*big.Int, error) {
	return _FarcasterKeyRegistry.Contract.MigratedAt(&_FarcasterKeyRegistry.CallOpts)
}

// MigratedAt is a free data retrieval call binding the contract method 0x8b21e484.
//
// Solidity: function migratedAt() view returns(uint40)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) MigratedAt() (*big.Int, error) {
	return _FarcasterKeyRegistry.Contract.MigratedAt(&_FarcasterKeyRegistry.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) Migrator() (common.Address, error) {
	return _FarcasterKeyRegistry.Contract.Migrator(&_FarcasterKeyRegistry.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) Migrator() (common.Address, error) {
	return _FarcasterKeyRegistry.Contract.Migrator(&_FarcasterKeyRegistry.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) Nonces(owner common.Address) (*big.Int, error) {
	return _FarcasterKeyRegistry.Contract.Nonces(&_FarcasterKeyRegistry.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _FarcasterKeyRegistry.Contract.Nonces(&_FarcasterKeyRegistry.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) Owner() (common.Address, error) {
	return _FarcasterKeyRegistry.Contract.Owner(&_FarcasterKeyRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) Owner() (common.Address, error) {
	return _FarcasterKeyRegistry.Contract.Owner(&_FarcasterKeyRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) Paused() (bool, error) {
	return _FarcasterKeyRegistry.Contract.Paused(&_FarcasterKeyRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) Paused() (bool, error) {
	return _FarcasterKeyRegistry.Contract.Paused(&_FarcasterKeyRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) PendingOwner() (common.Address, error) {
	return _FarcasterKeyRegistry.Contract.PendingOwner(&_FarcasterKeyRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) PendingOwner() (common.Address, error) {
	return _FarcasterKeyRegistry.Contract.PendingOwner(&_FarcasterKeyRegistry.CallOpts)
}

// TotalKeys is a free data retrieval call binding the contract method 0x6840b75e.
//
// Solidity: function totalKeys(uint256 fid, uint8 state) view returns(uint256)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) TotalKeys(opts *bind.CallOpts, fid *big.Int, state uint8) (*big.Int, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "totalKeys", fid, state)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalKeys is a free data retrieval call binding the contract method 0x6840b75e.
//
// Solidity: function totalKeys(uint256 fid, uint8 state) view returns(uint256)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) TotalKeys(fid *big.Int, state uint8) (*big.Int, error) {
	return _FarcasterKeyRegistry.Contract.TotalKeys(&_FarcasterKeyRegistry.CallOpts, fid, state)
}

// TotalKeys is a free data retrieval call binding the contract method 0x6840b75e.
//
// Solidity: function totalKeys(uint256 fid, uint8 state) view returns(uint256)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) TotalKeys(fid *big.Int, state uint8) (*big.Int, error) {
	return _FarcasterKeyRegistry.Contract.TotalKeys(&_FarcasterKeyRegistry.CallOpts, fid, state)
}

// Validators is a free data retrieval call binding the contract method 0xd8810395.
//
// Solidity: function validators(uint32 keyType, uint8 metadataType) view returns(address validator)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCaller) Validators(opts *bind.CallOpts, keyType uint32, metadataType uint8) (common.Address, error) {
	var out []interface{}
	err := _FarcasterKeyRegistry.contract.Call(opts, &out, "validators", keyType, metadataType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0xd8810395.
//
// Solidity: function validators(uint32 keyType, uint8 metadataType) view returns(address validator)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) Validators(keyType uint32, metadataType uint8) (common.Address, error) {
	return _FarcasterKeyRegistry.Contract.Validators(&_FarcasterKeyRegistry.CallOpts, keyType, metadataType)
}

// Validators is a free data retrieval call binding the contract method 0xd8810395.
//
// Solidity: function validators(uint32 keyType, uint8 metadataType) view returns(address validator)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryCallerSession) Validators(keyType uint32, metadataType uint8) (common.Address, error) {
	return _FarcasterKeyRegistry.Contract.Validators(&_FarcasterKeyRegistry.CallOpts, keyType, metadataType)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.AcceptOwnership(&_FarcasterKeyRegistry.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.AcceptOwnership(&_FarcasterKeyRegistry.TransactOpts)
}

// Add is a paid mutator transaction binding the contract method 0x207e3d69.
//
// Solidity: function add(address fidOwner, uint32 keyType, bytes key, uint8 metadataType, bytes metadata) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) Add(opts *bind.TransactOpts, fidOwner common.Address, keyType uint32, key []byte, metadataType uint8, metadata []byte) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "add", fidOwner, keyType, key, metadataType, metadata)
}

// Add is a paid mutator transaction binding the contract method 0x207e3d69.
//
// Solidity: function add(address fidOwner, uint32 keyType, bytes key, uint8 metadataType, bytes metadata) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) Add(fidOwner common.Address, keyType uint32, key []byte, metadataType uint8, metadata []byte) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.Add(&_FarcasterKeyRegistry.TransactOpts, fidOwner, keyType, key, metadataType, metadata)
}

// Add is a paid mutator transaction binding the contract method 0x207e3d69.
//
// Solidity: function add(address fidOwner, uint32 keyType, bytes key, uint8 metadataType, bytes metadata) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) Add(fidOwner common.Address, keyType uint32, key []byte, metadataType uint8, metadata []byte) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.Add(&_FarcasterKeyRegistry.TransactOpts, fidOwner, keyType, key, metadataType, metadata)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) AddGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "addGuardian", guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.AddGuardian(&_FarcasterKeyRegistry.TransactOpts, guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xa526d83b.
//
// Solidity: function addGuardian(address guardian) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) AddGuardian(guardian common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.AddGuardian(&_FarcasterKeyRegistry.TransactOpts, guardian)
}

// BulkAddKeysForMigration is a paid mutator transaction binding the contract method 0x708e9c70.
//
// Solidity: function bulkAddKeysForMigration((uint256,(bytes,bytes)[])[] items) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) BulkAddKeysForMigration(opts *bind.TransactOpts, items []IKeyRegistryBulkAddData) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "bulkAddKeysForMigration", items)
}

// BulkAddKeysForMigration is a paid mutator transaction binding the contract method 0x708e9c70.
//
// Solidity: function bulkAddKeysForMigration((uint256,(bytes,bytes)[])[] items) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) BulkAddKeysForMigration(items []IKeyRegistryBulkAddData) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.BulkAddKeysForMigration(&_FarcasterKeyRegistry.TransactOpts, items)
}

// BulkAddKeysForMigration is a paid mutator transaction binding the contract method 0x708e9c70.
//
// Solidity: function bulkAddKeysForMigration((uint256,(bytes,bytes)[])[] items) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) BulkAddKeysForMigration(items []IKeyRegistryBulkAddData) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.BulkAddKeysForMigration(&_FarcasterKeyRegistry.TransactOpts, items)
}

// BulkResetKeysForMigration is a paid mutator transaction binding the contract method 0x46b3f429.
//
// Solidity: function bulkResetKeysForMigration((uint256,bytes[])[] items) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) BulkResetKeysForMigration(opts *bind.TransactOpts, items []IKeyRegistryBulkResetData) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "bulkResetKeysForMigration", items)
}

// BulkResetKeysForMigration is a paid mutator transaction binding the contract method 0x46b3f429.
//
// Solidity: function bulkResetKeysForMigration((uint256,bytes[])[] items) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) BulkResetKeysForMigration(items []IKeyRegistryBulkResetData) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.BulkResetKeysForMigration(&_FarcasterKeyRegistry.TransactOpts, items)
}

// BulkResetKeysForMigration is a paid mutator transaction binding the contract method 0x46b3f429.
//
// Solidity: function bulkResetKeysForMigration((uint256,bytes[])[] items) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) BulkResetKeysForMigration(items []IKeyRegistryBulkResetData) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.BulkResetKeysForMigration(&_FarcasterKeyRegistry.TransactOpts, items)
}

// FreezeKeyGateway is a paid mutator transaction binding the contract method 0x47cf8d4d.
//
// Solidity: function freezeKeyGateway() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) FreezeKeyGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "freezeKeyGateway")
}

// FreezeKeyGateway is a paid mutator transaction binding the contract method 0x47cf8d4d.
//
// Solidity: function freezeKeyGateway() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) FreezeKeyGateway() (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.FreezeKeyGateway(&_FarcasterKeyRegistry.TransactOpts)
}

// FreezeKeyGateway is a paid mutator transaction binding the contract method 0x47cf8d4d.
//
// Solidity: function freezeKeyGateway() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) FreezeKeyGateway() (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.FreezeKeyGateway(&_FarcasterKeyRegistry.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) Migrate() (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.Migrate(&_FarcasterKeyRegistry.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) Migrate() (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.Migrate(&_FarcasterKeyRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) Pause() (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.Pause(&_FarcasterKeyRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.Pause(&_FarcasterKeyRegistry.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0x58edef4c.
//
// Solidity: function remove(bytes key) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) Remove(opts *bind.TransactOpts, key []byte) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "remove", key)
}

// Remove is a paid mutator transaction binding the contract method 0x58edef4c.
//
// Solidity: function remove(bytes key) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) Remove(key []byte) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.Remove(&_FarcasterKeyRegistry.TransactOpts, key)
}

// Remove is a paid mutator transaction binding the contract method 0x58edef4c.
//
// Solidity: function remove(bytes key) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) Remove(key []byte) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.Remove(&_FarcasterKeyRegistry.TransactOpts, key)
}

// RemoveFor is a paid mutator transaction binding the contract method 0x787bd966.
//
// Solidity: function removeFor(address fidOwner, bytes key, uint256 deadline, bytes sig) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) RemoveFor(opts *bind.TransactOpts, fidOwner common.Address, key []byte, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "removeFor", fidOwner, key, deadline, sig)
}

// RemoveFor is a paid mutator transaction binding the contract method 0x787bd966.
//
// Solidity: function removeFor(address fidOwner, bytes key, uint256 deadline, bytes sig) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) RemoveFor(fidOwner common.Address, key []byte, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.RemoveFor(&_FarcasterKeyRegistry.TransactOpts, fidOwner, key, deadline, sig)
}

// RemoveFor is a paid mutator transaction binding the contract method 0x787bd966.
//
// Solidity: function removeFor(address fidOwner, bytes key, uint256 deadline, bytes sig) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) RemoveFor(fidOwner common.Address, key []byte, deadline *big.Int, sig []byte) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.RemoveFor(&_FarcasterKeyRegistry.TransactOpts, fidOwner, key, deadline, sig)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) RemoveGuardian(opts *bind.TransactOpts, guardian common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "removeGuardian", guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.RemoveGuardian(&_FarcasterKeyRegistry.TransactOpts, guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x71404156.
//
// Solidity: function removeGuardian(address guardian) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) RemoveGuardian(guardian common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.RemoveGuardian(&_FarcasterKeyRegistry.TransactOpts, guardian)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.RenounceOwnership(&_FarcasterKeyRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.RenounceOwnership(&_FarcasterKeyRegistry.TransactOpts)
}

// SetIdRegistry is a paid mutator transaction binding the contract method 0x81749468.
//
// Solidity: function setIdRegistry(address _idRegistry) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) SetIdRegistry(opts *bind.TransactOpts, _idRegistry common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "setIdRegistry", _idRegistry)
}

// SetIdRegistry is a paid mutator transaction binding the contract method 0x81749468.
//
// Solidity: function setIdRegistry(address _idRegistry) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) SetIdRegistry(_idRegistry common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.SetIdRegistry(&_FarcasterKeyRegistry.TransactOpts, _idRegistry)
}

// SetIdRegistry is a paid mutator transaction binding the contract method 0x81749468.
//
// Solidity: function setIdRegistry(address _idRegistry) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) SetIdRegistry(_idRegistry common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.SetIdRegistry(&_FarcasterKeyRegistry.TransactOpts, _idRegistry)
}

// SetKeyGateway is a paid mutator transaction binding the contract method 0xb221dac4.
//
// Solidity: function setKeyGateway(address _keyGateway) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) SetKeyGateway(opts *bind.TransactOpts, _keyGateway common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "setKeyGateway", _keyGateway)
}

// SetKeyGateway is a paid mutator transaction binding the contract method 0xb221dac4.
//
// Solidity: function setKeyGateway(address _keyGateway) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) SetKeyGateway(_keyGateway common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.SetKeyGateway(&_FarcasterKeyRegistry.TransactOpts, _keyGateway)
}

// SetKeyGateway is a paid mutator transaction binding the contract method 0xb221dac4.
//
// Solidity: function setKeyGateway(address _keyGateway) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) SetKeyGateway(_keyGateway common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.SetKeyGateway(&_FarcasterKeyRegistry.TransactOpts, _keyGateway)
}

// SetMaxKeysPerFid is a paid mutator transaction binding the contract method 0xd4c04809.
//
// Solidity: function setMaxKeysPerFid(uint256 _maxKeysPerFid) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) SetMaxKeysPerFid(opts *bind.TransactOpts, _maxKeysPerFid *big.Int) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "setMaxKeysPerFid", _maxKeysPerFid)
}

// SetMaxKeysPerFid is a paid mutator transaction binding the contract method 0xd4c04809.
//
// Solidity: function setMaxKeysPerFid(uint256 _maxKeysPerFid) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) SetMaxKeysPerFid(_maxKeysPerFid *big.Int) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.SetMaxKeysPerFid(&_FarcasterKeyRegistry.TransactOpts, _maxKeysPerFid)
}

// SetMaxKeysPerFid is a paid mutator transaction binding the contract method 0xd4c04809.
//
// Solidity: function setMaxKeysPerFid(uint256 _maxKeysPerFid) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) SetMaxKeysPerFid(_maxKeysPerFid *big.Int) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.SetMaxKeysPerFid(&_FarcasterKeyRegistry.TransactOpts, _maxKeysPerFid)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.SetMigrator(&_FarcasterKeyRegistry.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.SetMigrator(&_FarcasterKeyRegistry.TransactOpts, _migrator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x368ab610.
//
// Solidity: function setValidator(uint32 keyType, uint8 metadataType, address validator) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) SetValidator(opts *bind.TransactOpts, keyType uint32, metadataType uint8, validator common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "setValidator", keyType, metadataType, validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x368ab610.
//
// Solidity: function setValidator(uint32 keyType, uint8 metadataType, address validator) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) SetValidator(keyType uint32, metadataType uint8, validator common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.SetValidator(&_FarcasterKeyRegistry.TransactOpts, keyType, metadataType, validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x368ab610.
//
// Solidity: function setValidator(uint32 keyType, uint8 metadataType, address validator) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) SetValidator(keyType uint32, metadataType uint8, validator common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.SetValidator(&_FarcasterKeyRegistry.TransactOpts, keyType, metadataType, validator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.TransferOwnership(&_FarcasterKeyRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.TransferOwnership(&_FarcasterKeyRegistry.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) Unpause() (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.Unpause(&_FarcasterKeyRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.Unpause(&_FarcasterKeyRegistry.TransactOpts)
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactor) UseNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterKeyRegistry.contract.Transact(opts, "useNonce")
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_FarcasterKeyRegistry *FarcasterKeyRegistrySession) UseNonce() (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.UseNonce(&_FarcasterKeyRegistry.TransactOpts)
}

// UseNonce is a paid mutator transaction binding the contract method 0x69615a4c.
//
// Solidity: function useNonce() returns(uint256)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryTransactorSession) UseNonce() (*types.Transaction, error) {
	return _FarcasterKeyRegistry.Contract.UseNonce(&_FarcasterKeyRegistry.TransactOpts)
}

// FarcasterKeyRegistryAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryAddIterator struct {
	Event *FarcasterKeyRegistryAdd // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistryAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistryAdd)
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
		it.Event = new(FarcasterKeyRegistryAdd)
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
func (it *FarcasterKeyRegistryAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistryAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistryAdd represents a Add event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryAdd struct {
	Fid          *big.Int
	KeyType      uint32
	Key          common.Hash
	KeyBytes     []byte
	MetadataType uint8
	Metadata     []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x7d285df41058466977811345cd453c0c52e8d841ffaabc74fc050f277ad4de02.
//
// Solidity: event Add(uint256 indexed fid, uint32 indexed keyType, bytes indexed key, bytes keyBytes, uint8 metadataType, bytes metadata)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterAdd(opts *bind.FilterOpts, fid []*big.Int, keyType []uint32, key [][]byte) (*FarcasterKeyRegistryAddIterator, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}
	var keyTypeRule []interface{}
	for _, keyTypeItem := range keyType {
		keyTypeRule = append(keyTypeRule, keyTypeItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "Add", fidRule, keyTypeRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistryAddIterator{contract: _FarcasterKeyRegistry.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x7d285df41058466977811345cd453c0c52e8d841ffaabc74fc050f277ad4de02.
//
// Solidity: event Add(uint256 indexed fid, uint32 indexed keyType, bytes indexed key, bytes keyBytes, uint8 metadataType, bytes metadata)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistryAdd, fid []*big.Int, keyType []uint32, key [][]byte) (event.Subscription, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}
	var keyTypeRule []interface{}
	for _, keyTypeItem := range keyType {
		keyTypeRule = append(keyTypeRule, keyTypeItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "Add", fidRule, keyTypeRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistryAdd)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "Add", log); err != nil {
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

// ParseAdd is a log parse operation binding the contract event 0x7d285df41058466977811345cd453c0c52e8d841ffaabc74fc050f277ad4de02.
//
// Solidity: event Add(uint256 indexed fid, uint32 indexed keyType, bytes indexed key, bytes keyBytes, uint8 metadataType, bytes metadata)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseAdd(log types.Log) (*FarcasterKeyRegistryAdd, error) {
	event := new(FarcasterKeyRegistryAdd)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistryAdd0Iterator is returned from FilterAdd0 and is used to iterate over the raw logs and unpacked data for Add0 events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryAdd0Iterator struct {
	Event *FarcasterKeyRegistryAdd0 // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistryAdd0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistryAdd0)
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
		it.Event = new(FarcasterKeyRegistryAdd0)
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
func (it *FarcasterKeyRegistryAdd0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistryAdd0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistryAdd0 represents a Add0 event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryAdd0 struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdd0 is a free log retrieval operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterAdd0(opts *bind.FilterOpts, guardian []common.Address) (*FarcasterKeyRegistryAdd0Iterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "Add0", guardianRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistryAdd0Iterator{contract: _FarcasterKeyRegistry.contract, event: "Add0", logs: logs, sub: sub}, nil
}

// WatchAdd0 is a free log subscription operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchAdd0(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistryAdd0, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "Add0", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistryAdd0)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "Add0", log); err != nil {
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

// ParseAdd0 is a log parse operation binding the contract event 0x87dc5eecd6d6bdeae407c426da6bfba5b7190befc554ed5d4d62dd5cf939fbae.
//
// Solidity: event Add(address indexed guardian)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseAdd0(log types.Log) (*FarcasterKeyRegistryAdd0, error) {
	event := new(FarcasterKeyRegistryAdd0)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "Add0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistryAdminResetIterator is returned from FilterAdminReset and is used to iterate over the raw logs and unpacked data for AdminReset events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryAdminResetIterator struct {
	Event *FarcasterKeyRegistryAdminReset // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistryAdminResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistryAdminReset)
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
		it.Event = new(FarcasterKeyRegistryAdminReset)
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
func (it *FarcasterKeyRegistryAdminResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistryAdminResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistryAdminReset represents a AdminReset event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryAdminReset struct {
	Fid      *big.Int
	Key      common.Hash
	KeyBytes []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminReset is a free log retrieval operation binding the contract event 0x1ecc1009ebad5d2fb61239462f4f9f6f152662defe1845fc87f07d96bd1c60b4.
//
// Solidity: event AdminReset(uint256 indexed fid, bytes indexed key, bytes keyBytes)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterAdminReset(opts *bind.FilterOpts, fid []*big.Int, key [][]byte) (*FarcasterKeyRegistryAdminResetIterator, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "AdminReset", fidRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistryAdminResetIterator{contract: _FarcasterKeyRegistry.contract, event: "AdminReset", logs: logs, sub: sub}, nil
}

// WatchAdminReset is a free log subscription operation binding the contract event 0x1ecc1009ebad5d2fb61239462f4f9f6f152662defe1845fc87f07d96bd1c60b4.
//
// Solidity: event AdminReset(uint256 indexed fid, bytes indexed key, bytes keyBytes)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchAdminReset(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistryAdminReset, fid []*big.Int, key [][]byte) (event.Subscription, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "AdminReset", fidRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistryAdminReset)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "AdminReset", log); err != nil {
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

// ParseAdminReset is a log parse operation binding the contract event 0x1ecc1009ebad5d2fb61239462f4f9f6f152662defe1845fc87f07d96bd1c60b4.
//
// Solidity: event AdminReset(uint256 indexed fid, bytes indexed key, bytes keyBytes)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseAdminReset(log types.Log) (*FarcasterKeyRegistryAdminReset, error) {
	event := new(FarcasterKeyRegistryAdminReset)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "AdminReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistryEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryEIP712DomainChangedIterator struct {
	Event *FarcasterKeyRegistryEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistryEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistryEIP712DomainChanged)
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
		it.Event = new(FarcasterKeyRegistryEIP712DomainChanged)
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
func (it *FarcasterKeyRegistryEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistryEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistryEIP712DomainChanged represents a EIP712DomainChanged event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*FarcasterKeyRegistryEIP712DomainChangedIterator, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistryEIP712DomainChangedIterator{contract: _FarcasterKeyRegistry.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistryEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistryEIP712DomainChanged)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseEIP712DomainChanged(log types.Log) (*FarcasterKeyRegistryEIP712DomainChanged, error) {
	event := new(FarcasterKeyRegistryEIP712DomainChanged)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistryFreezeKeyGatewayIterator is returned from FilterFreezeKeyGateway and is used to iterate over the raw logs and unpacked data for FreezeKeyGateway events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryFreezeKeyGatewayIterator struct {
	Event *FarcasterKeyRegistryFreezeKeyGateway // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistryFreezeKeyGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistryFreezeKeyGateway)
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
		it.Event = new(FarcasterKeyRegistryFreezeKeyGateway)
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
func (it *FarcasterKeyRegistryFreezeKeyGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistryFreezeKeyGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistryFreezeKeyGateway represents a FreezeKeyGateway event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryFreezeKeyGateway struct {
	KeyGateway common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFreezeKeyGateway is a free log retrieval operation binding the contract event 0xcb685c7ba5a65fe9e6be9b7decbb5dc8ebba92bcae3cb09fc2a5a075b1eb5772.
//
// Solidity: event FreezeKeyGateway(address keyGateway)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterFreezeKeyGateway(opts *bind.FilterOpts) (*FarcasterKeyRegistryFreezeKeyGatewayIterator, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "FreezeKeyGateway")
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistryFreezeKeyGatewayIterator{contract: _FarcasterKeyRegistry.contract, event: "FreezeKeyGateway", logs: logs, sub: sub}, nil
}

// WatchFreezeKeyGateway is a free log subscription operation binding the contract event 0xcb685c7ba5a65fe9e6be9b7decbb5dc8ebba92bcae3cb09fc2a5a075b1eb5772.
//
// Solidity: event FreezeKeyGateway(address keyGateway)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchFreezeKeyGateway(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistryFreezeKeyGateway) (event.Subscription, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "FreezeKeyGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistryFreezeKeyGateway)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "FreezeKeyGateway", log); err != nil {
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

// ParseFreezeKeyGateway is a log parse operation binding the contract event 0xcb685c7ba5a65fe9e6be9b7decbb5dc8ebba92bcae3cb09fc2a5a075b1eb5772.
//
// Solidity: event FreezeKeyGateway(address keyGateway)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseFreezeKeyGateway(log types.Log) (*FarcasterKeyRegistryFreezeKeyGateway, error) {
	event := new(FarcasterKeyRegistryFreezeKeyGateway)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "FreezeKeyGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistryMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryMigratedIterator struct {
	Event *FarcasterKeyRegistryMigrated // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistryMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistryMigrated)
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
		it.Event = new(FarcasterKeyRegistryMigrated)
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
func (it *FarcasterKeyRegistryMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistryMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistryMigrated represents a Migrated event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryMigrated struct {
	MigratedAt *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xe4a25c0c2cbe89d6ad8b64c61a7dbdd20d1f781f6023f1ab94ebb7fe0aef6ab8.
//
// Solidity: event Migrated(uint256 indexed migratedAt)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterMigrated(opts *bind.FilterOpts, migratedAt []*big.Int) (*FarcasterKeyRegistryMigratedIterator, error) {

	var migratedAtRule []interface{}
	for _, migratedAtItem := range migratedAt {
		migratedAtRule = append(migratedAtRule, migratedAtItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "Migrated", migratedAtRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistryMigratedIterator{contract: _FarcasterKeyRegistry.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xe4a25c0c2cbe89d6ad8b64c61a7dbdd20d1f781f6023f1ab94ebb7fe0aef6ab8.
//
// Solidity: event Migrated(uint256 indexed migratedAt)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistryMigrated, migratedAt []*big.Int) (event.Subscription, error) {

	var migratedAtRule []interface{}
	for _, migratedAtItem := range migratedAt {
		migratedAtRule = append(migratedAtRule, migratedAtItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "Migrated", migratedAtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistryMigrated)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "Migrated", log); err != nil {
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
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseMigrated(log types.Log) (*FarcasterKeyRegistryMigrated, error) {
	event := new(FarcasterKeyRegistryMigrated)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "Migrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistryOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryOwnershipTransferStartedIterator struct {
	Event *FarcasterKeyRegistryOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistryOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistryOwnershipTransferStarted)
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
		it.Event = new(FarcasterKeyRegistryOwnershipTransferStarted)
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
func (it *FarcasterKeyRegistryOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistryOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistryOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FarcasterKeyRegistryOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistryOwnershipTransferStartedIterator{contract: _FarcasterKeyRegistry.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistryOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistryOwnershipTransferStarted)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseOwnershipTransferStarted(log types.Log) (*FarcasterKeyRegistryOwnershipTransferStarted, error) {
	event := new(FarcasterKeyRegistryOwnershipTransferStarted)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryOwnershipTransferredIterator struct {
	Event *FarcasterKeyRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistryOwnershipTransferred)
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
		it.Event = new(FarcasterKeyRegistryOwnershipTransferred)
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
func (it *FarcasterKeyRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FarcasterKeyRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistryOwnershipTransferredIterator{contract: _FarcasterKeyRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistryOwnershipTransferred)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*FarcasterKeyRegistryOwnershipTransferred, error) {
	event := new(FarcasterKeyRegistryOwnershipTransferred)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryPausedIterator struct {
	Event *FarcasterKeyRegistryPaused // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistryPaused)
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
		it.Event = new(FarcasterKeyRegistryPaused)
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
func (it *FarcasterKeyRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistryPaused represents a Paused event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*FarcasterKeyRegistryPausedIterator, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistryPausedIterator{contract: _FarcasterKeyRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistryPaused)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParsePaused(log types.Log) (*FarcasterKeyRegistryPaused, error) {
	event := new(FarcasterKeyRegistryPaused)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistryRemoveIterator is returned from FilterRemove and is used to iterate over the raw logs and unpacked data for Remove events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryRemoveIterator struct {
	Event *FarcasterKeyRegistryRemove // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistryRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistryRemove)
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
		it.Event = new(FarcasterKeyRegistryRemove)
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
func (it *FarcasterKeyRegistryRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistryRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistryRemove represents a Remove event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryRemove struct {
	Fid      *big.Int
	Key      common.Hash
	KeyBytes []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemove is a free log retrieval operation binding the contract event 0x09e77066e0155f46785be12f6938a6b2e4be4381e59058129ce15f355cb96958.
//
// Solidity: event Remove(uint256 indexed fid, bytes indexed key, bytes keyBytes)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterRemove(opts *bind.FilterOpts, fid []*big.Int, key [][]byte) (*FarcasterKeyRegistryRemoveIterator, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "Remove", fidRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistryRemoveIterator{contract: _FarcasterKeyRegistry.contract, event: "Remove", logs: logs, sub: sub}, nil
}

// WatchRemove is a free log subscription operation binding the contract event 0x09e77066e0155f46785be12f6938a6b2e4be4381e59058129ce15f355cb96958.
//
// Solidity: event Remove(uint256 indexed fid, bytes indexed key, bytes keyBytes)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchRemove(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistryRemove, fid []*big.Int, key [][]byte) (event.Subscription, error) {

	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "Remove", fidRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistryRemove)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "Remove", log); err != nil {
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

// ParseRemove is a log parse operation binding the contract event 0x09e77066e0155f46785be12f6938a6b2e4be4381e59058129ce15f355cb96958.
//
// Solidity: event Remove(uint256 indexed fid, bytes indexed key, bytes keyBytes)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseRemove(log types.Log) (*FarcasterKeyRegistryRemove, error) {
	event := new(FarcasterKeyRegistryRemove)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "Remove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistryRemove0Iterator is returned from FilterRemove0 and is used to iterate over the raw logs and unpacked data for Remove0 events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryRemove0Iterator struct {
	Event *FarcasterKeyRegistryRemove0 // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistryRemove0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistryRemove0)
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
		it.Event = new(FarcasterKeyRegistryRemove0)
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
func (it *FarcasterKeyRegistryRemove0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistryRemove0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistryRemove0 represents a Remove0 event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryRemove0 struct {
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemove0 is a free log retrieval operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterRemove0(opts *bind.FilterOpts, guardian []common.Address) (*FarcasterKeyRegistryRemove0Iterator, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "Remove0", guardianRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistryRemove0Iterator{contract: _FarcasterKeyRegistry.contract, event: "Remove0", logs: logs, sub: sub}, nil
}

// WatchRemove0 is a free log subscription operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchRemove0(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistryRemove0, guardian []common.Address) (event.Subscription, error) {

	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "Remove0", guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistryRemove0)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "Remove0", log); err != nil {
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

// ParseRemove0 is a log parse operation binding the contract event 0xbe7c7ac3248df4581c206a84aab3cb4e7d521b5398b42b681757f78a5a7d411e.
//
// Solidity: event Remove(address indexed guardian)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseRemove0(log types.Log) (*FarcasterKeyRegistryRemove0, error) {
	event := new(FarcasterKeyRegistryRemove0)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "Remove0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistrySetIdRegistryIterator is returned from FilterSetIdRegistry and is used to iterate over the raw logs and unpacked data for SetIdRegistry events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistrySetIdRegistryIterator struct {
	Event *FarcasterKeyRegistrySetIdRegistry // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistrySetIdRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistrySetIdRegistry)
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
		it.Event = new(FarcasterKeyRegistrySetIdRegistry)
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
func (it *FarcasterKeyRegistrySetIdRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistrySetIdRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistrySetIdRegistry represents a SetIdRegistry event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistrySetIdRegistry struct {
	OldIdRegistry common.Address
	NewIdRegistry common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetIdRegistry is a free log retrieval operation binding the contract event 0x940dcf34ec2e245e837ee4599997e577ce274d7b87c73238e2878ac7ea1af2f1.
//
// Solidity: event SetIdRegistry(address oldIdRegistry, address newIdRegistry)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterSetIdRegistry(opts *bind.FilterOpts) (*FarcasterKeyRegistrySetIdRegistryIterator, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "SetIdRegistry")
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistrySetIdRegistryIterator{contract: _FarcasterKeyRegistry.contract, event: "SetIdRegistry", logs: logs, sub: sub}, nil
}

// WatchSetIdRegistry is a free log subscription operation binding the contract event 0x940dcf34ec2e245e837ee4599997e577ce274d7b87c73238e2878ac7ea1af2f1.
//
// Solidity: event SetIdRegistry(address oldIdRegistry, address newIdRegistry)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchSetIdRegistry(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistrySetIdRegistry) (event.Subscription, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "SetIdRegistry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistrySetIdRegistry)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "SetIdRegistry", log); err != nil {
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

// ParseSetIdRegistry is a log parse operation binding the contract event 0x940dcf34ec2e245e837ee4599997e577ce274d7b87c73238e2878ac7ea1af2f1.
//
// Solidity: event SetIdRegistry(address oldIdRegistry, address newIdRegistry)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseSetIdRegistry(log types.Log) (*FarcasterKeyRegistrySetIdRegistry, error) {
	event := new(FarcasterKeyRegistrySetIdRegistry)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "SetIdRegistry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistrySetKeyGatewayIterator is returned from FilterSetKeyGateway and is used to iterate over the raw logs and unpacked data for SetKeyGateway events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistrySetKeyGatewayIterator struct {
	Event *FarcasterKeyRegistrySetKeyGateway // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistrySetKeyGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistrySetKeyGateway)
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
		it.Event = new(FarcasterKeyRegistrySetKeyGateway)
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
func (it *FarcasterKeyRegistrySetKeyGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistrySetKeyGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistrySetKeyGateway represents a SetKeyGateway event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistrySetKeyGateway struct {
	OldKeyGateway common.Address
	NewKeyGateway common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetKeyGateway is a free log retrieval operation binding the contract event 0x56785750704201befc0a27dae1e5d37835a8ad6e35affc87136ed24d1ac694ac.
//
// Solidity: event SetKeyGateway(address oldKeyGateway, address newKeyGateway)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterSetKeyGateway(opts *bind.FilterOpts) (*FarcasterKeyRegistrySetKeyGatewayIterator, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "SetKeyGateway")
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistrySetKeyGatewayIterator{contract: _FarcasterKeyRegistry.contract, event: "SetKeyGateway", logs: logs, sub: sub}, nil
}

// WatchSetKeyGateway is a free log subscription operation binding the contract event 0x56785750704201befc0a27dae1e5d37835a8ad6e35affc87136ed24d1ac694ac.
//
// Solidity: event SetKeyGateway(address oldKeyGateway, address newKeyGateway)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchSetKeyGateway(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistrySetKeyGateway) (event.Subscription, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "SetKeyGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistrySetKeyGateway)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "SetKeyGateway", log); err != nil {
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

// ParseSetKeyGateway is a log parse operation binding the contract event 0x56785750704201befc0a27dae1e5d37835a8ad6e35affc87136ed24d1ac694ac.
//
// Solidity: event SetKeyGateway(address oldKeyGateway, address newKeyGateway)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseSetKeyGateway(log types.Log) (*FarcasterKeyRegistrySetKeyGateway, error) {
	event := new(FarcasterKeyRegistrySetKeyGateway)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "SetKeyGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistrySetMaxKeysPerFidIterator is returned from FilterSetMaxKeysPerFid and is used to iterate over the raw logs and unpacked data for SetMaxKeysPerFid events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistrySetMaxKeysPerFidIterator struct {
	Event *FarcasterKeyRegistrySetMaxKeysPerFid // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistrySetMaxKeysPerFidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistrySetMaxKeysPerFid)
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
		it.Event = new(FarcasterKeyRegistrySetMaxKeysPerFid)
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
func (it *FarcasterKeyRegistrySetMaxKeysPerFidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistrySetMaxKeysPerFidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistrySetMaxKeysPerFid represents a SetMaxKeysPerFid event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistrySetMaxKeysPerFid struct {
	OldMax *big.Int
	NewMax *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMaxKeysPerFid is a free log retrieval operation binding the contract event 0x6c336d0e5e74bb26a5e2d4646801801837b0fdbaddd9131923fd42d740449731.
//
// Solidity: event SetMaxKeysPerFid(uint256 oldMax, uint256 newMax)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterSetMaxKeysPerFid(opts *bind.FilterOpts) (*FarcasterKeyRegistrySetMaxKeysPerFidIterator, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "SetMaxKeysPerFid")
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistrySetMaxKeysPerFidIterator{contract: _FarcasterKeyRegistry.contract, event: "SetMaxKeysPerFid", logs: logs, sub: sub}, nil
}

// WatchSetMaxKeysPerFid is a free log subscription operation binding the contract event 0x6c336d0e5e74bb26a5e2d4646801801837b0fdbaddd9131923fd42d740449731.
//
// Solidity: event SetMaxKeysPerFid(uint256 oldMax, uint256 newMax)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchSetMaxKeysPerFid(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistrySetMaxKeysPerFid) (event.Subscription, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "SetMaxKeysPerFid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistrySetMaxKeysPerFid)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "SetMaxKeysPerFid", log); err != nil {
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

// ParseSetMaxKeysPerFid is a log parse operation binding the contract event 0x6c336d0e5e74bb26a5e2d4646801801837b0fdbaddd9131923fd42d740449731.
//
// Solidity: event SetMaxKeysPerFid(uint256 oldMax, uint256 newMax)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseSetMaxKeysPerFid(log types.Log) (*FarcasterKeyRegistrySetMaxKeysPerFid, error) {
	event := new(FarcasterKeyRegistrySetMaxKeysPerFid)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "SetMaxKeysPerFid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistrySetMigratorIterator is returned from FilterSetMigrator and is used to iterate over the raw logs and unpacked data for SetMigrator events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistrySetMigratorIterator struct {
	Event *FarcasterKeyRegistrySetMigrator // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistrySetMigratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistrySetMigrator)
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
		it.Event = new(FarcasterKeyRegistrySetMigrator)
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
func (it *FarcasterKeyRegistrySetMigratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistrySetMigratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistrySetMigrator represents a SetMigrator event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistrySetMigrator struct {
	OldMigrator common.Address
	NewMigrator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetMigrator is a free log retrieval operation binding the contract event 0xd8ad954fe808212ab9ed7139873e40807dff7995fe36e3d6cdeb8fa00fcebf10.
//
// Solidity: event SetMigrator(address oldMigrator, address newMigrator)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterSetMigrator(opts *bind.FilterOpts) (*FarcasterKeyRegistrySetMigratorIterator, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "SetMigrator")
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistrySetMigratorIterator{contract: _FarcasterKeyRegistry.contract, event: "SetMigrator", logs: logs, sub: sub}, nil
}

// WatchSetMigrator is a free log subscription operation binding the contract event 0xd8ad954fe808212ab9ed7139873e40807dff7995fe36e3d6cdeb8fa00fcebf10.
//
// Solidity: event SetMigrator(address oldMigrator, address newMigrator)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchSetMigrator(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistrySetMigrator) (event.Subscription, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "SetMigrator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistrySetMigrator)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "SetMigrator", log); err != nil {
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
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseSetMigrator(log types.Log) (*FarcasterKeyRegistrySetMigrator, error) {
	event := new(FarcasterKeyRegistrySetMigrator)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "SetMigrator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistrySetValidatorIterator is returned from FilterSetValidator and is used to iterate over the raw logs and unpacked data for SetValidator events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistrySetValidatorIterator struct {
	Event *FarcasterKeyRegistrySetValidator // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistrySetValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistrySetValidator)
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
		it.Event = new(FarcasterKeyRegistrySetValidator)
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
func (it *FarcasterKeyRegistrySetValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistrySetValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistrySetValidator represents a SetValidator event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistrySetValidator struct {
	KeyType      uint32
	MetadataType uint8
	OldValidator common.Address
	NewValidator common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetValidator is a free log retrieval operation binding the contract event 0xd964242236f6208120d76a25cd886db49c82403f50d88dfd1bc865ee60ad462d.
//
// Solidity: event SetValidator(uint32 keyType, uint8 metadataType, address oldValidator, address newValidator)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterSetValidator(opts *bind.FilterOpts) (*FarcasterKeyRegistrySetValidatorIterator, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "SetValidator")
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistrySetValidatorIterator{contract: _FarcasterKeyRegistry.contract, event: "SetValidator", logs: logs, sub: sub}, nil
}

// WatchSetValidator is a free log subscription operation binding the contract event 0xd964242236f6208120d76a25cd886db49c82403f50d88dfd1bc865ee60ad462d.
//
// Solidity: event SetValidator(uint32 keyType, uint8 metadataType, address oldValidator, address newValidator)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchSetValidator(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistrySetValidator) (event.Subscription, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "SetValidator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistrySetValidator)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "SetValidator", log); err != nil {
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

// ParseSetValidator is a log parse operation binding the contract event 0xd964242236f6208120d76a25cd886db49c82403f50d88dfd1bc865ee60ad462d.
//
// Solidity: event SetValidator(uint32 keyType, uint8 metadataType, address oldValidator, address newValidator)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseSetValidator(log types.Log) (*FarcasterKeyRegistrySetValidator, error) {
	event := new(FarcasterKeyRegistrySetValidator)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "SetValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterKeyRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryUnpausedIterator struct {
	Event *FarcasterKeyRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *FarcasterKeyRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterKeyRegistryUnpaused)
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
		it.Event = new(FarcasterKeyRegistryUnpaused)
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
func (it *FarcasterKeyRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterKeyRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterKeyRegistryUnpaused represents a Unpaused event raised by the FarcasterKeyRegistry contract.
type FarcasterKeyRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*FarcasterKeyRegistryUnpausedIterator, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &FarcasterKeyRegistryUnpausedIterator{contract: _FarcasterKeyRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *FarcasterKeyRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _FarcasterKeyRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterKeyRegistryUnpaused)
				if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_FarcasterKeyRegistry *FarcasterKeyRegistryFilterer) ParseUnpaused(log types.Log) (*FarcasterKeyRegistryUnpaused, error) {
	event := new(FarcasterKeyRegistryUnpaused)
	if err := _FarcasterKeyRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
