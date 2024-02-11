package example

import (
	"github.com/eryajf/xirang/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type CloudAccount struct{}

func (s CloudAccount) InitExamleRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	cloudAccount := r.Group("/cloudaccount")
	// 开启jwt认证中间件
	cloudAccount.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	cloudAccount.Use(middleware.CasbinMiddleware())

	{
		cloudAccount.POST("/add", CloudAccountController.Add)
		cloudAccount.GET("/list", CloudAccountController.List)
		cloudAccount.POST("/update", CloudAccountController.Update)
		cloudAccount.POST("/delete", CloudAccountController.Delete)
	}

	return r
}
