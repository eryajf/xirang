package example

import "github.com/eryajf/xirang/service"

type LogicGroup struct {
	DomainLogic
}

// 初始化 service
var (
	DomainService = service.ServiceGroupApp.ExampleServiceGroup.DomainService
)
