// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// AccountABI is the input ABI used to generate the binding from.
const AccountABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountClass\",\"type\":\"string\"}],\"name\":\"addAccountClassEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountClass\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createtime\",\"type\":\"uint256\"}],\"name\":\"addKeyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"borrowKeyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountClass\",\"type\":\"string\"}],\"name\":\"deleteAccountClassEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountClass\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"deleteKeyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"doKeyReturnEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountClass\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldAccountClass\",\"type\":\"string\"}],\"name\":\"modifyAccountClassEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"modifyKeyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountClass\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldAccountClass\",\"type\":\"string\"}],\"name\":\"moveKeyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumKeyType\",\"name\":\"\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"setKeyByIdEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountClassListMap\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountClass\",\"type\":\"string\"}],\"name\":\"addAccountClassFunc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_accountClass\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_pass\",\"type\":\"string\"}],\"name\":\"addKeyFunc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"borrowKeyByIdFunc\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_oldClass\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_newClass\",\"type\":\"string\"}],\"name\":\"changeClass\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountClass\",\"type\":\"string\"}],\"name\":\"deleteAccountClassFunc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"deleteKeyByIdFunc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountClass\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"deleteKeyFunc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"doKeyReturnByIdFunc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"idKeyMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountClass\",\"type\":\"string\"},{\"internalType\":\"enumKeyType\",\"name\":\"keyType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isBorrow\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"ethPledge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountClass\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"isKeyFunc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isKeyIDFunc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"keyIdMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"keyIndexMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"keyListMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_oldClass\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_newClass\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"moveFunc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"returnClassAllFunc\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountClass\",\"type\":\"string\"}],\"name\":\"returnClassKeyIdAllFunc\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ethPledge\",\"type\":\"uint256\"}],\"name\":\"setKeyByIdFunc\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_pass\",\"type\":\"string\"}],\"name\":\"updateKeyByIdFunc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountClass\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_pass\",\"type\":\"string\"}],\"name\":\"updateKeyFunc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Account is an auto generated Go binding around an Ethereum contract.
type Account struct {
	AccountCaller     // Read-only binding to the contract
	AccountTransactor // Write-only binding to the contract
	AccountFilterer   // Log filterer for contract events
}

// AccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountSession struct {
	Contract     *Account          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountCallerSession struct {
	Contract *AccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountTransactorSession struct {
	Contract     *AccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountRaw struct {
	Contract *Account // Generic contract binding to access the raw methods on
}

// AccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountCallerRaw struct {
	Contract *AccountCaller // Generic read-only contract binding to access the raw methods on
}

// AccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountTransactorRaw struct {
	Contract *AccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccount creates a new instance of Account, bound to a specific deployed contract.
func NewAccount(address common.Address, backend bind.ContractBackend) (*Account, error) {
	contract, err := bindAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Account{AccountCaller: AccountCaller{contract: contract}, AccountTransactor: AccountTransactor{contract: contract}, AccountFilterer: AccountFilterer{contract: contract}}, nil
}

// NewAccountCaller creates a new read-only instance of Account, bound to a specific deployed contract.
func NewAccountCaller(address common.Address, caller bind.ContractCaller) (*AccountCaller, error) {
	contract, err := bindAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountCaller{contract: contract}, nil
}

// NewAccountTransactor creates a new write-only instance of Account, bound to a specific deployed contract.
func NewAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountTransactor, error) {
	contract, err := bindAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountTransactor{contract: contract}, nil
}

// NewAccountFilterer creates a new log filterer instance of Account, bound to a specific deployed contract.
func NewAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountFilterer, error) {
	contract, err := bindAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountFilterer{contract: contract}, nil
}

// bindAccount binds a generic wrapper to an already deployed contract.
func bindAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Account *AccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Account.Contract.AccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Account *AccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.Contract.AccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Account *AccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Account.Contract.AccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Account *AccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Account.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Account *AccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Account *AccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Account.Contract.contract.Transact(opts, method, params...)
}

// AccountClassListMap is a free data retrieval call binding the contract method 0xbbafddfb.
//
// Solidity: function accountClassListMap(address , uint256 ) view returns(string)
func (_Account *AccountCaller) AccountClassListMap(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "accountClassListMap", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AccountClassListMap is a free data retrieval call binding the contract method 0xbbafddfb.
//
// Solidity: function accountClassListMap(address , uint256 ) view returns(string)
func (_Account *AccountSession) AccountClassListMap(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Account.Contract.AccountClassListMap(&_Account.CallOpts, arg0, arg1)
}

// AccountClassListMap is a free data retrieval call binding the contract method 0xbbafddfb.
//
// Solidity: function accountClassListMap(address , uint256 ) view returns(string)
func (_Account *AccountCallerSession) AccountClassListMap(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Account.Contract.AccountClassListMap(&_Account.CallOpts, arg0, arg1)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_Account *AccountCaller) GetBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "getBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_Account *AccountSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _Account.Contract.GetBalance(&_Account.CallOpts, addr)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_Account *AccountCallerSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _Account.Contract.GetBalance(&_Account.CallOpts, addr)
}

// GetCurrentBalance is a free data retrieval call binding the contract method 0xa5749710.
//
// Solidity: function getCurrentBalance() view returns(uint256)
func (_Account *AccountCaller) GetCurrentBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "getCurrentBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBalance is a free data retrieval call binding the contract method 0xa5749710.
//
// Solidity: function getCurrentBalance() view returns(uint256)
func (_Account *AccountSession) GetCurrentBalance() (*big.Int, error) {
	return _Account.Contract.GetCurrentBalance(&_Account.CallOpts)
}

// GetCurrentBalance is a free data retrieval call binding the contract method 0xa5749710.
//
// Solidity: function getCurrentBalance() view returns(uint256)
func (_Account *AccountCallerSession) GetCurrentBalance() (*big.Int, error) {
	return _Account.Contract.GetCurrentBalance(&_Account.CallOpts)
}

// IdKeyMap is a free data retrieval call binding the contract method 0x5c65ee52.
//
// Solidity: function idKeyMap(uint256 ) view returns(uint256 id, address owner, address borrower, string key, string password, string accountClass, uint8 keyType, bool isBorrow, uint256 ethPledge, uint256 time)
func (_Account *AccountCaller) IdKeyMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id           *big.Int
	Owner        common.Address
	Borrower     common.Address
	Key          string
	Password     string
	AccountClass string
	KeyType      uint8
	IsBorrow     bool
	EthPledge    *big.Int
	Time         *big.Int
}, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "idKeyMap", arg0)

	outstruct := new(struct {
		Id           *big.Int
		Owner        common.Address
		Borrower     common.Address
		Key          string
		Password     string
		AccountClass string
		KeyType      uint8
		IsBorrow     bool
		EthPledge    *big.Int
		Time         *big.Int
	})

	outstruct.Id = out[0].(*big.Int)
	outstruct.Owner = out[1].(common.Address)
	outstruct.Borrower = out[2].(common.Address)
	outstruct.Key = out[3].(string)
	outstruct.Password = out[4].(string)
	outstruct.AccountClass = out[5].(string)
	outstruct.KeyType = out[6].(uint8)
	outstruct.IsBorrow = out[7].(bool)
	outstruct.EthPledge = out[8].(*big.Int)
	outstruct.Time = out[9].(*big.Int)

	return *outstruct, err

}

// IdKeyMap is a free data retrieval call binding the contract method 0x5c65ee52.
//
// Solidity: function idKeyMap(uint256 ) view returns(uint256 id, address owner, address borrower, string key, string password, string accountClass, uint8 keyType, bool isBorrow, uint256 ethPledge, uint256 time)
func (_Account *AccountSession) IdKeyMap(arg0 *big.Int) (struct {
	Id           *big.Int
	Owner        common.Address
	Borrower     common.Address
	Key          string
	Password     string
	AccountClass string
	KeyType      uint8
	IsBorrow     bool
	EthPledge    *big.Int
	Time         *big.Int
}, error) {
	return _Account.Contract.IdKeyMap(&_Account.CallOpts, arg0)
}

// IdKeyMap is a free data retrieval call binding the contract method 0x5c65ee52.
//
// Solidity: function idKeyMap(uint256 ) view returns(uint256 id, address owner, address borrower, string key, string password, string accountClass, uint8 keyType, bool isBorrow, uint256 ethPledge, uint256 time)
func (_Account *AccountCallerSession) IdKeyMap(arg0 *big.Int) (struct {
	Id           *big.Int
	Owner        common.Address
	Borrower     common.Address
	Key          string
	Password     string
	AccountClass string
	KeyType      uint8
	IsBorrow     bool
	EthPledge    *big.Int
	Time         *big.Int
}, error) {
	return _Account.Contract.IdKeyMap(&_Account.CallOpts, arg0)
}

// IsKeyFunc is a free data retrieval call binding the contract method 0x5f5264b6.
//
// Solidity: function isKeyFunc(string _accountClass, string _key) view returns(bool)
func (_Account *AccountCaller) IsKeyFunc(opts *bind.CallOpts, _accountClass string, _key string) (bool, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "isKeyFunc", _accountClass, _key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKeyFunc is a free data retrieval call binding the contract method 0x5f5264b6.
//
// Solidity: function isKeyFunc(string _accountClass, string _key) view returns(bool)
func (_Account *AccountSession) IsKeyFunc(_accountClass string, _key string) (bool, error) {
	return _Account.Contract.IsKeyFunc(&_Account.CallOpts, _accountClass, _key)
}

// IsKeyFunc is a free data retrieval call binding the contract method 0x5f5264b6.
//
// Solidity: function isKeyFunc(string _accountClass, string _key) view returns(bool)
func (_Account *AccountCallerSession) IsKeyFunc(_accountClass string, _key string) (bool, error) {
	return _Account.Contract.IsKeyFunc(&_Account.CallOpts, _accountClass, _key)
}

// IsKeyIDFunc is a free data retrieval call binding the contract method 0xf8464cda.
//
// Solidity: function isKeyIDFunc(uint256 _id) view returns(bool)
func (_Account *AccountCaller) IsKeyIDFunc(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "isKeyIDFunc", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKeyIDFunc is a free data retrieval call binding the contract method 0xf8464cda.
//
// Solidity: function isKeyIDFunc(uint256 _id) view returns(bool)
func (_Account *AccountSession) IsKeyIDFunc(_id *big.Int) (bool, error) {
	return _Account.Contract.IsKeyIDFunc(&_Account.CallOpts, _id)
}

// IsKeyIDFunc is a free data retrieval call binding the contract method 0xf8464cda.
//
// Solidity: function isKeyIDFunc(uint256 _id) view returns(bool)
func (_Account *AccountCallerSession) IsKeyIDFunc(_id *big.Int) (bool, error) {
	return _Account.Contract.IsKeyIDFunc(&_Account.CallOpts, _id)
}

// KeyIdMap is a free data retrieval call binding the contract method 0x753671ce.
//
// Solidity: function keyIdMap(address , string , string ) view returns(uint256)
func (_Account *AccountCaller) KeyIdMap(opts *bind.CallOpts, arg0 common.Address, arg1 string, arg2 string) (*big.Int, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "keyIdMap", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KeyIdMap is a free data retrieval call binding the contract method 0x753671ce.
//
// Solidity: function keyIdMap(address , string , string ) view returns(uint256)
func (_Account *AccountSession) KeyIdMap(arg0 common.Address, arg1 string, arg2 string) (*big.Int, error) {
	return _Account.Contract.KeyIdMap(&_Account.CallOpts, arg0, arg1, arg2)
}

// KeyIdMap is a free data retrieval call binding the contract method 0x753671ce.
//
// Solidity: function keyIdMap(address , string , string ) view returns(uint256)
func (_Account *AccountCallerSession) KeyIdMap(arg0 common.Address, arg1 string, arg2 string) (*big.Int, error) {
	return _Account.Contract.KeyIdMap(&_Account.CallOpts, arg0, arg1, arg2)
}

// KeyIndexMap is a free data retrieval call binding the contract method 0x47c73524.
//
// Solidity: function keyIndexMap(address , string , string ) view returns(uint256)
func (_Account *AccountCaller) KeyIndexMap(opts *bind.CallOpts, arg0 common.Address, arg1 string, arg2 string) (*big.Int, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "keyIndexMap", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KeyIndexMap is a free data retrieval call binding the contract method 0x47c73524.
//
// Solidity: function keyIndexMap(address , string , string ) view returns(uint256)
func (_Account *AccountSession) KeyIndexMap(arg0 common.Address, arg1 string, arg2 string) (*big.Int, error) {
	return _Account.Contract.KeyIndexMap(&_Account.CallOpts, arg0, arg1, arg2)
}

// KeyIndexMap is a free data retrieval call binding the contract method 0x47c73524.
//
// Solidity: function keyIndexMap(address , string , string ) view returns(uint256)
func (_Account *AccountCallerSession) KeyIndexMap(arg0 common.Address, arg1 string, arg2 string) (*big.Int, error) {
	return _Account.Contract.KeyIndexMap(&_Account.CallOpts, arg0, arg1, arg2)
}

// KeyListMap is a free data retrieval call binding the contract method 0x61d83e0c.
//
// Solidity: function keyListMap(address , string , uint256 ) view returns(uint256)
func (_Account *AccountCaller) KeyListMap(opts *bind.CallOpts, arg0 common.Address, arg1 string, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "keyListMap", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KeyListMap is a free data retrieval call binding the contract method 0x61d83e0c.
//
// Solidity: function keyListMap(address , string , uint256 ) view returns(uint256)
func (_Account *AccountSession) KeyListMap(arg0 common.Address, arg1 string, arg2 *big.Int) (*big.Int, error) {
	return _Account.Contract.KeyListMap(&_Account.CallOpts, arg0, arg1, arg2)
}

// KeyListMap is a free data retrieval call binding the contract method 0x61d83e0c.
//
// Solidity: function keyListMap(address , string , uint256 ) view returns(uint256)
func (_Account *AccountCallerSession) KeyListMap(arg0 common.Address, arg1 string, arg2 *big.Int) (*big.Int, error) {
	return _Account.Contract.KeyListMap(&_Account.CallOpts, arg0, arg1, arg2)
}

// ReturnClassAllFunc is a free data retrieval call binding the contract method 0x1f86a59b.
//
// Solidity: function returnClassAllFunc() view returns(string[])
func (_Account *AccountCaller) ReturnClassAllFunc(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "returnClassAllFunc")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ReturnClassAllFunc is a free data retrieval call binding the contract method 0x1f86a59b.
//
// Solidity: function returnClassAllFunc() view returns(string[])
func (_Account *AccountSession) ReturnClassAllFunc() ([]string, error) {
	return _Account.Contract.ReturnClassAllFunc(&_Account.CallOpts)
}

// ReturnClassAllFunc is a free data retrieval call binding the contract method 0x1f86a59b.
//
// Solidity: function returnClassAllFunc() view returns(string[])
func (_Account *AccountCallerSession) ReturnClassAllFunc() ([]string, error) {
	return _Account.Contract.ReturnClassAllFunc(&_Account.CallOpts)
}

// ReturnClassKeyIdAllFunc is a free data retrieval call binding the contract method 0x7a8a159b.
//
// Solidity: function returnClassKeyIdAllFunc(string _accountClass) view returns(uint256[])
func (_Account *AccountCaller) ReturnClassKeyIdAllFunc(opts *bind.CallOpts, _accountClass string) ([]*big.Int, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "returnClassKeyIdAllFunc", _accountClass)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ReturnClassKeyIdAllFunc is a free data retrieval call binding the contract method 0x7a8a159b.
//
// Solidity: function returnClassKeyIdAllFunc(string _accountClass) view returns(uint256[])
func (_Account *AccountSession) ReturnClassKeyIdAllFunc(_accountClass string) ([]*big.Int, error) {
	return _Account.Contract.ReturnClassKeyIdAllFunc(&_Account.CallOpts, _accountClass)
}

// ReturnClassKeyIdAllFunc is a free data retrieval call binding the contract method 0x7a8a159b.
//
// Solidity: function returnClassKeyIdAllFunc(string _accountClass) view returns(uint256[])
func (_Account *AccountCallerSession) ReturnClassKeyIdAllFunc(_accountClass string) ([]*big.Int, error) {
	return _Account.Contract.ReturnClassKeyIdAllFunc(&_Account.CallOpts, _accountClass)
}

// AddAccountClassFunc is a paid mutator transaction binding the contract method 0xb8b897ee.
//
// Solidity: function addAccountClassFunc(string _accountClass) returns()
func (_Account *AccountTransactor) AddAccountClassFunc(opts *bind.TransactOpts, _accountClass string) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "addAccountClassFunc", _accountClass)
}

// AddAccountClassFunc is a paid mutator transaction binding the contract method 0xb8b897ee.
//
// Solidity: function addAccountClassFunc(string _accountClass) returns()
func (_Account *AccountSession) AddAccountClassFunc(_accountClass string) (*types.Transaction, error) {
	return _Account.Contract.AddAccountClassFunc(&_Account.TransactOpts, _accountClass)
}

// AddAccountClassFunc is a paid mutator transaction binding the contract method 0xb8b897ee.
//
// Solidity: function addAccountClassFunc(string _accountClass) returns()
func (_Account *AccountTransactorSession) AddAccountClassFunc(_accountClass string) (*types.Transaction, error) {
	return _Account.Contract.AddAccountClassFunc(&_Account.TransactOpts, _accountClass)
}

// AddKeyFunc is a paid mutator transaction binding the contract method 0x6c457baf.
//
// Solidity: function addKeyFunc(uint256 _id, string _accountClass, string _key, string _pass) returns()
func (_Account *AccountTransactor) AddKeyFunc(opts *bind.TransactOpts, _id *big.Int, _accountClass string, _key string, _pass string) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "addKeyFunc", _id, _accountClass, _key, _pass)
}

// AddKeyFunc is a paid mutator transaction binding the contract method 0x6c457baf.
//
// Solidity: function addKeyFunc(uint256 _id, string _accountClass, string _key, string _pass) returns()
func (_Account *AccountSession) AddKeyFunc(_id *big.Int, _accountClass string, _key string, _pass string) (*types.Transaction, error) {
	return _Account.Contract.AddKeyFunc(&_Account.TransactOpts, _id, _accountClass, _key, _pass)
}

// AddKeyFunc is a paid mutator transaction binding the contract method 0x6c457baf.
//
// Solidity: function addKeyFunc(uint256 _id, string _accountClass, string _key, string _pass) returns()
func (_Account *AccountTransactorSession) AddKeyFunc(_id *big.Int, _accountClass string, _key string, _pass string) (*types.Transaction, error) {
	return _Account.Contract.AddKeyFunc(&_Account.TransactOpts, _id, _accountClass, _key, _pass)
}

// BorrowKeyByIdFunc is a paid mutator transaction binding the contract method 0xb1ddfbe8.
//
// Solidity: function borrowKeyByIdFunc(uint256 _id) payable returns()
func (_Account *AccountTransactor) BorrowKeyByIdFunc(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "borrowKeyByIdFunc", _id)
}

// BorrowKeyByIdFunc is a paid mutator transaction binding the contract method 0xb1ddfbe8.
//
// Solidity: function borrowKeyByIdFunc(uint256 _id) payable returns()
func (_Account *AccountSession) BorrowKeyByIdFunc(_id *big.Int) (*types.Transaction, error) {
	return _Account.Contract.BorrowKeyByIdFunc(&_Account.TransactOpts, _id)
}

// BorrowKeyByIdFunc is a paid mutator transaction binding the contract method 0xb1ddfbe8.
//
// Solidity: function borrowKeyByIdFunc(uint256 _id) payable returns()
func (_Account *AccountTransactorSession) BorrowKeyByIdFunc(_id *big.Int) (*types.Transaction, error) {
	return _Account.Contract.BorrowKeyByIdFunc(&_Account.TransactOpts, _id)
}

// ChangeClass is a paid mutator transaction binding the contract method 0xecc8a70f.
//
// Solidity: function changeClass(string _oldClass, string _newClass) returns()
func (_Account *AccountTransactor) ChangeClass(opts *bind.TransactOpts, _oldClass string, _newClass string) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "changeClass", _oldClass, _newClass)
}

// ChangeClass is a paid mutator transaction binding the contract method 0xecc8a70f.
//
// Solidity: function changeClass(string _oldClass, string _newClass) returns()
func (_Account *AccountSession) ChangeClass(_oldClass string, _newClass string) (*types.Transaction, error) {
	return _Account.Contract.ChangeClass(&_Account.TransactOpts, _oldClass, _newClass)
}

// ChangeClass is a paid mutator transaction binding the contract method 0xecc8a70f.
//
// Solidity: function changeClass(string _oldClass, string _newClass) returns()
func (_Account *AccountTransactorSession) ChangeClass(_oldClass string, _newClass string) (*types.Transaction, error) {
	return _Account.Contract.ChangeClass(&_Account.TransactOpts, _oldClass, _newClass)
}

// DeleteAccountClassFunc is a paid mutator transaction binding the contract method 0x59b38991.
//
// Solidity: function deleteAccountClassFunc(string _accountClass) returns()
func (_Account *AccountTransactor) DeleteAccountClassFunc(opts *bind.TransactOpts, _accountClass string) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "deleteAccountClassFunc", _accountClass)
}

// DeleteAccountClassFunc is a paid mutator transaction binding the contract method 0x59b38991.
//
// Solidity: function deleteAccountClassFunc(string _accountClass) returns()
func (_Account *AccountSession) DeleteAccountClassFunc(_accountClass string) (*types.Transaction, error) {
	return _Account.Contract.DeleteAccountClassFunc(&_Account.TransactOpts, _accountClass)
}

// DeleteAccountClassFunc is a paid mutator transaction binding the contract method 0x59b38991.
//
// Solidity: function deleteAccountClassFunc(string _accountClass) returns()
func (_Account *AccountTransactorSession) DeleteAccountClassFunc(_accountClass string) (*types.Transaction, error) {
	return _Account.Contract.DeleteAccountClassFunc(&_Account.TransactOpts, _accountClass)
}

// DeleteKeyByIdFunc is a paid mutator transaction binding the contract method 0x05463bd8.
//
// Solidity: function deleteKeyByIdFunc(uint256 _id) returns()
func (_Account *AccountTransactor) DeleteKeyByIdFunc(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "deleteKeyByIdFunc", _id)
}

// DeleteKeyByIdFunc is a paid mutator transaction binding the contract method 0x05463bd8.
//
// Solidity: function deleteKeyByIdFunc(uint256 _id) returns()
func (_Account *AccountSession) DeleteKeyByIdFunc(_id *big.Int) (*types.Transaction, error) {
	return _Account.Contract.DeleteKeyByIdFunc(&_Account.TransactOpts, _id)
}

// DeleteKeyByIdFunc is a paid mutator transaction binding the contract method 0x05463bd8.
//
// Solidity: function deleteKeyByIdFunc(uint256 _id) returns()
func (_Account *AccountTransactorSession) DeleteKeyByIdFunc(_id *big.Int) (*types.Transaction, error) {
	return _Account.Contract.DeleteKeyByIdFunc(&_Account.TransactOpts, _id)
}

// DeleteKeyFunc is a paid mutator transaction binding the contract method 0x072d2606.
//
// Solidity: function deleteKeyFunc(string _accountClass, string _key) returns()
func (_Account *AccountTransactor) DeleteKeyFunc(opts *bind.TransactOpts, _accountClass string, _key string) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "deleteKeyFunc", _accountClass, _key)
}

// DeleteKeyFunc is a paid mutator transaction binding the contract method 0x072d2606.
//
// Solidity: function deleteKeyFunc(string _accountClass, string _key) returns()
func (_Account *AccountSession) DeleteKeyFunc(_accountClass string, _key string) (*types.Transaction, error) {
	return _Account.Contract.DeleteKeyFunc(&_Account.TransactOpts, _accountClass, _key)
}

// DeleteKeyFunc is a paid mutator transaction binding the contract method 0x072d2606.
//
// Solidity: function deleteKeyFunc(string _accountClass, string _key) returns()
func (_Account *AccountTransactorSession) DeleteKeyFunc(_accountClass string, _key string) (*types.Transaction, error) {
	return _Account.Contract.DeleteKeyFunc(&_Account.TransactOpts, _accountClass, _key)
}

// DoKeyReturnByIdFunc is a paid mutator transaction binding the contract method 0x2c551d7c.
//
// Solidity: function doKeyReturnByIdFunc(uint256 _id) returns()
func (_Account *AccountTransactor) DoKeyReturnByIdFunc(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "doKeyReturnByIdFunc", _id)
}

// DoKeyReturnByIdFunc is a paid mutator transaction binding the contract method 0x2c551d7c.
//
// Solidity: function doKeyReturnByIdFunc(uint256 _id) returns()
func (_Account *AccountSession) DoKeyReturnByIdFunc(_id *big.Int) (*types.Transaction, error) {
	return _Account.Contract.DoKeyReturnByIdFunc(&_Account.TransactOpts, _id)
}

// DoKeyReturnByIdFunc is a paid mutator transaction binding the contract method 0x2c551d7c.
//
// Solidity: function doKeyReturnByIdFunc(uint256 _id) returns()
func (_Account *AccountTransactorSession) DoKeyReturnByIdFunc(_id *big.Int) (*types.Transaction, error) {
	return _Account.Contract.DoKeyReturnByIdFunc(&_Account.TransactOpts, _id)
}

// MoveFunc is a paid mutator transaction binding the contract method 0x18928e38.
//
// Solidity: function moveFunc(string _oldClass, string _newClass, string _key) returns()
func (_Account *AccountTransactor) MoveFunc(opts *bind.TransactOpts, _oldClass string, _newClass string, _key string) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "moveFunc", _oldClass, _newClass, _key)
}

// MoveFunc is a paid mutator transaction binding the contract method 0x18928e38.
//
// Solidity: function moveFunc(string _oldClass, string _newClass, string _key) returns()
func (_Account *AccountSession) MoveFunc(_oldClass string, _newClass string, _key string) (*types.Transaction, error) {
	return _Account.Contract.MoveFunc(&_Account.TransactOpts, _oldClass, _newClass, _key)
}

// MoveFunc is a paid mutator transaction binding the contract method 0x18928e38.
//
// Solidity: function moveFunc(string _oldClass, string _newClass, string _key) returns()
func (_Account *AccountTransactorSession) MoveFunc(_oldClass string, _newClass string, _key string) (*types.Transaction, error) {
	return _Account.Contract.MoveFunc(&_Account.TransactOpts, _oldClass, _newClass, _key)
}

// SetKeyByIdFunc is a paid mutator transaction binding the contract method 0x50ecb2ba.
//
// Solidity: function setKeyByIdFunc(uint256 _id, uint256 _ethPledge) returns(uint8)
func (_Account *AccountTransactor) SetKeyByIdFunc(opts *bind.TransactOpts, _id *big.Int, _ethPledge *big.Int) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "setKeyByIdFunc", _id, _ethPledge)
}

// SetKeyByIdFunc is a paid mutator transaction binding the contract method 0x50ecb2ba.
//
// Solidity: function setKeyByIdFunc(uint256 _id, uint256 _ethPledge) returns(uint8)
func (_Account *AccountSession) SetKeyByIdFunc(_id *big.Int, _ethPledge *big.Int) (*types.Transaction, error) {
	return _Account.Contract.SetKeyByIdFunc(&_Account.TransactOpts, _id, _ethPledge)
}

// SetKeyByIdFunc is a paid mutator transaction binding the contract method 0x50ecb2ba.
//
// Solidity: function setKeyByIdFunc(uint256 _id, uint256 _ethPledge) returns(uint8)
func (_Account *AccountTransactorSession) SetKeyByIdFunc(_id *big.Int, _ethPledge *big.Int) (*types.Transaction, error) {
	return _Account.Contract.SetKeyByIdFunc(&_Account.TransactOpts, _id, _ethPledge)
}

// UpdateKeyByIdFunc is a paid mutator transaction binding the contract method 0x049b29fe.
//
// Solidity: function updateKeyByIdFunc(uint256 _id, string _pass) returns()
func (_Account *AccountTransactor) UpdateKeyByIdFunc(opts *bind.TransactOpts, _id *big.Int, _pass string) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "updateKeyByIdFunc", _id, _pass)
}

// UpdateKeyByIdFunc is a paid mutator transaction binding the contract method 0x049b29fe.
//
// Solidity: function updateKeyByIdFunc(uint256 _id, string _pass) returns()
func (_Account *AccountSession) UpdateKeyByIdFunc(_id *big.Int, _pass string) (*types.Transaction, error) {
	return _Account.Contract.UpdateKeyByIdFunc(&_Account.TransactOpts, _id, _pass)
}

// UpdateKeyByIdFunc is a paid mutator transaction binding the contract method 0x049b29fe.
//
// Solidity: function updateKeyByIdFunc(uint256 _id, string _pass) returns()
func (_Account *AccountTransactorSession) UpdateKeyByIdFunc(_id *big.Int, _pass string) (*types.Transaction, error) {
	return _Account.Contract.UpdateKeyByIdFunc(&_Account.TransactOpts, _id, _pass)
}

// UpdateKeyFunc is a paid mutator transaction binding the contract method 0x3f2bf9f5.
//
// Solidity: function updateKeyFunc(string _accountClass, string _key, string _pass) returns()
func (_Account *AccountTransactor) UpdateKeyFunc(opts *bind.TransactOpts, _accountClass string, _key string, _pass string) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "updateKeyFunc", _accountClass, _key, _pass)
}

// UpdateKeyFunc is a paid mutator transaction binding the contract method 0x3f2bf9f5.
//
// Solidity: function updateKeyFunc(string _accountClass, string _key, string _pass) returns()
func (_Account *AccountSession) UpdateKeyFunc(_accountClass string, _key string, _pass string) (*types.Transaction, error) {
	return _Account.Contract.UpdateKeyFunc(&_Account.TransactOpts, _accountClass, _key, _pass)
}

// UpdateKeyFunc is a paid mutator transaction binding the contract method 0x3f2bf9f5.
//
// Solidity: function updateKeyFunc(string _accountClass, string _key, string _pass) returns()
func (_Account *AccountTransactorSession) UpdateKeyFunc(_accountClass string, _key string, _pass string) (*types.Transaction, error) {
	return _Account.Contract.UpdateKeyFunc(&_Account.TransactOpts, _accountClass, _key, _pass)
}

// AccountAddAccountClassEventIterator is returned from FilterAddAccountClassEvent and is used to iterate over the raw logs and unpacked data for AddAccountClassEvent events raised by the Account contract.
type AccountAddAccountClassEventIterator struct {
	Event *AccountAddAccountClassEvent // Event containing the contract specifics and raw log

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
func (it *AccountAddAccountClassEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountAddAccountClassEvent)
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
		it.Event = new(AccountAddAccountClassEvent)
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
func (it *AccountAddAccountClassEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountAddAccountClassEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountAddAccountClassEvent represents a AddAccountClassEvent event raised by the Account contract.
type AccountAddAccountClassEvent struct {
	UserAddr     common.Address
	AccountClass string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddAccountClassEvent is a free log retrieval operation binding the contract event 0xff9c526a54aa4cc5ca817392703c2ab34e2165831a71c7a72bce59cce5d99d1b.
//
// Solidity: event addAccountClassEvent(address indexed userAddr, string accountClass)
func (_Account *AccountFilterer) FilterAddAccountClassEvent(opts *bind.FilterOpts, userAddr []common.Address) (*AccountAddAccountClassEventIterator, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Account.contract.FilterLogs(opts, "addAccountClassEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return &AccountAddAccountClassEventIterator{contract: _Account.contract, event: "addAccountClassEvent", logs: logs, sub: sub}, nil
}

// WatchAddAccountClassEvent is a free log subscription operation binding the contract event 0xff9c526a54aa4cc5ca817392703c2ab34e2165831a71c7a72bce59cce5d99d1b.
//
// Solidity: event addAccountClassEvent(address indexed userAddr, string accountClass)
func (_Account *AccountFilterer) WatchAddAccountClassEvent(opts *bind.WatchOpts, sink chan<- *AccountAddAccountClassEvent, userAddr []common.Address) (event.Subscription, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Account.contract.WatchLogs(opts, "addAccountClassEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountAddAccountClassEvent)
				if err := _Account.contract.UnpackLog(event, "addAccountClassEvent", log); err != nil {
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

// ParseAddAccountClassEvent is a log parse operation binding the contract event 0xff9c526a54aa4cc5ca817392703c2ab34e2165831a71c7a72bce59cce5d99d1b.
//
// Solidity: event addAccountClassEvent(address indexed userAddr, string accountClass)
func (_Account *AccountFilterer) ParseAddAccountClassEvent(log types.Log) (*AccountAddAccountClassEvent, error) {
	event := new(AccountAddAccountClassEvent)
	if err := _Account.contract.UnpackLog(event, "addAccountClassEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountAddKeyEventIterator is returned from FilterAddKeyEvent and is used to iterate over the raw logs and unpacked data for AddKeyEvent events raised by the Account contract.
type AccountAddKeyEventIterator struct {
	Event *AccountAddKeyEvent // Event containing the contract specifics and raw log

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
func (it *AccountAddKeyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountAddKeyEvent)
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
		it.Event = new(AccountAddKeyEvent)
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
func (it *AccountAddKeyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountAddKeyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountAddKeyEvent represents a AddKeyEvent event raised by the Account contract.
type AccountAddKeyEvent struct {
	UserAddr     common.Address
	AccountClass string
	Key          string
	Createtime   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddKeyEvent is a free log retrieval operation binding the contract event 0x89d013651b97ee3324a9472ad9ad46b7a7de6a77a8f7b055864ffadcabeb76a8.
//
// Solidity: event addKeyEvent(address indexed userAddr, string accountClass, string key, uint256 createtime)
func (_Account *AccountFilterer) FilterAddKeyEvent(opts *bind.FilterOpts, userAddr []common.Address) (*AccountAddKeyEventIterator, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Account.contract.FilterLogs(opts, "addKeyEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return &AccountAddKeyEventIterator{contract: _Account.contract, event: "addKeyEvent", logs: logs, sub: sub}, nil
}

// WatchAddKeyEvent is a free log subscription operation binding the contract event 0x89d013651b97ee3324a9472ad9ad46b7a7de6a77a8f7b055864ffadcabeb76a8.
//
// Solidity: event addKeyEvent(address indexed userAddr, string accountClass, string key, uint256 createtime)
func (_Account *AccountFilterer) WatchAddKeyEvent(opts *bind.WatchOpts, sink chan<- *AccountAddKeyEvent, userAddr []common.Address) (event.Subscription, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Account.contract.WatchLogs(opts, "addKeyEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountAddKeyEvent)
				if err := _Account.contract.UnpackLog(event, "addKeyEvent", log); err != nil {
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

// ParseAddKeyEvent is a log parse operation binding the contract event 0x89d013651b97ee3324a9472ad9ad46b7a7de6a77a8f7b055864ffadcabeb76a8.
//
// Solidity: event addKeyEvent(address indexed userAddr, string accountClass, string key, uint256 createtime)
func (_Account *AccountFilterer) ParseAddKeyEvent(log types.Log) (*AccountAddKeyEvent, error) {
	event := new(AccountAddKeyEvent)
	if err := _Account.contract.UnpackLog(event, "addKeyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountBorrowKeyEventIterator is returned from FilterBorrowKeyEvent and is used to iterate over the raw logs and unpacked data for BorrowKeyEvent events raised by the Account contract.
type AccountBorrowKeyEventIterator struct {
	Event *AccountBorrowKeyEvent // Event containing the contract specifics and raw log

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
func (it *AccountBorrowKeyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountBorrowKeyEvent)
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
		it.Event = new(AccountBorrowKeyEvent)
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
func (it *AccountBorrowKeyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountBorrowKeyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountBorrowKeyEvent represents a BorrowKeyEvent event raised by the Account contract.
type AccountBorrowKeyEvent struct {
	Indexd *big.Int
	Arg1   *big.Int
	Arg2   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBorrowKeyEvent is a free log retrieval operation binding the contract event 0x4458841029f40dc19d4df036baabe668ed4b8bae222b6ff72b75ae4908bb6b12.
//
// Solidity: event borrowKeyEvent(uint256 indexd, uint256 arg1, uint256 arg2)
func (_Account *AccountFilterer) FilterBorrowKeyEvent(opts *bind.FilterOpts) (*AccountBorrowKeyEventIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "borrowKeyEvent")
	if err != nil {
		return nil, err
	}
	return &AccountBorrowKeyEventIterator{contract: _Account.contract, event: "borrowKeyEvent", logs: logs, sub: sub}, nil
}

// WatchBorrowKeyEvent is a free log subscription operation binding the contract event 0x4458841029f40dc19d4df036baabe668ed4b8bae222b6ff72b75ae4908bb6b12.
//
// Solidity: event borrowKeyEvent(uint256 indexd, uint256 arg1, uint256 arg2)
func (_Account *AccountFilterer) WatchBorrowKeyEvent(opts *bind.WatchOpts, sink chan<- *AccountBorrowKeyEvent) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "borrowKeyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountBorrowKeyEvent)
				if err := _Account.contract.UnpackLog(event, "borrowKeyEvent", log); err != nil {
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

// ParseBorrowKeyEvent is a log parse operation binding the contract event 0x4458841029f40dc19d4df036baabe668ed4b8bae222b6ff72b75ae4908bb6b12.
//
// Solidity: event borrowKeyEvent(uint256 indexd, uint256 arg1, uint256 arg2)
func (_Account *AccountFilterer) ParseBorrowKeyEvent(log types.Log) (*AccountBorrowKeyEvent, error) {
	event := new(AccountBorrowKeyEvent)
	if err := _Account.contract.UnpackLog(event, "borrowKeyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountDeleteAccountClassEventIterator is returned from FilterDeleteAccountClassEvent and is used to iterate over the raw logs and unpacked data for DeleteAccountClassEvent events raised by the Account contract.
type AccountDeleteAccountClassEventIterator struct {
	Event *AccountDeleteAccountClassEvent // Event containing the contract specifics and raw log

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
func (it *AccountDeleteAccountClassEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountDeleteAccountClassEvent)
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
		it.Event = new(AccountDeleteAccountClassEvent)
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
func (it *AccountDeleteAccountClassEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountDeleteAccountClassEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountDeleteAccountClassEvent represents a DeleteAccountClassEvent event raised by the Account contract.
type AccountDeleteAccountClassEvent struct {
	UserAddr     common.Address
	AccountClass string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeleteAccountClassEvent is a free log retrieval operation binding the contract event 0x716211ec08d3c407d6e6cb80002fe85986af54c54c12ca51fffe268117e9d683.
//
// Solidity: event deleteAccountClassEvent(address indexed userAddr, string accountClass)
func (_Account *AccountFilterer) FilterDeleteAccountClassEvent(opts *bind.FilterOpts, userAddr []common.Address) (*AccountDeleteAccountClassEventIterator, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Account.contract.FilterLogs(opts, "deleteAccountClassEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return &AccountDeleteAccountClassEventIterator{contract: _Account.contract, event: "deleteAccountClassEvent", logs: logs, sub: sub}, nil
}

// WatchDeleteAccountClassEvent is a free log subscription operation binding the contract event 0x716211ec08d3c407d6e6cb80002fe85986af54c54c12ca51fffe268117e9d683.
//
// Solidity: event deleteAccountClassEvent(address indexed userAddr, string accountClass)
func (_Account *AccountFilterer) WatchDeleteAccountClassEvent(opts *bind.WatchOpts, sink chan<- *AccountDeleteAccountClassEvent, userAddr []common.Address) (event.Subscription, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Account.contract.WatchLogs(opts, "deleteAccountClassEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountDeleteAccountClassEvent)
				if err := _Account.contract.UnpackLog(event, "deleteAccountClassEvent", log); err != nil {
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

// ParseDeleteAccountClassEvent is a log parse operation binding the contract event 0x716211ec08d3c407d6e6cb80002fe85986af54c54c12ca51fffe268117e9d683.
//
// Solidity: event deleteAccountClassEvent(address indexed userAddr, string accountClass)
func (_Account *AccountFilterer) ParseDeleteAccountClassEvent(log types.Log) (*AccountDeleteAccountClassEvent, error) {
	event := new(AccountDeleteAccountClassEvent)
	if err := _Account.contract.UnpackLog(event, "deleteAccountClassEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountDeleteKeyEventIterator is returned from FilterDeleteKeyEvent and is used to iterate over the raw logs and unpacked data for DeleteKeyEvent events raised by the Account contract.
type AccountDeleteKeyEventIterator struct {
	Event *AccountDeleteKeyEvent // Event containing the contract specifics and raw log

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
func (it *AccountDeleteKeyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountDeleteKeyEvent)
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
		it.Event = new(AccountDeleteKeyEvent)
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
func (it *AccountDeleteKeyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountDeleteKeyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountDeleteKeyEvent represents a DeleteKeyEvent event raised by the Account contract.
type AccountDeleteKeyEvent struct {
	UserAddr     common.Address
	AccountClass string
	Key          string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeleteKeyEvent is a free log retrieval operation binding the contract event 0x2563b537df5e859d4928c3f28220f6c5e6aab2948198de0bd936fc5e7750eed1.
//
// Solidity: event deleteKeyEvent(address indexed userAddr, string accountClass, string key)
func (_Account *AccountFilterer) FilterDeleteKeyEvent(opts *bind.FilterOpts, userAddr []common.Address) (*AccountDeleteKeyEventIterator, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Account.contract.FilterLogs(opts, "deleteKeyEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return &AccountDeleteKeyEventIterator{contract: _Account.contract, event: "deleteKeyEvent", logs: logs, sub: sub}, nil
}

// WatchDeleteKeyEvent is a free log subscription operation binding the contract event 0x2563b537df5e859d4928c3f28220f6c5e6aab2948198de0bd936fc5e7750eed1.
//
// Solidity: event deleteKeyEvent(address indexed userAddr, string accountClass, string key)
func (_Account *AccountFilterer) WatchDeleteKeyEvent(opts *bind.WatchOpts, sink chan<- *AccountDeleteKeyEvent, userAddr []common.Address) (event.Subscription, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Account.contract.WatchLogs(opts, "deleteKeyEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountDeleteKeyEvent)
				if err := _Account.contract.UnpackLog(event, "deleteKeyEvent", log); err != nil {
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

// ParseDeleteKeyEvent is a log parse operation binding the contract event 0x2563b537df5e859d4928c3f28220f6c5e6aab2948198de0bd936fc5e7750eed1.
//
// Solidity: event deleteKeyEvent(address indexed userAddr, string accountClass, string key)
func (_Account *AccountFilterer) ParseDeleteKeyEvent(log types.Log) (*AccountDeleteKeyEvent, error) {
	event := new(AccountDeleteKeyEvent)
	if err := _Account.contract.UnpackLog(event, "deleteKeyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountDoKeyReturnEventIterator is returned from FilterDoKeyReturnEvent and is used to iterate over the raw logs and unpacked data for DoKeyReturnEvent events raised by the Account contract.
type AccountDoKeyReturnEventIterator struct {
	Event *AccountDoKeyReturnEvent // Event containing the contract specifics and raw log

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
func (it *AccountDoKeyReturnEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountDoKeyReturnEvent)
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
		it.Event = new(AccountDoKeyReturnEvent)
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
func (it *AccountDoKeyReturnEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountDoKeyReturnEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountDoKeyReturnEvent represents a DoKeyReturnEvent event raised by the Account contract.
type AccountDoKeyReturnEvent struct {
	Indexd *big.Int
	Arg1   *big.Int
	Arg2   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDoKeyReturnEvent is a free log retrieval operation binding the contract event 0x3b80e5d256e1465fb2a7903bc318c449f50cc9f187a7ae3817a4bf16f7b7b682.
//
// Solidity: event doKeyReturnEvent(uint256 indexd, uint256 arg1, uint256 arg2)
func (_Account *AccountFilterer) FilterDoKeyReturnEvent(opts *bind.FilterOpts) (*AccountDoKeyReturnEventIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "doKeyReturnEvent")
	if err != nil {
		return nil, err
	}
	return &AccountDoKeyReturnEventIterator{contract: _Account.contract, event: "doKeyReturnEvent", logs: logs, sub: sub}, nil
}

// WatchDoKeyReturnEvent is a free log subscription operation binding the contract event 0x3b80e5d256e1465fb2a7903bc318c449f50cc9f187a7ae3817a4bf16f7b7b682.
//
// Solidity: event doKeyReturnEvent(uint256 indexd, uint256 arg1, uint256 arg2)
func (_Account *AccountFilterer) WatchDoKeyReturnEvent(opts *bind.WatchOpts, sink chan<- *AccountDoKeyReturnEvent) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "doKeyReturnEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountDoKeyReturnEvent)
				if err := _Account.contract.UnpackLog(event, "doKeyReturnEvent", log); err != nil {
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

// ParseDoKeyReturnEvent is a log parse operation binding the contract event 0x3b80e5d256e1465fb2a7903bc318c449f50cc9f187a7ae3817a4bf16f7b7b682.
//
// Solidity: event doKeyReturnEvent(uint256 indexd, uint256 arg1, uint256 arg2)
func (_Account *AccountFilterer) ParseDoKeyReturnEvent(log types.Log) (*AccountDoKeyReturnEvent, error) {
	event := new(AccountDoKeyReturnEvent)
	if err := _Account.contract.UnpackLog(event, "doKeyReturnEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountModifyAccountClassEventIterator is returned from FilterModifyAccountClassEvent and is used to iterate over the raw logs and unpacked data for ModifyAccountClassEvent events raised by the Account contract.
type AccountModifyAccountClassEventIterator struct {
	Event *AccountModifyAccountClassEvent // Event containing the contract specifics and raw log

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
func (it *AccountModifyAccountClassEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountModifyAccountClassEvent)
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
		it.Event = new(AccountModifyAccountClassEvent)
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
func (it *AccountModifyAccountClassEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountModifyAccountClassEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountModifyAccountClassEvent represents a ModifyAccountClassEvent event raised by the Account contract.
type AccountModifyAccountClassEvent struct {
	UserAddr        common.Address
	AccountClass    string
	OldAccountClass string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterModifyAccountClassEvent is a free log retrieval operation binding the contract event 0xd09f39d7f08c053a18b667dcd5ad857074a15b6c9ca620906f712ab601478ff7.
//
// Solidity: event modifyAccountClassEvent(address indexed userAddr, string accountClass, string oldAccountClass)
func (_Account *AccountFilterer) FilterModifyAccountClassEvent(opts *bind.FilterOpts, userAddr []common.Address) (*AccountModifyAccountClassEventIterator, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Account.contract.FilterLogs(opts, "modifyAccountClassEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return &AccountModifyAccountClassEventIterator{contract: _Account.contract, event: "modifyAccountClassEvent", logs: logs, sub: sub}, nil
}

// WatchModifyAccountClassEvent is a free log subscription operation binding the contract event 0xd09f39d7f08c053a18b667dcd5ad857074a15b6c9ca620906f712ab601478ff7.
//
// Solidity: event modifyAccountClassEvent(address indexed userAddr, string accountClass, string oldAccountClass)
func (_Account *AccountFilterer) WatchModifyAccountClassEvent(opts *bind.WatchOpts, sink chan<- *AccountModifyAccountClassEvent, userAddr []common.Address) (event.Subscription, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Account.contract.WatchLogs(opts, "modifyAccountClassEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountModifyAccountClassEvent)
				if err := _Account.contract.UnpackLog(event, "modifyAccountClassEvent", log); err != nil {
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

// ParseModifyAccountClassEvent is a log parse operation binding the contract event 0xd09f39d7f08c053a18b667dcd5ad857074a15b6c9ca620906f712ab601478ff7.
//
// Solidity: event modifyAccountClassEvent(address indexed userAddr, string accountClass, string oldAccountClass)
func (_Account *AccountFilterer) ParseModifyAccountClassEvent(log types.Log) (*AccountModifyAccountClassEvent, error) {
	event := new(AccountModifyAccountClassEvent)
	if err := _Account.contract.UnpackLog(event, "modifyAccountClassEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountModifyKeyEventIterator is returned from FilterModifyKeyEvent and is used to iterate over the raw logs and unpacked data for ModifyKeyEvent events raised by the Account contract.
type AccountModifyKeyEventIterator struct {
	Event *AccountModifyKeyEvent // Event containing the contract specifics and raw log

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
func (it *AccountModifyKeyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountModifyKeyEvent)
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
		it.Event = new(AccountModifyKeyEvent)
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
func (it *AccountModifyKeyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountModifyKeyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountModifyKeyEvent represents a ModifyKeyEvent event raised by the Account contract.
type AccountModifyKeyEvent struct {
	UserAddr common.Address
	Id       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterModifyKeyEvent is a free log retrieval operation binding the contract event 0x7b787125ec9764f4b01855985392fa1906a2e75956951aec6b158299b2344411.
//
// Solidity: event modifyKeyEvent(address indexed userAddr, uint256 indexed id)
func (_Account *AccountFilterer) FilterModifyKeyEvent(opts *bind.FilterOpts, userAddr []common.Address, id []*big.Int) (*AccountModifyKeyEventIterator, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Account.contract.FilterLogs(opts, "modifyKeyEvent", userAddrRule, idRule)
	if err != nil {
		return nil, err
	}
	return &AccountModifyKeyEventIterator{contract: _Account.contract, event: "modifyKeyEvent", logs: logs, sub: sub}, nil
}

// WatchModifyKeyEvent is a free log subscription operation binding the contract event 0x7b787125ec9764f4b01855985392fa1906a2e75956951aec6b158299b2344411.
//
// Solidity: event modifyKeyEvent(address indexed userAddr, uint256 indexed id)
func (_Account *AccountFilterer) WatchModifyKeyEvent(opts *bind.WatchOpts, sink chan<- *AccountModifyKeyEvent, userAddr []common.Address, id []*big.Int) (event.Subscription, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Account.contract.WatchLogs(opts, "modifyKeyEvent", userAddrRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountModifyKeyEvent)
				if err := _Account.contract.UnpackLog(event, "modifyKeyEvent", log); err != nil {
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

// ParseModifyKeyEvent is a log parse operation binding the contract event 0x7b787125ec9764f4b01855985392fa1906a2e75956951aec6b158299b2344411.
//
// Solidity: event modifyKeyEvent(address indexed userAddr, uint256 indexed id)
func (_Account *AccountFilterer) ParseModifyKeyEvent(log types.Log) (*AccountModifyKeyEvent, error) {
	event := new(AccountModifyKeyEvent)
	if err := _Account.contract.UnpackLog(event, "modifyKeyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountMoveKeyEventIterator is returned from FilterMoveKeyEvent and is used to iterate over the raw logs and unpacked data for MoveKeyEvent events raised by the Account contract.
type AccountMoveKeyEventIterator struct {
	Event *AccountMoveKeyEvent // Event containing the contract specifics and raw log

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
func (it *AccountMoveKeyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountMoveKeyEvent)
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
		it.Event = new(AccountMoveKeyEvent)
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
func (it *AccountMoveKeyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountMoveKeyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountMoveKeyEvent represents a MoveKeyEvent event raised by the Account contract.
type AccountMoveKeyEvent struct {
	UserAddr        common.Address
	AccountClass    string
	Key             string
	OldAccountClass string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMoveKeyEvent is a free log retrieval operation binding the contract event 0x5d7df306a6bbb7e121f443296aed1fcaabee5bbfddf586b3a720e19832a4661b.
//
// Solidity: event moveKeyEvent(address indexed userAddr, string accountClass, string key, string oldAccountClass)
func (_Account *AccountFilterer) FilterMoveKeyEvent(opts *bind.FilterOpts, userAddr []common.Address) (*AccountMoveKeyEventIterator, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Account.contract.FilterLogs(opts, "moveKeyEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return &AccountMoveKeyEventIterator{contract: _Account.contract, event: "moveKeyEvent", logs: logs, sub: sub}, nil
}

// WatchMoveKeyEvent is a free log subscription operation binding the contract event 0x5d7df306a6bbb7e121f443296aed1fcaabee5bbfddf586b3a720e19832a4661b.
//
// Solidity: event moveKeyEvent(address indexed userAddr, string accountClass, string key, string oldAccountClass)
func (_Account *AccountFilterer) WatchMoveKeyEvent(opts *bind.WatchOpts, sink chan<- *AccountMoveKeyEvent, userAddr []common.Address) (event.Subscription, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Account.contract.WatchLogs(opts, "moveKeyEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountMoveKeyEvent)
				if err := _Account.contract.UnpackLog(event, "moveKeyEvent", log); err != nil {
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

// ParseMoveKeyEvent is a log parse operation binding the contract event 0x5d7df306a6bbb7e121f443296aed1fcaabee5bbfddf586b3a720e19832a4661b.
//
// Solidity: event moveKeyEvent(address indexed userAddr, string accountClass, string key, string oldAccountClass)
func (_Account *AccountFilterer) ParseMoveKeyEvent(log types.Log) (*AccountMoveKeyEvent, error) {
	event := new(AccountMoveKeyEvent)
	if err := _Account.contract.UnpackLog(event, "moveKeyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountSetKeyByIdEventIterator is returned from FilterSetKeyByIdEvent and is used to iterate over the raw logs and unpacked data for SetKeyByIdEvent events raised by the Account contract.
type AccountSetKeyByIdEventIterator struct {
	Event *AccountSetKeyByIdEvent // Event containing the contract specifics and raw log

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
func (it *AccountSetKeyByIdEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountSetKeyByIdEvent)
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
		it.Event = new(AccountSetKeyByIdEvent)
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
func (it *AccountSetKeyByIdEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountSetKeyByIdEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountSetKeyByIdEvent represents a SetKeyByIdEvent event raised by the Account contract.
type AccountSetKeyByIdEvent struct {
	Arg0 *big.Int
	Arg1 uint8
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetKeyByIdEvent is a free log retrieval operation binding the contract event 0xde46eaa05c7c64d37d45cfc07876bcde59560680a5193df830fd5d25ace97308.
//
// Solidity: event setKeyByIdEvent(uint256 arg0, uint8 arg1, uint256 arg2)
func (_Account *AccountFilterer) FilterSetKeyByIdEvent(opts *bind.FilterOpts) (*AccountSetKeyByIdEventIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "setKeyByIdEvent")
	if err != nil {
		return nil, err
	}
	return &AccountSetKeyByIdEventIterator{contract: _Account.contract, event: "setKeyByIdEvent", logs: logs, sub: sub}, nil
}

// WatchSetKeyByIdEvent is a free log subscription operation binding the contract event 0xde46eaa05c7c64d37d45cfc07876bcde59560680a5193df830fd5d25ace97308.
//
// Solidity: event setKeyByIdEvent(uint256 arg0, uint8 arg1, uint256 arg2)
func (_Account *AccountFilterer) WatchSetKeyByIdEvent(opts *bind.WatchOpts, sink chan<- *AccountSetKeyByIdEvent) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "setKeyByIdEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountSetKeyByIdEvent)
				if err := _Account.contract.UnpackLog(event, "setKeyByIdEvent", log); err != nil {
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

// ParseSetKeyByIdEvent is a log parse operation binding the contract event 0xde46eaa05c7c64d37d45cfc07876bcde59560680a5193df830fd5d25ace97308.
//
// Solidity: event setKeyByIdEvent(uint256 arg0, uint8 arg1, uint256 arg2)
func (_Account *AccountFilterer) ParseSetKeyByIdEvent(log types.Log) (*AccountSetKeyByIdEvent, error) {
	event := new(AccountSetKeyByIdEvent)
	if err := _Account.contract.UnpackLog(event, "setKeyByIdEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
