// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kvdb

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// KvdbABI is the input ABI used to generate the binding from.
const KvdbABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KvdbBin is the compiled bytecode used for deploying new contracts.
var KvdbBin = "0x608060405234801561001057600080fd5b5060df8061001f6000396000f3006080604052600436106049576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632e64cec114604e5780636057361d146076575b600080fd5b348015605957600080fd5b50606060a0565b6040518082815260200191505060405180910390f35b348015608157600080fd5b50609e6004803603810190808035906020019092919050505060a9565b005b60008054905090565b80600081905550505600a165627a7a723058209588bad8d19b8bd74babe043b24cf839c951bfc1ff8a98a2e9dd0087184bf1980029"

// DeployKvdb deploys a new Ethereum contract, binding an instance of Kvdb to it.
func DeployKvdb(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Kvdb, error) {
	parsed, err := abi.JSON(strings.NewReader(KvdbABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KvdbBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Kvdb{KvdbCaller: KvdbCaller{contract: contract}, KvdbTransactor: KvdbTransactor{contract: contract}, KvdbFilterer: KvdbFilterer{contract: contract}}, nil
}

// Kvdb is an auto generated Go binding around an Ethereum contract.
type Kvdb struct {
	KvdbCaller     // Read-only binding to the contract
	KvdbTransactor // Write-only binding to the contract
	KvdbFilterer   // Log filterer for contract events
}

// KvdbCaller is an auto generated read-only Go binding around an Ethereum contract.
type KvdbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KvdbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KvdbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KvdbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KvdbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KvdbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KvdbSession struct {
	Contract     *Kvdb             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KvdbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KvdbCallerSession struct {
	Contract *KvdbCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KvdbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KvdbTransactorSession struct {
	Contract     *KvdbTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KvdbRaw is an auto generated low-level Go binding around an Ethereum contract.
type KvdbRaw struct {
	Contract *Kvdb // Generic contract binding to access the raw methods on
}

// KvdbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KvdbCallerRaw struct {
	Contract *KvdbCaller // Generic read-only contract binding to access the raw methods on
}

// KvdbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KvdbTransactorRaw struct {
	Contract *KvdbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKvdb creates a new instance of Kvdb, bound to a specific deployed contract.
func NewKvdb(address common.Address, backend bind.ContractBackend) (*Kvdb, error) {
	contract, err := bindKvdb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kvdb{KvdbCaller: KvdbCaller{contract: contract}, KvdbTransactor: KvdbTransactor{contract: contract}, KvdbFilterer: KvdbFilterer{contract: contract}}, nil
}

// NewKvdbCaller creates a new read-only instance of Kvdb, bound to a specific deployed contract.
func NewKvdbCaller(address common.Address, caller bind.ContractCaller) (*KvdbCaller, error) {
	contract, err := bindKvdb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KvdbCaller{contract: contract}, nil
}

// NewKvdbTransactor creates a new write-only instance of Kvdb, bound to a specific deployed contract.
func NewKvdbTransactor(address common.Address, transactor bind.ContractTransactor) (*KvdbTransactor, error) {
	contract, err := bindKvdb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KvdbTransactor{contract: contract}, nil
}

// NewKvdbFilterer creates a new log filterer instance of Kvdb, bound to a specific deployed contract.
func NewKvdbFilterer(address common.Address, filterer bind.ContractFilterer) (*KvdbFilterer, error) {
	contract, err := bindKvdb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KvdbFilterer{contract: contract}, nil
}

// bindKvdb binds a generic wrapper to an already deployed contract.
func bindKvdb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KvdbABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kvdb *KvdbRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kvdb.Contract.KvdbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kvdb *KvdbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kvdb.Contract.KvdbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kvdb *KvdbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kvdb.Contract.KvdbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kvdb *KvdbCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kvdb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kvdb *KvdbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kvdb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kvdb *KvdbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kvdb.Contract.contract.Transact(opts, method, params...)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Kvdb *KvdbCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kvdb.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Kvdb *KvdbSession) Retrieve() (*big.Int, error) {
	return _Kvdb.Contract.Retrieve(&_Kvdb.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Kvdb *KvdbCallerSession) Retrieve() (*big.Int, error) {
	return _Kvdb.Contract.Retrieve(&_Kvdb.CallOpts)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_Kvdb *KvdbTransactor) Store(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _Kvdb.contract.Transact(opts, "store", num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_Kvdb *KvdbSession) Store(num *big.Int) (*types.Transaction, error) {
	return _Kvdb.Contract.Store(&_Kvdb.TransactOpts, num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_Kvdb *KvdbTransactorSession) Store(num *big.Int) (*types.Transaction, error) {
	return _Kvdb.Contract.Store(&_Kvdb.TransactOpts, num)
}
