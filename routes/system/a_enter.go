package system

import "github.com/eryajf/xirang/controller"

type RouterGroup struct {
	ApiRouter
	BaseRouter
	DeptRouter
	MenuRouter
	OperationLogRouter
	RoleRouter
	UserRouter
}

var (
	ApiController          = controller.ControllerGroupApp.SystemControllerGroup.ApiController
	BaseController         = controller.ControllerGroupApp.SystemControllerGroup.BaseController
	DeptController         = controller.ControllerGroupApp.SystemControllerGroup.DeptController
	MenuController         = controller.ControllerGroupApp.SystemControllerGroup.MenuController
	OperationLogController = controller.ControllerGroupApp.SystemControllerGroup.OperationLogController
	RoleController         = controller.ControllerGroupApp.SystemControllerGroup.RoleController
	UserController         = controller.ControllerGroupApp.SystemControllerGroup.UserController
)
