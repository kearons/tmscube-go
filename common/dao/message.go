package dao

import (
	"time"
	"tmscube-go/common/model"
)

type MessageDao struct{}

func (d *MessageDao) InsertMessageBatch(parity string, models *[]model.MsgModel) {
	mMessage.DB().
		Table(d.getTableName(parity)).
		CreateInBatches(&models, 250)

	return
}

func (d *MessageDao) getTableName(parity string) string {
	return parity + "_phone_tms_message_" + time.Now().Format("200601")
}
