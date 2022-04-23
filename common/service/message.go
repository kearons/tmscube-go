package service

import (
	"tmscube-go/common/model"
)

type MessageService struct{}

type ListUserMessageVin struct {
	*BaseVin
	MessageId     int    `form:"message_id"`
	MessageType   string `form:"message_type"`
	MessageStatus uint8  `form:"message_status" binding:"oneof:0,1,2"`
	Limit         int    `form:"limit"`
	StartTime     int    `form:"start_time"`
	EndTime       int    `form:"end_time"`
}
type ListUserMessageVout struct {
	ID           uint32      `json:"id"`
	SID          string      `json:"sid"`
	Title        string      `json:"title"`
	Content      string      `json:"content"`
	TemplateType uint8       `json:"template_type"`
	TemplateKey  string      `json:"template_key"`
	PicUrl       string      `json:"pic_url"`
	JumpSup      uint16      `json:"jump_sup"`
	JumpUrl      string      `json:"jump_url"`
	IsShare      uint8       `json:"is_share"`
	BtnLabel     string      `json:"btn_label"`
	IsRead       uint8       `json:"is_read"`
	AuthMode     uint8       `json:"auth_mode"`
	CreateTime   uint32      `json:"create_time"`
	Extra        interface{} `json:"extra"`
}

func (s *MessageService) ListUserMessage() []ListUserMessageVout {
	var (
		list []ListUserMessageVout
	)

	//rows := d.

	return list
}

func (s *MessageService) CreateMessageBatch(parity string, m *[]model.MsgModel) {
	if len(*m) == 0 {
		return
	}
	messageDao.InsertMessageBatch(parity, m)
	return
}
