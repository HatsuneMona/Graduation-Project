/**
 * @Author HatsuneMona
 * @Date  2021-02-05 17:05
 * @Description 初始化 MySQL database
 **/
package MySQL

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type connInfo struct {
	addr     string
	user     string
	pw       string
	database string
}

var DB *gorm.DB

func init() {
	c := connInfo{
		addr:     "172.20.0.2",
		user:     "hospital",
		pw:       "62QppvxZjCGm7c9jdUyQ",
		database: "hospital",
	}
	connString := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		c.user, c.pw, c.addr, c.database)
	var err error
	DB, err = gorm.Open("mysql", connString)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}
