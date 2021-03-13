// @Title       Models
// @Description 用户数据库操作函数测试
// @Author      HatsuneMona
// @CreateTime  2021/3/13 18:09

package Models

import (
	"service/Models"
	"testing"
)

func Test_GetHandler(t *testing.T) {
	t.Run("UserID读取用户测试", func(t *testing.T) {
		user := new(Models.User)
		err := user.GetUserByID(19999)
		if err != nil {
			t.Fatal(err)
		} else {
			t.Logf("读取user信息成功！读取到的信息如下：%v", user)
			if user.Phone != "19912345678" {
				t.Fatal("预计读取到的user信息错误")
			}
		}
	})

	t.Run("UserPhone读取用户测试", func(t *testing.T) {
		user := new(Models.User)
		err := user.GetUserByPhone("19912345678")
		if err != nil {
			t.Fatal(err)
		} else {
			t.Logf("读取user信息成功！读取到的信息如下：%v", user)
			if user.Nickname != "常驻测试" {
				t.Fatal("预计读取到的user信息错误")
			}
		}
	})
}

func Test_UpdateHandler(t *testing.T) {

	//TODO 暂未编写相关测试。

}
