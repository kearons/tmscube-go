package service

import (
	"fmt"
	"reflect"
	"strings"
	"tmscube-go/common/model"
	"tmscube-go/common/widget"
)

type RedRemindService struct{}

type CountUserUnReadMessageVin struct {
	*BaseVin
	Include string `form:"include"`
}

func (s *RedRemindService) CountUserUnReadMessage(userId int, from string, tabs []string) int {
	var columns []string

	for _, v := range tabs {
		columns = append(columns, from+"_"+v+"_num")
	}

	return d.GetUserUnReadMessageCount(userId, columns)
}

type PeekUserUnReadMessageVin struct {
	*CountUserUnReadMessageVin
}
type PeekListVout struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Count       int    `json:"count"`
	LastTime    int    `json:"last_time"`
	LastMessage string `json:"last_message"`
}

func (s *RedRemindService) PeekUserUnReadMessage(userId int, from string, tabs []string) []PeekListVout {
	var (
		columns []string
		list    []PeekListVout
		tabZh   = map[string]string{"course": "课程", "grab": "抢单", "invk": "邀课", "edu": "教务", "other": "其他"}
	)

	for _, v := range tabs {
		columns = append(columns, from+"_"+v+"_num", from+"_"+v+"_message_time", from+"_"+v+"_message")
	}

	t, err := d.GetUserUnReadMessagePeek(userId, columns)

	rf := reflect.ValueOf(t)
	for _, v := range tabs {
		if err == nil {
			var (
				tmp      reflect.Value
				count    = 0
				lastTime = 0
			)
			if tmp = rf.FieldByName(widget.SnakeCase([]string{from, v, "num"})); !tmp.IsZero() {
				count = int(tmp.Int())
			}
			if tmp = rf.FieldByName(widget.SnakeCase([]string{from, v, "message", "time"})); !tmp.IsZero() {
				lastTime = int(tmp.Int())
			}
			lastMessage := rf.FieldByName(widget.SnakeCase([]string{from, v, "message"})).String()
			list = append(list, PeekListVout{Type: v, Name: tabZh[v], Count: count, LastTime: lastTime, LastMessage: lastMessage})
		} else {
			list = append(list, PeekListVout{Type: v, Name: tabZh[v], Count: 0, LastTime: 0, LastMessage: ""})
		}
	}

	return list
}

func (s *RedRemindService) UpdateTeacherRedRemindBatch(m *[]model.MsgModel) bool {
	var columns = make(map[string][]string, 2)

	for _, v := range *m {
		sqlVal := fmt.Sprintf("(%d, 1, \"%s\", %d)", v.MessageModel.UserId, v.MessageModel.Content, v.MessageModel.CreateTime)
		if j, ok := columns[v.Template.MessageType]; ok {
			columns[v.Template.MessageType] = append(j, sqlVal)
		} else {
			columns[v.Template.MessageType] = []string{sqlVal}
		}
	}

	for k, v := range columns {
		d.UpsertUserRedmindRaw(k, strings.Join(v, ","))
	}

	return true
}
