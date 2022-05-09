package router

import (
	"github.com/eryajf/xirang/controller"
	"github.com/eryajf/xirang/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(viper.GetString("SESSION_SECRET")))
	// 跨域的处理
	r.Use(middleware.Cors())
	// 获取用户身份
	r.Use(middleware.CurrentUser())

	r.POST("/api/captcha/check", controller.CheckToken)
	r.POST("/api/user/check_token", controller.CheckToken)
	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", controller.Ping)
		// 用户登录
		v1.POST("/user/login", controller.UserLogin)
		v1.GET("/getCaptcha", controller.GenerateCaptchaHandler)
	}
	// v1.Use(middleware.JWTAuth())   // 测试场景禁用jwt校验，便于接口测试
	{
		// 用户管理
		v1.POST("/user", controller.AddUser)                             // 注册用户
		v1.GET("/user/:user_id", controller.GetUser)                     // 查看用户
		v1.PUT("/user/:user_id", controller.UpdataUser)                  // 修改用户
		v1.DELETE("/user/:user_id", controller.DeleteUser)               // 删除用户
		v1.GET("/users", controller.ListUser)                            // 查看所有用户
		v1.PUT("/user/:user_id/state/:type", controller.ChangeUserState) // 修改用户状态
		// v1.PUT("user/:id/password", controller.ChangeUserPassword) // 修改用户密码

		// 用户组管理
		v1.POST("group", controller.AddGroup)                              // 新增用户组
		v1.GET("group/:group_id", controller.GetGroup)                     // 查看用户组
		v1.PUT("group/:group_id", controller.UpdataGroup)                  // 修改用户组
		v1.DELETE("group/:group_id", controller.DeleteGroup)               // 删除用户组
		v1.GET("groups", controller.ListGroup)                             // 查看所有用户组
		v1.PUT("group/:group_id/state/:type", controller.ChangeGroupState) // 修改用户状态

		// 用户与组
		v1.POST("groupuser/:group_id", controller.AddUToG)         //添加用户到分组
		v1.GET("groupuser/:group_id", controller.GetGroupUser)     //获取组内用户
		v1.GET("groupnouser/:group_id", controller.GetGroupNoUser) //获取不在组内的用户
	}
	return r
}
