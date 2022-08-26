package controller

import (
	"github.com/eryajf/xirang/logic"
	"github.com/eryajf/xirang/model/request"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

// Dashboard 系统首页展示数据
func (m *BaseController) Dashboard(c *gin.Context) {
	req := new(request.BaseDashboardReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Base.Dashboard(c, req)
	})
}

// GetPasswd 生成加密密码
func (m *BaseController) GetPasswd(c *gin.Context) {
	req := new(request.GetPasswdReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Base.GetPasswd(c, req)
	})
}
