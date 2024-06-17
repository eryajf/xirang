package system

import "github.com/eryajf/xirang/logic"

type ControllerGroup struct {
	ApiController
	DeptController
	MenuController
	RoleController
	UserController
	OperationLogController
	BaseController
}

var (
	ApiLogic          = logic.LogicGroupApp.SystemLogicGroup.ApiLogic
	BaseLogic         = logic.LogicGroupApp.SystemLogicGroup.BaseLogic
	DeptLogic         = logic.LogicGroupApp.SystemLogicGroup.DeptLogic
	MenuLogic         = logic.LogicGroupApp.SystemLogicGroup.MenuLogic
	OperationLogLogic = logic.LogicGroupApp.SystemLogicGroup.OperationLogLogic
	RoleLogic         = logic.LogicGroupApp.SystemLogicGroup.RoleLogic
	UserLogic         = logic.LogicGroupApp.SystemLogicGroup.UserLogic
)
