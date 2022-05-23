package controllers

import (
	"app/logic"
	"app/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 首页
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, "请求参数错误")
			return
		}
		ResponseError(c, removeTopStruct(errs.Translate(trans)))
		return
	}
	fmt.Printf("%#v\n", p)
	// &models.ParamSignUp{Address:"0xf4EF873f6F0feE8b1DFe83C33386AFDB448D46Bf", Account:"youngh", Email:"111@qq.com", Password:"qqqqqq", RePassword:"qqqqqq", VerifyCode:"qqqq"}SignUpH

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		ResponseError(c, err)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}
