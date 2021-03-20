//Package Databases 用来初始化MySQL和Redis
package Databases

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
	"service/pkg/config"
)

//DB 是MySQL的连接池
var DB *gorm.DB

func init() {

	connString := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySQLConfig.User,
		config.MySQLConfig.Password,
		config.MySQLConfig.Addr,
		config.MySQLConfig.Port,
		config.MySQLConfig.Database)
	var err error
	DB, err = gorm.Open("mysql", connString)
	if err != nil {
		fmt.Printf("mysql connect error %v\n", err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v\n", DB.Error)
	}
}
