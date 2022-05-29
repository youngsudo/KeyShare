package middlewares

import (
	"app/controllers"
	"app/pkg/jwt"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CookieMiddlewares() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			fmt.Printf("cookie err:%#v", err)
			controllers.ResponseError(c, "cookie err")
		}
		mc, err := jwt.ParseToken(cookie)
		if err != nil {
			controllers.ResponseError(c, "无效的Token")
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(controllers.CtxUserIDKey, mc.UserID)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
