package dao

import (
	"errors"
	"strings"
	"tmscube-go/common/model"
)

var (
	m        model.TeacherRedRemindModel
	mMessage model.MessageModel
)

type RedRemindDao struct{}

func (d *RedRemindDao) GetUserUnReadMessageCount(userId int, columns []string) int {
	num := 0

	m.DB().Select(strings.Join(columns, "+")).
		Table(m.TableName()).
		Where(&model.TeacherRedRemindModel{UserId: userId}).
		Scan(&num)

	return num
}

func (d *RedRemindDao) GetUserUnReadMessagePeek(userId int, columns []string) (model.TeacherRedRemindModel, error) {
	var (
		t   model.TeacherRedRemindModel
		err error
	)

	r := m.DB().Select(columns).
		Table(m.TableName()).
		Where(&model.TeacherRedRemindModel{UserId: userId}).
		First(&t)

	if r.RowsAffected < 1 {
		err = errors.New("no rows")
	}
	return t, err
}

func (d *RedRemindDao) UpsertUserRedmindRaw(t string, values string) bool {
	sql := "INSERT INTO `teacher_red_remind` (user_id, phone_" + t + "_num, phone_" + t + "_message, phone_" + t + "_message_time) VALUES " +
		values + " ON DUPLICATE KEY UPDATE phone_" + t + "_num = phone_" + t + "_num + VALUES(phone_" + t + "_num), phone_" + t + "_message = VALUES(phone_" + t + "_message), phone_" + t + "_message_time = VALUES(phone_" + t + "_message_time);"

	m.DB().Exec(sql)
	return true
}
