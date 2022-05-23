package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

/*	16
	View
	不用地址和签名
*/
// 判断用户地址是否存在(是否已经注册)
func IsExitUserAddressView(addr string) (bool, error) {
	return KeyShare.IsExitUserAddressFunc(nil, common.HexToAddress(addr))
}

// 判断 用户账号 是否存在
func IsExitUserAccountView(account string) (bool, error) {
	return KeyShare.IsExitUserAccountFunc(nil, account)
}

// 判断 用户邮箱 是否存在
func IsExitUserEmailView(email string) (bool, error) {
	return KeyShare.IsExitUserEmailFunc(nil, email)
}

// 获取管理员总数量,包括超级管理员
func GetAdminNumView() (uint8, error) {
	return KeyShare.AdminNum(nil)
}

// 获取管理员地址
func AdministratorView(index uint8) (string, error) {
	addr, err := KeyShare.Administrator(nil)
	return addr.Hex(), err
}

// 判断用户是否为管理员 0 是administrator 超级管理员; 1 admin 管理员; 2 是普通用户
func IsAdminView(addr string) (bool, error) {
	return KeyShare.IsAdminFunc(nil, common.HexToAddress(addr))
}

// 登录
func Login(addr, account, password string) (bool, uint8, error) {
	return KeyShare.LoginFunc(nil, common.HexToAddress(addr), account, password)
}

/*
	View
	msg.sender
*/
// 查看自己的助记符
func GetSymbolView() (string, error) {
	return KeyShare.GetSymbolFunc(nil)
}

// 通过助记符查找地址
func GetAddressView(symbol string) (string, error) {
	addr, err := KeyShare.Administrator(nil)
	return addr.Hex(), err
}

// 用户查看自己的信息
func GetUserInfoView(addr string) (common.Address, string, string, *big.Int, uint8, error) {
	return KeyShare.GetMyselfInfoFunc(nil)
}

/*
	View
	admin
*/
// 管理员查看用户的总数量
func GetSumView() (*big.Int, error) {
	return KeyShare.GetSumFunc(nil)
}

// 查看所有管理员的地址 onlyAdmin
func GetAllAdminView() ([]string, error) {
	var address = make([]string, 0, 10)
	addr, err := KeyShare.GetAdminListFunc(nil) // []common.Address
	if err != nil {
		return address, err
	}
	for _, v := range addr {
		address = append(address, v.Hex())
	}
	return address, nil
}

//  查看用户的地址 onlyAdmin
func GetUserListView(index *big.Int) (string, error) {
	addr, err := KeyShare.GetUserListFunc(nil, index)
	return addr.Hex(), err
}

/*
	橙色
*/
// 创建用户
func AddUserPublic(opts *bind.TransactOpts, addr, account, password, email, mnemonic string) (*types.Transaction, error) {
	return KeyShare.AddUserFunc(opts, common.HexToAddress(addr), account, password, email, mnemonic)
}

/*
	橙色
	administrator
*/
// 将用户升级为管理员
func UpgradePublic(opts *bind.TransactOpts, addr string) (*types.Transaction, error) {
	return KeyShare.AddAdminFunc(opts, common.HexToAddress(addr))
}

// 将管理员移除
func RemoveAdminPublic(opts *bind.TransactOpts, addr string) (*types.Transaction, error) {
	return KeyShare.RemoveAdminFunc(opts, common.HexToAddress(addr))
}
