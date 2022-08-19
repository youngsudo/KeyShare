package eth

import (
	"app/eth/contract"
	"app/setting"
	"context"
	"fmt"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

var (
	KeyShare *contract.KeyShare
	Account  *contract.Account
	client   *ethclient.Client
)

func Init(cfg *setting.ETHConfig) error {
	var err error
	client, err = ethclient.Dial(cfg.Ethclient)
	if err != nil {
		return ErrorClientConnect
	}
	// defer client.Close()

	// 创建绑定客户端 合约地址
	b, err := IsContractAddr(cfg.Contract)
	if err != nil {
		zap.L().Warn("判断合约地址是否是合约地址失败", zap.Error(err), zap.String("contract", cfg.Contract))
		return err
	}
	if !b {
		zap.L().Warn("合约地址不是合约地址", zap.String("contract", cfg.Contract))
		return ErrorIsNotContract
	}
	// ins, err := contract.NewKeyShare(common.HexToAddress(cfg.Contract), client)
	// if err != nil {
	// 	return ErrorKeyStore
	// }
	ins, err := contract.NewAccount(common.HexToAddress(cfg.Contract), client)
	if err != nil {
		return ErrorKeyStore
	}
	// fmt.Println(ins)
	// KeyShare = ins
	Account = ins
	return nil
}

// 判断传入的地址是否是合约地址
// true 表示是合约地址, false 表示不是合约地址,而是用户地址
func IsContractAddr(addr string) (bool, error) {

	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$") // 正则表达式
	// 判断是否是 一个地址字符串
	isAddrStr := re.MatchString(addr)

	if !isAddrStr {
		return false, ErrorAddr
	}
	// 当地址上没有字节码时，我们知道它不是一个智能合约，它是一个标准的以太坊账户。
	address := common.HexToAddress(addr)
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		zap.L().Fatal("获取合约代码失败", zap.Error(err))
	}

	return len(bytecode) > 0, nil
}

// 获取chainID
func GetChainID() (*big.Int, error) {
	chainID, err := client.NetworkID(context.Background()) // 获取链ID
	if err != nil {
		return big.NewInt(0), ErrorGetChainID
	}
	fmt.Println("chainID:", chainID)
	return chainID, nil
}

// 获取gasPrice
func GetGasPrice() (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background()) // 获取最低价格
	if err != nil {
		return big.NewInt(0), ErrorGetGasPrice
	}

	fmt.Println("gasPrice", gasPrice)
	return gasPrice, err
}

// 传入执行合约的账户的私钥
func GetOpts(privateKey string) (*bind.TransactOpts, error) {
	// 获取私钥
	// HexToECDSA解析secp256k1私钥。
	key, err := crypto.HexToECDSA(privateKey) // *ecdsa.PrivateKey类型
	if err != nil {
		zap.L().Warn("获取私钥失败", zap.Error(err))
	}
	chainID, err := GetChainID()
	if err != nil {
		zap.L().Error("获取链ID失败", zap.Error(err))
	}
	zap.L().Debug("chainID", zap.Any("chainID", chainID))
	// key *ecdsa.PrivateKey, chainID *big.Int
	opts, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		zap.L().Info("NewTransactor", zap.Error(err))
	}
	return opts, err
}

// 关闭连接
func Close() {
	client.Close()
}
