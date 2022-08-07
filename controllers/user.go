package controllers

import (
	"app/dao/mysql"
	"app/dao/redis"
	"app/logic"
	"app/models"
	"app/pkg/jwt"
	"app/setting"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 发送验证码
func SendVerifyCodeHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	var email = c.PostForm("email")
	if email == "" {
		ResponseError(c, "请输入邮箱")
		return
	}
	// 2. 业务处理
	code, err := logic.SendVerify(email)
	if err != nil {
		zap.L().Error("logic.SendVerifyCode failed", zap.Error(err))
		ResponseError(c, err)
		return
	}
	// 3. 存入缓存
	fmt.Println("发送验证码:", email, code)
	redis.SetVerify(email, code)
	// 4. 返回响应
	ResponseSuccess(c, "发送成功")
}

// 注册
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

// 登录
func LoginHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, "请求参数错误")
			return
		}
		ResponseError(c, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 业务处理
	userType, address, err := logic.Login(p)
	if err != nil {
		zap.L().Error("登录失败", zap.Error(err))
		ResponseError(c, "登录失败")
		return
	}

	//3. 都正确则查询数据库,获取用户信息
	user, err := mysql.GetUserByUserIdOrAddress(address)
	fmt.Printf("=====>%#v\n", user)
	userId := user.UserID
	fmt.Println(userId)
	if err != nil {
		fmt.Println("查询用户信息失败", err)
		ResponseError(c, "查询用户信息失败")
	}

	// 存入缓存
	err = redis.SetUser(userId, user)
	if err != nil {
		zap.L().Error("redis.SetUser failed", zap.Error(err))
	}

	// 生成Token
	tokenString, _ := jwt.GenToken(userId, setting.Conf.TokenExpiration)
	// 使用 cookie
	// c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true) // 3600 秒 = 1小时

	// 返回响应
	// ResponseSuccess(c, map[string]interface{}{
	// 	"userType": userType,
	// 	"token":    tokenString,
	// })
	ResponseSuccess(c, gin.H{
		"userType": userType,
		"id":       userId,
		"token":    tokenString,
	})
}

// 联系我们
func ContactHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamContract)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Contaxt with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, "联系我们的请求参数错误")
			return
		}
		ResponseError(c, removeTopStruct(errs.Translate(trans)))
		return
	}
	fmt.Printf("%#v\n", p)
	// 2. 业务处理
	if err := logic.ContactUS(p); err != nil {
		zap.L().Error("logic.ContactUS failed", zap.Error(err))
		ResponseError(c, err)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, "发送成功")
}
