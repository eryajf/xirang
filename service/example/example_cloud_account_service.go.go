package example

import (
	"errors"
	"fmt"
	"strings"

	"github.com/eryajf/xirang/model/example"
	"github.com/eryajf/xirang/model/example/request"
	"github.com/eryajf/xirang/public/common"
	"github.com/eryajf/xirang/public/tools"
	"gorm.io/gorm"
)

type CloudAccountService struct{}

// List 获取数据列表
func (s CloudAccountService) List(req *request.CloudAccountListReq) ([]*example.CloudAccount, error) {
	var list []*example.CloudAccount
	db := common.DB.Model(&example.CloudAccount{}).Order("created_at DESC")

	cloudName := strings.TrimSpace(req.CloudName)
	if cloudName != "" {
		db = db.Where("cloud_name LIKE ?", fmt.Sprintf("%%%s%%", cloudName))
	}
	cloudType := strings.TrimSpace(req.CloudType)
	if cloudType != "" {
		db = db.Where("cloud_type LIKE ?", fmt.Sprintf("%%%s%%", cloudType))
	}
	secretId := strings.TrimSpace(req.SecretId)
	if secretId != "" {
		db = db.Where("secret_id LIKE ?", fmt.Sprintf("%%%s%%", secretId))
	}

	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Find(&list).Error
	return list, err
}

// List 获取数据列表
func (s CloudAccountService) ListAll() (list []*example.CloudAccount, err error) {
	err = common.DB.Model(&example.CloudAccount{}).Order("created_at DESC").Find(&list).Error

	return list, err
}

// ListCount 获取数据总数
func (s CloudAccountService) ListCount(req *request.CloudAccountListReq) (int64, error) {
	var count int64
	db := common.DB.Model(&example.CloudAccount{}).Order("created_at DESC")

	cloudName := strings.TrimSpace(req.CloudName)
	if cloudName != "" {
		db = db.Where("cloud_name LIKE ?", fmt.Sprintf("%%%s%%", cloudName))
	}
	cloudType := strings.TrimSpace(req.CloudType)
	if cloudType != "" {
		db = db.Where("cloud_type LIKE ?", fmt.Sprintf("%%%s%%", cloudType))
	}
	secretId := strings.TrimSpace(req.SecretId)
	if secretId != "" {
		db = db.Where("secret_id LIKE ?", fmt.Sprintf("%%%s%%", secretId))
	}
	err := db.Count(&count).Error
	return count, err
}

// Count 获取数据总数
func (s CloudAccountService) Count() (int64, error) {
	var count int64
	db := common.DB.Model(&example.CloudAccount{}).Order("created_at DESC")

	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s CloudAccountService) Add(dataObj *example.CloudAccount) error {
	return common.DB.Create(dataObj).Error
}

// Update 更新资源
func (s CloudAccountService) Update(dataObj *example.CloudAccount) error {
	return common.DB.Model(dataObj).Where("id = ?", dataObj.ID).Updates(dataObj).Error
}

// Find 获取单个资源
func (s CloudAccountService) Find(filter map[string]interface{}, data *example.CloudAccount) error {
	return common.DB.Where(filter).First(&data).Error
}

// Exist 判断资源是否存在
func (s CloudAccountService) Exist(filter map[string]interface{}) bool {
	var dataObj example.CloudAccount
	err := common.DB.Debug().Order("created_at DESC").Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Delete 批量删除
func (s CloudAccountService) Delete(ids []uint) error {
	return common.DB.Debug().Where("id IN (?)", ids).Delete(&example.CloudAccount{}).Error
}

// GetCloudAccountsById 根据接口ID获取接口列表
func (s CloudAccountService) GetCloudAccountsById(cloudAccountIds []uint) ([]*example.CloudAccount, error) {
	var cloudAccounts []*example.CloudAccount
	err := common.DB.Where("id IN (?)", cloudAccountIds).Find(&cloudAccounts).Error
	return cloudAccounts, err
}
