package routes

import (
	"app/controllers"
	"app/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) // 日志中间件
	gin.SetMode(gin.DebugMode)                          //设置模式 ReleaseMode生产模式,DebugMode开发模式

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
	}

	return r
}
