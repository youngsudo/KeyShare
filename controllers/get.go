package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 首页
func GetIndexHandler(c *gin.Context) {
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
