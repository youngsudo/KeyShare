package password

// AES算法 + CBC模式 对称加密
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// AES算法 + CTR模式 对称加密
// 输入明文，输出密文
func AesCtrEncrypt(plainText, key []byte) ([]byte, error) {
	// 第一步：创建aes密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//第二步：创建分组模式ctr
	// iv 要与算法长度一致，16字节
	// 使用bytes.Repeat创建一个切片，长度为blockSize()，16个字符"1"
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)

	//第三步：加密
	dst := make([]byte, len(plainText))
	stream.XORKeyStream(dst, plainText)

	return dst, nil
}

//输入密文，得到明文
func AesCtrDecrypt(encryptData, key []byte) ([]byte, error) {
	return AesCtrEncrypt(encryptData, key)
}

// func main() {
// 	//明文，需要加密的数据
// 	src := "Stream接口代表一个流模式的加/解密器。"

// 	//对称秘钥，aes，16字节
// 	key := "1234567887654321" //16
// 	// key := "12345678876543210" //17, 无效的
// 	encryptData, err := aesCtrEncrypt([]byte(src), []byte(key)) //加密 []bytes
// 	if err != nil {
// 		fmt.Println("加密出错err:", err)
// 		return
// 	}

// 	fmt.Printf("encryptData： %x\n", encryptData)

// 	//调用解密函数
// 	plainText, err := aesCtrDecrypt(encryptData, []byte(key))
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}

// 	fmt.Printf("解密后的数据: %s\n", plainText)
// }
