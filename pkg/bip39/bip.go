// 生成随机的助记词
package bip39

import (
	"fmt"

	"github.com/tyler-smith/go-bip39"
)

func bip() (*string, error) {
	// entropy 生成, y = 32 * x , 128 <= y <= 256
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("entropy:", entropy)
	// 生成随机的助记词
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("助记符mnemonic:", mnemonic)
	return &mnemonic, nil
}
