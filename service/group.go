package service

import (
	"errors"
	"xirang/model"
	"xirang/public/tools"

	"gorm.io/gorm"
)

// AddGroup 添加分组
func AddGroup(groupObj model.Group) tools.Response {
	err := model.DB.Table(groupObj.TableName()).Where("group_name = ?", groupObj.GroupName).First(&groupObj).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return tools.HasError(tools.AlreadyExistErr, "分组已经注册", nil)
	}
	if err := model.DB.Table(groupObj.TableName()).Create(&groupObj).Error; err != nil {
		return tools.HasError(tools.CodeDBError, tools.CreatedFail, err)
	}
	return tools.Response{
		Code: 200,
		Msg:  tools.CreatedSuccess,
	}
}

// GetGroup 查询单个分组
func GetGroup(groupObj model.Group) tools.Response {
	if err := model.DB.Table(groupObj.TableName()).First(&groupObj, groupObj.GroupID).Error; err != nil {
		return tools.HasError(tools.NotExistErr, tools.NotFound, err)
	}
	return tools.Response{
		Code: 200,
		Data: groupObj,
		Msg:  tools.GetSuccess,
	}
}

// UpdataGroup 更新分组信息
func UpdataGroup(groupObj model.Group) tools.Response {
	if err := model.DB.Table(groupObj.TableName()).Where("group_id = ?", groupObj.GroupID).Save(&groupObj).Error; err != nil {
		return tools.HasError(tools.CodeDBError, tools.UpdatedFail, err)
	}
	return tools.Response{
		Code: 200,
		Data: groupObj,
		Msg:  tools.UpdatedSuccess,
	}
}

// DeleteGroup 删除分组
func DeleteGroup(groupObj model.Group) tools.Response {
	if err := model.DB.Table(groupObj.TableName()).First(&groupObj, groupObj.GroupID).Error; err != nil {
		return tools.HasError(tools.NotExistErr, tools.NotFound, err)
	}
	if err := model.DB.Table(groupObj.TableName()).Where("group_id = ?", groupObj.GroupID).Delete(&groupObj).Error; err != nil {
		return tools.HasError(tools.CodeDBError, tools.DeletedFail, err)
	}
	return tools.Response{
		Code: 200,
		Msg:  tools.DeletedSuccess,
	}
}

// ListGroup 分组列表
func ListGroup(groupObj model.QueryDataList) tools.Response {
	var groups []model.Group
	var total int64 = 0
	if groupObj.Limit == 0 {
		groupObj.Limit = 3
	}

	if err := model.DB.Table("groups").Where("deleted_at is null").Count(&total).Error; err != nil {
		return tools.HasError(tools.CodeDBError, tools.NotFound, err)
	}

	if err := model.DB.Table("groups").Limit(groupObj.Limit).Offset(groupObj.Start).
		Where("group_name LIKE ?", "%"+groupObj.Query+"%").Find(&groups).Error; err != nil {
		return tools.HasError(tools.CodeDBError, tools.NotFound, err)
	}
	return tools.BuildListResponse(200, groups, uint(total))
}

// ChangeGroupState 更新状态
func ChangeGroupState(groupObj model.ChangeGroupState) tools.Response {
	var group model.Group
	if err := model.DB.Table("groups").First(&group, groupObj.GroupID).Error; err != nil {
		return tools.HasError(tools.NotExistErr, "分组不存在", err)
	}
	if err := model.DB.Table("groups").Where("group_id = ?", groupObj.GroupID).Update("state", groupObj.State).Error; err != nil {
		return tools.HasError(tools.CodeDBError, tools.UpdatedFail, err)
	}
	return tools.Response{
		Code: 200,
		Msg:  tools.UpdatedSuccess,
	}
}

// AddUToG 添加用户到分组
func AddUToG(dataObj model.AddUToG) tools.Response {
	for _, v := range dataObj.Users {
		user, _ := model.GetUserbyUserName(v)
		relation := model.UserGroupRelation{
			GroupID: dataObj.GroupID,
			UserID:  user.UserID,
		}
		// 因为提供给前端的数据已经把添加过的用户过滤了，因此这里不再做判断
		// notRegister := model.DB.Where("group_id = ? and user_id = ?", relation.GroupID, relation.UserID).First(&relation).RecordNotFound()
		// if !notRegister {
		// 	return tools.Response{Code: 40001, Msg: "请勿重复添加"}
		// }
		if err := model.DB.Create(&relation).Error; err != nil {
			return tools.HasError(tools.CodeDBError, tools.CreatedFail, err)
		}
	}
	return tools.Response{
		Code: 200,
		Msg:  tools.CreatedSuccess,
	}
}

// GetGroupUser 查询单个分组内成员
func GetGroupUser(objid int) tools.Response {
	var groupusers []model.GroupUsers
	err := model.DB.Table("groups").Select("groups.group_name,users.username,users.nick_name").
		Joins("left join user_group_relations on user_group_relations.group_id=groups.group_id").
		Joins("left join users on user_group_relations.user_id=users.user_id and users.deleted_at is null").
		Where("groups.group_id = ? and nick_name is not null", objid).Find(&groupusers).Error
	if err != nil {
		return tools.HasError(tools.NotExistErr, tools.NotFound, err)
	}
	return tools.Response{
		Code: 200,
		Data: groupusers,
		Msg:  tools.GetSuccess,
	}
}

// GetGroupNoUser 获取不在组内的用户
func GetGroupNoUser(objid int) tools.Response {
	var groupnousers []model.GroupUsers
	err := model.DB.Table("users").Select("users.username,users.nick_name").
		Joins("left join user_group_relations on users.user_id = user_group_relations.user_id and  user_group_relations.group_id = ?", objid).
		Where("user_group_relations.user_id is null").Find(&groupnousers).Error
	if err != nil {
		return tools.HasError(tools.CodeDBError, tools.NotFound, err)
	}
	return tools.Response{
		Code: 200,
		Data: groupnousers,
		Msg:  tools.GetSuccess,
	}
}
