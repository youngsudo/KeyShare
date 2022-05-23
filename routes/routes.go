package routes

import (
	"app/controllers"
	"app/logger"

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
	app := r.Group("/api/v1")
	{
		app.GET("/", controllers.GetIndexHandler)
		app.GET("/index", controllers.GetIndexHandler)
		app.GET("/login", controllers.GetLoginHandler)
		app.GET("/register", controllers.GetRegisterHandler)
		app.GET("/registerToAddress", controllers.GetRegisterToAddressHandler)
	}
	{
		app.POST("/register", controllers.SignUpHandler)
		app.POST("/login", controllers.LoginHandler)
	}

	return r
}
