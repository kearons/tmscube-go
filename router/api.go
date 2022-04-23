package router

import (
	"github.com/gin-gonic/gin"
	"tmscube-go/controller"
)

func Load(router *gin.Engine) {
	//router.NoRoute()

	teacherApi := router.Group("/common/teacher-api")
	{
		teacherApi.GET("/get-unread-message-count", controller.GetUnreadMessageCount)
		teacherApi.GET("/get-unread-message-peek", controller.GetUnreadMessagePeek)
	}
}