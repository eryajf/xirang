package routes

import (
	"fmt"
	"time"

	"github.com/eryajf/xirang/config"
	"github.com/eryajf/xirang/middleware"
	"github.com/eryajf/xirang/public/common"
	"github.com/eryajf/xirang/routes/example"
	"github.com/eryajf/xirang/routes/system"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

// 初始化
func InitRoutes() *gin.Engine {
	//设置模式
	gin.SetMode(config.Conf.System.Mode)

	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	r := gin.Default()
	// 创建不带中间件的路由:
	// r := gin.New()
	// r.Use(gin.Recovery())

	// 启用限流中间件
	// 默认每50毫秒填充一个令牌，最多填充200个
	fillInterval := time.Duration(config.Conf.RateLimit.FillInterval)
	capacity := config.Conf.RateLimit.Capacity
	r.Use(middleware.RateLimitMiddleware(time.Millisecond*fillInterval, capacity))

	// 启用全局跨域中间件
	r.Use(middleware.CORSMiddleware())

	// 启用操作日志中间件
	r.Use(middleware.OperationLogMiddleware())

	// 初始化JWT认证中间件
	authMiddleware, err := middleware.InitAuth()
	if err != nil {
		common.Log.Panicf("初始化JWT中间件失败：%v", err)
		panic(fmt.Sprintf("初始化JWT中间件失败：%v", err))
	}

	// 基础路由分组
	apiGroup := r.Group("/" + config.Conf.System.UrlPathPrefix)

	// 路由分组
	systemApiGroup := apiGroup.Group("/system/v1")

	// 注册路由
	RouterGroupApp.System.InitBaseRoutes(systemApiGroup, authMiddleware)         // 注册基础路由, 不需要jwt认证中间件,不需要casbin中间件
	RouterGroupApp.System.InitUserRoutes(systemApiGroup, authMiddleware)         // 注册用户路由, jwt认证中间件,casbin鉴权中间件
	RouterGroupApp.System.InitGroupRoutes(systemApiGroup, authMiddleware)        // 注册分组路由, jwt认证中间件,casbin鉴权中间件
	RouterGroupApp.System.InitRoleRoutes(systemApiGroup, authMiddleware)         // 注册角色路由, jwt认证中间件,casbin鉴权中间件
	RouterGroupApp.System.InitMenuRoutes(systemApiGroup, authMiddleware)         // 注册菜单路由, jwt认证中间件,casbin鉴权中间件
	RouterGroupApp.System.InitApiRoutes(systemApiGroup, authMiddleware)          // 注册接口路由, jwt认证中间件,casbin鉴权中间件
	RouterGroupApp.System.InitOperationLogRoutes(systemApiGroup, authMiddleware) // 注册操作日志路由, jwt认证中间件,casbin鉴权中间件

	// 路由分组
	cmdbApiGroup := apiGroup.Group("/example/v1")
	RouterGroupApp.Example.InitExamleRoutes(cmdbApiGroup, authMiddleware)

	common.Log.Info("初始化路由完成！")
	return r
}
