package system

import (
	systemReq "github.com/eryajf/xirang/model/system/request"
	"github.com/eryajf/xirang/public/tools"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

// List 记录列表
func (m *RoleController) List(c *gin.Context) {
	req := new(systemReq.RoleListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return RoleLogic.List(c, req)
	})
}

// Add 新建
func (m *RoleController) Add(c *gin.Context) {
	req := new(systemReq.RoleAddReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return RoleLogic.Add(c, req)
	})
}

// Update 更新记录
func (m *RoleController) Update(c *gin.Context) {
	req := new(systemReq.RoleUpdateReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return RoleLogic.Update(c, req)
	})
}

// Delete 删除记录
func (m *RoleController) Delete(c *gin.Context) {
	req := new(systemReq.RoleDeleteReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return RoleLogic.Delete(c, req)
	})
}

// GetMenuList 获取菜单列表
func (m *RoleController) GetMenuList(c *gin.Context) {
	req := new(systemReq.RoleGetMenuListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return RoleLogic.GetMenuList(c, req)
	})
}

// GetApiList 获取接口列表
func (m *RoleController) GetApiList(c *gin.Context) {
	req := new(systemReq.RoleGetApiListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return RoleLogic.GetApiList(c, req)
	})
}

// UpdateMenus 更新菜单
func (m *RoleController) UpdateMenus(c *gin.Context) {
	req := new(systemReq.RoleUpdateMenusReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return RoleLogic.UpdateMenus(c, req)
	})
}

// UpdateApis 更新接口
func (m *RoleController) UpdateApis(c *gin.Context) {
	req := new(systemReq.RoleUpdateApisReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return RoleLogic.UpdateApis(c, req)
	})
}
