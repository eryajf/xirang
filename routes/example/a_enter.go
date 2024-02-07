package example

import "github.com/eryajf/xirang/controller"

type RouterGroup struct {
	ExamleDomainRouter
}

// 初始化 cmdb Controller
var (
	DomainController = controller.ControllerGroupApp.ExampleControllerGroup.DomainController
)
