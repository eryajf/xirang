package system

import (
	systemReq "github.com/eryajf/xirang/model/system/request"
	"github.com/eryajf/xirang/public/tools"

	"github.com/gin-gonic/gin"
)

type OperationLogController struct{}

// List 记录列表
func (m *OperationLogController) List(c *gin.Context) {
	req := new(systemReq.OperationLogListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return OperationLogLogic.List(c, req)
	})
}

// Delete 删除记录
func (m *OperationLogController) Delete(c *gin.Context) {
	req := new(systemReq.OperationLogDeleteReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return OperationLogLogic.Delete(c, req)
	})
}
