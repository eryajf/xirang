package system

import (
	systemReq "github.com/eryajf/xirang/model/system/request"
	"github.com/eryajf/xirang/public/tools"

	"github.com/gin-gonic/gin"
)

type MenuController struct{}

// GetTree 菜单树
func (m *MenuController) GetTree(c *gin.Context) {
	req := new(systemReq.MenuGetTreeReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return MenuLogic.GetTree(c, req)
	})
}

// GetUserMenuTreeByUserId 获取用户菜单树
func (m *MenuController) GetAccessTree(c *gin.Context) {
	req := new(systemReq.MenuGetAccessTreeReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return MenuLogic.GetAccessTree(c, req)
	})
}

// Add 新建
func (m *MenuController) Add(c *gin.Context) {
	req := new(systemReq.MenuAddReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return MenuLogic.Add(c, req)
	})
}

// Update 更新记录
func (m *MenuController) Update(c *gin.Context) {
	req := new(systemReq.MenuUpdateReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return MenuLogic.Update(c, req)
	})
}

// Delete 删除记录
func (m *MenuController) Delete(c *gin.Context) {
	req := new(systemReq.MenuDeleteReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return MenuLogic.Delete(c, req)
	})
}
