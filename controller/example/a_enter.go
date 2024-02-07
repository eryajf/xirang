package example

import "github.com/eryajf/xirang/logic"

type ControllerGroup struct {
	DomainController
}

var (
	DomainLogic = logic.LogicGroupApp.ExampleLogicGroup.DomainLogic
)
