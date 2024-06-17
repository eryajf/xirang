package system

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type DeptRouter struct{}

func (s DeptRouter) InitGroupRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	dept := r.Group("/dept")
	// 开启jwt认证中间件
	// dept.Use(authMiddleware.MiddlewareFunc())
	// // 开启casbin鉴权中间件
	// dept.Use(middleware.CasbinMiddleware())

	{
		dept.GET("/list", DeptController.List)
		dept.GET("/tree", DeptController.GetTree)
		dept.POST("/add", DeptController.Add)
		dept.POST("/update", DeptController.Update)
		dept.POST("/delete", DeptController.Delete)
		dept.POST("/adduser", DeptController.AddUser)
		dept.POST("/removeuser", DeptController.RemoveUser)
	}

	return r
}
