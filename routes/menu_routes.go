package routes

import (
	"github.com/eryajf/xirang/controller"
	"github.com/eryajf/xirang/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitMenuRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	menu := r.Group("/menu")
	// 开启jwt认证中间件
	menu.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	menu.Use(middleware.CasbinMiddleware())
	{
		menu.GET("/tree", controller.Menu.GetTree)
		// menu.GET("/list", controller.Menu.List)
		menu.POST("/add", controller.Menu.Add)
		menu.POST("/update", controller.Menu.Update)
		menu.POST("/delete", controller.Menu.Delete)
	}

	return r
}
