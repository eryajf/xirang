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

type DeptService struct{}

// List 获取数据列表
func (s DeptService) List(req *systemReq.DeptListReq) ([]*system.Dept, error) {
	var list []*system.Dept
	db := common.DB.Model(&system.Dept{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	remark := strings.TrimSpace(req.Remark)
	if remark != "" {
		db = db.Where("remark LIKE ?", fmt.Sprintf("%%%s%%", remark))
	}

	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Find(&list).Error
	return list, err
}

// List 获取数据列表
func (s DeptService) ListTree(req *systemReq.DeptListReq) ([]*system.Dept, error) {
	var list []*system.Dept
	db := common.DB.Model(&system.Dept{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	remark := strings.TrimSpace(req.Remark)
	if remark != "" {
		db = db.Where("remark LIKE ?", fmt.Sprintf("%%%s%%", remark))
	}

	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Find(&list).Error
	return list, err
}

// List 获取数据列表
func (s DeptService) ListAll(req *systemReq.DeptListAllReq) ([]*system.Dept, error) {
	var list []*system.Dept
	db := common.DB.Model(&system.Dept{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	remark := strings.TrimSpace(req.Remark)
	if remark != "" {
		db = db.Where("remark LIKE ?", fmt.Sprintf("%%%s%%", remark))
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
func (s DeptService) Count() (int64, error) {
	var count int64
	err := common.DB.Model(&system.Dept{}).Count(&count).Error
	return count, err
}

// ListCount 获取附和条件的数据总数
func (s DeptService) ListCount(req *systemReq.DeptListReq) (int64, error) {
	var count int64
	db := common.DB.Model(&system.Dept{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	remark := strings.TrimSpace(req.Remark)
	if remark != "" {
		db = db.Where("remark LIKE ?", fmt.Sprintf("%%%s%%", remark))
	}
	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s DeptService) Add(data *system.Dept) error {
	return common.DB.Create(data).Error
}

// Update 更新资源
func (s DeptService) Update(dataObj *system.Dept) error {
	return common.DB.Model(dataObj).Where("id = ?", dataObj.ID).Updates(dataObj).Error
}

// Find 获取单个资源
func (s DeptService) Find(filter map[string]interface{}, data *system.Dept, args ...interface{}) error {
	return common.DB.Where(filter, args).Preload("Users").First(&data).Error
}

// Exist 判断资源是否存在
func (s DeptService) Exist(filter map[string]interface{}) bool {
	var dataObj system.Dept
	err := common.DB.Debug().Order("created_at DESC").Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Delete 批量删除
func (s DeptService) Delete(depts []*system.Dept) error {
	return common.DB.Debug().Select("Users").Unscoped().Delete(&depts).Error
}

// GetApisById 根据接口ID获取接口列表
func (s DeptService) GetDeptByIds(ids []uint) (datas []*system.Dept, err error) {
	err = common.DB.Where("id IN (?)", ids).Preload("Users").Find(&datas).Error
	return datas, err
}

// AddUserToDept 添加用户到分组
func (s DeptService) AddUserToDept(dept *system.Dept, users []system.User) error {
	return common.DB.Model(&dept).Association("Users").Append(users)
}

// RemoveUserFromDept 将用户从分组移除
func (s DeptService) RemoveUserFromDept(dept *system.Dept, users []system.User) error {
	return common.DB.Model(&dept).Association("Users").Delete(users)
}

// DeptIdsToDeptIds 将企业IM部门id转换为MySQL分组id
func (s DeptService) DeptIdsToDeptIds(ids []string) (deptIds []uint, err error) {
	var tempDepts []system.Dept
	err = common.DB.Model(&system.Dept{}).Where("source_dept_id IN (?)", ids).Find(&tempDepts).Error
	if err != nil {
		return nil, err
	}
	var tempDeptIds []uint
	for _, g := range tempDepts {
		tempDeptIds = append(tempDeptIds, g.ID)
	}
	return tempDeptIds, nil
}
