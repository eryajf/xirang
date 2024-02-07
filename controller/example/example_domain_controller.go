package example

import (
	exampleReq "github.com/eryajf/xirang/model/example/request"
	"github.com/eryajf/xirang/public/tools"
	"github.com/gin-gonic/gin"
)

type DomainController struct{}

// List 记录列表
func (m *DomainController) List(c *gin.Context) {
	req := new(exampleReq.DomainListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return DomainLogic.List(c, req)
	})
}
