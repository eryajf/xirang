package example

import "github.com/eryajf/xirang/controller"

type RouterGroup struct {
	CloudAccount
}

// 初始化 cmdb Controller
var (
	CloudAccountController = controller.ControllerGroupApp.ExampleControllerGroup.CloudAccountController
)
