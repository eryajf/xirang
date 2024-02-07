package system

import (
	"github.com/eryajf/xirang/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

// 注册用户路由
func (s UserRouter) InitUserRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	user := r.Group("/user")
	// 开启jwt认证中间件
	user.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	user.Use(middleware.CasbinMiddleware())
	{
		user.GET("/info", UserController.GetUserInfo)                   // 暂时未完成
		user.GET("/list", UserController.List)                          // 用户列表
		user.POST("/add", UserController.Add)                           // 添加用户
		user.POST("/update", UserController.Update)                     // 更新用户
		user.POST("/delete", UserController.Delete)                     // 删除用户
		user.POST("/changePwd", UserController.ChangePwd)               // 修改用户密码
		user.POST("/changeUserStatus", UserController.ChangeUserStatus) // 修改用户状态
	}
	return r
}
