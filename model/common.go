package model

import (
	"eryajfgo/public"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

// BaseModel 基础字段
type BaseModel struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// UserGroupRelation 用户与分组映射
type UserGroupRelation struct {
	BaseModel
	ID      int `json:"id" gorm:"primaryKey;AUTO_INCREMENT;comment:'ID'"`
	GroupID int `json:"group_id" gorm:"type:int;unsigned;comment:'分组ID'"`
	UserID  int `json:"user_id" gorm:"type:int;unsigned;comment:'用户ID'"`
}

// QueryDataList 主机列表
type QueryDataList struct {
	Limit int    `form:"limit" json:"limit"`
	Start int    `form:"start" json:"start"`
	Query string `form:"query" json:"query"`
}

// DB 数据库链接单例
var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	// fmt.Println(viper.GetString(fmt.Sprintf("%s.%s", public.GetRunEvn(), "DB_NAME"))) // 输出 eryajfgo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", viper.GetString(fmt.Sprintf("%s.%s", public.GetRunEvn(), "DB_USER")),
		viper.GetString(fmt.Sprintf("%s.%s", public.GetRunEvn(), "DB_PASS")), viper.GetString(fmt.Sprintf("%s.%s", public.GetRunEvn(), "DB_HOST")),
		viper.GetString(fmt.Sprintf("%s.%s", public.GetRunEvn(), "DB_PORT")), viper.GetString(fmt.Sprintf("%s.%s", public.GetRunEvn(), "DB_NAME")))
	// 不输出数据库日志
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		msg := fmt.Errorf("conn db %s failed, err: %v", viper.GetString(fmt.Sprintf("%s.%s", public.GetRunEvn(), "DB_HOST")), err)
		panic(msg)
	}
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(20)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	DB = db
	// 根据结构体自动同步表
	migration()
}

// migration 执行数据自同步
func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Group{})
	DB.AutoMigrate(&UserGroupRelation{})
}
