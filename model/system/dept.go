package system

import (
	"gorm.io/gorm"
)

type Dept struct {
	gorm.Model
	Name      string `gorm:"type:varchar(128);comment:'名称'" json:"name"`
	ParentId  uint   `gorm:"default:0;comment:'上一级编号(编号为0时表示根)'" json:"parentId"`
	Sort      uint   `gorm:"default:0;comment:'排序'" json:"sort"`
	Status    uint   `gorm:"type:tinyint(1);default:1;comment:'状态:1 启用, 0停用'" json:"status"` // 状态
	Creator   string `gorm:"type:varchar(20);comment:'创建人'" json:"creator"`
	Remark    string `gorm:"type:varchar(128);comment:'说明'" json:"remark"`
	Principal string `gorm:"type:varchar(128);comment:'负责人'" json:"principal"` // 负责人
	Email     string `gorm:"type:varchar(100);comment:'邮箱'" json:"email"`      // 邮箱
	Phone     string `gorm:"type:varchar(15);comment:'手机号'" json:"phone"`      // 手机号

	// Source   string  `gorm:"type:varchar(20);comment:'来源：dingTalk、weCom、ldap、platform'" json:"source"`
	// Children []*Dept `gorm:"-" json:"children"`
	// Users    []*User `gorm:"many2many:group_users" json:"users"`
}
