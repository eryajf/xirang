package system

import (
	systemReq "github.com/eryajf/xirang/model/system/request"
	"github.com/eryajf/xirang/public/tools"

	"github.com/gin-gonic/gin"
)

type DeptController struct{}

// List 记录列表
func (m *DeptController) List(c *gin.Context) {
	req := new(systemReq.DeptListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return DeptLogic.List(c, req)
	})
}

// UserInDept 在分组内的用户
func (m *DeptController) UserInDept(c *gin.Context) {
	req := new(systemReq.DeptInDeptReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return DeptLogic.UserInDept(c, req)
	})
}

// UserNoInDept 不在分组的用户
func (m *DeptController) UserNoInDept(c *gin.Context) {
	req := new(systemReq.UserNoInDeptReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return DeptLogic.UserNoInDept(c, req)
	})
}

// GetTree 接口树
func (m *DeptController) GetTree(c *gin.Context) {
	req := new(systemReq.DeptListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return DeptLogic.GetTree(c, req)
	})
}

// Add 新建记录
func (m *DeptController) Add(c *gin.Context) {
	req := new(systemReq.DeptAddReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return DeptLogic.Add(c, req)
	})
}

// Update 更新记录
func (m *DeptController) Update(c *gin.Context) {
	req := new(systemReq.DeptUpdateReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return DeptLogic.Update(c, req)
	})
}

// Delete 删除记录
func (m *DeptController) Delete(c *gin.Context) {
	req := new(systemReq.DeptDeleteReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return DeptLogic.Delete(c, req)
	})
}

// AddUser 添加用户
func (m *DeptController) AddUser(c *gin.Context) {
	req := new(systemReq.DeptAddUserReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return DeptLogic.AddUser(c, req)
	})
}

// RemoveUser 移除用户
func (m *DeptController) RemoveUser(c *gin.Context) {
	req := new(systemReq.DeptRemoveUserReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return DeptLogic.RemoveUser(c, req)
	})
}
