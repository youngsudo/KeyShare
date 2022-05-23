package eth

import (
	"app/eth/contract"
	"app/setting"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

var (
	KeyShare          *contract.KeyShare
	ChainID, GasPrice *big.Int
	client            *ethclient.Client
)

func Init(cfg *setting.ETHConfig) error {
	var err error
	client, err = ethclient.Dial(cfg.Ethclient)
	if err != nil {
		return ErrorClientConnect
	}
	// defer client.Close()

	ChainID, err = client.NetworkID(context.Background()) // 获取链ID
	if err != nil {
		return ErrorGetChainID
	}
	fmt.Println("chainID:", ChainID)

	GasPrice, err = client.SuggestGasPrice(context.Background()) // 获取最低价格
	if err != nil {
		return ErrorGetGasPrice
	}

	fmt.Println("gasPrice", GasPrice)

	// 创建绑定客户端 合约地址
	ins, err := contract.NewKeyShare(common.HexToAddress(cfg.Contract), client)
	if err != nil {
		return ErrorKeyStore
	}
	// fmt.Println(ins)
	KeyShare = ins
	return nil
}

// 传入执行合约的账户的私钥
func GetOpts(privateKey string) (*bind.TransactOpts, error) {
	// 获取私钥
	// HexToECDSA解析secp256k1私钥。
	key, err := crypto.HexToECDSA(privateKey) // *ecdsa.PrivateKey类型
	if err != nil {
		zap.L().Warn("获取私钥失败", zap.Error(err))
	}
	// key *ecdsa.PrivateKey, chainID *big.Int
	opts, err := bind.NewKeyedTransactorWithChainID(key, ChainID)
	if err != nil {
		zap.L().Info("NewTransactor", zap.Error(err))
	}
	return opts, err
}

// 关闭连接
func Close() {
	client.Close()
}
