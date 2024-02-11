package example

import (
	"gorm.io/gorm"
)

type CloudAccount struct {
	gorm.Model
	CloudName string `gorm:"size:24;column:cloud_name;comment:云账号名称" json:"cloudName"`
	CloudType string `gorm:"size:10;column:cloud_type;default:tx;comment:云厂商名称" json:"cloudType"` // 只能是 腾讯云/阿里云
	SecretId  string `gorm:"size:64;column:secret_id;comment:访问秘钥ID" json:"secretId"`
	SecretKey string `gorm:"size:255;column:secret_key;comment:访问秘钥key" json:"secretKey"`
	Remark    string `gorm:"size:128;comment:说明" json:"remark"`
}

func (m *CloudAccount) TableName() string {
	return "cloud_account"
}
