package model

import (
	"time"
)

type Device struct {
	Id       int       `gorm:"primary_key;auto_increment"`
	Status   uint8     `gorm:"type:tinyint(4);not null;comment:状态信息 0 为未签到 1为已经签到 2为已经签退"`
	Ip       string    `gorm:"comment:ip地址"`
	Mac      string    `gorm:"comment:mac地址"`
	StuName  string    `gorm:"comment:学生姓名"`
	StuNum   string    `gorm:"comment:学生学号"`
	CheckIn  time.Time `gorm:"comment:签到时间"`
	CheckOut time.Time `gorm:"comment:签退时间"`
}
