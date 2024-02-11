package system

import (
	systemReq "github.com/eryajf/xirang/model/system/request"
	"github.com/eryajf/xirang/public/tools"

	"github.com/gin-gonic/gin"
)

type GroupController struct{}

// List 记录列表
func (m *GroupController) List(c *gin.Context) {
	req := new(systemReq.GroupListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return GroupLogic.List(c, req)
	})
}

// UserInGroup 在分组内的用户
func (m *GroupController) UserInGroup(c *gin.Context) {
	req := new(systemReq.UserInGroupReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return GroupLogic.UserInGroup(c, req)
	})
}

// UserNoInGroup 不在分组的用户
func (m *GroupController) UserNoInGroup(c *gin.Context) {
	req := new(systemReq.UserNoInGroupReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return GroupLogic.UserNoInGroup(c, req)
	})
}

// GetTree 接口树
func (m *GroupController) GetTree(c *gin.Context) {
	req := new(systemReq.GroupListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return GroupLogic.GetTree(c, req)
	})
}

// Add 新建记录
func (m *GroupController) Add(c *gin.Context) {
	req := new(systemReq.GroupAddReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return GroupLogic.Add(c, req)
	})
}

// Update 更新记录
func (m *GroupController) Update(c *gin.Context) {
	req := new(systemReq.GroupUpdateReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return GroupLogic.Update(c, req)
	})
}

// Delete 删除记录
func (m *GroupController) Delete(c *gin.Context) {
	req := new(systemReq.GroupDeleteReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return GroupLogic.Delete(c, req)
	})
}

// AddUser 添加用户
func (m *GroupController) AddUser(c *gin.Context) {
	req := new(systemReq.GroupAddUserReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return GroupLogic.AddUser(c, req)
	})
}

// RemoveUser 移除用户
func (m *GroupController) RemoveUser(c *gin.Context) {
	req := new(systemReq.GroupRemoveUserReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return GroupLogic.RemoveUser(c, req)
	})
}
