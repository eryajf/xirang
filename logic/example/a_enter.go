package example

import "github.com/eryajf/xirang/service"

type LogicGroup struct {
	CloudAccountLogic
}

// 初始化 service
var (
	CloudAccountService = service.ServiceGroupApp.ExampleServiceGroup.CloudAccountService
)
