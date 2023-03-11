package models

import (
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `vaild:"matches(^1[3-9]{1}\\d(9)$)"` //正则表达式  校验是否为电话号码
	Email         string `vaild:"email"`
	Avater        string //头像
	Identity      string
	ClientIP      string
	ClintPort     uint64
	Salt          string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LogOutTime    time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
