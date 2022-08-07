package redis

import (
	"app/models"
	"app/setting"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var ErrorNotData = errors.New("没有数据")

func SetUser(userId int64, user *models.DBUser) (err error) {
	userbyte, err := user.MarshalBinary() // 序列化
	if err != nil {
		zap.L().Warn("MarshalBinary failed", zap.Error(err))
	}
	fmt.Println("===========================\nredis TokenExpiration:", setting.Conf.TokenExpiration)
	err = rdb.Set(fmt.Sprintf("token:%v", userId), userbyte, setting.Conf.TokenExpiration*time.Minute).Err()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetUserById(userId int64) (user *models.DBUser, err error) {
	// rdb.HGetAll()
	// rdb.HMGet()
	// rdb.HGet()
	s, err := rdb.Get(fmt.Sprintf("token:%v", userId)).Result()
	if err == redis.Nil {
		zap.L().Error("email verify does not exist", zap.Error(err))
		return nil, ErrorNotData
	} else if err != nil {
		zap.L().Error("get verify failed", zap.Error(err))
		return
	} else {
		// 反序列化 s
		// 创建一个变量
		user = &models.DBUser{}
		// 传入json字符串，和指针
		err = json.Unmarshal([]byte(s), user)
		if err != nil {
			zap.L().Warn("反序列化失败", zap.Error(err))
		}
		return
	}
}
