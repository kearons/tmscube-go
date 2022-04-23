package rabbitmq

import "github.com/streadway/amqp"

type Receiver interface {
	QueueName() string            // 获取接收者需要监听的队列
	RoutingKey() string           // 这个队列绑定的路由
	GetPrefetchCount() int        // 绑定
	OnError(error)                // 处理遇到的错误，当RabbitMQ对象发生了错误，他需要告诉接收者处理错误
	OnReceive([]byte) bool        // 处理收到的消息, 这里需要告知RabbitMQ对象消息是否处理成功
	OnReceiveBatch([]amqp.Delivery) bool // 处理收到的消息, 这里需要告知RabbitMQ对象消息是否处理成功
}
