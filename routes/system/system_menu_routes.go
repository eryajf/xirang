package system

import (
	"github.com/eryajf/xirang/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (s MenuRouter) InitMenuRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	menu := r.Group("/menu")
	// 开启jwt认证中间件
	menu.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	menu.Use(middleware.CasbinMiddleware())
	{
		menu.GET("/tree", MenuController.GetTree)
		menu.GET("/access/tree", MenuController.GetAccessTree)
		menu.POST("/add", MenuController.Add)
		menu.POST("/update", MenuController.Update)
		menu.POST("/delete", MenuController.Delete)
	}

	return r
}
