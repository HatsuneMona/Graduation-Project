/**
 * @Author HatsuneMona
 * @Date  2021-03-08 20:13
 * @Description //TODO
 **/
package Databases_test

import (
	"service/Databases"
	"testing"
)

func Test_MySQLConnection(t *testing.T) {
	t.Run("MySQL连接测试", func(t *testing.T) {
		if err := Databases.DB.DB().Ping(); err != nil {
			t.Fatalf("%v失败，Ping失败，错误信息：%v", t.Name(), err)
		}
	})
}
