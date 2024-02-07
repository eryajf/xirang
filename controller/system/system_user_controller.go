package system

import (
	systemReq "github.com/eryajf/xirang/model/system/request"
	"github.com/eryajf/xirang/public/tools"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// Add 添加记录
func (m *UserController) Add(c *gin.Context) {
	req := new(systemReq.UserAddReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return UserLogic.Add(c, req)
	})
}

// Update 更新记录
func (m *UserController) Update(c *gin.Context) {
	req := new(systemReq.UserUpdateReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return UserLogic.Update(c, req)
	})
}

// List 记录列表
func (m *UserController) List(c *gin.Context) {
	req := new(systemReq.UserListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return UserLogic.List(c, req)
	})
}

// Delete 删除记录
func (m UserController) Delete(c *gin.Context) {
	req := new(systemReq.UserDeleteReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return UserLogic.Delete(c, req)
	})
}

// ChangePwd 更新密码
func (m UserController) ChangePwd(c *gin.Context) {
	req := new(systemReq.UserChangePwdReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return UserLogic.ChangePwd(c, req)
	})
}

// ChangeUserStatus 更改用户状态
func (m UserController) ChangeUserStatus(c *gin.Context) {
	req := new(systemReq.UserChangeUserStatusReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return UserLogic.ChangeUserStatus(c, req)
	})
}

// GetUserInfo 获取当前登录用户信息
func (uc UserController) GetUserInfo(c *gin.Context) {
	req := new(systemReq.UserGetUserInfoReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return UserLogic.GetUserInfo(c, req)
	})
}
