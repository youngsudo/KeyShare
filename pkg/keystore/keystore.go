package keystore

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
)

// 导入keystore
func ImportKs(fileName string) (privateKey, publicKey, address string, err error) {
	dir, err := os.Getwd() //  (main.go所在目录)
	if err != nil {
		zap.L().Fatal("获取当前目录错误", zap.Error(err))
		return
	}
	file := filepath.Join(dir, "pkg", "keystore", "wallets", fileName)
	jsonBytes, err := ioutil.ReadFile(file) // 读取文件
	if err != nil {
		zap.L().Fatal("读取文件Keystore文件错误", zap.Error(err))
		return
	}

	password := "secret"
	Key, err := keystore.DecryptKey(jsonBytes, password) // *keystore.Key	// 解密keystore文件
	if err != nil {
		zap.L().Fatal("解密keystore文件错误", zap.Error(err))
		return
	}
	// 生成私钥
	privateKey = hexutil.Encode(crypto.FromECDSA(Key.PrivateKey))
	// 生成公钥
	publicKey = hexutil.Encode(crypto.FromECDSAPub(&Key.PrivateKey.PublicKey))

	// 生成地址
	address = crypto.PubkeyToAddress(Key.PrivateKey.PublicKey).Hex()
	return
}
