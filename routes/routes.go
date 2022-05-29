package routes

import (
	"app/controllers"
	"app/logger"
	"app/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) // 日志中间件
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")

	// 路由组
	login := r.Group("/api/v1") // 不需要鉴权的路由
	{
		login.GET("/login", controllers.GetLoginHandler)
		login.GET("/register", controllers.GetRegisterHandler)
		login.GET("/registerToAddress", controllers.GetRegisterToAddressHandler)
		login.GET("/forgotPassword", controllers.GetForgetPasswordHandler) // 忘记密码

		login.POST("/login", controllers.LoginHandler)
		login.POST("/register", controllers.SignUpHandler)
		login.POST("/sendVerifyCode", controllers.SendVerifyCodeHandler)
	}

	app := r.Group("/api/v1", middlewares.CookieMiddlewares()) // 需要鉴权的路由
	{
		app.GET("/", controllers.GetIndexHandler)
		app.GET("/index", controllers.GetIndexHandler)             // 首页
		app.GET("/key", controllers.GetKeyHandler)                 // key 钥匙
		app.GET("/information", controllers.GetInformationHandler) // 个人中心
		app.GET("/contact", controllers.GetContactUsHandler)       // 联系我们
	}

	{
		app.POST("/contact", controllers.ContactHandler)
	}

	return r
}
