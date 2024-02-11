package system

import "github.com/eryajf/xirang/controller"

type RouterGroup struct {
	ApiRouter
	BaseRouter
	GroupRouter
	MenuRouter
	OperationLogRouter
	RoleRouter
	UserRouter
}

var (
	ApiController          = controller.ControllerGroupApp.SystemControllerGroup.ApiController
	BaseController         = controller.ControllerGroupApp.SystemControllerGroup.BaseController
	GroupController        = controller.ControllerGroupApp.SystemControllerGroup.GroupController
	MenuController         = controller.ControllerGroupApp.SystemControllerGroup.MenuController
	OperationLogController = controller.ControllerGroupApp.SystemControllerGroup.OperationLogController
	RoleController         = controller.ControllerGroupApp.SystemControllerGroup.RoleController
	UserController         = controller.ControllerGroupApp.SystemControllerGroup.UserController
)
