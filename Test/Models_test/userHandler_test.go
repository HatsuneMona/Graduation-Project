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

	t.Run("测试Add/Delete新用户", func(t *testing.T) {
		tUser := Models.User{
			ID:                0,
			Phone:             "13366669999",
			Nickname:          "testUpdate",
			Password:          "testUpdateHandler",
			BindPatientIDMeta: "",
			BindPatientID:     nil,
		}
		err := tUser.AddNewUser()
		if err != nil {
			t.Fatal(err)
		}
		t.Run("删除用户", func(t *testing.T) {
			err := tUser.Delete()
			if err != nil {
				t.Fatal(err)
			}
		})
	})

	tUser := Models.User{ID: 19999}
	_ = tUser.GetUserByID()

	t.Run("测试添加关联病人", func(t *testing.T) {
		t.Run("重复添加", func(t *testing.T) {
			err := tUser.AddPatient(123456)
			if err != nil {
				t.Logf(err.Error())
			} else {
				t.Fatalf("关联病人成功，%v", tUser)
			}
		})
		t.Run("正常添加", func(t *testing.T) {
			err := tUser.AddPatient(321654)
			if err != nil {
				t.Fatal(err)
			} else {
				t.Logf("关联病人成功，%v", tUser)
				_ = tUser.DeletePatient(321654)
			}
		})
	})

	t.Run("测试删除关联病人", func(t *testing.T) {
		err := tUser.DeletePatient(123456)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("删除病人后信息如下：%v", tUser)
		_ = tUser.AddPatient(321654)
	})
}
