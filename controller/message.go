package controller

import (
	"github.com/gin-gonic/gin"
	"tmscube-go/common/service"
)

func GetMessageList(c *gin.Context) {
	var (
		s service.MessageService
		p service.ListUserMessageVin
	)
	if err := c.ShouldBindQuery(&p); err != nil {
		Error(c, 500, err.Error())
		return
	}

	result := s.ListUserMessage()
	Success(c, result)
}
