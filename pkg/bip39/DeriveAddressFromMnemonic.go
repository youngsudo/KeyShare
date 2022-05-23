package bip39

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/tyler-smith/go-bip39"
)

func DerivePrivateKey(path accounts.DerivationPath, masterKey *hdkeychain.ExtendedKey) (*ecdsa.PrivateKey, error) {
	var err error
	key := masterKey
	// 将 key 转换为 ecdsa 私钥
	privateKey, err := key.ECPrivKey()

	privateKeyECDSA := privateKey.ToECDSA()

	if err != nil {
		return nil, err
	}
	return privateKeyECDSA, nil

}

// 推导私钥
func DeriveAddressFromMnemonic(mnemonic *string) (*ecdsa.PrivateKey, *string, error) {
	var nm = new(string)
	// 1.先推导路径
	path, err := accounts.ParseDerivationPath("m/44'/60'/0'/0/1")
	if err != nil {
		panic(err)
	}
	// 2.获得种子
	if *mnemonic != "" { // 用户上传助记词
		nm = mnemonic
	} else {
		nm, err = bip()
		if err != nil {
			return nil, nm, err
		}
	}
	// 生成私钥  *mnemonic 助记词
	seed, err := bip39.NewSeedWithErrorChecking(*nm, "")
	if err != nil {
		return nil, nm, err
	}
	// 3.获得主key
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Printf("Failed to NewMaster : %v\n", err)
		return nil, nm, err
	}
	// 4.推导私钥 *invalid type
	privateKey, err := DerivePrivateKey(path, masterKey) // *ecdsa.PrivateKey
	if err != nil {
		return nil, nm, err
	}
	// fmt.Println("privateKey:", privateKey)
	// 5.推导公钥
	// 6.推导地址
	return privateKey, nm, nil
}
