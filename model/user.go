package model

import (
	"golang.org/x/crypto/bcrypt"
)

// UserName 用户名
type UserName struct {
	Username string `json:"username" gorm:"type:varchar(64);comment:'用户名'"`
}

// PassWord 用户密码
type PassWord struct {
	Password string `json:"password" gorm:"type:varchar(128);comment:'用户密码'"`
}

// UserLogin User login structure
type UserLogin struct {
	UserName
	PassWord
}

// UserInfo 用户信息
type UserInfo struct {
	UserID    int    `json:"user_id"    gorm:"primaryKey;AUTO_INCREMENT;comment:'用户ID'"`
	JobNumber string `json:"job_number" gorm:"type:varchar(64);comment:'工号'"`
	NickName  string `json:"nick_name"  gorm:"min=2,max=30;comment:'昵称'"`
	Email     string `json:"email"      gorm:"min=2,max=25;comment:'邮箱'"`
	Phone     string `json:"phone"      gorm:"type:char(11);comment:'手机号'"`
	Gender    string `json:"gender"     gorm:"type:varchar(8);comment:'性别'"`
	State     int    `json:"state"      gorm:"default:0;comment:'状态'"`
	AvatarURL string `json:"avatar_url" gorm:"default:'https://tvax1.sinaimg.cn/large/008k1Yt0ly1gxde1sj29mj308c08cdhd.jpg';comment:'头像'"`
	RoleID    int    `json:"role_id"    gorm:"type:int(11);comment:'角色编码'"`
	DeptID    int    `json:"dept_id"    gorm:"type:int(11);comment:'部门编码'"`
}

// User 用户模型
type User struct {
	UserInfo
	UserLogin
	BaseModel
}

// UserPage 用户以及分组
type UserPage struct {
	User
	GroupName string `json:"group_name"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// GetUser 用ID获取用户
func GetUser(objid interface{}) (User, error) {
	var user User
	result := DB.Where("user_id = ?", objid).First(&user)
	return user, result.Error
}

// GetUserbyNickName 用名字获取用户
func GetUserbyNickName(objname string) (User, error) {
	var user User
	result := DB.Where("nick_name = ?", objname).First(&user)
	return user, result.Error
}

// GetUserbyUserName 通过用户名获取名字
func GetUserbyUserName(objname string) (User, error) {
	var user User
	result := DB.Where("username = ?", objname).First(&user)
	return user, result.Error
}

// ChangeUserState 修改用户状态
type ChangeUserState struct {
	UserID int `json:"user_id"`
	State  int `json:"state"`
}

// ChangeUserPassword 修改用户密码
type ChangeUserPassword struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// SetPassword 密码加密
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
