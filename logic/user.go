package logic

import (
	"app/dao/mysql"
	"app/eth"
	"app/models"
	"app/pkg/aesctr"
	bip "app/pkg/bip39"
	"app/pkg/snowflake"
	"app/setting"
	"crypto/ecdsa"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

var (
	ErrorIsExitUserAddress = errors.New("地址已存在")
	ErrorIsExitUserAccount = errors.New("用户名已存在")
	ErrorIsExitUserEmail   = errors.New("邮箱已存在")
	ErrorAesCtrPass        = errors.New("加密出错了")
)

func SignUp(p *models.ParamSignUp) error {

	// 1.加载管理员账户
	opts, err := eth.GetOpts(setting.Conf.ETHConfig.PrivateKey)
	if err != nil {
		return err
	}
	opts.GasLimit = 6721975
	opts.GasPrice = eth.GasPrice

	// 1.1判断地址是否存在
	if p.Address != "" { // 如果地址不为空，则判断地址是否存在
		b, err := eth.IsExitUserAddressView(p.Address)
		if err != nil {
			fmt.Println(err)
			return err
		} else if b {
			return ErrorIsExitUserAddress
		}
		// 继续向下执行
		fmt.Println("1判断地址是否存在", b)
	}

	// 1.2 判断用户名是否存在
	b, err := eth.IsExitUserAccountView(p.Account)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("用户名是否存在", b)

	if b {
		return ErrorIsExitUserAccount
	}

	// 1.3 判断邮箱是否存在
	b, err = eth.IsExitUserEmailView(p.Email)
	if err != nil {
		return err
	} else if b {
		return ErrorIsExitUserEmail
	}
	fmt.Println("邮箱是否存在", b)
	// 2. 生成用户ID
	userID := snowflake.GenID()
	fmt.Println("userID:", userID)

	// 3. 密码加密
	// pass := p.Password
	key := "1234567887654321"                                                 // 密钥
	encryptData, err := aesctr.AesCtrEncrypt([]byte(p.Password), []byte(key)) //加密 []bytes
	if err != nil {
		return ErrorAesCtrPass
	}

	pass := fmt.Sprintf("%x\n", encryptData)
	// 4. 创建用户
	// 如果没有传入地址，则生成地址,如果传入地址，则使用传入的地址,同时也将用户上传的助记词(任意字符串)保存到合约中
	ecdsaPrivateKey := new(ecdsa.PrivateKey) // 生成私钥
	mnemonic := new(string)                  // 生成助记词

	if p.Address == "" && p.Mnemonic == "" {
		ecdsaPrivateKey, mnemonic, err = bip.DeriveAddressFromMnemonic(new(string)) //  *ecdsa.PrivateKey
		if err != nil {
			fmt.Println(err)
			zap.L().Error("bip.DeriveAddressFromMnemonic", zap.Error(err))
			return err
		}
	}
	// 用户瞎传助记词

	// 私钥,公钥,地址
	privateKey, publicKey, address := bip.GenerateKey(ecdsaPrivateKey)
	fmt.Println("privateKey:", privateKey)
	fmt.Println("publicKey:", publicKey)
	fmt.Println("address:", address)

	fmt.Println("助记词", *mnemonic)
	zap.L().Info("生成地址", zap.String("address", address))

	p.Address = address
	// 将用户信息添加到合约
	tx, err := eth.AddUserPublic(opts, p.Address, p.Account, p.Password, p.Email, p.Mnemonic)
	if err != nil {
		zap.L().Info("AddUserPublic", zap.Error(err))
		fmt.Println(err)
		return err
	}

	zap.L().Info("添加一个用户", zap.String("address", p.Address), zap.String("tx", tx.Hash().Hex()))
	fmt.Println("添加一个用户", p.Address, tx.Hash().Hex())
	// 4. 将用户信息写入数据库
	err = mysql.InsertUser(userID, p.Address, p.Account, p.Email, pass, privateKey)
	if err != nil {
		fmt.Println("MySQL添加用户错误", err)
		zap.L().Warn("MySQL添加用户错误", zap.Error(err))
		return err
	}
	fmt.Println("MySQL添加用户成功")
	// 5. 将用户信息写入缓存
	// 6. 返回响应
	return nil
}
