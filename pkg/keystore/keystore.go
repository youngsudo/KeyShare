package keystore

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
)

// 导入keystore
func ImportKs(fileName string) (string, string, string, error) {
	// fmt.Println(_address, _path)
	dir, err := os.Getwd() //  (main.go所在目录)
	if err != nil {
		zap.L().Fatal("获取当前目录错误", zap.Error(err))
		return "", "", "", err
	}
	file := filepath.Join(dir, "pkg", "keystore", "wallets", fileName)
	jsonBytes, err := ioutil.ReadFile(file) // 读取文件
	if err != nil {
		zap.L().Fatal("读取文件Keystore文件错误", zap.Error(err))
		return "", "", "", err
	}

	password := "secret"
	Key, err := keystore.DecryptKey(jsonBytes, password) // *keystore.Key	// 解密keystore文件
	if err != nil {
		zap.L().Fatal("解密keystore文件错误", zap.Error(err))
		return "", "", "", err
		// err: 无法使用给定密码解密密钥
	}
	// fmt.Printf("Key:%#v,%[1]T\n", Key)

	// 生成私钥
	privateKey := crypto.FromECDSA(Key.PrivateKey) //[]byte
	fmt.Printf("privateKey:%#v\n", hexutil.Encode(privateKey))
	// 生成公钥
	publicKey := crypto.FromECDSAPub(&Key.PrivateKey.PublicKey) //[]byte
	fmt.Printf("publicKey:%#v\n", hexutil.Encode(publicKey))    // hexutil.Encode(publicKey) []byte ===> string

	// 生成地址
	address := crypto.PubkeyToAddress(Key.PrivateKey.PublicKey)
	return hexutil.Encode(privateKey), hexutil.Encode(publicKey), address.Hex(), err
}
