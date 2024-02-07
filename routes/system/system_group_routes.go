package system

import (
	"github.com/eryajf/xirang/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type GroupRouter struct{}

func (s GroupRouter) InitGroupRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	group := r.Group("/group")
	// 开启jwt认证中间件
	group.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	group.Use(middleware.CasbinMiddleware())

	{
		group.GET("/list", GroupController.List)
		group.GET("/tree", GroupController.GetTree)
		group.POST("/add", GroupController.Add)
		group.POST("/update", GroupController.Update)
		group.POST("/delete", GroupController.Delete)
		group.POST("/adduser", GroupController.AddUser)
		group.POST("/removeuser", GroupController.RemoveUser)

		group.GET("/useringroup", GroupController.UserInGroup)
		group.GET("/usernoingroup", GroupController.UserNoInGroup)
	}

	return r
}
