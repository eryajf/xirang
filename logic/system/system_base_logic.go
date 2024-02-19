package system

import (
	"fmt"

	systemReq "github.com/eryajf/xirang/model/system/request"
	systemRsp "github.com/eryajf/xirang/model/system/response"
	"github.com/eryajf/xirang/public/tools"

	"github.com/gin-gonic/gin"
)

type BaseLogic struct{}

// Dashboard 仪表盘
func (l BaseLogic) Dashboard(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	_, ok := req.(*systemReq.BaseDashboardReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	userCount, err := userService.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取用户总数失败"))
	}
	groupCount, err := groupService.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取分组总数失败"))
	}
	roleCount, err := roleService.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取角色总数失败"))
	}
	menuCount, err := menuService.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取菜单总数失败"))
	}
	apiCount, err := apiService.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取接口总数失败"))
	}
	logCount, err := operationLogService.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取日志总数失败"))
	}

	rst := make([]*systemRsp.DashboardList, 0)

	rst = append(rst,
		&systemRsp.DashboardList{
			DataType:  "user",
			DataName:  "用户",
			DataCount: userCount,
			Icon:      "user",
			Path:      "#/system/user",
		},
		&systemRsp.DashboardList{
			DataType:  "group",
			DataName:  "分组",
			DataCount: groupCount,
			Icon:      "peoples",
			Path:      "#/system/group",
		},
		&systemRsp.DashboardList{
			DataType:  "role",
			DataName:  "角色",
			DataCount: roleCount,
			Icon:      "eye-open",
			Path:      "#/system/role",
		},
		&systemRsp.DashboardList{
			DataType:  "menu",
			DataName:  "菜单",
			DataCount: menuCount,
			Icon:      "tree-table",
			Path:      "#/system/menu",
		},
		&systemRsp.DashboardList{
			DataType:  "api",
			DataName:  "接口",
			DataCount: apiCount,
			Icon:      "tree",
			Path:      "#/system/api",
		},
		&systemRsp.DashboardList{
			DataType:  "log",
			DataName:  "日志",
			DataCount: logCount,
			Icon:      "documentation",
			Path:      "#/system/log/operationLog",
		},
	)

	return rst, nil
}

// GetPasswd
func (l BaseLogic) GetPasswd(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.GetPasswdReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	return tools.NewGenPasswd(r.Passwd), nil
}
