package example

import (
	"fmt"

	"github.com/eryajf/xirang/model/example"
	exampleReq "github.com/eryajf/xirang/model/example/request"
	"github.com/eryajf/xirang/model/example/response"
	"github.com/eryajf/xirang/public/tools"

	"github.com/gin-gonic/gin"
)

type CloudAccountLogic struct{}

// Add 添加数据
func (l CloudAccountLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*exampleReq.CloudAccountAddReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	cloudAccount := example.CloudAccount{
		CloudName: r.CloudName,
		CloudType: r.CloudType,
		SecretId:  r.SecretId,
		SecretKey: tools.NewGenPasswd(r.SecretKey),
		Remark:    r.Remark,
	}

	// 创建数据
	err := CloudAccountService.Add(&cloudAccount)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("创建数据失败: %s", err.Error()))
	}

	return nil, nil
}

// List 数据列表
func (l CloudAccountLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*exampleReq.CloudAccountListReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取数据列表
	cloudAccounts, err := CloudAccountService.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据列表失败: %s", err.Error()))
	}

	rets := make([]example.CloudAccount, 0)
	for _, cloudAccount := range cloudAccounts {
		cloudAccount.SecretKey = "******"
		rets = append(rets, *cloudAccount)
	}
	count, err := CloudAccountService.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取数据总数失败"))
	}

	return response.CloudAccountListRsp{
		Total:         count,
		CloudAccounts: rets,
	}, nil
}

// Update 更新数据
func (l CloudAccountLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*exampleReq.CloudAccountUpdateReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": int(r.ID)}
	if !CloudAccountService.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("数据不存在"))
	}

	oldData := new(example.CloudAccount)
	err := CloudAccountService.Find(filter, oldData)
	if err != nil {
		return nil, tools.NewMySqlError(err)
	}

	cloudAccount := example.CloudAccount{
		Model:     oldData.Model,
		CloudName: r.CloudName,
		CloudType: r.CloudType,
		SecretId:  r.SecretId,
		SecretKey: tools.NewGenPasswd(r.SecretKey),
		Remark:    r.Remark,
	}
	err = CloudAccountService.Update(&cloudAccount)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("更新数据失败: %s", err.Error()))
	}
	return nil, nil
}

// Delete 删除数据
func (l CloudAccountLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*exampleReq.CloudAccountDeleteReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	for _, id := range r.CloudAccountIds {
		filter := tools.H{"id": int(id)}
		if !CloudAccountService.Exist(filter) {
			return nil, tools.NewMySqlError(fmt.Errorf("数据不存在"))
		}
	}
	// 删除数据
	err := CloudAccountService.Delete(r.CloudAccountIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除数据失败: %s", err.Error()))
	}
	return nil, nil
}
