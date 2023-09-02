// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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

// RICRegistryStatus is an auto generated low-level Go binding around an user-defined struct.
type RICRegistryStatus struct {
	Status          uint8
	Provider        common.Address
	QueuedTimestamp *big.Int
	ChainID         *big.Int
	Config          []byte
}

// RICRegistryMetaData contains all meta data concerning the RICRegistry contract.
var RICRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_queueTimeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_providerStakeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"}],\"name\":\"providerSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"providerStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"providerUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"rollupActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestedTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"name\":\"rollupQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"rollupRequested\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activatedRollupsL1Addresses\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"l1Addresses_\",\"type\":\"bytes\"}],\"name\":\"deployRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID_\",\"type\":\"uint256\"}],\"name\":\"getRollupStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumRICRegistry.RollupStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"queuedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"internalType\":\"structRICRegistry.Status\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID_\",\"type\":\"uint256\"}],\"name\":\"queueRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueTimeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"name\":\"requestRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rollupStatus\",\"outputs\":[{\"internalType\":\"enumRICRegistry.RollupStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"queuedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeAsProvider\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// RICRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RICRegistryMetaData.ABI instead.
var RICRegistryABI = RICRegistryMetaData.ABI

// RICRegistry is an auto generated Go binding around an Ethereum contract.
type RICRegistry struct {
	RICRegistryCaller     // Read-only binding to the contract
	RICRegistryTransactor // Write-only binding to the contract
	RICRegistryFilterer   // Log filterer for contract events
}

// RICRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RICRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RICRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RICRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RICRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RICRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RICRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RICRegistrySession struct {
	Contract     *RICRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RICRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RICRegistryCallerSession struct {
	Contract *RICRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RICRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RICRegistryTransactorSession struct {
	Contract     *RICRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RICRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RICRegistryRaw struct {
	Contract *RICRegistry // Generic contract binding to access the raw methods on
}

// RICRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RICRegistryCallerRaw struct {
	Contract *RICRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RICRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RICRegistryTransactorRaw struct {
	Contract *RICRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRICRegistry creates a new instance of RICRegistry, bound to a specific deployed contract.
func NewRICRegistry(address common.Address, backend bind.ContractBackend) (*RICRegistry, error) {
	contract, err := bindRICRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RICRegistry{RICRegistryCaller: RICRegistryCaller{contract: contract}, RICRegistryTransactor: RICRegistryTransactor{contract: contract}, RICRegistryFilterer: RICRegistryFilterer{contract: contract}}, nil
}

// NewRICRegistryCaller creates a new read-only instance of RICRegistry, bound to a specific deployed contract.
func NewRICRegistryCaller(address common.Address, caller bind.ContractCaller) (*RICRegistryCaller, error) {
	contract, err := bindRICRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RICRegistryCaller{contract: contract}, nil
}

// NewRICRegistryTransactor creates a new write-only instance of RICRegistry, bound to a specific deployed contract.
func NewRICRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RICRegistryTransactor, error) {
	contract, err := bindRICRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RICRegistryTransactor{contract: contract}, nil
}

// NewRICRegistryFilterer creates a new log filterer instance of RICRegistry, bound to a specific deployed contract.
func NewRICRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RICRegistryFilterer, error) {
	contract, err := bindRICRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RICRegistryFilterer{contract: contract}, nil
}

// bindRICRegistry binds a generic wrapper to an already deployed contract.
func bindRICRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RICRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RICRegistry *RICRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RICRegistry.Contract.RICRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RICRegistry *RICRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RICRegistry.Contract.RICRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RICRegistry *RICRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RICRegistry.Contract.RICRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RICRegistry *RICRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RICRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RICRegistry *RICRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RICRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RICRegistry *RICRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RICRegistry.Contract.contract.Transact(opts, method, params...)
}

// ActivatedRollupsL1Addresses is a free data retrieval call binding the contract method 0x93178857.
//
// Solidity: function activatedRollupsL1Addresses(uint256 ) view returns(bytes)
func (_RICRegistry *RICRegistryCaller) ActivatedRollupsL1Addresses(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _RICRegistry.contract.Call(opts, &out, "activatedRollupsL1Addresses", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ActivatedRollupsL1Addresses is a free data retrieval call binding the contract method 0x93178857.
//
// Solidity: function activatedRollupsL1Addresses(uint256 ) view returns(bytes)
func (_RICRegistry *RICRegistrySession) ActivatedRollupsL1Addresses(arg0 *big.Int) ([]byte, error) {
	return _RICRegistry.Contract.ActivatedRollupsL1Addresses(&_RICRegistry.CallOpts, arg0)
}

// ActivatedRollupsL1Addresses is a free data retrieval call binding the contract method 0x93178857.
//
// Solidity: function activatedRollupsL1Addresses(uint256 ) view returns(bytes)
func (_RICRegistry *RICRegistryCallerSession) ActivatedRollupsL1Addresses(arg0 *big.Int) ([]byte, error) {
	return _RICRegistry.Contract.ActivatedRollupsL1Addresses(&_RICRegistry.CallOpts, arg0)
}

// GetRollupStatus is a free data retrieval call binding the contract method 0xc76ec393.
//
// Solidity: function getRollupStatus(uint256 chainID_) view returns((uint8,address,uint256,uint256,bytes))
func (_RICRegistry *RICRegistryCaller) GetRollupStatus(opts *bind.CallOpts, chainID_ *big.Int) (RICRegistryStatus, error) {
	var out []interface{}
	err := _RICRegistry.contract.Call(opts, &out, "getRollupStatus", chainID_)

	if err != nil {
		return *new(RICRegistryStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(RICRegistryStatus)).(*RICRegistryStatus)

	return out0, err

}

// GetRollupStatus is a free data retrieval call binding the contract method 0xc76ec393.
//
// Solidity: function getRollupStatus(uint256 chainID_) view returns((uint8,address,uint256,uint256,bytes))
func (_RICRegistry *RICRegistrySession) GetRollupStatus(chainID_ *big.Int) (RICRegistryStatus, error) {
	return _RICRegistry.Contract.GetRollupStatus(&_RICRegistry.CallOpts, chainID_)
}

// GetRollupStatus is a free data retrieval call binding the contract method 0xc76ec393.
//
// Solidity: function getRollupStatus(uint256 chainID_) view returns((uint8,address,uint256,uint256,bytes))
func (_RICRegistry *RICRegistryCallerSession) GetRollupStatus(chainID_ *big.Int) (RICRegistryStatus, error) {
	return _RICRegistry.Contract.GetRollupStatus(&_RICRegistry.CallOpts, chainID_)
}

// ProviderStake is a free data retrieval call binding the contract method 0x13bd0523.
//
// Solidity: function providerStake(address ) view returns(uint256)
func (_RICRegistry *RICRegistryCaller) ProviderStake(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RICRegistry.contract.Call(opts, &out, "providerStake", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProviderStake is a free data retrieval call binding the contract method 0x13bd0523.
//
// Solidity: function providerStake(address ) view returns(uint256)
func (_RICRegistry *RICRegistrySession) ProviderStake(arg0 common.Address) (*big.Int, error) {
	return _RICRegistry.Contract.ProviderStake(&_RICRegistry.CallOpts, arg0)
}

// ProviderStake is a free data retrieval call binding the contract method 0x13bd0523.
//
// Solidity: function providerStake(address ) view returns(uint256)
func (_RICRegistry *RICRegistryCallerSession) ProviderStake(arg0 common.Address) (*big.Int, error) {
	return _RICRegistry.Contract.ProviderStake(&_RICRegistry.CallOpts, arg0)
}

// ProviderStakeAmount is a free data retrieval call binding the contract method 0x6854d20f.
//
// Solidity: function providerStakeAmount() view returns(uint256)
func (_RICRegistry *RICRegistryCaller) ProviderStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RICRegistry.contract.Call(opts, &out, "providerStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProviderStakeAmount is a free data retrieval call binding the contract method 0x6854d20f.
//
// Solidity: function providerStakeAmount() view returns(uint256)
func (_RICRegistry *RICRegistrySession) ProviderStakeAmount() (*big.Int, error) {
	return _RICRegistry.Contract.ProviderStakeAmount(&_RICRegistry.CallOpts)
}

// ProviderStakeAmount is a free data retrieval call binding the contract method 0x6854d20f.
//
// Solidity: function providerStakeAmount() view returns(uint256)
func (_RICRegistry *RICRegistryCallerSession) ProviderStakeAmount() (*big.Int, error) {
	return _RICRegistry.Contract.ProviderStakeAmount(&_RICRegistry.CallOpts)
}

// QueueTimeout is a free data retrieval call binding the contract method 0x6d2ddb0c.
//
// Solidity: function queueTimeout() view returns(uint256)
func (_RICRegistry *RICRegistryCaller) QueueTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RICRegistry.contract.Call(opts, &out, "queueTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueueTimeout is a free data retrieval call binding the contract method 0x6d2ddb0c.
//
// Solidity: function queueTimeout() view returns(uint256)
func (_RICRegistry *RICRegistrySession) QueueTimeout() (*big.Int, error) {
	return _RICRegistry.Contract.QueueTimeout(&_RICRegistry.CallOpts)
}

// QueueTimeout is a free data retrieval call binding the contract method 0x6d2ddb0c.
//
// Solidity: function queueTimeout() view returns(uint256)
func (_RICRegistry *RICRegistryCallerSession) QueueTimeout() (*big.Int, error) {
	return _RICRegistry.Contract.QueueTimeout(&_RICRegistry.CallOpts)
}

// RollupStatus is a free data retrieval call binding the contract method 0x9d29ca27.
//
// Solidity: function rollupStatus(uint256 ) view returns(uint8 status, address provider, uint256 queuedTimestamp, uint256 chainID, bytes config)
func (_RICRegistry *RICRegistryCaller) RollupStatus(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status          uint8
	Provider        common.Address
	QueuedTimestamp *big.Int
	ChainID         *big.Int
	Config          []byte
}, error) {
	var out []interface{}
	err := _RICRegistry.contract.Call(opts, &out, "rollupStatus", arg0)

	outstruct := new(struct {
		Status          uint8
		Provider        common.Address
		QueuedTimestamp *big.Int
		ChainID         *big.Int
		Config          []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Provider = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.QueuedTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ChainID = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Config = *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return *outstruct, err

}

// RollupStatus is a free data retrieval call binding the contract method 0x9d29ca27.
//
// Solidity: function rollupStatus(uint256 ) view returns(uint8 status, address provider, uint256 queuedTimestamp, uint256 chainID, bytes config)
func (_RICRegistry *RICRegistrySession) RollupStatus(arg0 *big.Int) (struct {
	Status          uint8
	Provider        common.Address
	QueuedTimestamp *big.Int
	ChainID         *big.Int
	Config          []byte
}, error) {
	return _RICRegistry.Contract.RollupStatus(&_RICRegistry.CallOpts, arg0)
}

// RollupStatus is a free data retrieval call binding the contract method 0x9d29ca27.
//
// Solidity: function rollupStatus(uint256 ) view returns(uint8 status, address provider, uint256 queuedTimestamp, uint256 chainID, bytes config)
func (_RICRegistry *RICRegistryCallerSession) RollupStatus(arg0 *big.Int) (struct {
	Status          uint8
	Provider        common.Address
	QueuedTimestamp *big.Int
	ChainID         *big.Int
	Config          []byte
}, error) {
	return _RICRegistry.Contract.RollupStatus(&_RICRegistry.CallOpts, arg0)
}

// DeployRollup is a paid mutator transaction binding the contract method 0xefe30aba.
//
// Solidity: function deployRollup(uint256 chainID_, bytes l1Addresses_) returns()
func (_RICRegistry *RICRegistryTransactor) DeployRollup(opts *bind.TransactOpts, chainID_ *big.Int, l1Addresses_ []byte) (*types.Transaction, error) {
	return _RICRegistry.contract.Transact(opts, "deployRollup", chainID_, l1Addresses_)
}

// DeployRollup is a paid mutator transaction binding the contract method 0xefe30aba.
//
// Solidity: function deployRollup(uint256 chainID_, bytes l1Addresses_) returns()
func (_RICRegistry *RICRegistrySession) DeployRollup(chainID_ *big.Int, l1Addresses_ []byte) (*types.Transaction, error) {
	return _RICRegistry.Contract.DeployRollup(&_RICRegistry.TransactOpts, chainID_, l1Addresses_)
}

// DeployRollup is a paid mutator transaction binding the contract method 0xefe30aba.
//
// Solidity: function deployRollup(uint256 chainID_, bytes l1Addresses_) returns()
func (_RICRegistry *RICRegistryTransactorSession) DeployRollup(chainID_ *big.Int, l1Addresses_ []byte) (*types.Transaction, error) {
	return _RICRegistry.Contract.DeployRollup(&_RICRegistry.TransactOpts, chainID_, l1Addresses_)
}

// ProviderUnstake is a paid mutator transaction binding the contract method 0x56340244.
//
// Solidity: function providerUnstake() returns()
func (_RICRegistry *RICRegistryTransactor) ProviderUnstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RICRegistry.contract.Transact(opts, "providerUnstake")
}

// ProviderUnstake is a paid mutator transaction binding the contract method 0x56340244.
//
// Solidity: function providerUnstake() returns()
func (_RICRegistry *RICRegistrySession) ProviderUnstake() (*types.Transaction, error) {
	return _RICRegistry.Contract.ProviderUnstake(&_RICRegistry.TransactOpts)
}

// ProviderUnstake is a paid mutator transaction binding the contract method 0x56340244.
//
// Solidity: function providerUnstake() returns()
func (_RICRegistry *RICRegistryTransactorSession) ProviderUnstake() (*types.Transaction, error) {
	return _RICRegistry.Contract.ProviderUnstake(&_RICRegistry.TransactOpts)
}

// QueueRollup is a paid mutator transaction binding the contract method 0x53219005.
//
// Solidity: function queueRollup(uint256 chainID_) returns()
func (_RICRegistry *RICRegistryTransactor) QueueRollup(opts *bind.TransactOpts, chainID_ *big.Int) (*types.Transaction, error) {
	return _RICRegistry.contract.Transact(opts, "queueRollup", chainID_)
}

// QueueRollup is a paid mutator transaction binding the contract method 0x53219005.
//
// Solidity: function queueRollup(uint256 chainID_) returns()
func (_RICRegistry *RICRegistrySession) QueueRollup(chainID_ *big.Int) (*types.Transaction, error) {
	return _RICRegistry.Contract.QueueRollup(&_RICRegistry.TransactOpts, chainID_)
}

// QueueRollup is a paid mutator transaction binding the contract method 0x53219005.
//
// Solidity: function queueRollup(uint256 chainID_) returns()
func (_RICRegistry *RICRegistryTransactorSession) QueueRollup(chainID_ *big.Int) (*types.Transaction, error) {
	return _RICRegistry.Contract.QueueRollup(&_RICRegistry.TransactOpts, chainID_)
}

// RequestRollup is a paid mutator transaction binding the contract method 0x96832d07.
//
// Solidity: function requestRollup(uint256 chainID_, bytes config) returns()
func (_RICRegistry *RICRegistryTransactor) RequestRollup(opts *bind.TransactOpts, chainID_ *big.Int, config []byte) (*types.Transaction, error) {
	return _RICRegistry.contract.Transact(opts, "requestRollup", chainID_, config)
}

// RequestRollup is a paid mutator transaction binding the contract method 0x96832d07.
//
// Solidity: function requestRollup(uint256 chainID_, bytes config) returns()
func (_RICRegistry *RICRegistrySession) RequestRollup(chainID_ *big.Int, config []byte) (*types.Transaction, error) {
	return _RICRegistry.Contract.RequestRollup(&_RICRegistry.TransactOpts, chainID_, config)
}

// RequestRollup is a paid mutator transaction binding the contract method 0x96832d07.
//
// Solidity: function requestRollup(uint256 chainID_, bytes config) returns()
func (_RICRegistry *RICRegistryTransactorSession) RequestRollup(chainID_ *big.Int, config []byte) (*types.Transaction, error) {
	return _RICRegistry.Contract.RequestRollup(&_RICRegistry.TransactOpts, chainID_, config)
}

// StakeAsProvider is a paid mutator transaction binding the contract method 0x8c374073.
//
// Solidity: function stakeAsProvider() payable returns()
func (_RICRegistry *RICRegistryTransactor) StakeAsProvider(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RICRegistry.contract.Transact(opts, "stakeAsProvider")
}

// StakeAsProvider is a paid mutator transaction binding the contract method 0x8c374073.
//
// Solidity: function stakeAsProvider() payable returns()
func (_RICRegistry *RICRegistrySession) StakeAsProvider() (*types.Transaction, error) {
	return _RICRegistry.Contract.StakeAsProvider(&_RICRegistry.TransactOpts)
}

// StakeAsProvider is a paid mutator transaction binding the contract method 0x8c374073.
//
// Solidity: function stakeAsProvider() payable returns()
func (_RICRegistry *RICRegistryTransactorSession) StakeAsProvider() (*types.Transaction, error) {
	return _RICRegistry.Contract.StakeAsProvider(&_RICRegistry.TransactOpts)
}

// RICRegistryProviderSlashedIterator is returned from FilterProviderSlashed and is used to iterate over the raw logs and unpacked data for ProviderSlashed events raised by the RICRegistry contract.
type RICRegistryProviderSlashedIterator struct {
	Event *RICRegistryProviderSlashed // Event containing the contract specifics and raw log

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
func (it *RICRegistryProviderSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RICRegistryProviderSlashed)
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
		it.Event = new(RICRegistryProviderSlashed)
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
func (it *RICRegistryProviderSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RICRegistryProviderSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RICRegistryProviderSlashed represents a ProviderSlashed event raised by the RICRegistry contract.
type RICRegistryProviderSlashed struct {
	Provider common.Address
	Slasher  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderSlashed is a free log retrieval operation binding the contract event 0x6224637401e13f0bde2c33c63b3f7b896286f5f25c031feced5a2a536b5d4a1a.
//
// Solidity: event providerSlashed(address provider, address slasher)
func (_RICRegistry *RICRegistryFilterer) FilterProviderSlashed(opts *bind.FilterOpts) (*RICRegistryProviderSlashedIterator, error) {

	logs, sub, err := _RICRegistry.contract.FilterLogs(opts, "providerSlashed")
	if err != nil {
		return nil, err
	}
	return &RICRegistryProviderSlashedIterator{contract: _RICRegistry.contract, event: "providerSlashed", logs: logs, sub: sub}, nil
}

// WatchProviderSlashed is a free log subscription operation binding the contract event 0x6224637401e13f0bde2c33c63b3f7b896286f5f25c031feced5a2a536b5d4a1a.
//
// Solidity: event providerSlashed(address provider, address slasher)
func (_RICRegistry *RICRegistryFilterer) WatchProviderSlashed(opts *bind.WatchOpts, sink chan<- *RICRegistryProviderSlashed) (event.Subscription, error) {

	logs, sub, err := _RICRegistry.contract.WatchLogs(opts, "providerSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RICRegistryProviderSlashed)
				if err := _RICRegistry.contract.UnpackLog(event, "providerSlashed", log); err != nil {
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

// ParseProviderSlashed is a log parse operation binding the contract event 0x6224637401e13f0bde2c33c63b3f7b896286f5f25c031feced5a2a536b5d4a1a.
//
// Solidity: event providerSlashed(address provider, address slasher)
func (_RICRegistry *RICRegistryFilterer) ParseProviderSlashed(log types.Log) (*RICRegistryProviderSlashed, error) {
	event := new(RICRegistryProviderSlashed)
	if err := _RICRegistry.contract.UnpackLog(event, "providerSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RICRegistryProviderStakedIterator is returned from FilterProviderStaked and is used to iterate over the raw logs and unpacked data for ProviderStaked events raised by the RICRegistry contract.
type RICRegistryProviderStakedIterator struct {
	Event *RICRegistryProviderStaked // Event containing the contract specifics and raw log

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
func (it *RICRegistryProviderStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RICRegistryProviderStaked)
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
		it.Event = new(RICRegistryProviderStaked)
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
func (it *RICRegistryProviderStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RICRegistryProviderStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RICRegistryProviderStaked represents a ProviderStaked event raised by the RICRegistry contract.
type RICRegistryProviderStaked struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderStaked is a free log retrieval operation binding the contract event 0x1736f0c38f824f3c01726a7af14e1f353c3266e388bdfc5ec1e4097d23e9e23e.
//
// Solidity: event providerStaked(address provider)
func (_RICRegistry *RICRegistryFilterer) FilterProviderStaked(opts *bind.FilterOpts) (*RICRegistryProviderStakedIterator, error) {

	logs, sub, err := _RICRegistry.contract.FilterLogs(opts, "providerStaked")
	if err != nil {
		return nil, err
	}
	return &RICRegistryProviderStakedIterator{contract: _RICRegistry.contract, event: "providerStaked", logs: logs, sub: sub}, nil
}

// WatchProviderStaked is a free log subscription operation binding the contract event 0x1736f0c38f824f3c01726a7af14e1f353c3266e388bdfc5ec1e4097d23e9e23e.
//
// Solidity: event providerStaked(address provider)
func (_RICRegistry *RICRegistryFilterer) WatchProviderStaked(opts *bind.WatchOpts, sink chan<- *RICRegistryProviderStaked) (event.Subscription, error) {

	logs, sub, err := _RICRegistry.contract.WatchLogs(opts, "providerStaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RICRegistryProviderStaked)
				if err := _RICRegistry.contract.UnpackLog(event, "providerStaked", log); err != nil {
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

// ParseProviderStaked is a log parse operation binding the contract event 0x1736f0c38f824f3c01726a7af14e1f353c3266e388bdfc5ec1e4097d23e9e23e.
//
// Solidity: event providerStaked(address provider)
func (_RICRegistry *RICRegistryFilterer) ParseProviderStaked(log types.Log) (*RICRegistryProviderStaked, error) {
	event := new(RICRegistryProviderStaked)
	if err := _RICRegistry.contract.UnpackLog(event, "providerStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RICRegistryProviderUnstakedIterator is returned from FilterProviderUnstaked and is used to iterate over the raw logs and unpacked data for ProviderUnstaked events raised by the RICRegistry contract.
type RICRegistryProviderUnstakedIterator struct {
	Event *RICRegistryProviderUnstaked // Event containing the contract specifics and raw log

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
func (it *RICRegistryProviderUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RICRegistryProviderUnstaked)
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
		it.Event = new(RICRegistryProviderUnstaked)
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
func (it *RICRegistryProviderUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RICRegistryProviderUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RICRegistryProviderUnstaked represents a ProviderUnstaked event raised by the RICRegistry contract.
type RICRegistryProviderUnstaked struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderUnstaked is a free log retrieval operation binding the contract event 0xc040f13890127304b60b2a24fa91a0e4877937f8426724681c2d0dfe7ff90207.
//
// Solidity: event providerUnstaked(address provider)
func (_RICRegistry *RICRegistryFilterer) FilterProviderUnstaked(opts *bind.FilterOpts) (*RICRegistryProviderUnstakedIterator, error) {

	logs, sub, err := _RICRegistry.contract.FilterLogs(opts, "providerUnstaked")
	if err != nil {
		return nil, err
	}
	return &RICRegistryProviderUnstakedIterator{contract: _RICRegistry.contract, event: "providerUnstaked", logs: logs, sub: sub}, nil
}

// WatchProviderUnstaked is a free log subscription operation binding the contract event 0xc040f13890127304b60b2a24fa91a0e4877937f8426724681c2d0dfe7ff90207.
//
// Solidity: event providerUnstaked(address provider)
func (_RICRegistry *RICRegistryFilterer) WatchProviderUnstaked(opts *bind.WatchOpts, sink chan<- *RICRegistryProviderUnstaked) (event.Subscription, error) {

	logs, sub, err := _RICRegistry.contract.WatchLogs(opts, "providerUnstaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RICRegistryProviderUnstaked)
				if err := _RICRegistry.contract.UnpackLog(event, "providerUnstaked", log); err != nil {
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

// ParseProviderUnstaked is a log parse operation binding the contract event 0xc040f13890127304b60b2a24fa91a0e4877937f8426724681c2d0dfe7ff90207.
//
// Solidity: event providerUnstaked(address provider)
func (_RICRegistry *RICRegistryFilterer) ParseProviderUnstaked(log types.Log) (*RICRegistryProviderUnstaked, error) {
	event := new(RICRegistryProviderUnstaked)
	if err := _RICRegistry.contract.UnpackLog(event, "providerUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RICRegistryRollupActivatedIterator is returned from FilterRollupActivated and is used to iterate over the raw logs and unpacked data for RollupActivated events raised by the RICRegistry contract.
type RICRegistryRollupActivatedIterator struct {
	Event *RICRegistryRollupActivated // Event containing the contract specifics and raw log

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
func (it *RICRegistryRollupActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RICRegistryRollupActivated)
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
		it.Event = new(RICRegistryRollupActivated)
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
func (it *RICRegistryRollupActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RICRegistryRollupActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RICRegistryRollupActivated represents a RollupActivated event raised by the RICRegistry contract.
type RICRegistryRollupActivated struct {
	ChainID  *big.Int
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRollupActivated is a free log retrieval operation binding the contract event 0x566940d2cdad4df4f1073f7b1e2e72d9b96d30c5a69432c1432059fcc0c9266f.
//
// Solidity: event rollupActivated(uint256 chainID, address provider)
func (_RICRegistry *RICRegistryFilterer) FilterRollupActivated(opts *bind.FilterOpts) (*RICRegistryRollupActivatedIterator, error) {

	logs, sub, err := _RICRegistry.contract.FilterLogs(opts, "rollupActivated")
	if err != nil {
		return nil, err
	}
	return &RICRegistryRollupActivatedIterator{contract: _RICRegistry.contract, event: "rollupActivated", logs: logs, sub: sub}, nil
}

// WatchRollupActivated is a free log subscription operation binding the contract event 0x566940d2cdad4df4f1073f7b1e2e72d9b96d30c5a69432c1432059fcc0c9266f.
//
// Solidity: event rollupActivated(uint256 chainID, address provider)
func (_RICRegistry *RICRegistryFilterer) WatchRollupActivated(opts *bind.WatchOpts, sink chan<- *RICRegistryRollupActivated) (event.Subscription, error) {

	logs, sub, err := _RICRegistry.contract.WatchLogs(opts, "rollupActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RICRegistryRollupActivated)
				if err := _RICRegistry.contract.UnpackLog(event, "rollupActivated", log); err != nil {
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

// ParseRollupActivated is a log parse operation binding the contract event 0x566940d2cdad4df4f1073f7b1e2e72d9b96d30c5a69432c1432059fcc0c9266f.
//
// Solidity: event rollupActivated(uint256 chainID, address provider)
func (_RICRegistry *RICRegistryFilterer) ParseRollupActivated(log types.Log) (*RICRegistryRollupActivated, error) {
	event := new(RICRegistryRollupActivated)
	if err := _RICRegistry.contract.UnpackLog(event, "rollupActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RICRegistryRollupQueuedIterator is returned from FilterRollupQueued and is used to iterate over the raw logs and unpacked data for RollupQueued events raised by the RICRegistry contract.
type RICRegistryRollupQueuedIterator struct {
	Event *RICRegistryRollupQueued // Event containing the contract specifics and raw log

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
func (it *RICRegistryRollupQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RICRegistryRollupQueued)
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
		it.Event = new(RICRegistryRollupQueued)
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
func (it *RICRegistryRollupQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RICRegistryRollupQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RICRegistryRollupQueued represents a RollupQueued event raised by the RICRegistry contract.
type RICRegistryRollupQueued struct {
	ChainID            *big.Int
	Provider           common.Address
	RequestedTimestamp *big.Int
	TimeoutTimestamp   *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRollupQueued is a free log retrieval operation binding the contract event 0xd470aeaba6a6d85fa12188085bef6f7834d5dee13359db2cf8a255c1825416b2.
//
// Solidity: event rollupQueued(uint256 chainID, address provider, uint256 requestedTimestamp, uint256 timeoutTimestamp)
func (_RICRegistry *RICRegistryFilterer) FilterRollupQueued(opts *bind.FilterOpts) (*RICRegistryRollupQueuedIterator, error) {

	logs, sub, err := _RICRegistry.contract.FilterLogs(opts, "rollupQueued")
	if err != nil {
		return nil, err
	}
	return &RICRegistryRollupQueuedIterator{contract: _RICRegistry.contract, event: "rollupQueued", logs: logs, sub: sub}, nil
}

// WatchRollupQueued is a free log subscription operation binding the contract event 0xd470aeaba6a6d85fa12188085bef6f7834d5dee13359db2cf8a255c1825416b2.
//
// Solidity: event rollupQueued(uint256 chainID, address provider, uint256 requestedTimestamp, uint256 timeoutTimestamp)
func (_RICRegistry *RICRegistryFilterer) WatchRollupQueued(opts *bind.WatchOpts, sink chan<- *RICRegistryRollupQueued) (event.Subscription, error) {

	logs, sub, err := _RICRegistry.contract.WatchLogs(opts, "rollupQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RICRegistryRollupQueued)
				if err := _RICRegistry.contract.UnpackLog(event, "rollupQueued", log); err != nil {
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

// ParseRollupQueued is a log parse operation binding the contract event 0xd470aeaba6a6d85fa12188085bef6f7834d5dee13359db2cf8a255c1825416b2.
//
// Solidity: event rollupQueued(uint256 chainID, address provider, uint256 requestedTimestamp, uint256 timeoutTimestamp)
func (_RICRegistry *RICRegistryFilterer) ParseRollupQueued(log types.Log) (*RICRegistryRollupQueued, error) {
	event := new(RICRegistryRollupQueued)
	if err := _RICRegistry.contract.UnpackLog(event, "rollupQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RICRegistryRollupRequestedIterator is returned from FilterRollupRequested and is used to iterate over the raw logs and unpacked data for RollupRequested events raised by the RICRegistry contract.
type RICRegistryRollupRequestedIterator struct {
	Event *RICRegistryRollupRequested // Event containing the contract specifics and raw log

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
func (it *RICRegistryRollupRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RICRegistryRollupRequested)
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
		it.Event = new(RICRegistryRollupRequested)
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
func (it *RICRegistryRollupRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RICRegistryRollupRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RICRegistryRollupRequested represents a RollupRequested event raised by the RICRegistry contract.
type RICRegistryRollupRequested struct {
	ChainID   *big.Int
	Requester common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRollupRequested is a free log retrieval operation binding the contract event 0x808f5da474ed309cd671fe67f68f30567ab860240c2c2a9449640cdc95b8f7cf.
//
// Solidity: event rollupRequested(uint256 chainID, address requester, uint256 timestamp)
func (_RICRegistry *RICRegistryFilterer) FilterRollupRequested(opts *bind.FilterOpts) (*RICRegistryRollupRequestedIterator, error) {

	logs, sub, err := _RICRegistry.contract.FilterLogs(opts, "rollupRequested")
	if err != nil {
		return nil, err
	}
	return &RICRegistryRollupRequestedIterator{contract: _RICRegistry.contract, event: "rollupRequested", logs: logs, sub: sub}, nil
}

// WatchRollupRequested is a free log subscription operation binding the contract event 0x808f5da474ed309cd671fe67f68f30567ab860240c2c2a9449640cdc95b8f7cf.
//
// Solidity: event rollupRequested(uint256 chainID, address requester, uint256 timestamp)
func (_RICRegistry *RICRegistryFilterer) WatchRollupRequested(opts *bind.WatchOpts, sink chan<- *RICRegistryRollupRequested) (event.Subscription, error) {

	logs, sub, err := _RICRegistry.contract.WatchLogs(opts, "rollupRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RICRegistryRollupRequested)
				if err := _RICRegistry.contract.UnpackLog(event, "rollupRequested", log); err != nil {
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

// ParseRollupRequested is a log parse operation binding the contract event 0x808f5da474ed309cd671fe67f68f30567ab860240c2c2a9449640cdc95b8f7cf.
//
// Solidity: event rollupRequested(uint256 chainID, address requester, uint256 timestamp)
func (_RICRegistry *RICRegistryFilterer) ParseRollupRequested(log types.Log) (*RICRegistryRollupRequested, error) {
	event := new(RICRegistryRollupRequested)
	if err := _RICRegistry.contract.UnpackLog(event, "rollupRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
