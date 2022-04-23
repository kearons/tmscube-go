package worker

import "fmt"

type MessageBatchReciever struct {
	QueueN  string
	RoutingK string
}

func (mbr MessageBatchReciever) QueueName() string {
	return mbr.QueueN
}

func (mbr MessageBatchReciever) RoutingKey() string {
	return mbr.RoutingK
}

func (mbr MessageBatchReciever) OnError(err error) {

}

func (mbr MessageBatchReciever) OnReceive(b []byte) bool {
	fmt.Println(string(b), "接收到数据了！")
	return true
}
