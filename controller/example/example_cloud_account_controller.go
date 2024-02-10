package example

import (
	exampleReq "github.com/eryajf/xirang/model/example/request"
	"github.com/eryajf/xirang/public/tools"
	"github.com/gin-gonic/gin"
)

type CloudAccountController struct{}

// List 记录列表
func (m *CloudAccountController) List(c *gin.Context) {
	req := new(exampleReq.CloudAccountListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return CloudAccountLogic.List(c, req)
	})
}

// Add 新建记录
func (m *CloudAccountController) Add(c *gin.Context) {
	req := new(exampleReq.CloudAccountAddReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return CloudAccountLogic.Add(c, req)
	})
}

// Update 更新记录
func (m *CloudAccountController) Update(c *gin.Context) {
	req := new(exampleReq.CloudAccountUpdateReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return CloudAccountLogic.Update(c, req)
	})
}

// Delete 删除记录
func (m *CloudAccountController) Delete(c *gin.Context) {
	req := new(exampleReq.CloudAccountDeleteReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return CloudAccountLogic.Delete(c, req)
	})
}
