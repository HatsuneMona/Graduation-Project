//Package Models 定义了数据库表的实例。
package Models

import (
	"errors"
	"fmt"
	"service/Databases"
	"service/Utils"
)

type Admin struct {
	ID           int    `json:"admin_id" gorm:"primary_key;column:admin_id;type:int(11) auto_increment;comment:'管理员id'"`                                                                      // 管理员id
	Username     string `json:"admin_username" gorm:"column:admin_username;type:char(20);not null;unique_index:hospital_admin_admin_username_uindex;comment:'管理员用户名'"`                        // 管理员用户名
	Password     string `json:"admin_password" gorm:"column:admin_password;type:char(30);not null;comment:'管理员密码'"`                                                                           // 管理员密码
	Name         string `json:"admin_name" gorm:"column:admin_name;type:varchar(10);not null;comment:'管理员名字'"`                                                                                // 管理员名字
	Phone        string `json:"admin_phone" gorm:"column:admin_phone;type:char(15);not null;comment:'管理员电话'"`                                                                                 // 管理员电话
	PermissionID int    `json:"admin_permission_id" gorm:"column:admin_permission_id;type:int(11);not null;index:hospital_admin_hospital_admin_permissions_permission_id_fk;comment:'管理员权限'"` // 管理员权限
	AdminHandlers
}

// TableName returns the table name of the Admin model
func (admin *Admin) TableName() string {
	return "hospital_admin"
}

/*
 * @description   关于Admin操作的相关接口，其中包括读取（查询）Admin信息接口，设置（更新、添加、删除）Admin信息接口）

 * @auther        HatsuneMona

 * @createTime    2021/3/10 16:45

 */
type AdminHandlers struct {
	adminReader
	adminSetter
}

type adminReader interface {
	GetAdminByID(id int) error
	GetAdminByUsername(username string) error
}

type adminSetter interface {
	UpdatePassword(newPassword string) error
	UpdateName(newName string) error
	UpdatePhone(newPhone string) error
	UpdatePermissionID(newPermissionID int) error
	AddNewAdmin() error
	Delete() error
}

/*
 * @description   通过给定adminID，从数据库中查询admin信息
 * @param         id int 需要查询的adminID
 * @return        查询成功：error=nil，失败则返回报错信息。
 * @auther        HatsuneMona
 * @createTime    2021/3/10 17:31
 */
func (admin *Admin) GetAdminByID(id int) error {
	admin.ID = id
	res := Databases.DB.Take(&admin)
	//if res.Value != 1 {
	//	log.Logger.Info(fmt.Sprintf("查询Admin错误，查询结果数量=%v", res.Value))
	//	return admin, errors.New(fmt.Sprintf("查询Admin错误，查询结果数量=%v", res.Value))
	//}
	if res.Error != nil {
		//log.Logger.Info(fmt.Sprintf("查询Admin错误，查询错误=%v", res.Error))
		return errors.New(fmt.Sprintf("查询Admin错误，查询错误=%v", res.Error))
	}
	return nil
}

/*
 * @description		通过给定adminID，从数据库中查询admin信息

 * @param
					`id` int 需要查询的adminID
 * @return
					查询成功：error=nil，失败则返回报错信息。

 * @auther			HatsuneMona

 * @createTime		2021/3/10 17:31
*/
func (admin *Admin) GetAdminByUsername(username string) error {
	res := Databases.DB.Where("admin_username = ?", username).Take(&admin)
	// if res.Value != 1 {
	//	log.Logger.Info(fmt.Sprintf("查询Admin错误，查询结果数量=%v", res.Value))
	//	return admin, errors.New(fmt.Sprintf("查询Admin错误，查询结果数量=%v", res.Value))
	//}
	if res.Error != nil {
		//log.Logger.Info(fmt.Sprintf("查询Admin错误，查询错误=%v", res.Error))
		return errors.New(fmt.Sprintf("查询Admin错误，查询错误=%v", res.Error))
	}
	return nil
}

/*
@Description	更新Admin用户密码

@Param
				`newPassword` string 新密码，输入的密码应该是明文密码，而非加密过的密码。

@Return
				成功：error=nil，失败：返回错误信息。

@Author        HatsuneMona

@CreateTime    2021/3/10 17:29
*/
func (admin *Admin) UpdatePassword(newPassword string) error {
	if same, _ := Utils.PasswordVerify(newPassword, admin.Password); same {
		return errors.New("新密码与旧密码相同，pass")
	}
	pwSHA := Utils.PasswordWithSaltGenToSHA(newPassword)
	result := Databases.DB.Model(&admin).Update("admin_password", pwSHA)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (admin *Admin) UpdateName(newName string) error {
	if admin.Name == newName {
		return errors.New("新名字与旧名字相同，pass")
	}
	result := Databases.DB.Model(&admin).Update("admin_name", newName)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (admin *Admin) UpdatePhone(newPhoneNum string) error {
	if admin.Phone == newPhoneNum {
		return errors.New("新手机号与旧手机号相同，pass")
	}
	result := Databases.DB.Model(&admin).Update("admin_phonenum", newPhoneNum)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (admin *Admin) UpdatePermissionID(newPermissionID int) error {
	if admin.PermissionID == newPermissionID {
		return errors.New("新权限与旧权限相同，pass")
	}
	result := Databases.DB.Model(&admin).Update("admin_permission_id", newPermissionID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/*
@Description	向数据库中添加一个新的Admin

@Param			调用该函数的admin，应该保持admin.ID==0，密码为明文密码

@Return
成功：error=nil，失败：返回错误信息。

@Author        HatsuneMona

@CreateTime    2021/3/18  09:44
*/
func (admin *Admin) AddNewAdmin() error {
	if admin.ID != 0 {
		return errors.New("禁止指定AdminID")
	}
	admin.Password = Utils.PasswordWithSaltGenToSHA(admin.Password)
	result := Databases.DB.Create(&admin)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (admin *Admin) Delete() error {
	if admin.ID == 0 {
		return errors.New("请提供正确的adminID")
	}
	result := Databases.DB.Delete(&admin)
	if result.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("删除失败，查无此ID（DeleteID=%v）", admin.ID))
	}
	if result.Error != nil {
		return result.Error
	}
	//原有数据是否需要手动覆盖？需要。
	*admin = Admin{
		ID:           -admin.ID,
		Username:     "(Deleted)",
		Password:     "(Deleted)",
		Name:         "(Deleted)",
		Phone:        "(Deleted)",
		PermissionID: -1, //系统保留权限值，空权限
	}
	return nil
}
