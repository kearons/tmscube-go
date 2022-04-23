package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tmscube-go/common/model"
	"tmscube-go/config"
	"tmscube-go/core"
	"tmscube-go/core/rabbitmq"
	"tmscube-go/router"
	"tmscube-go/worker"
)

func init() {
	config.Load()
}

func main1() {
	route := gin.Default()

	router.Load(route)

	route.Run()
}

func main3() {
	users := []model.MessageModel{{UserId: 1024520, Content: "test"}, {UserId: 1024520, Content: "test2"}}
	core.MessageCenterDB.Table("even_phone_tms_message").CreateInBatches(&users, 2)
	for _, v := range users {
		fmt.Println(v.ID)
	}
}

func main() {
	mq := rabbitmq.New()
	singleReciever := worker.MessageSingleReciever{QueueN: "queue_cube_test", RoutingK: "cube_key", PrefetchCount: 250}
	mq.RegisterReceiver(singleReciever)
	mq.Start()
}
