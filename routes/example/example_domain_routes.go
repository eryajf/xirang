package example

import (
	"github.com/eryajf/xirang/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type ExamleDomainRouter struct{}

func (s ExamleDomainRouter) InitExamleRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	domain := r.Group("/domain")
	// 开启jwt认证中间件
	domain.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	domain.Use(middleware.CasbinMiddleware())

	{
		domain.GET("/list", DomainController.List) // 标签键列表
	}

	return r
}
