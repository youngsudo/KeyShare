package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 首页
func GetIndexHandler(c *gin.Context) {
	// id, err := strconv.Atoi(c.Query("id"))
	// if err != nil {
	// 	zap.L().Debug("url无参数", zap.Error(err))
	// }

	// user, err := GetUserInformationById(int64(id))
	// if err != nil {
	// 	zap.Error(err)
	// }
	// fmt.Println(user)
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// 登陆
func GetLoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

// 注册
func GetRegisterHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
func GetRegisterToAddressHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "registerByAddress.html", gin.H{})
}

// key 钥匙
func GetKeyHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "key.html", gin.H{})
}

// 个人中心
func GetInformationHandler(c *gin.Context) {
	// 获取当前用户的信息
	c.HTML(http.StatusOK, "information.html", gin.H{})
}

// 联系我们
func GetContactUsHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.html", gin.H{})
}

// 忘记密码
func GetForgetPasswordHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "forgetPassword.html", gin.H{})
}

//  返回用户信息的路由
func GetUserInformationHandler(c *gin.Context) {
	user, err := GetUserInformation(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	fmt.Println(user.ID)
	fmt.Println(user.ID == 0)
	if user.ID == 0 {
		c.Redirect(http.StatusMovedPermanently, "/api/v1/login") // 跳转到登陆页面
	} else {
		c.JSON(http.StatusOK, user)
	}
}
