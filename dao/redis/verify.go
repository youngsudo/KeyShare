package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

// 设置验证码
func SetVerify(email, verify string) (err error) {
	err = rdb.Set(fmt.Sprintf("email:%s", email), verify, 60*time.Second).Err()
	if err != nil {
		zap.L().Error("set email verify failed", zap.Error(err))
		return
	}
	return
}

// 获取验证码
func GetVerify(email string) (verify string, err error) {
	verify, err = rdb.Get(fmt.Sprintf("email:%s", email)).Result()
	if err == redis.Nil {
		zap.L().Error("email verify does not exist", zap.Error(err))
		return
	} else if err != nil {
		zap.L().Error("get verify failed", zap.Error(err))
		return
	} else {
		return verify, nil
	}
}
