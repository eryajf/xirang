package controller

import (
	"encoding/json"
	"fmt"

	"github.com/eryajf/xirang/model"
	"github.com/eryajf/xirang/public/captcha"
	"github.com/eryajf/xirang/public/tools"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v8"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, tools.Response{
		Code: 0,
		Msg:  "Pong",
	})
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) tools.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			return tools.HasError(tools.CodeParamErr, fmt.Sprintf("%s%s", e.Field, e.Tag), err)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return tools.HasError(tools.CodeParamErr, "JSON类型不匹配", err)
	}

	return tools.HasError(tools.CodeParamErr, "参数错误", err)
}

// CheckToken test
func CheckToken(c *gin.Context) {
	var a interface{}
	fmt.Println(c.ShouldBind(&a))
}

func GenerateCaptchaHandler(c *gin.Context) {
	id, b64s, err := captcha.DriverDigitFunc()
	tools.HasError(tools.NotExistErr, "验证码获取失败", err)
	tools.Custum(c, gin.H{
		"code": 200,
		"data": b64s,
		"id":   id,
		"msg":  "success",
	})
}
