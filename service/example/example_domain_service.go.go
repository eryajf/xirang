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
	"gorm.io/gorm/clause"
)

type DomainService struct{}

// UpSerts 存在则更新，不存在则创建
func (s DomainService) UpSerts(dataObj []example.Domain) error {
	return common.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "domain_id"}},
		UpdateAll: true,
	}).Create(dataObj).Error
}

func getDomainListDb(req *request.DomainListReq) *gorm.DB {
	db := common.DB.Model(&example.Domain{}).Order("created_at DESC")
	domainId := strings.TrimSpace(req.DomainID)
	if domainId != "" {
		db = db.Where("domain_id LIKE ?", fmt.Sprintf("%%%s%%", domainId))
	}
	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	remark := strings.TrimSpace(req.Remark)
	if remark != "" {
		db = db.Where("remark LIKE ?", fmt.Sprintf("%%%s%%", remark))
	}
	return db
}

// List 获取数据列表
func (s DomainService) List(req *request.DomainListReq) ([]*example.Domain, error) {
	var list []*example.Domain
	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := getDomainListDb(req).Preload("CloudAccount").Offset(pageReq.PageNum).Limit(pageReq.PageSize).Find(&list).Error
	return list, err
}

// List 获取数据列表
func (s DomainService) ListAll() (list []*example.Domain, err error) {
	err = common.DB.Model(&example.Domain{}).Preload("CloudAccount").Order("created_at DESC").Find(&list).Error
	return list, err
}

// ListCount 获取数据总数
func (s DomainService) ListCount(req *request.DomainListReq) (int64, error) {
	var count int64
	err := getDomainListDb(req).Count(&count).Error
	return count, err
}

// Count 获取数据总数
func (s DomainService) Count() (int64, error) {
	var count int64
	db := common.DB.Model(&example.Domain{}).Order("created_at DESC")

	err := db.Count(&count).Error
	return count, err
}

// Add 添加资源
func (s DomainService) Add(dataObj *example.Domain) error {
	return common.DB.Create(dataObj).Error
}

// Update 更新资源
func (s DomainService) Update(dataObj *example.Domain) error {
	return common.DB.Model(dataObj).Where("id = ?", dataObj.ID).Updates(dataObj).Error
}

// Find 获取单个资源
func (s DomainService) Find(filter map[string]interface{}, data *example.Domain) error {
	return common.DB.Where(filter).First(&data).Error
}

// Exist 判断资源是否存在
func (s DomainService) Exist(filter map[string]interface{}) bool {
	var dataObj example.Domain
	err := common.DB.Debug().Order("created_at DESC").Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Delete 批量删除
func (s DomainService) Delete(ids []uint) error {
	return common.DB.Debug().Where("id IN (?)", ids).Delete(&example.Domain{}).Error
}

// GetDomainsById 根据接口ID获取接口列表
func (s DomainService) GetDomainsById(domainIds []uint) ([]*example.Domain, error) {
	var domains []*example.Domain
	err := common.DB.Where("id IN (?)", domainIds).Find(&domains).Error
	return domains, err
}
