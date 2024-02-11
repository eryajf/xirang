package system

import "github.com/eryajf/xirang/logic"

type ControllerGroup struct {
	ApiController
	GroupController
	MenuController
	RoleController
	UserController
	OperationLogController
	BaseController
}

var (
	ApiLogic          = logic.LogicGroupApp.SystemLogicGroup.ApiLogic
	BaseLogic         = logic.LogicGroupApp.SystemLogicGroup.BaseLogic
	GroupLogic        = logic.LogicGroupApp.SystemLogicGroup.GroupLogic
	MenuLogic         = logic.LogicGroupApp.SystemLogicGroup.MenuLogic
	OperationLogLogic = logic.LogicGroupApp.SystemLogicGroup.OperationLogLogic
	RoleLogic         = logic.LogicGroupApp.SystemLogicGroup.RoleLogic
	UserLogic         = logic.LogicGroupApp.SystemLogicGroup.UserLogic
)
