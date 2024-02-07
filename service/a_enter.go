package service

import (
	"github.com/eryajf/xirang/service/example"
	"github.com/eryajf/xirang/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
