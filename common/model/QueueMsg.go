package model

import "encoding/json"

type SingleMsg struct {
	UserId      uint32            `json:"user_id"`
	TemplateKey string            `json:"template_key"`
	Keywords    map[string]string `json:"keywords,omitempty"`
	Extra       json.RawMessage   `json:"extra,omitempty"`
	Push        struct {
		Send       uint8                  `json:"send"`                  //是否需要发送, 若需要发送,则此处必需填写1
		Title      string                 `json:"title"`                 //push标题, 空值或不传则默认使用站内信的标题 (会截取32字)
		Content    string                 `json:"content"`               //push文案，若空或不传会使用站内信的文案（原因是因为push文案可能跟站内信的文案不一致）
		Payload    map[string]interface{} `json:"payload"`               //需要传递APP的参数, 具体详见push推送文档 (注意其中的type不用传)
		NoticeType uint8                  `json:"notice_type"`           //0 为通知栏 , 1为透传 (默认为0)
		BusinessId uint64                 `json:"business_id,string,omitempty"` //业务id, 不传则默认雪花id
		Expiration uint32                 `json:"expiration"`            //push推送截止时间 (不会在指定时间后推送) 默认72小时
	} `json:"push,omitempty"`
	NotificationBar struct {
		Content string `json:"content,omitempty"` //通知栏内容（当且仅当几种增强消息时可用）
	} `json:"notification_bar,omitempty"`
}

type MsgModel struct {
	MessageModel MessageModel      `gorm:"embedded"`
	SingleMsg    *SingleMsg        `gorm:"-"`
	Template     *TemplateDicModel `gorm:"-"`
}
