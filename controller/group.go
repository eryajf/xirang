package controller

import (
	"net/http"
	"strconv"

	"github.com/eryajf/xirang/model"
	"github.com/eryajf/xirang/service"

	"github.com/gin-gonic/gin"
)

// AddGroup 用户组注册接口
func AddGroup(c *gin.Context) {
	var groupObj model.Group
	if err := c.ShouldBind(&groupObj); err == nil {
		groupObj.CreateBy = CurrentUser(c).NickName
		groupObj.UpdateBy = CurrentUser(c).NickName
		res := service.AddGroup(groupObj)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetGroup 获取用户组
func GetGroup(c *gin.Context) {
	var groupObj model.Group
	if err := c.ShouldBind(&groupObj); err == nil {
		groupObj.GroupID, _ = strconv.Atoi(c.Param("group_id"))
		res := service.GetGroup(groupObj)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UpdataGroup 修改用户组信息
func UpdataGroup(c *gin.Context) {
	var groupObj model.Group
	if err := c.ShouldBind(&groupObj); err == nil {
		groupObj.UpdateBy = CurrentUser(c).NickName
		res := service.UpdataGroup(groupObj)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// DeleteGroup 删除用户组
func DeleteGroup(c *gin.Context) {
	var groupObj model.Group
	if err := c.ShouldBind(&groupObj); err == nil {
		groupObj.GroupID, _ = strconv.Atoi(c.Param("group_id"))
		res := service.DeleteGroup(groupObj)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// ListGroup 获取所有用户组
func ListGroup(c *gin.Context) {
	var groupObj model.QueryDataList
	if err := c.ShouldBind(&groupObj); err == nil {
		res := service.ListGroup(groupObj)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// ChangeGroupState 修改用户组状态
func ChangeGroupState(c *gin.Context) {
	groupObj := model.ChangeGroupState{}
	if err := c.ShouldBind(&groupObj); err == nil {
		groupObj.GroupID, _ = strconv.Atoi(c.Param("group_id"))
		groupObj.State, _ = strconv.Atoi(c.Param("type"))
		res := service.ChangeGroupState(groupObj)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// AddUToG 添加用户到分组
func AddUToG(c *gin.Context) {
	var dataObj model.AddUToG
	if err := c.ShouldBind(&dataObj); err == nil {
		dataObj.GroupID, _ = strconv.Atoi(c.Param("group_id"))
		// dataObj.Users = c.Bind("users")
		res := service.AddUToG(dataObj)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// GetGroupUser 获取组内用户
func GetGroupUser(c *gin.Context) {
	var groupObj model.Group
	if err := c.ShouldBind(&groupObj); err == nil {
		objid, _ := strconv.Atoi(c.Param("group_id"))
		res := service.GetGroupUser(objid)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// GetGroupNoUser 获取不在组内的用户
func GetGroupNoUser(c *gin.Context) {
	var groupObj model.Group
	if err := c.ShouldBind(&groupObj); err == nil {
		objid, _ := strconv.Atoi(c.Param("group_id"))
		res := service.GetGroupNoUser(objid)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
