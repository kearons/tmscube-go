package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tmscube-go/config"
)


func ConnectDB() *gorm.DB {

		//fmt.Println(getDsn(config.Mysql))
		conn, err := gorm.Open(mysql.Open(getMysqlDsn(config.GetConf().Mysql)), &gorm.Config{SkipDefaultTransaction: true})

		if err != nil {
			panic(err)
		}
		//error := db.Use(
		//	dbresolver.Register(dbresolver.Config{
		//		//Sources: []gorm.Dialector{},
		//		Replicas: []gorm.Dialector{mysql.Open("2"), mysql.Open("3")},
		//	}),
		//)

		//if error != nil {
		//	panic(error)
		//}
		return conn
}


func getMysqlDsn(c config.Mysql) string {
	//return fmt.Sprintf("mysql:host=%s;port=%d;dbname=%s", c.Host, c.Port, c.Db) //php style
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Username, c.Password, c.Host, c.Port, c.Db)
}
