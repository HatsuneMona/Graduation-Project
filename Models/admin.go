package Models

import (
	"errors"
	"fmt"
	"service/Databases/DBPool"
)

type Admin struct {
	AdminID           int    `json:"admin_id" gorm:"primary_key;column:admin_id;type:int(11) auto_increment;comment:'管理员id'"`                                                                      // 管理员id
	AdminUsername     string `json:"admin_username" gorm:"column:admin_username;type:char(20);not null;unique_index:hospital_admin_admin_username_uindex;comment:'管理员用户名'"`                        // 管理员用户名
	AdminPassword     string `json:"admin_password" gorm:"column:admin_password;type:char(30);not null;comment:'管理员密码'"`                                                                           // 管理员密码
	AdminName         string `json:"admin_name" gorm:"column:admin_name;type:varchar(10);not null;comment:'管理员名字'"`                                                                                // 管理员名字
	AdminPhonenum     string `json:"admin_phoneNum" gorm:"column:admin_phoneNum;type:char(15);not null;comment:'管理员电话'"`                                                                           // 管理员电话
	AdminPermissionID int    `json:"admin_permission_id" gorm:"column:admin_permission_id;type:int(11);not null;index:hospital_admin_hospital_admin_permissions_permission_id_fk;comment:'管理员权限'"` // 管理员权限
	AdminHandlers
}

// TableName returns the table name of the Admin model
func (admin *Admin) TableName() string {
	return "hospital_admin"
}

type AdminHandlers interface {
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
	UpdatePhonenum(newPhoneNum string) error
	UpdatePermissionID(newPermissionID int) error
}

func (admin *Admin) GetAdminByID(id int) error {
	admin.AdminID = id
	res := DBPool.DB.First(&admin)
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

func (admin *Admin) GetAdminByUsername(username string) error {
	res := DBPool.DB.Where("admin_username = ?", username).First(&admin)
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

func (admin *Admin) UpdatePassword(newPassword string) error {
	if admin.AdminPassword == newPassword {
		return errors.New("新密码与旧密码相同，pass")
	}
	result := DBPool.DB.Model(&admin).Update("admin_password", newPassword)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (admin *Admin) UpdateName(newName string) error {
	if admin.AdminName == newName {
		return errors.New("新名字与旧名字相同，pass")
	}
	result := DBPool.DB.Model(&admin).Update("admin_name", newName)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (admin *Admin) UpdatePhonenum(newPhoneNum string) error {
	if admin.AdminPhonenum == newPhoneNum {
		return errors.New("新手机号与旧手机号相同，pass")
	}
	result := DBPool.DB.Model(&admin).Update("admin_phonenum", newPhoneNum)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (admin *Admin) UpdatePermissionID(newPermissionID int) error {
	if admin.AdminPermissionID == newPermissionID {
		return errors.New("新权限与旧权限相同，pass")
	}
	result := DBPool.DB.Model(&admin).Update("admin_permission_id", newPermissionID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
