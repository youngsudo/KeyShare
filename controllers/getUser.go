package controllers

import (
	"app/dao/mysql"
	"app/dao/redis"
	"app/models"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var ErrorUserNotLogin = errors.New("用户未登录")

const CtxUserIDKey = "userID"
const CtxInformationKey = "information"

// 获取当前登陆用户ID
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	userID, ok = uid.(int64) // 强制类型
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// 获取当前登陆用户信息
func GetUserInformation(c *gin.Context) (user *models.DBUser, err error) {
	userID, err := getCurrentUser(c)
	if err != nil {
		fmt.Println("获取当前登陆用户信息错误:", err)
		ResponseError(c, ErrorUserNotLogin)
		return
	}
	fmt.Println("userID------>", userID)
	// 先从redis中获取
	user, err = redis.GetUserById(userID)
	fmt.Println("redis中获取", user)

	if errors.Is(err, redis.ErrorNotData) { // 缓存中没有数据
		// 从mysql中获取
		user, err = mysql.GetUserByUserIdOrAddress(userID)
		if err != nil {
			fmt.Println(err)
			zap.L().Debug("get user failed", zap.Error(err))
			ResponseError(c, "获取用户信息失败")
			return
		}
		// 写入redis
		err = redis.SetUser(userID, user)
		if err != nil {
			zap.L().Debug("redis.SetUser failed", zap.Error(err))
			ResponseError(c, "login redis.SetUser failed")
			return
		}
		return user, nil
	} else if err != nil {
		fmt.Println(err)
		ResponseError(c, err)
		return nil, err
	} else {
		fmt.Printf("获取当前登陆用户信息 redis----------->%#v\n", user)
		return
	}
}
