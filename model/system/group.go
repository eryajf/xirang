package system

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	GroupName string   `gorm:"type:varchar(128);comment:'分组名称'" json:"groupName"`
	Remark    string   `gorm:"type:varchar(128);comment:'分组中文说明'" json:"remark"`
	Creator   string   `gorm:"type:varchar(20);comment:'创建人'" json:"creator"`
	Users     []*User  `gorm:"many2many:group_users" json:"users"`
	ParentId  uint     `gorm:"default:0;comment:'父组编号(编号为0时表示根组)'" json:"parentId"`
	Source    string   `gorm:"type:varchar(20);comment:'来源：dingTalk、weCom、ldap、platform'" json:"source"`
	Children  []*Group `gorm:"-" json:"children"`
}
