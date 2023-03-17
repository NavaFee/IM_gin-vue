package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	OwnerId  uint //谁的关系信息
	TargetId uint //对应的是谁  /群ID
	Type     int  //对应的类型  1 好友  2 群组  3
	Desc     string
}
