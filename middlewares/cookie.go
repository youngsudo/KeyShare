package middlewares

import (
	"app/controllers"
	"app/pkg/jwt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CookieMiddlewares() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			zap.L().Info("cookie err", zap.Error(err))
			controllers.ResponseError(c, "无效的Token")
			// c.JSON(http.StatusUnauthorized, controllers.ErrorUserNotLogin)
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(cookie)
		if err != nil {
			zap.L().Info("cookie err", zap.Error(err))
			controllers.ResponseError(c, "无效的Token")
			// c.Redirect(http.StatusMovedPermanently, "/api/v1/login")
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(controllers.CtxUserIDKey, mc.UserID)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
