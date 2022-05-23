package models

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Address    string `json:"address"`                                         // 地址,允许为空
	Mnemonic   string `json:"mnemonic"`                                        // 助记词
	Account    string `json:"account" binding:"required"`                      // 用户名,必填
	Email      string `json:"email" binding:"required"`                        // 邮箱,必填
	Password   string `json:"password" binding:"required"`                     // 密码,必填
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 确认密码,必填,与密码一致
	VerifyCode string `json:"verifyCode" binding:"required"`                   // 验证码,必填
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
