package Models

import (
	"errors"
	"fmt"
	"service/Databases"
)

type User struct {
	ID       int    `json:"user_id" gorm:"primary_key;column:user_id;type:int(11) auto_increment;comment:'注册用户id'"`                                                                  // 注册用户id
	Phone    string `json:"user_phone" gorm:"column:user_phone;type:char(16);not null;index:public_user_user_nickname_user_phoneNum_index;comment:'注册用户手机号'"`                        // 注册用户手机号
	Nickname string `json:"user_nickname" gorm:"column:user_nickname;type:varchar(16);not null;index:public_user_user_nickname_user_phoneNum_index;comment:'注册用户昵称，可以重复，不作为登录名使用。'"` // 注册用户昵称，可以重复，不作为登录名使用。
	Password string `json:"user_password" gorm:"column:user_password;type:char(32);not null;comment:'注册用户密码'"`                                                                       // 注册用户密码
}

// TableName returns the table name of the User model
func (u *User) TableName() string {
	return "public_user"
}

type UserHandlers interface {
	userReader
	userSetter
}

type userReader interface {
	GetUserByID(id int) error
	//GetUserByNickname(username string) error
	GetUserByPhone(phone string) error
}

type userSetter interface {
	UpdatePassword(newPassword string) error
	UpdateNickName(newNickName string) error
	UpdatePhone(newPhone string) error
	AddNewUser() error
	Delete() error
}

func (u *User) GetUserByID(id int) error {
	u.ID = id
	res := Databases.DB.Take(&u)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

//func (u *User) GetUserByNickname(username string) error {
//	res := Databases.DB.Model(&u).Where("user_nickname = ?", username).Take(&u)
//	if res.Error != nil {
//		return res.Error
//	} else {
//		return nil
//	}
//}

func (u *User) GetUserByPhone(phone string) error {
	if len(phone) != 11 {
		return errors.New(fmt.Sprintf("非法手机号：%v", phone))
	}
	res := Databases.DB.Model(&u).Where("user_phone = ?", phone).Take(&u)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

func (u *User) UpdatePassword(newPassword string) error {
	if u.Password == newPassword {
		return errors.New("新旧密码相同，pass")
	}
	result := Databases.DB.Model(&u).Update("user_password", newPassword)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) UpdateNickName(newNickName string) error {
	if u.Nickname == newNickName {
		return errors.New("新旧用户名相同，pass")
	}
	result := Databases.DB.Model(&u).Update("user_nickname", newNickName)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) UpdatePhone(newPhoneNum string) error {
	if u.Phone == newPhoneNum {
		return errors.New("新旧手机号相同，pass")
	}
	result := Databases.DB.Model(&u).Update("user_phone", newPhoneNum)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) AddNewUser() error {
	if u.ID != 0 {
		return errors.New("禁止指定UserID")
	}
	result := Databases.DB.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) Delete() error {
	if u.ID == 0 {
		return errors.New("请提供正确的userID")
	}
	result := Databases.DB.Delete(&u)
	if result.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("删除失败，查无此ID（DeleteID=%v）", u.ID))
	}
	if result.Error != nil {
		return result.Error
	}
	//覆盖内存中原有的数据
	*u = User{
		ID:       -u.ID,
		Phone:    "(Deleted)",
		Nickname: "(Deleted)",
		Password: "(Deleted)",
	}
	return nil
}
