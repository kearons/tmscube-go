package worker

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/streadway/amqp"
	"os"
	"strings"
	"tmscube-go/common/model"
	"tmscube-go/common/service"
	"tmscube-go/constant"
)

type MessageSingleReciever struct {
	QueueN        string
	RoutingK      string
	PrefetchCount int
}

func (msr MessageSingleReciever) QueueName() string {
	return msr.QueueN
}

func (msr MessageSingleReciever) RoutingKey() string {
	return msr.RoutingK
}

func (msr MessageSingleReciever) GetPrefetchCount() int {
	if msr.PrefetchCount > 1 {
		return msr.PrefetchCount
	}
	return 1
}

func (msr MessageSingleReciever) OnError(err error) {
	fmt.Println("system error: ", err)
}

func (msr MessageSingleReciever) OnReceive(b []byte) bool {
	fmt.Println(string(b), "接收到数据了！")
	return true
}

var (
	sT   service.TemplateService
	sM   service.MessageService
	sR   service.RedRemindService
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func init() {
	extra.RegisterFuzzyDecoders()
}

func (msr MessageSingleReciever) OnReceiveBatch(msgs []amqp.Delivery) bool {

	msgStr := "["
	for i := 0; i < len(msgs); i++ {
		msgStr += string(msgs[i].Body) + ","
	}
	msgStr = strings.TrimRight(msgStr, ",")
	msgStr += "]"

	//fmt.Println(msgStr)

	var (
		msgsFormatted []model.SingleMsg
		odds          []model.MsgModel
		evens         []model.MsgModel
	)

	err := json.Unmarshal([]byte(msgStr), &msgsFormatted)
	if err != nil {
		fmt.Println(err)
		return true
	}

	for _, v := range msgsFormatted {

		template, err := sT.CompileTemplate(v.TemplateKey, v.Keywords)
		if err != nil {
			fmt.Println(err)
			os.Exit(123)
			continue
		}

		m := model.MsgModel{
			MessageModel: model.MessageModel{
				UserId:      v.UserId,
				TemplateKey: v.TemplateKey,
				Content:     template.CompiledContent,
				ExtraData:   string(v.Extra),
			},
			SingleMsg: &v,
			Template:  template,
		}

		if (v.UserId & 1) == 0 {
			evens = append(evens, m)
		} else {
			odds = append(odds, m)
		}
	}

	//插入消息
	sM.CreateMessageBatch("even", &evens)
	sM.CreateMessageBatch("odd", &odds)

	odds = append(odds, evens...)

	//更新红点
	sR.UpdateTeacherRedRemindBatch(&odds)

	//发送push
	sendNotify(&odds)

	return true
}

type PushBody struct {
	BusinessID     uint64                 `json:"businessId"`
	SrcType        int                    `json:"srcType"` //调用方类型 0:APP;1:CRM;2:MMS;3:TMS;4:MINI-ADMIN;5:MMS-ADMIN
	UserID         uint32                 `json:"userId"`
	ClientType     int                    `json:"clientType"` //客户端类型 0:学生端;1:家长端;2:老师端;3:小猪爱画画
	Title          string                 `json:"title"`
	Content        string                 `json:"content"`
	PayloadVO      map[string]interface{} `json:"payloadVO"`
	DelayType      int                    `json:"delayType"` //消息推送延迟类型,0:及时推送;1:延时推送
	ExpirationTime int                    `json:"expirationTime"`
	NoticeType     int                    `json:"noticeType"` //通知类型 0:通知栏;1:透传
}

func sendNotify(m *[]model.MsgModel) {
	var pushBodies []PushBody

	for _, v := range *m {
		if v.SingleMsg.Push.Send != 1 {
			continue
		}
		pushBody := PushBody{
			SrcType:    1,
			UserID:     v.MessageModel.UserId,
			ClientType: 2,
			Title:      v.Template.Title,
			Content:    v.MessageModel.Content,
			DelayType:  0,
			NoticeType: 0,
		}

		if len(v.SingleMsg.Push.Title) > 0 {
			pushBody.Title = v.SingleMsg.Push.Title
		}

		if len(v.SingleMsg.Push.Content) > 0 {
			pushBody.Content = v.SingleMsg.Push.Content
		}

		if v.SingleMsg.Push.BusinessId > 0 {
			pushBody.BusinessID = v.SingleMsg.Push.BusinessId
		}

		pushBody.PayloadVO = make(map[string]interface{})

		defaultPushType := constant.DefaultPushType[v.Template.MessageType]
		pushBody.PayloadVO["type"] = defaultPushType
		for k, z := range v.SingleMsg.Push.Payload {
			pushBody.PayloadVO[k] = z
		}

		pushBodies = append(pushBodies, pushBody)
		jso, _ := json.Marshal(pushBody)
		fmt.Println(string(jso))

		//特殊的push消息
		if sp, ok := constant.SpecialList[v.MessageModel.TemplateKey]; ok {
			pushBody.NoticeType = 1 //透传
			pushBody.PayloadVO["type"] = sp["type"]
			pushBody.PayloadVO["showType"] = sp["showType"]
			pushBody.PayloadVO["msgId"] = fmt.Sprintf("%d", v.MessageModel.ID)
			pushBody.PayloadVO["msgType"] = v.Template.MessageType
			pushBody.PayloadVO["popupTitle"] = v.Template.Title
			pushBody.PayloadVO["popupContent"] = v.MessageModel.Content

			pushBodies = append(pushBodies, pushBody)

			jso, _ := json.Marshal(pushBody)
			fmt.Println(string(jso))
		}
		fmt.Println(pushBody)
		os.Exit(1)
	}
	//todo send kafka
}
