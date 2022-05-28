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

	addr, ok := c.Get("account")
	if !ok {
		fmt.Println(err)
	}
	fmt.Println("Account:------->", addr)

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
