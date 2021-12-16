package model

// Group 用户组模型
type Group struct {
	GroupID   int    `json:"group_id"   gorm:"primaryKey;AUTO_INCREMENT;comment:'用户组ID'"`
	GroupName string `json:"group_name" gorm:"min=2,max=30;comment:'用户组名'"`
	State     int    `json:"state"      gorm:"default:0;comment:'状态: 1 启用 0 禁用'"`
	Desc      string `json:"desc"       gorm:"type:varchar(128);comment:'备注'"`
	CreateBy  string `json:"create_by"  gorm:"type:varchar(128);comment:'谁创建'"`
	UpdateBy  string `json:"update_by"  gorm:"type:varchar(128);comment:'谁更新'"`
	BaseModel
}

// TableName 指定表名
func (Group) TableName() string {
	return "groups"
}

// GetGroup 用ID获取用户组
func GetGroup(ID interface{}) (Group, error) {
	var group Group
	result := DB.First(&group, ID)
	return group, result.Error
}

// ChangeGroupState 修改用户状态
type ChangeGroupState struct {
	GroupID int `json:"group_id"`
	State   int `json:"state"`
}

// AddUToG 添加用户到分组
type AddUToG struct {
	GroupID int      `json:"group_id"`
	Users   []string `json:"users"`
}

// GroupUsers 组内用户
type GroupUsers struct {
	GroupName string `json:"group_name"`
	Username  string `json:"username"`
	NickName  string `json:"nick_name"`
}
