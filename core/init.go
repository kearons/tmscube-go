package core

import (
	"gorm.io/gorm"
)

var (
	MessageCenterDB *gorm.DB
)

func init() {
	MessageCenterDB = ConnectDB()
}