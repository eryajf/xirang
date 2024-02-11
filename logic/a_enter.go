package logic

import (
	"github.com/eryajf/xirang/logic/example"
	"github.com/eryajf/xirang/logic/system"
)

type LogicGroup struct {
	SystemLogicGroup  system.LogicGroup
	ExampleLogicGroup example.LogicGroup
}

var LogicGroupApp = new(LogicGroup)
