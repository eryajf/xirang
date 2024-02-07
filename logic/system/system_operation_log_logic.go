package system

import (
	"fmt"

	"github.com/eryajf/xirang/model/system"
	systemReq "github.com/eryajf/xirang/model/system/request"
	systemRsp "github.com/eryajf/xirang/model/system/response"
	"github.com/eryajf/xirang/public/tools"

	"github.com/gin-gonic/gin"
)

type OperationLogLogic struct{}

// List 数据列表
func (l OperationLogLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.OperationLogListReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取数据列表
	logs, err := operationLogService.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取接口列表失败: %s", err.Error()))
	}

	rets := make([]system.OperationLog, 0)
	for _, log := range logs {
		rets = append(rets, *log)
	}
	count, err := operationLogService.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取接口总数失败"))
	}

	return systemRsp.LogListRsp{
		Total: count,
		Logs:  rets,
	}, nil

	// 获取
	// logs, err := operationLogService.List(&r)
	// if err != nil {
	// 	systemRsp.Fail(c, nil, "获取操作日志列表失败: "+err.Error())
	// 	return
	// }
	// return nil, nil
}

// Delete 删除数据
func (l OperationLogLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.OperationLogDeleteReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	for _, id := range r.OperationLogIds {
		filter := tools.H{"id": int(id)}
		if !operationLogService.Exist(filter) {
			return nil, tools.NewMySqlError(fmt.Errorf("该条记录不存在"))
		}
	}
	// 删除接口
	err := operationLogService.Delete(r.OperationLogIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除该条记录失败: %s", err.Error()))
	}
	return nil, nil
}
