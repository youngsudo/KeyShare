package eth

import "errors"

var (
	ErrorClientConnect = errors.New("连接到以太坊链失败")
	ErrorGetChainID    = errors.New("获取链ID失败")
	ErrorGetGasPrice   = errors.New("获取最低价格失败")
	ErrorKeyStore      = errors.New("创建绑定客户端失败")
	ErrorAddr          = errors.New("地址不合法")
	ErrorIsNotContract = errors.New("不是合约地址")
	ErrorIsContract    = errors.New("这是合约地址")
)
