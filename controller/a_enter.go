package controller

import (
	"net/http"

	"github.com/eryajf/xirang/controller/example"
	"github.com/eryajf/xirang/controller/system"
	"github.com/eryajf/xirang/public/tools"

	"github.com/gin-gonic/gin"
)

type ControllerGroup struct {
	SystemControllerGroup  system.ControllerGroup
	ExampleControllerGroup example.ControllerGroup
}

var ControllerGroupApp = new(ControllerGroup)

func Demo(c *gin.Context) {
	CodeDebug()
	c.JSON(http.StatusOK, tools.H{"code": 200, "msg": "ok", "data": "pong"})
}

func CodeDebug() {
}
