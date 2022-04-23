package service

import (
	"tmscube-go/common/dao"
)

var (
	d           dao.RedRemindDao
	templateDao dao.TemplateDao
	messageDao  dao.MessageDao
)

type BaseVin struct {
	From   string `form:"from" binding:"required,eq=phone`
	UserId int    `form:"user_id" binding:"required"`
}