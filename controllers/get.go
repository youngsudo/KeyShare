package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 首页
func GetIndexHandler(c *gin.Context) {

	userID, err := getCurrentUser(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("userID------>", userID)

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
