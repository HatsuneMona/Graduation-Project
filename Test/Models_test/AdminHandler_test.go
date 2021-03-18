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
		if admin.Username != "TestAdmin" {
			t.Fatalf("%v错误，错误信息：预期AdminName值不正确=%v", t.Name(), admin.Username)
		}
		t.Log(admin)
	})

	t.Run("测试GetAdminByID，无结果（错误结果测试）", func(t *testing.T) {
		admin := new(Models.Admin)
		err := admin.GetAdminByID(9900)
		if err != nil {
			t.Logf("%v得到原函数传来的，错误信息：%v", t.Name(), err)
		} else {
			t.Fatalf("未得到错误，测试未通过。")
		}
		t.Log(admin)
	})

	t.Run("测试GetAdminByUsername", func(t *testing.T) {
		admin := new(Models.Admin)
		err := admin.GetAdminByUsername("TestAdmin")
		if err != nil {
			t.Fatalf("%v错误，错误信息：%v", t.Name(), err)
		}
		if admin.ID != 9999 {
			t.Fatalf("%v错误，错误信息：预期AdminID值不正确=%v", t.Name(), admin.ID)
		}
		t.Log(admin)
	})

}

func Test_AdminSetter(t *testing.T) {
	newAdmin := Models.Admin{
		ID:           0,
		Username:     "TestAddAdmin",
		Password:     "TestAddAdmin",
		Name:         "添加测试",
		Phone:        "TestAddAdmin",
		PermissionID: -2,
	}

	t.Run("添加、删除用户测试（添加成功测试）", func(t *testing.T) {
		err := newAdmin.AddNewAdmin()
		if err != nil {
			t.Fatal(err)
		} else {
			t.Logf("添加成功后，newAdmin信息如下：%v", newAdmin)
		}
		err = newAdmin.Delete()
		if err != nil {
			t.Fatal(err)
		} else {
			t.Logf("删除成功后，newAdmin信息如下：%v", newAdmin)
		}
	})

	t.Run("删除失败", func(t *testing.T) {
		deleteErrAdmin := Models.Admin{
			ID:           9000,
			Username:     "nothing",
			Password:     "nothing",
			Name:         "nothing",
			Phone:        "nothing",
			PermissionID: -2,
		}
		err := deleteErrAdmin.Delete()
		if err == nil {
			t.Fatal("删除不存在的数据时未报错")
		} else {
			t.Logf("删除不存在的数据时成功报错，报错信息：%v", err)
		}
	})
}
