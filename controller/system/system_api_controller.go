package system

import (
	systemReq "github.com/eryajf/xirang/model/system/request"
	"github.com/eryajf/xirang/public/tools"

	"github.com/gin-gonic/gin"
)

type ApiController struct{}

// List 记录列表
func (m *ApiController) List(c *gin.Context) {
	req := new(systemReq.ApiListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ApiLogic.List(c, req)
	})
}

// GetTree 接口树
func (m *ApiController) GetTree(c *gin.Context) {
	req := new(systemReq.ApiGetTreeReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ApiLogic.GetTree(c, req)
	})
}

// Add 新建记录
func (m *ApiController) Add(c *gin.Context) {
	req := new(systemReq.ApiAddReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ApiLogic.Add(c, req)
	})
}

// Update 更新记录
func (m *ApiController) Update(c *gin.Context) {
	req := new(systemReq.ApiUpdateReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ApiLogic.Update(c, req)
	})
}

// Delete 删除记录
func (m *ApiController) Delete(c *gin.Context) {
	req := new(systemReq.ApiDeleteReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return ApiLogic.Delete(c, req)
	})
}
