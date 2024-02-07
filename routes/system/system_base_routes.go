package system

import (
	"github.com/eryajf/xirang/controller"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

// 注册基础路由
func (s BaseRouter) InitBaseRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	base := r.Group("/base")
	{
		base.GET("ping", controller.Demo)
		base.GET("getpasswd", BaseController.GetPasswd)  // 将明文字符串转为MySQL识别的密码
		base.GET("/dashboard", BaseController.Dashboard) // 系统首页展示数据
		// 登录登出刷新token无需鉴权
		base.POST("/login", authMiddleware.LoginHandler)
		base.POST("/logout", authMiddleware.LogoutHandler)
		base.POST("/refreshToken", authMiddleware.RefreshHandler)
	}
	return r
}
