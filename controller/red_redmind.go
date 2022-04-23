package controller

import (
	"github.com/gin-gonic/gin"
	"strings"
	"tmscube-go/common/service"
)

func GetUnreadMessageCount(c *gin.Context) {
	var (
		count = 0
		s     service.RedRemindService
		p     service.CountUserUnReadMessageVin
	)
	if err := c.ShouldBindQuery(&p); err != nil {
		Error(c, 500, err.Error())
		return
	}
	tabs := []string{"course", "grab", "edu", "other"}

	if strings.Contains(p.Include, "invk") {
		tabs = []string{"course", "grab", "edu", "other", "invk"}
	}

	count = s.CountUserUnReadMessage(p.UserId, p.From, tabs)
	Success(c, map[string]int{"count": count})
}

func GetUnreadMessagePeek(c *gin.Context) {
	var (
		s service.RedRemindService
		p service.PeekUserUnReadMessageVin
	)
	if err := c.ShouldBindQuery(&p); err != nil {
		Error(c, 500, err.Error())
		return
	}
	tabs := []string{"course", "grab", "edu", "other"}

	if strings.Contains(p.Include, "invk") {
		tabs = []string{"course", "grab", "edu", "other", "invk"}
	}

	result := s.PeekUserUnReadMessage(p.UserId, p.From, tabs)
	Success(c, result)
}
