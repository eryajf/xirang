package example

import "github.com/eryajf/xirang/logic"

type ControllerGroup struct {
	CloudAccountController
}

var (
	CloudAccountLogic = logic.LogicGroupApp.ExampleLogicGroup.CloudAccountLogic
)
