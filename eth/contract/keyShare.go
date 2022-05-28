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

// KeyShareABI is the input ABI used to generate the binding from.
const KeyShareABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_adminNum\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"}],\"name\":\"addUserEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"removeAdminEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"upgradeAdminEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"addAdminFunc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_account\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_password\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"addUserFunc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminNum\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"administrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"getAddressFunc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdminListFunc\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMyselfInfoFunc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumUserType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSumFunc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSymbolFunc\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getUserListFunc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isAdminFunc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_userAddress\",\"type\":\"string\"}],\"name\":\"isExitUserAccountFunc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"isExitUserAddressFunc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"}],\"name\":\"isExitUserEmailFunc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_account\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_password\",\"type\":\"string\"}],\"name\":\"loginFunc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"enumUserType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"removeAdminFunc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KeyShare is an auto generated Go binding around an Ethereum contract.
type KeyShare struct {
	KeyShareCaller     // Read-only binding to the contract
	KeyShareTransactor // Write-only binding to the contract
	KeyShareFilterer   // Log filterer for contract events
}

// KeyShareCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyShareCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyShareTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyShareTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyShareFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyShareFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyShareSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyShareSession struct {
	Contract     *KeyShare         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeyShareCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyShareCallerSession struct {
	Contract *KeyShareCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// KeyShareTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyShareTransactorSession struct {
	Contract     *KeyShareTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KeyShareRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyShareRaw struct {
	Contract *KeyShare // Generic contract binding to access the raw methods on
}

// KeyShareCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyShareCallerRaw struct {
	Contract *KeyShareCaller // Generic read-only contract binding to access the raw methods on
}

// KeyShareTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyShareTransactorRaw struct {
	Contract *KeyShareTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyShare creates a new instance of KeyShare, bound to a specific deployed contract.
func NewKeyShare(address common.Address, backend bind.ContractBackend) (*KeyShare, error) {
	contract, err := bindKeyShare(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyShare{KeyShareCaller: KeyShareCaller{contract: contract}, KeyShareTransactor: KeyShareTransactor{contract: contract}, KeyShareFilterer: KeyShareFilterer{contract: contract}}, nil
}

// NewKeyShareCaller creates a new read-only instance of KeyShare, bound to a specific deployed contract.
func NewKeyShareCaller(address common.Address, caller bind.ContractCaller) (*KeyShareCaller, error) {
	contract, err := bindKeyShare(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyShareCaller{contract: contract}, nil
}

// NewKeyShareTransactor creates a new write-only instance of KeyShare, bound to a specific deployed contract.
func NewKeyShareTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyShareTransactor, error) {
	contract, err := bindKeyShare(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyShareTransactor{contract: contract}, nil
}

// NewKeyShareFilterer creates a new log filterer instance of KeyShare, bound to a specific deployed contract.
func NewKeyShareFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyShareFilterer, error) {
	contract, err := bindKeyShare(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyShareFilterer{contract: contract}, nil
}

// bindKeyShare binds a generic wrapper to an already deployed contract.
func bindKeyShare(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyShareABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyShare *KeyShareRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyShare.Contract.KeyShareCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyShare *KeyShareRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyShare.Contract.KeyShareTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyShare *KeyShareRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyShare.Contract.KeyShareTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyShare *KeyShareCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyShare.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyShare *KeyShareTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyShare.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyShare *KeyShareTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyShare.Contract.contract.Transact(opts, method, params...)
}

// AdminNum is a free data retrieval call binding the contract method 0x7e12f6b3.
//
// Solidity: function adminNum() view returns(uint8)
func (_KeyShare *KeyShareCaller) AdminNum(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _KeyShare.contract.Call(opts, &out, "adminNum")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AdminNum is a free data retrieval call binding the contract method 0x7e12f6b3.
//
// Solidity: function adminNum() view returns(uint8)
func (_KeyShare *KeyShareSession) AdminNum() (uint8, error) {
	return _KeyShare.Contract.AdminNum(&_KeyShare.CallOpts)
}

// AdminNum is a free data retrieval call binding the contract method 0x7e12f6b3.
//
// Solidity: function adminNum() view returns(uint8)
func (_KeyShare *KeyShareCallerSession) AdminNum() (uint8, error) {
	return _KeyShare.Contract.AdminNum(&_KeyShare.CallOpts)
}

// Administrator is a free data retrieval call binding the contract method 0xf53d0a8e.
//
// Solidity: function administrator() view returns(address)
func (_KeyShare *KeyShareCaller) Administrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyShare.contract.Call(opts, &out, "administrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Administrator is a free data retrieval call binding the contract method 0xf53d0a8e.
//
// Solidity: function administrator() view returns(address)
func (_KeyShare *KeyShareSession) Administrator() (common.Address, error) {
	return _KeyShare.Contract.Administrator(&_KeyShare.CallOpts)
}

// Administrator is a free data retrieval call binding the contract method 0xf53d0a8e.
//
// Solidity: function administrator() view returns(address)
func (_KeyShare *KeyShareCallerSession) Administrator() (common.Address, error) {
	return _KeyShare.Contract.Administrator(&_KeyShare.CallOpts)
}

// GetAddressFunc is a free data retrieval call binding the contract method 0xaa05e42b.
//
// Solidity: function getAddressFunc(string _symbol) view returns(address)
func (_KeyShare *KeyShareCaller) GetAddressFunc(opts *bind.CallOpts, _symbol string) (common.Address, error) {
	var out []interface{}
	err := _KeyShare.contract.Call(opts, &out, "getAddressFunc", _symbol)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressFunc is a free data retrieval call binding the contract method 0xaa05e42b.
//
// Solidity: function getAddressFunc(string _symbol) view returns(address)
func (_KeyShare *KeyShareSession) GetAddressFunc(_symbol string) (common.Address, error) {
	return _KeyShare.Contract.GetAddressFunc(&_KeyShare.CallOpts, _symbol)
}

// GetAddressFunc is a free data retrieval call binding the contract method 0xaa05e42b.
//
// Solidity: function getAddressFunc(string _symbol) view returns(address)
func (_KeyShare *KeyShareCallerSession) GetAddressFunc(_symbol string) (common.Address, error) {
	return _KeyShare.Contract.GetAddressFunc(&_KeyShare.CallOpts, _symbol)
}

// GetAdminListFunc is a free data retrieval call binding the contract method 0x0b653732.
//
// Solidity: function getAdminListFunc() view returns(address[])
func (_KeyShare *KeyShareCaller) GetAdminListFunc(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _KeyShare.contract.Call(opts, &out, "getAdminListFunc")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdminListFunc is a free data retrieval call binding the contract method 0x0b653732.
//
// Solidity: function getAdminListFunc() view returns(address[])
func (_KeyShare *KeyShareSession) GetAdminListFunc() ([]common.Address, error) {
	return _KeyShare.Contract.GetAdminListFunc(&_KeyShare.CallOpts)
}

// GetAdminListFunc is a free data retrieval call binding the contract method 0x0b653732.
//
// Solidity: function getAdminListFunc() view returns(address[])
func (_KeyShare *KeyShareCallerSession) GetAdminListFunc() ([]common.Address, error) {
	return _KeyShare.Contract.GetAdminListFunc(&_KeyShare.CallOpts)
}

// GetMyselfInfoFunc is a free data retrieval call binding the contract method 0x25bf257c.
//
// Solidity: function getMyselfInfoFunc() view returns(address, string, string, uint256, uint8)
func (_KeyShare *KeyShareCaller) GetMyselfInfoFunc(opts *bind.CallOpts) (common.Address, string, string, *big.Int, uint8, error) {
	var out []interface{}
	err := _KeyShare.contract.Call(opts, &out, "getMyselfInfoFunc")

	if err != nil {
		return *new(common.Address), *new(string), *new(string), *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return out0, out1, out2, out3, out4, err

}

// GetMyselfInfoFunc is a free data retrieval call binding the contract method 0x25bf257c.
//
// Solidity: function getMyselfInfoFunc() view returns(address, string, string, uint256, uint8)
func (_KeyShare *KeyShareSession) GetMyselfInfoFunc() (common.Address, string, string, *big.Int, uint8, error) {
	return _KeyShare.Contract.GetMyselfInfoFunc(&_KeyShare.CallOpts)
}

// GetMyselfInfoFunc is a free data retrieval call binding the contract method 0x25bf257c.
//
// Solidity: function getMyselfInfoFunc() view returns(address, string, string, uint256, uint8)
func (_KeyShare *KeyShareCallerSession) GetMyselfInfoFunc() (common.Address, string, string, *big.Int, uint8, error) {
	return _KeyShare.Contract.GetMyselfInfoFunc(&_KeyShare.CallOpts)
}

// GetSumFunc is a free data retrieval call binding the contract method 0x9fa893ec.
//
// Solidity: function getSumFunc() view returns(uint256)
func (_KeyShare *KeyShareCaller) GetSumFunc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KeyShare.contract.Call(opts, &out, "getSumFunc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSumFunc is a free data retrieval call binding the contract method 0x9fa893ec.
//
// Solidity: function getSumFunc() view returns(uint256)
func (_KeyShare *KeyShareSession) GetSumFunc() (*big.Int, error) {
	return _KeyShare.Contract.GetSumFunc(&_KeyShare.CallOpts)
}

// GetSumFunc is a free data retrieval call binding the contract method 0x9fa893ec.
//
// Solidity: function getSumFunc() view returns(uint256)
func (_KeyShare *KeyShareCallerSession) GetSumFunc() (*big.Int, error) {
	return _KeyShare.Contract.GetSumFunc(&_KeyShare.CallOpts)
}

// GetSymbolFunc is a free data retrieval call binding the contract method 0x7e2cd852.
//
// Solidity: function getSymbolFunc() view returns(string)
func (_KeyShare *KeyShareCaller) GetSymbolFunc(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KeyShare.contract.Call(opts, &out, "getSymbolFunc")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSymbolFunc is a free data retrieval call binding the contract method 0x7e2cd852.
//
// Solidity: function getSymbolFunc() view returns(string)
func (_KeyShare *KeyShareSession) GetSymbolFunc() (string, error) {
	return _KeyShare.Contract.GetSymbolFunc(&_KeyShare.CallOpts)
}

// GetSymbolFunc is a free data retrieval call binding the contract method 0x7e2cd852.
//
// Solidity: function getSymbolFunc() view returns(string)
func (_KeyShare *KeyShareCallerSession) GetSymbolFunc() (string, error) {
	return _KeyShare.Contract.GetSymbolFunc(&_KeyShare.CallOpts)
}

// GetUserListFunc is a free data retrieval call binding the contract method 0x72124951.
//
// Solidity: function getUserListFunc(uint256 _index) view returns(address)
func (_KeyShare *KeyShareCaller) GetUserListFunc(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KeyShare.contract.Call(opts, &out, "getUserListFunc", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserListFunc is a free data retrieval call binding the contract method 0x72124951.
//
// Solidity: function getUserListFunc(uint256 _index) view returns(address)
func (_KeyShare *KeyShareSession) GetUserListFunc(_index *big.Int) (common.Address, error) {
	return _KeyShare.Contract.GetUserListFunc(&_KeyShare.CallOpts, _index)
}

// GetUserListFunc is a free data retrieval call binding the contract method 0x72124951.
//
// Solidity: function getUserListFunc(uint256 _index) view returns(address)
func (_KeyShare *KeyShareCallerSession) GetUserListFunc(_index *big.Int) (common.Address, error) {
	return _KeyShare.Contract.GetUserListFunc(&_KeyShare.CallOpts, _index)
}

// IsAdminFunc is a free data retrieval call binding the contract method 0x842ca1d4.
//
// Solidity: function isAdminFunc(address _addr) view returns(bool)
func (_KeyShare *KeyShareCaller) IsAdminFunc(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _KeyShare.contract.Call(opts, &out, "isAdminFunc", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdminFunc is a free data retrieval call binding the contract method 0x842ca1d4.
//
// Solidity: function isAdminFunc(address _addr) view returns(bool)
func (_KeyShare *KeyShareSession) IsAdminFunc(_addr common.Address) (bool, error) {
	return _KeyShare.Contract.IsAdminFunc(&_KeyShare.CallOpts, _addr)
}

// IsAdminFunc is a free data retrieval call binding the contract method 0x842ca1d4.
//
// Solidity: function isAdminFunc(address _addr) view returns(bool)
func (_KeyShare *KeyShareCallerSession) IsAdminFunc(_addr common.Address) (bool, error) {
	return _KeyShare.Contract.IsAdminFunc(&_KeyShare.CallOpts, _addr)
}

// IsExitUserAccountFunc is a free data retrieval call binding the contract method 0xffca01de.
//
// Solidity: function isExitUserAccountFunc(string _userAddress) view returns(bool)
func (_KeyShare *KeyShareCaller) IsExitUserAccountFunc(opts *bind.CallOpts, _userAddress string) (bool, error) {
	var out []interface{}
	err := _KeyShare.contract.Call(opts, &out, "isExitUserAccountFunc", _userAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExitUserAccountFunc is a free data retrieval call binding the contract method 0xffca01de.
//
// Solidity: function isExitUserAccountFunc(string _userAddress) view returns(bool)
func (_KeyShare *KeyShareSession) IsExitUserAccountFunc(_userAddress string) (bool, error) {
	return _KeyShare.Contract.IsExitUserAccountFunc(&_KeyShare.CallOpts, _userAddress)
}

// IsExitUserAccountFunc is a free data retrieval call binding the contract method 0xffca01de.
//
// Solidity: function isExitUserAccountFunc(string _userAddress) view returns(bool)
func (_KeyShare *KeyShareCallerSession) IsExitUserAccountFunc(_userAddress string) (bool, error) {
	return _KeyShare.Contract.IsExitUserAccountFunc(&_KeyShare.CallOpts, _userAddress)
}

// IsExitUserAddressFunc is a free data retrieval call binding the contract method 0x7d03aed8.
//
// Solidity: function isExitUserAddressFunc(address _userAddress) view returns(bool)
func (_KeyShare *KeyShareCaller) IsExitUserAddressFunc(opts *bind.CallOpts, _userAddress common.Address) (bool, error) {
	var out []interface{}
	err := _KeyShare.contract.Call(opts, &out, "isExitUserAddressFunc", _userAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExitUserAddressFunc is a free data retrieval call binding the contract method 0x7d03aed8.
//
// Solidity: function isExitUserAddressFunc(address _userAddress) view returns(bool)
func (_KeyShare *KeyShareSession) IsExitUserAddressFunc(_userAddress common.Address) (bool, error) {
	return _KeyShare.Contract.IsExitUserAddressFunc(&_KeyShare.CallOpts, _userAddress)
}

// IsExitUserAddressFunc is a free data retrieval call binding the contract method 0x7d03aed8.
//
// Solidity: function isExitUserAddressFunc(address _userAddress) view returns(bool)
func (_KeyShare *KeyShareCallerSession) IsExitUserAddressFunc(_userAddress common.Address) (bool, error) {
	return _KeyShare.Contract.IsExitUserAddressFunc(&_KeyShare.CallOpts, _userAddress)
}

// IsExitUserEmailFunc is a free data retrieval call binding the contract method 0x4710a0dc.
//
// Solidity: function isExitUserEmailFunc(string _email) view returns(bool)
func (_KeyShare *KeyShareCaller) IsExitUserEmailFunc(opts *bind.CallOpts, _email string) (bool, error) {
	var out []interface{}
	err := _KeyShare.contract.Call(opts, &out, "isExitUserEmailFunc", _email)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExitUserEmailFunc is a free data retrieval call binding the contract method 0x4710a0dc.
//
// Solidity: function isExitUserEmailFunc(string _email) view returns(bool)
func (_KeyShare *KeyShareSession) IsExitUserEmailFunc(_email string) (bool, error) {
	return _KeyShare.Contract.IsExitUserEmailFunc(&_KeyShare.CallOpts, _email)
}

// IsExitUserEmailFunc is a free data retrieval call binding the contract method 0x4710a0dc.
//
// Solidity: function isExitUserEmailFunc(string _email) view returns(bool)
func (_KeyShare *KeyShareCallerSession) IsExitUserEmailFunc(_email string) (bool, error) {
	return _KeyShare.Contract.IsExitUserEmailFunc(&_KeyShare.CallOpts, _email)
}

// LoginFunc is a free data retrieval call binding the contract method 0x11a59dee.
//
// Solidity: function loginFunc(address _addr, string _account, string _password) view returns(bool, uint8, address)
func (_KeyShare *KeyShareCaller) LoginFunc(opts *bind.CallOpts, _addr common.Address, _account string, _password string) (bool, uint8, common.Address, error) {
	var out []interface{}
	err := _KeyShare.contract.Call(opts, &out, "loginFunc", _addr, _account, _password)

	if err != nil {
		return *new(bool), *new(uint8), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// LoginFunc is a free data retrieval call binding the contract method 0x11a59dee.
//
// Solidity: function loginFunc(address _addr, string _account, string _password) view returns(bool, uint8, address)
func (_KeyShare *KeyShareSession) LoginFunc(_addr common.Address, _account string, _password string) (bool, uint8, common.Address, error) {
	return _KeyShare.Contract.LoginFunc(&_KeyShare.CallOpts, _addr, _account, _password)
}

// LoginFunc is a free data retrieval call binding the contract method 0x11a59dee.
//
// Solidity: function loginFunc(address _addr, string _account, string _password) view returns(bool, uint8, address)
func (_KeyShare *KeyShareCallerSession) LoginFunc(_addr common.Address, _account string, _password string) (bool, uint8, common.Address, error) {
	return _KeyShare.Contract.LoginFunc(&_KeyShare.CallOpts, _addr, _account, _password)
}

// AddAdminFunc is a paid mutator transaction binding the contract method 0x74062a52.
//
// Solidity: function addAdminFunc(address _addr) returns()
func (_KeyShare *KeyShareTransactor) AddAdminFunc(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _KeyShare.contract.Transact(opts, "addAdminFunc", _addr)
}

// AddAdminFunc is a paid mutator transaction binding the contract method 0x74062a52.
//
// Solidity: function addAdminFunc(address _addr) returns()
func (_KeyShare *KeyShareSession) AddAdminFunc(_addr common.Address) (*types.Transaction, error) {
	return _KeyShare.Contract.AddAdminFunc(&_KeyShare.TransactOpts, _addr)
}

// AddAdminFunc is a paid mutator transaction binding the contract method 0x74062a52.
//
// Solidity: function addAdminFunc(address _addr) returns()
func (_KeyShare *KeyShareTransactorSession) AddAdminFunc(_addr common.Address) (*types.Transaction, error) {
	return _KeyShare.Contract.AddAdminFunc(&_KeyShare.TransactOpts, _addr)
}

// AddUserFunc is a paid mutator transaction binding the contract method 0x2456b8b0.
//
// Solidity: function addUserFunc(address _addr, string _account, string _password, string _email, string _symbol) returns()
func (_KeyShare *KeyShareTransactor) AddUserFunc(opts *bind.TransactOpts, _addr common.Address, _account string, _password string, _email string, _symbol string) (*types.Transaction, error) {
	return _KeyShare.contract.Transact(opts, "addUserFunc", _addr, _account, _password, _email, _symbol)
}

// AddUserFunc is a paid mutator transaction binding the contract method 0x2456b8b0.
//
// Solidity: function addUserFunc(address _addr, string _account, string _password, string _email, string _symbol) returns()
func (_KeyShare *KeyShareSession) AddUserFunc(_addr common.Address, _account string, _password string, _email string, _symbol string) (*types.Transaction, error) {
	return _KeyShare.Contract.AddUserFunc(&_KeyShare.TransactOpts, _addr, _account, _password, _email, _symbol)
}

// AddUserFunc is a paid mutator transaction binding the contract method 0x2456b8b0.
//
// Solidity: function addUserFunc(address _addr, string _account, string _password, string _email, string _symbol) returns()
func (_KeyShare *KeyShareTransactorSession) AddUserFunc(_addr common.Address, _account string, _password string, _email string, _symbol string) (*types.Transaction, error) {
	return _KeyShare.Contract.AddUserFunc(&_KeyShare.TransactOpts, _addr, _account, _password, _email, _symbol)
}

// RemoveAdminFunc is a paid mutator transaction binding the contract method 0xbd0c732e.
//
// Solidity: function removeAdminFunc(address _addr) returns()
func (_KeyShare *KeyShareTransactor) RemoveAdminFunc(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _KeyShare.contract.Transact(opts, "removeAdminFunc", _addr)
}

// RemoveAdminFunc is a paid mutator transaction binding the contract method 0xbd0c732e.
//
// Solidity: function removeAdminFunc(address _addr) returns()
func (_KeyShare *KeyShareSession) RemoveAdminFunc(_addr common.Address) (*types.Transaction, error) {
	return _KeyShare.Contract.RemoveAdminFunc(&_KeyShare.TransactOpts, _addr)
}

// RemoveAdminFunc is a paid mutator transaction binding the contract method 0xbd0c732e.
//
// Solidity: function removeAdminFunc(address _addr) returns()
func (_KeyShare *KeyShareTransactorSession) RemoveAdminFunc(_addr common.Address) (*types.Transaction, error) {
	return _KeyShare.Contract.RemoveAdminFunc(&_KeyShare.TransactOpts, _addr)
}

// KeyShareAddUserEventIterator is returned from FilterAddUserEvent and is used to iterate over the raw logs and unpacked data for AddUserEvent events raised by the KeyShare contract.
type KeyShareAddUserEventIterator struct {
	Event *KeyShareAddUserEvent // Event containing the contract specifics and raw log

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
func (it *KeyShareAddUserEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyShareAddUserEvent)
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
		it.Event = new(KeyShareAddUserEvent)
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
func (it *KeyShareAddUserEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyShareAddUserEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyShareAddUserEvent represents a AddUserEvent event raised by the KeyShare contract.
type KeyShareAddUserEvent struct {
	UserAddr common.Address
	Account  common.Hash
	Email    common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddUserEvent is a free log retrieval operation binding the contract event 0x5ae6671f20b51a868ae522f9841ff46bb92cc988a8f2f502deba9d90b8cd8693.
//
// Solidity: event addUserEvent(address indexed userAddr, string indexed account, string indexed email)
func (_KeyShare *KeyShareFilterer) FilterAddUserEvent(opts *bind.FilterOpts, userAddr []common.Address, account []string, email []string) (*KeyShareAddUserEventIterator, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var emailRule []interface{}
	for _, emailItem := range email {
		emailRule = append(emailRule, emailItem)
	}

	logs, sub, err := _KeyShare.contract.FilterLogs(opts, "addUserEvent", userAddrRule, accountRule, emailRule)
	if err != nil {
		return nil, err
	}
	return &KeyShareAddUserEventIterator{contract: _KeyShare.contract, event: "addUserEvent", logs: logs, sub: sub}, nil
}

// WatchAddUserEvent is a free log subscription operation binding the contract event 0x5ae6671f20b51a868ae522f9841ff46bb92cc988a8f2f502deba9d90b8cd8693.
//
// Solidity: event addUserEvent(address indexed userAddr, string indexed account, string indexed email)
func (_KeyShare *KeyShareFilterer) WatchAddUserEvent(opts *bind.WatchOpts, sink chan<- *KeyShareAddUserEvent, userAddr []common.Address, account []string, email []string) (event.Subscription, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var emailRule []interface{}
	for _, emailItem := range email {
		emailRule = append(emailRule, emailItem)
	}

	logs, sub, err := _KeyShare.contract.WatchLogs(opts, "addUserEvent", userAddrRule, accountRule, emailRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyShareAddUserEvent)
				if err := _KeyShare.contract.UnpackLog(event, "addUserEvent", log); err != nil {
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

// ParseAddUserEvent is a log parse operation binding the contract event 0x5ae6671f20b51a868ae522f9841ff46bb92cc988a8f2f502deba9d90b8cd8693.
//
// Solidity: event addUserEvent(address indexed userAddr, string indexed account, string indexed email)
func (_KeyShare *KeyShareFilterer) ParseAddUserEvent(log types.Log) (*KeyShareAddUserEvent, error) {
	event := new(KeyShareAddUserEvent)
	if err := _KeyShare.contract.UnpackLog(event, "addUserEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KeyShareRemoveAdminEventIterator is returned from FilterRemoveAdminEvent and is used to iterate over the raw logs and unpacked data for RemoveAdminEvent events raised by the KeyShare contract.
type KeyShareRemoveAdminEventIterator struct {
	Event *KeyShareRemoveAdminEvent // Event containing the contract specifics and raw log

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
func (it *KeyShareRemoveAdminEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyShareRemoveAdminEvent)
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
		it.Event = new(KeyShareRemoveAdminEvent)
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
func (it *KeyShareRemoveAdminEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyShareRemoveAdminEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyShareRemoveAdminEvent represents a RemoveAdminEvent event raised by the KeyShare contract.
type KeyShareRemoveAdminEvent struct {
	UserAddr common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemoveAdminEvent is a free log retrieval operation binding the contract event 0xfd6adb633c1ac063aa19914bb6588330135eff0659b9f22b920a8cdb55b311cf.
//
// Solidity: event removeAdminEvent(address indexed userAddr)
func (_KeyShare *KeyShareFilterer) FilterRemoveAdminEvent(opts *bind.FilterOpts, userAddr []common.Address) (*KeyShareRemoveAdminEventIterator, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _KeyShare.contract.FilterLogs(opts, "removeAdminEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return &KeyShareRemoveAdminEventIterator{contract: _KeyShare.contract, event: "removeAdminEvent", logs: logs, sub: sub}, nil
}

// WatchRemoveAdminEvent is a free log subscription operation binding the contract event 0xfd6adb633c1ac063aa19914bb6588330135eff0659b9f22b920a8cdb55b311cf.
//
// Solidity: event removeAdminEvent(address indexed userAddr)
func (_KeyShare *KeyShareFilterer) WatchRemoveAdminEvent(opts *bind.WatchOpts, sink chan<- *KeyShareRemoveAdminEvent, userAddr []common.Address) (event.Subscription, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _KeyShare.contract.WatchLogs(opts, "removeAdminEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyShareRemoveAdminEvent)
				if err := _KeyShare.contract.UnpackLog(event, "removeAdminEvent", log); err != nil {
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

// ParseRemoveAdminEvent is a log parse operation binding the contract event 0xfd6adb633c1ac063aa19914bb6588330135eff0659b9f22b920a8cdb55b311cf.
//
// Solidity: event removeAdminEvent(address indexed userAddr)
func (_KeyShare *KeyShareFilterer) ParseRemoveAdminEvent(log types.Log) (*KeyShareRemoveAdminEvent, error) {
	event := new(KeyShareRemoveAdminEvent)
	if err := _KeyShare.contract.UnpackLog(event, "removeAdminEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KeyShareUpgradeAdminEventIterator is returned from FilterUpgradeAdminEvent and is used to iterate over the raw logs and unpacked data for UpgradeAdminEvent events raised by the KeyShare contract.
type KeyShareUpgradeAdminEventIterator struct {
	Event *KeyShareUpgradeAdminEvent // Event containing the contract specifics and raw log

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
func (it *KeyShareUpgradeAdminEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyShareUpgradeAdminEvent)
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
		it.Event = new(KeyShareUpgradeAdminEvent)
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
func (it *KeyShareUpgradeAdminEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyShareUpgradeAdminEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyShareUpgradeAdminEvent represents a UpgradeAdminEvent event raised by the KeyShare contract.
type KeyShareUpgradeAdminEvent struct {
	UserAddr common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpgradeAdminEvent is a free log retrieval operation binding the contract event 0x1183b3c9db5f38951a3165df7a6e9842ae0ce34d47d33a271cbd70ddde178ea2.
//
// Solidity: event upgradeAdminEvent(address indexed userAddr)
func (_KeyShare *KeyShareFilterer) FilterUpgradeAdminEvent(opts *bind.FilterOpts, userAddr []common.Address) (*KeyShareUpgradeAdminEventIterator, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _KeyShare.contract.FilterLogs(opts, "upgradeAdminEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return &KeyShareUpgradeAdminEventIterator{contract: _KeyShare.contract, event: "upgradeAdminEvent", logs: logs, sub: sub}, nil
}

// WatchUpgradeAdminEvent is a free log subscription operation binding the contract event 0x1183b3c9db5f38951a3165df7a6e9842ae0ce34d47d33a271cbd70ddde178ea2.
//
// Solidity: event upgradeAdminEvent(address indexed userAddr)
func (_KeyShare *KeyShareFilterer) WatchUpgradeAdminEvent(opts *bind.WatchOpts, sink chan<- *KeyShareUpgradeAdminEvent, userAddr []common.Address) (event.Subscription, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _KeyShare.contract.WatchLogs(opts, "upgradeAdminEvent", userAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyShareUpgradeAdminEvent)
				if err := _KeyShare.contract.UnpackLog(event, "upgradeAdminEvent", log); err != nil {
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

// ParseUpgradeAdminEvent is a log parse operation binding the contract event 0x1183b3c9db5f38951a3165df7a6e9842ae0ce34d47d33a271cbd70ddde178ea2.
//
// Solidity: event upgradeAdminEvent(address indexed userAddr)
func (_KeyShare *KeyShareFilterer) ParseUpgradeAdminEvent(log types.Log) (*KeyShareUpgradeAdminEvent, error) {
	event := new(KeyShareUpgradeAdminEvent)
	if err := _KeyShare.contract.UnpackLog(event, "upgradeAdminEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
