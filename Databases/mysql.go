//Package Databases 用来初始化MySQL和Redis
package Databases

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

type connInfo struct {
	addr     string
	port     int
	user     string
	pw       string
	database string
}

//DB 是MySQL的连接池
var DB *gorm.DB

func init() {
	c := connInfo{
		addr:     "172.20.0.2",
		port:     3306,
		user:     "hospital",
		pw:       "62QppvxZjCGm7c9jdUyQ",
		database: "hospital",
	}
	connString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		c.user, c.pw, c.addr, c.port, c.database)
	var err error
	DB, err = gorm.Open("mysql", connString)
	if err != nil {
		fmt.Printf("mysql connect error %v\n", err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v\n", DB.Error)
	}
}
