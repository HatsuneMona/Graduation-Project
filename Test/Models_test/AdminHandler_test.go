/**
 * @Author HatsuneMona
 * @Date  2021-03-09 09:11
 * @Description 测试Models中的Admin
 **/
package Models_test

import (
	"service/Models"
	"testing"
)

func Test_AdminReader(t *testing.T) {
	t.Run("测试GetAdminByID", func(t *testing.T) {
		admin := new(Models.Admin)
		err := admin.GetAdminByID(9999)
		if err != nil {
			t.Fatalf("%v错误，错误信息：%v", t.Name(), err)
		}
		if admin.AdminUsername != "TestAdmin" {
			t.Fatalf("%v错误，错误信息：预期AdminName值不正确=%v", t.Name(), admin.AdminUsername)
		}
		t.Log(admin)
	})

	t.Run("测试GetAdminByUsername", func(t *testing.T) {
		admin := new(Models.Admin)
		err := admin.GetAdminByUsername("TestAdmin")
		if err != nil {
			t.Fatalf("%v错误，错误信息：%v", t.Name(), err)
		}
		if admin.AdminID != 9999 {
			t.Fatalf("%v错误，错误信息：预期AdminID值不正确=%v", t.Name(), admin.AdminID)
		}
		t.Log(admin)
	})
}

func Test_AdminSetter(t *testing.T) {
	//略
}
