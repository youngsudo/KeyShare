package main

import (
	"app/controllers"
	"app/dao/mysql"
	"app/dao/redis"
	"app/eth"
	"app/logger"
	"app/pkg/keystore"
	"app/pkg/port"
	"app/pkg/snowflake"
	"app/routes"
	"app/setting"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {

	// 1. 加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Printf("加载配置文件失败,err: %v\n", err)
		os.Exit(1)
	}

	// 2. 初始化日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.App.Mode); err != nil {
		fmt.Printf("初始化日志库错误,err :%v\n", err)
		os.Exit(1)
	}
	zap.L().Debug("初始化日志完成")
	zap.L().Sync() // 将日志写入磁盘

	// 3. 初始化 MySQL 数据库
	if err := mysql.Init(setting.Conf.MysqlConfig); err != nil {
		zap.L().Fatal("初始化 MySQL 失败", zap.Error(err))
	}
	zap.L().Debug("初始化 MySQL 完成")
	defer mysql.Close()

	// 4. 初始化Redis连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		zap.L().Fatal("初始化Redis失败", zap.Error(err))
	}
	zap.L().Debug("初始化 Redis 完成")
	defer redis.Close()

	// 5.雪花算法
	if err := snowflake.Init(setting.Conf.App.StartTime, setting.Conf.App.MachineID); err != nil {
		zap.L().Fatal("init snowflake failed", zap.Error(err))
	}

	// // 初始化gin框架内置的校验器使用的翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		zap.L().Fatal("init validator trans failed", zap.Error(err))
	}

	// 初始化 ETH
	if err := eth.Init(setting.Conf.ETHConfig); err != nil {
		zap.L().Fatal("init ETH failed", zap.Error(err))
	}

	// 获取管理员私钥
	// 1.如果直接在配置文件中写有管理员的私钥

	// 2.在 pkg 的 keystore 目录下 wallets 中解密出管理员的私钥
	if setting.Conf.ETHConfig.PrivateKey == "" {
		privateKey, _, _, err := keystore.ImportKs(setting.Conf.ETHConfig.KeystoreName)
		if err != nil {
			zap.L().Fatal("import ks failed", zap.Error(err))
		}
		setting.Conf.ETHConfig.PrivateKey = privateKey
	}

	// 6. 注册路由
	r := routes.Setup(setting.Conf.App.Mode)

	// 7. 启动服务,优雅关机
	port := port.FindPort(setting.Conf.App.Port)
	fmt.Println("服务启动成功,端口号:", port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("ListenAndServe", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1)                      // 创建一个接收信号的通道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
