package routes

import (
	"github.com/eryajf/xirang/controller"
	"github.com/eryajf/xirang/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitGroupRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	group := r.Group("/group")
	// 开启jwt认证中间件
	group.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	group.Use(middleware.CasbinMiddleware())
	{
		group.GET("/list", controller.Group.List)
		group.GET("/tree", controller.Group.GetTree)
		group.POST("/add", controller.Group.Add)
		group.POST("/update", controller.Group.Update)
		group.POST("/delete", controller.Group.Delete)
		group.POST("/adduser", controller.Group.AddUser)
		group.POST("/removeuser", controller.Group.RemoveUser)

		group.GET("/useringroup", controller.Group.UserInGroup)
		group.GET("/usernoingroup", controller.Group.UserNoInGroup)
	}

	return r
}
