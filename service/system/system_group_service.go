package system

import (
	"errors"
	"fmt"
	"strings"

	"github.com/eryajf/xirang/model/system"
	systemReq "github.com/eryajf/xirang/model/system/request"
	"github.com/eryajf/xirang/public/common"
	"github.com/eryajf/xirang/public/tools"

	"gorm.io/gorm"
)

type GroupService struct{}

// List 获取数据列表
func (s GroupService) List(req *systemReq.GroupListReq) ([]*system.Group, error) {
	var list []*system.Group
	db := common.DB.Model(&system.Group{}).Order("created_at DESC")

	groupName := strings.TrimSpace(req.GroupName)
	if groupName != "" {
		db = db.Where("group_name LIKE ?", fmt.Sprintf("%%%s%%", groupName))
	}
	groupRemark := strings.TrimSpace(req.Remark)
	if groupRemark != "" {
		db = db.Where("remark LIKE ?", fmt.Sprintf("%%%s%%", groupRemark))
	}

	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Preload("Users").Find(&list).Error
	return list, err
}

// List 获取数据列表
func (s GroupService) ListTree(req *systemReq.GroupListReq) ([]*system.Group, error) {
	var list []*system.Group
	db := common.DB.Model(&system.Group{}).Order("created_at DESC")

	groupName := strings.TrimSpace(req.GroupName)
	if groupName != "" {
		db = db.Where("group_name LIKE ?", fmt.Sprintf("%%%s%%", groupName))
	}
	groupRemark := strings.TrimSpace(req.Remark)
	if groupRemark != "" {
		db = db.Where("remark LIKE ?", fmt.Sprintf("%%%s%%", groupRemark))
	}

	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Find(&list).Error
	return list, err
}

// List 获取数据列表
func (s GroupService) ListAll(req *systemReq.GroupListAllReq) ([]*system.Group, error) {
	var list []*system.Group
	db := common.DB.Model(&system.Group{}).Order("created_at DESC")

	groupName := strings.TrimSpace(req.GroupName)
	if groupName != "" {
		db = db.Where("group_name LIKE ?", fmt.Sprintf("%%%s%%", groupName))
	}
	groupRemark := strings.TrimSpace(req.Remark)
	if groupRemark != "" {
		db = db.Where("remark LIKE ?", fmt.Sprintf("%%%s%%", groupRemark))
	}
	groupType := strings.TrimSpace(req.GroupType)
	if groupType != "" {
		db = db.Where("group_type = ?", groupType)
	}
	source := strings.TrimSpace(req.Source)
	if source != "" {
		db = db.Where("source = ?", source)
	}
	sourceDeptId := strings.TrimSpace(req.SourceDeptId)
	if sourceDeptId != "" {
		db = db.Where("source_dept_id = ?", sourceDeptId)
	}
	sourceDeptParentId := strings.TrimSpace(req.SourceDeptParentId)
	if sourceDeptParentId != "" {
		db = db.Where("source_dept_parent_id = ?", sourceDeptParentId)
	}

	err := db.Find(&list).Error
	return list, err
}

// Count 获取数据总数
func (s GroupService) Count() (int64, error) {
	var count int64
	err := common.DB.Model(&system.Group{}).Count(&count).Error
	return count, err
}

// ListCount 获取附和条件的数据总数
func (s GroupService) ListCount(req *systemReq.GroupListReq) (int64, error) {
	var count int64
	db := common.DB.Model(&system.Group{}).Order("created_at DESC")

	groupName := strings.TrimSpace(req.GroupName)
	if groupName != "" {
		db = db.Where("group_name LIKE ?", fmt.Sprintf("%%%s%%", groupName))
	}
	groupRemark := strings.TrimSpace(req.Remark)
	if groupRemark != "" {
		db = db.Where("remark LIKE ?", fmt.Sprintf("%%%s%%", groupRemark))
	}
	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s GroupService) Add(data *system.Group) error {
	return common.DB.Create(data).Error
}

// Update 更新资源
func (s GroupService) Update(dataObj *system.Group) error {
	return common.DB.Model(dataObj).Where("id = ?", dataObj.ID).Updates(dataObj).Error
}

// Find 获取单个资源
func (s GroupService) Find(filter map[string]interface{}, data *system.Group, args ...interface{}) error {
	return common.DB.Where(filter, args).Preload("Users").First(&data).Error
}

// Exist 判断资源是否存在
func (s GroupService) Exist(filter map[string]interface{}) bool {
	var dataObj system.Group
	err := common.DB.Debug().Order("created_at DESC").Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Delete 批量删除
func (s GroupService) Delete(groups []*system.Group) error {
	return common.DB.Debug().Select("Users").Unscoped().Delete(&groups).Error
}

// GetApisById 根据接口ID获取接口列表
func (s GroupService) GetGroupByIds(ids []uint) (datas []*system.Group, err error) {
	err = common.DB.Where("id IN (?)", ids).Preload("Users").Find(&datas).Error
	return datas, err
}

// AddUserToGroup 添加用户到分组
func (s GroupService) AddUserToGroup(group *system.Group, users []system.User) error {
	return common.DB.Model(&group).Association("Users").Append(users)
}

// RemoveUserFromGroup 将用户从分组移除
func (s GroupService) RemoveUserFromGroup(group *system.Group, users []system.User) error {
	return common.DB.Model(&group).Association("Users").Delete(users)
}

// DeptIdsToGroupIds 将企业IM部门id转换为MySQL分组id
func (s GroupService) DeptIdsToGroupIds(ids []string) (groupIds []uint, err error) {
	var tempGroups []system.Group
	err = common.DB.Model(&system.Group{}).Where("source_dept_id IN (?)", ids).Find(&tempGroups).Error
	if err != nil {
		return nil, err
	}
	var tempGroupIds []uint
	for _, g := range tempGroups {
		tempGroupIds = append(tempGroupIds, g.ID)
	}
	return tempGroupIds, nil
}
