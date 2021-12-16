package controller

import (
	"eryajfgo/model"
	"eryajfgo/public/tools"
	"eryajfgo/service"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AddUser 用户注册接口
func AddUser(c *gin.Context) {
	var userObj model.User
	if err := c.ShouldBind(&userObj); err == nil {
		res := service.AddUser(userObj)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetUser 获取用户
func GetUser(c *gin.Context) {
	userObj := model.UserPage{}
	if err := c.ShouldBind(&userObj); err == nil {
		userObj.UserID, _ = strconv.Atoi(c.Param("user_id"))
		res := service.GetUser(userObj)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UpdataUser 修改用户信息
func UpdataUser(c *gin.Context) {
	var userObj model.User
	if err := c.ShouldBind(&userObj); err == nil {
		userObj.UserID, _ = strconv.Atoi(c.Param("user_id"))
		res := service.UpdataUser(userObj)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	var userObj model.User
	if err := c.ShouldBind(&userObj); err == nil {
		userObj.UserID, _ = strconv.Atoi(c.Param("user_id"))
		res := service.DeleteUser(userObj)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// ListUser 获取所有用户
func ListUser(c *gin.Context) {
	var userObj model.QueryDataList
	if err := c.ShouldBind(&userObj); err == nil {
		res := service.ListUser(userObj)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// ChangeUserState 修改用户状态
func ChangeUserState(c *gin.Context) {
	var userObj model.ChangeUserState
	if err := c.ShouldBind(&userObj); err == nil {
		userObj.UserID, _ = strconv.Atoi(c.Param("user_id"))
		userObj.State, _ = strconv.Atoi(c.Param("type"))
		res := service.ChangeUserState(userObj)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var userObj model.UserLogin
	if err := c.ShouldBind(&userObj); err == nil {
		if user, err := service.UserLogin(userObj); err == nil {
			s := sessions.Default(c)
			s.Clear()
			s.Set("user_id", user.UserID)
			s.Save()
			c.JSON(200, userObj)
		} else {
			c.JSON(200, err)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, tools.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
