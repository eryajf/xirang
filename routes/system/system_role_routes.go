package system

import (
	"github.com/eryajf/xirang/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (s RoleRouter) InitRoleRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	role := r.Group("/role")
	// 开启jwt认证中间件
	role.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	role.Use(middleware.CasbinMiddleware())
	{
		role.GET("/list", RoleController.List)
		role.POST("/add", RoleController.Add)
		role.POST("/update", RoleController.Update)
		role.POST("/delete", RoleController.Delete)

		role.GET("/getmenulist", RoleController.GetMenuList)
		role.GET("/getapilist", RoleController.GetApiList)
		role.POST("/updatemenus", RoleController.UpdateMenus)
		role.POST("/updateapis", RoleController.UpdateApis)
	}
	return r
}
