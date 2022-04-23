package model

import (
	"gorm.io/gorm"
	"tmscube-go/core"
)

type TeacherRedRemindModel struct {
	UserId int `gorm:"primaryKey;autoIncrement:false"`
	PhoneCourseNum int `gorm:"column:phone_course_num"`
	PhoneCourseMessage string `gorm:"column:phone_course_message"`
	PhoneCourseMessageTime int `gorm:"column:phone_course_message_time"`

	PhoneGrabNum int `gorm:"column:phone_grab_num"`
	PhoneGrabMessage string `gorm:"column:phone_grab_message"`
	PhoneGrabMessageTime int `gorm:"column:phone_grab_message_time"`

	PhoneEduNum int `gorm:"column:phone_edu_num"`
	PhoneEduMessage string `gorm:"column:phone_edu_message"`
	PhoneEduMessageTime int `gorm:"column:phone_edu_message_time"`

	PhoneOtherNum int `gorm:"column:phone_other_num"`
	PhoneOtherMessage string `gorm:"column:phone_other_message"`
	PhoneOtherMessageTime int `gorm:"column:phone_other_message_time"`

	PhoneInvkNum int `gorm:"column:phone_invk_num"`
	PhoneInvkMessage string `gorm:"column:phone_invk_message"`
	PhoneInvkMessageTime int `gorm:"column:phone_invk_message_time"`
}

func (TeacherRedRemindModel) TableName() string {
	return "teacher_red_remind"
}

func (TeacherRedRemindModel) DB() *gorm.DB {
	return core.MessageCenterDB
}