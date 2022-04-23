package model

import (
	"gorm.io/gorm"
	"tmscube-go/core"
)

type MessageModel struct {
	ID           uint32 `gorm:"column:id;primaryKey"`
	UserId       uint32 `gorm:"column:user_id"`
	Content      string `gorm:"column:content"`
	ExtraData    string `gorm:"column:extra_data"`
	TemplateKey  string `gorm:"column:template_key"`
	ActionUserId int    `gorm:"column:action_user_id"`
	IsRead       uint8  `gorm:"column:action_user_id"`
	CreateTime   int    `gorm:"column:create_time;autoCreateTime"`
	//UpdateTime    time.Time `gorm:"-"`
	IsDelete      uint8  `gorm:"column:is_delete"`
	DeleteComment string `gorm:"column:delete_comment"`
}

func (m *MessageModel) TableName() string {
	return ""
}

func (m *MessageModel) DB() *gorm.DB {
	return core.MessageCenterDB
}
