package logic

import (
	"app/dao/mysql"
	"app/eth"
	"app/models"
	bip "app/pkg/bip39"
	"app/pkg/password"
	"app/pkg/snowflake"
	"app/setting"
	"errors"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
)

var (
	ErrorIsExitUserAddress = errors.New("地址已存在")
	ErrorIsExitUserAccount = errors.New("用户名已存在")
	ErrorIsExitUserEmail   = errors.New("邮箱已存在")
	ErrorAesCtrPass        = errors.New("加密出错了")
	ErrorLogin             = errors.New("登陆失败")
	ErrorAccountPassword   = errors.New("账号或密码错误")
)

// 注册
func SignUp(p *models.ParamSignUp) error {

	// 1.加载管理员账户
	opts, err := eth.GetOpts(setting.Conf.ETHConfig.PrivateKey)
	if err != nil {
		return err
	}
	gasPrice, err := eth.GetGasPrice()
	if err != nil {
		zap.L().Error("获取gasPrice失败", zap.Error(err))
	}
	opts.GasLimit = 6721975
	opts.GasPrice = gasPrice

	// 1.1判断地址是否存在
	if p.Address != "" { // 如果地址不为空，则判断地址是否存在
		b, err := eth.IsExitUserAddressView(p.Address)
		if err != nil {
			zap.L().Debug("eth.IsExitUserAddressView failed", zap.Error(err))
			return err
		} else if b {
			return ErrorIsExitUserAddress
		}
		// 继续向下执行
	}

	// 1.2 判断用户名是否存在
	b, err := eth.IsExitUserAccountView(p.Account)
	if err != nil {
		zap.L().Debug("eth.IsExitUserAccountView failed", zap.Error(err))
		return err
	}

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
	zap.L().Debug("邮箱不存在")
	fmt.Println("邮箱不存在")
	// 2. 生成用户ID
	userID := snowflake.GenID()
	zap.L().Debug("userID", zap.Int64("userID", userID))

	// 3. 密码加密
	// pass := p.Password
	pass, err := password.HashPassword(p.Password)
	if err != nil {
		return ErrorAesCtrPass
	}
	// 4. 创建用户
	// 如果没有传入地址，则生成地址,如果传入地址，则使用传入的地址,同时也将用户上传的助记词(任意字符串)保存到合约中
	// ecdsaPrivateKey := new(ecdsa.PrivateKey) // 生成私钥
	mnemonic := new(string) // 生成助记词
	var privateKey string   // 生成私钥
	var publicKey string    // 生成公钥
	var address string      // 生成地址

	if p.Address == "" && p.Mnemonic == "" {
		ecdsaPrivateKey, _mnemonic, err := bip.DeriveAddressFromMnemonic(new(string)) //  *ecdsa.PrivateKey
		if err != nil {
			zap.L().Error("bip.DeriveAddressFromMnemonic", zap.Error(err))
			return err
		}
		mnemonic = _mnemonic
		privateKey, publicKey, address = bip.GenerateKey(ecdsaPrivateKey)

	} else if p.Address != "" && p.Mnemonic != "" {
		ecdsaPrivateKey, err := crypto.HexToECDSA(p.Mnemonic)
		if err != nil {
			zap.L().Error("bip.DeriveAddressFromMnemonic", zap.Error(err))
			return err
		}
		privateKey, publicKey, address = bip.GenerateKey(ecdsaPrivateKey)
		if p.Address != address {
			return errors.New("地址不匹配")
		}

	} else {
		return errors.New("地址和私钥(助记词)不能同时为空")
	}
	// 生成私钥,公钥,地址
	zap.L().Debug("privateKey", zap.String("privateKey", privateKey))
	zap.L().Debug("publicKey", zap.String("publicKey", publicKey))
	zap.L().Debug("address", zap.String("address", address))
	zap.L().Debug("助记词", zap.String("mnemonic", *mnemonic))

	fmt.Println("地址", address)
	fmt.Printf("助记词:%v\n", *mnemonic)

	// 判断地址是否正确
	b, err = eth.IsContractAddr(address)
	if err != nil {
		zap.L().Error("eth.IsContractAddr", zap.Error(err))
		return eth.ErrorAddr
	}
	if b {
		return eth.ErrorIsContract
	}

	p.Address = address
	// 如果用户传了私钥(助记词)，则 moemonic 为空,反之则在上面代码中生成了 moemonic,即不为空
	if *mnemonic != "" {
		p.Mnemonic = *mnemonic
	}

	// 将用户信息添加到合约,不传加密的密码
	tx, err := eth.AddUserPublic(opts, p.Address, p.Account, p.Password, p.Email, p.Mnemonic)
	if err != nil {
		zap.L().Info("AddUserPublic", zap.Error(err))
		fmt.Println(err)
		return err
	}
	zap.L().Info("添加一个用户", zap.String("address", p.Address), zap.String("tx", tx.Hash().Hex()))

	// 4. 将用户信息写入数据库
	key := "1234567887654321"                                                   // 加密私钥                                               // 密钥
	encryptData, err := password.AesCtrEncrypt([]byte(privateKey), []byte(key)) //加密 []bytes
	if err != nil {
		return ErrorAesCtrPass
	}
	private_key := fmt.Sprintf("%x\n", encryptData)
	err = mysql.InsertUser(userID, p.Address, p.Account, p.Email, pass, private_key)
	if err != nil {
		zap.L().Warn("MySQL添加用户错误", zap.Error(err))
		return err
	}
	zap.L().Debug("MySQL添加用户成功")
	return nil
}

// 登陆
func Login(p *models.ParamLogin) (uint8, string, error) {
	zap.L().Debug("登陆", zap.String("account", p.Account), zap.String("password", p.Password))
	fmt.Println("登陆", p.Account, p.Password)
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$") // 正则表达式

	var _address string
	var _account string
	// re.MatchString() 方法返回 true 如果字符串 s 匹配正则表达式 re。
	if b := re.MatchString(p.Account); !b {
		// 地址为空则给一个初始值
		_address = "0x0000000000000000000000000000000000000000"
		_account = p.Account
	} else {
		// 判断地址是否正确？不用
		_address = p.Account
		// 账号为空则给一个初始值
		_account = ""
	}
	fmt.Println(_address, _account)

	// userType 用户身份, 0为超级管理员,1为管理员,2为普通用户,3 即为错误
	b, userType, address, err := eth.Login(_address, _account, p.Password)
	if err != nil {
		zap.L().Debug("eth.IsExitUserAddressView failed", zap.Error(err))
	}
	if !b {
		return 3, "", ErrorAccountPassword
	}
	// fmt.Println(userType, address)
	return userType, address, nil
}
