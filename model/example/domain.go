package example

import "gorm.io/gorm"

// Domain 示例域名模型
type Domain struct {
	gorm.Model
	DomainID    string `gorm:"size:64;uniqueIndex;column:domain_id;comment:域名ID" json:"domainId,omitempty"`
	Name        string `gorm:"size:64;column:name;comment:域名" json:"name,omitempty"`
	Remark      string `gorm:"size:64;comment:说明" json:"remark,omitempty"`
	CreateTime  string `gorm:"size:64;column:update_time;comment:创建时间" json:"createTime,omitempty"`
	RecordCount uint64 `gorm:"column:record_count;comment:记录总数" json:"record_count,omitempty"`
}
