package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

// 1个红色方法
// 11个橙色方法
// 11个蓝色方法

/*
	红色
*/
// 借入钥匙
func BorrowKey(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return Account.BorrowKeyByIdFunc(opts, id)
}

/**
橙色 11个
*/
// 添加钥匙分类
func AddAccountClass(opts *bind.TransactOpts, class string) (*types.Transaction, error) {
	return Account.AddAccountClassFunc(opts, class)
}
func ChangeClass(opts *bind.TransactOpts, _oldClass string, _newClass string) (*types.Transaction, error) {
	return Account.ChangeClass(opts, _oldClass, _newClass)
}
