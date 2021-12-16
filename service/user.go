package service

import (
	"errors"
	"eryajfgo/model"
	"eryajfgo/public/tools"

	"gorm.io/gorm"
)

// AddUser 新增用户
func AddUser(userObj model.User) tools.Response {
	err := model.DB.Table(userObj.TableName()).Where("username = ?", userObj.Username).First(&userObj).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return tools.HasError(tools.AlreadyExistErr, "用户名已经注册", nil)
	}
	// 加密密码
	if err := userObj.SetPassword(userObj.Password); err != nil {
		return tools.HasError(tools.CodeEncryptError, "密码加密失败", err)
	}
	// 创建用户
	if err := model.DB.Table(userObj.TableName()).Create(&userObj).Error; err != nil {
		return tools.HasError(tools.CodeDBError, tools.CreatedFail, err)
	}
	return tools.Response{
		Code: 200,
		Msg:  tools.CreatedSuccess,
	}
}

// GetUser 查询单个用户
func GetUser(userObj model.UserPage) tools.Response {
	if err := model.DB.Table(userObj.TableName()).First(&userObj, userObj.UserID).Error; err != nil {
		return tools.HasError(tools.CodeDBError, tools.NotFound, err)
	}
	return tools.Response{
		Code: 200,
		Data: userObj,
		Msg:  tools.GetSuccess,
	}
}

// UpdataUser 更新用户信息
func UpdataUser(userObj model.User) tools.Response {
	if err := model.DB.Table(userObj.TableName()).Where("user_id = ?", userObj.UserID).Save(&userObj).Error; err != nil {
		return tools.HasError(tools.CodeDBError, tools.UpdatedFail, err)
	}
	return tools.Response{
		Code: 200,
		Msg:  tools.UpdatedSuccess,
	}
}

// DeleteUser 删除用户
func DeleteUser(userObj model.User) tools.Response {
	if err := model.DB.Table(userObj.TableName()).First(&userObj, userObj.UserID).Error; err != nil {
		return tools.HasError(tools.CodeDBError, tools.DeletedFail, err)
	}

	if err := model.DB.Table(userObj.TableName()).Where("user_id = ?", userObj.UserID).Delete(&userObj).Error; err != nil {
		return tools.HasError(tools.CodeDBError, tools.DeletedFail, err)
	}
	return tools.Response{
		Code: 200,
		Msg:  tools.DeletedSuccess,
	}
}

// ListUser 用户列表
func ListUser(userObj model.QueryDataList) tools.Response {
	var users []model.UserPage
	var total int64 = 0
	if userObj.Limit == 0 {
		userObj.Limit = 10
	}
	if err := model.DB.Table("users").Where("deleted_at is null and username LIKE ? OR nick_name LIKE ?", "%"+userObj.Query+"%", "%"+userObj.Query+"%").Count(&total).Error; err != nil {
		return tools.HasError(tools.CodeDBError, tools.NotFound, err)
	}
	if err := model.DB.Table("users").Limit(userObj.Limit).Offset(userObj.Start).
		Where("username LIKE ? OR nick_name LIKE ?", "%"+userObj.Query+"%", "%"+userObj.Query+"%").
		Find(&users).Error; err != nil {
		return tools.HasError(tools.CodeDBError, tools.NotFound, err)
	}
	return tools.BuildListResponse(200, users, uint(total))
}

// ChangeUserState 更新状态
func ChangeUserState(userObj model.ChangeUserState) tools.Response {
	var user model.User
	if err := model.DB.Table(user.TableName()).First(&user, userObj.UserID).Error; err != nil {
		return tools.HasError(tools.NotExistErr, tools.NotFound, err)
	}
	if err := model.DB.Table(user.TableName()).Where("user_id = ?", userObj.UserID).Update("state", userObj.State).Error; err != nil {
		return tools.HasError(tools.CodeDBError, tools.UpdatedFail, err)
	}
	return tools.Response{
		Code: 200,
		Msg:  tools.UpdatedSuccess,
	}
}

// UserLogin 用户登录函数
func UserLogin(userObj model.UserLogin) (model.UserPage, *tools.Response) {
	var user model.UserPage
	if err := model.DB.Table(user.TableName()).Where("username = ?", userObj.Username).First(&user).Error; err != nil {
		return user, &tools.Response{Code: tools.NotExistErr, Msg: tools.NotFound}
	}
	if !user.CheckPassword(userObj.Password) {
		return user, &tools.Response{Code: tools.CodeParamErr, Msg: "密码错误"}
	}
	return user, nil
}
