package example

import (
	"fmt"

	"github.com/eryajf/xirang/model/example"
	exampleReq "github.com/eryajf/xirang/model/example/request"
	exampleRsp "github.com/eryajf/xirang/model/example/response"
	"github.com/eryajf/xirang/public/tools"

	"github.com/gin-gonic/gin"
)

type DomainLogic struct{}

// List 数据列表
func (l DomainLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*exampleReq.DomainListReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取数据列表
	domains, err := DomainService.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据列表失败: %s", err.Error()))
	}

	rets := make([]example.Domain, 0)
	for _, domain := range domains {
		rets = append(rets, *domain)
	}
	count, err := DomainService.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据总数失败"))
	}

	return exampleRsp.DomainListRsp{
		Total:   count,
		Domains: rets,
	}, nil
}
