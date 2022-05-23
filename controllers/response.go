package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	// "code": 10000, // 程序中的错误码
	"msg": xx,     // 提示信息
	"data": {},    // 数据
}

*/
// 定义返回数据结构
type ResponseData struct {
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// 返回错误和数据
func ResponseError(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Msg:  msg,
		Data: nil,
	})
}

// 返回成功
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Msg:  "success",
		Data: data, // 要返回的数据
	})
}
