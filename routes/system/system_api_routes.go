package system

import (
	"github.com/eryajf/xirang/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (s ApiRouter) InitApiRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	api := r.Group("/api")
	// 开启jwt认证中间件
	api.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	api.Use(middleware.CasbinMiddleware())
	{
		api.GET("/tree", ApiController.GetTree)
		api.GET("/list", ApiController.List)
		api.POST("/add", ApiController.Add)
		api.POST("/update", ApiController.Update)
		api.POST("/delete", ApiController.Delete)
	}

	return r
}
