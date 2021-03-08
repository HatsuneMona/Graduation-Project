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

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "wuyu:MIDSUMMERfish0@/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}
