package model

import (
	"gorm.io/gorm"
	"tmscube-go/core"
)

type TemplateDicModel struct {
	TemplateKey  string `gorm:"column:template_key;unique;"`
	TemplateType uint8  `gorm:"column:template_type"`
	MessageType  string `gorm:"column:message_type"`
	Title        string `gorm:"column:title"`
	ImageUrl     string `gorm:"column:image_url"`
	Content      string `gorm:"column:content"`
	Keywords     string `gorm:"column:keywords"`

	IsPhonePush   uint8  `gorm:"column:is_phone_push"`
	IsPhoneUrl    uint8  `gorm:"column:is_phone_url"`
	IsPhoneShare  uint8  `gorm:"column:is_phone_share"`
	PhoneUrl      string `gorm:"column:phone_h5_url"`
	PhoneLinkName string `gorm:"column:phone_link_name"`

	VerificationMode string    `gorm:"column:verification_mode"`
	CreateTime       int       `gorm:"column:create_time;autoCreateTime"`
	//UpdateTime       time.Time `gorm:"column:update_time"`

	Status uint8 `gorm:"column:status"` //模板状态，0-正常发送，10-禁用模板

	CompiledContent string `gorm:"-"` //替换变量后的模板
}

func (m *TemplateDicModel) TableName() string {
	return "tms_template_dic"
}

func (m *TemplateDicModel) DB() *gorm.DB {
	return core.MessageCenterDB
}
