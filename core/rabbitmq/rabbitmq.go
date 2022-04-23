package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"sync"
	"time"
	"tmscube-go/config"
)

// RabbitMQ 用于管理和维护rabbitmq的对象
type RabbitMQ struct {
	wg sync.WaitGroup

	channel      *amqp.Channel
	exchangeName string // exchange的名称
	exchangeType string // exchange的类型
	receivers    []Receiver
}

// New 创建一个新的操作RabbitMQ的对象
func New() *RabbitMQ {
	// 这里可以根据自己的需要去定义
	return &RabbitMQ{
		exchangeName: "tms_task",
		exchangeType: "direct",
	}
}

// prepareExchange 准备rabbitmq的Exchange
func (mq *RabbitMQ) prepareExchange() error {
	// 申明Exchange
	err := mq.channel.ExchangeDeclare(
		mq.exchangeName, // exchange
		mq.exchangeType, // type
		true,            // durable
		false,           // autoDelete
		false,           // internal
		false,           // noWait
		nil,             // args
	)

	if nil != err {
		return err
	}

	return nil
}

// run 开始获取连接并初始化相关操作
func (mq *RabbitMQ) run() {
	//if !config.Global.RabbitMQ.Refresh() {
	//	log.Errorf("rabbit刷新连接失败，将要重连: %s", config.Global.RabbitMQ.URL)
	//	return
	//}

	// 获取新的channel对象
	//mq.channel = config.Global.RabbitMQ.Channel()
	conn := ConnectMQ()

	mq.channel, _ = conn.Channel()

	// 初始化Exchange
	mq.prepareExchange()

	for _, receiver := range mq.receivers {
		mq.wg.Add(1)
		go mq.listen(receiver) // 每个接收者单独启动一个goroutine用来初始化queue并接收消息
	}

	mq.wg.Wait()

	//log.Errorf("所有处理queue的任务都意外退出了")

	// 理论上mq.run()在程序的执行过程中是不会结束的
	// 一旦结束就说明所有的接收者都退出了，那么意味着程序与rabbitmq的连接断开
	// 那么则需要重新连接，这里尝试销毁当前连接
	//config.Global.RabbitMQ.Distory()
}

// Start 启动Rabbitmq的客户端
func (mq *RabbitMQ) Start() {
	for {
		mq.run()
		fmt.Println("连接断了")

		// 一旦连接断开，那么需要隔一段时间去重连
		// 这里最好有一个时间间隔
		time.Sleep(3 * time.Second)
	}
}

// RegisterReceiver 注册一个用于接收指定队列指定路由的数据接收者
func (mq *RabbitMQ) RegisterReceiver(receiver Receiver) {
	mq.receivers = append(mq.receivers, receiver)
}

// Listen 监听指定路由发来的消息
// 这里需要针对每一个接收者启动一个goroutine来执行listen
// 该方法负责从每一个接收者监听的队列中获取数据，并负责重试
func (mq *RabbitMQ) listen(receiver Receiver) {
	defer mq.wg.Done()

	// 这里获取每个接收者需要监听的队列和路由
	queueName := receiver.QueueName()
	routerKey := receiver.RoutingKey()

	// 申明Queue
	_, err := mq.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when usused
		false,     // exclusive(排他性队列)
		false,     // no-wait
		nil,       // arguments
	)
	if nil != err {
		// 当队列初始化失败的时候，需要告诉这个接收者相应的错误
		receiver.OnError(fmt.Errorf("初始化队列 %s 失败: %s", queueName, err.Error()))
	}

	// 将Queue绑定到Exchange上去
	err = mq.channel.QueueBind(
		queueName,       // queue name
		routerKey,       // routing key
		mq.exchangeName, // exchange
		false,           // no-wait
		nil,
	)
	if nil != err {
		receiver.OnError(fmt.Errorf("绑定队列 [%s - %s] 到交换机失败: %s", queueName, routerKey, err.Error()))
	}

	prefetchCount := receiver.GetPrefetchCount()

	// 获取消费通道
	mq.channel.Qos(prefetchCount, 0, false)
	delivery, err := mq.channel.Consume(
		queueName, // queue
		"tag-go",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if nil != err {
		receiver.OnError(fmt.Errorf("获取队列 %s 的消费通道失败: %s", queueName, err.Error()))
	}

	var (
		count    = 0
		msgBatch []amqp.Delivery
	)

	for {
		select {
		case msg := <-delivery:
			msgBatch = append(msgBatch, msg)
			count += 1
			if count >= prefetchCount {
				receiver.OnReceiveBatch(msgBatch)
				//for _, v := range msgBatch {
				//	v.Ack(false)
				//}
				msgBatch = append([]amqp.Delivery{})
				count = 0
			}
		case <-time.After(50 * time.Millisecond):
			if len(msgBatch) > 0 {
				receiver.OnReceiveBatch(msgBatch)
				//for _, v := range msgBatch {
				//	v.Ack(false)
				//}
				msgBatch = append([]amqp.Delivery{})
				count = 0
			}
		}
	}
}

func ConnectMQ() *amqp.Connection {
	conn, err := amqp.Dial(getMqDsn(config.Rabbitmq{
		Host:     "test.rabbitmq.miaokecloud.com",
		Port:     5672,
		Login:    "mqadmin",
		Password: "mqadmin",
	}))

	if err != nil {
		panic(err)
	}
	return conn
}

func getMqDsn(c config.Rabbitmq) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", c.Login, c.Password, c.Host, c.Port)
}
