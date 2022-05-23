// 推导私钥
package bip39

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
)

// 将私钥转换成 私钥,公钥,地址字符串
func GenerateKey(privateKey *ecdsa.PrivateKey) (string, string, string) {
	privateKeyBytes := crypto.FromECDSA(privateKey) // 使用FromECDSA方法将其转换为字节
	// fmt.Println("私钥", hexutil.Encode(privateKeyBytes)[2:]) // 0x
	// 公钥
	publicKey := privateKey.Public()
	// 将其转换为十六进制的过程与我们使用转化私钥的过程类似。 我们剥离了0x和前2个字符04，它始终是EC前缀，不是必需的。
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // 类型断言 *ecdsa.PublicKey
	if !ok {
		zap.L().Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA) // 使用FromECDSAPub方法将其转换为字节
	// fmt.Println("公钥", hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	// fmt.Println("地址", address) // 0x96216849c49358B10257cb55b28eA603c874b05E
	return hexutil.Encode(privateKeyBytes)[2:], hexutil.Encode(publicKeyBytes)[4:], address
}
