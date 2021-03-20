package Models

import (
	"errors"
	"fmt"
	"service/Databases"
	"service/pkg/password"
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
	GetUserByID(id ...int) error
	//GetUserByNickname(username string) error
	GetUserByPhone(phone ...string) error
}

type userSetter interface {
	UpdatePassword(newPassword string) error
	UpdateNickName(newNickName string) error
	UpdatePhone(newPhone string) error
	AddNewUser() error
	Delete() error
}

//GetUserByID
//
//@Description	通过给定用户ID，来返回用户的所有信息。
//
//@Param
//				`id ...int` 需要查询的ID，可选参数。
//				若 `id ...int` 保持为空，则从调用者的结构体中获取要查询的用户ID。
//
//@Return
//				查询到的信息将写入调用者的结构体中，并返回。
//				error 查询正确则返回nil，否则返回Error
func (u *User) GetUserByID(id ...int) error {
	if len(id) > 0 {
		u.ID = id[0]
	}
	res := Databases.DB.Take(&u)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

//GetUserByPhone
//
//@Description	通过给定用户的电话号码，来返回用户的所有信息。
//
//@Param
//				`phone ...string` 需要查询的电话号码，可选参数。（优先）（若提供多个值，则仅查询第一个值）
//				若 `phone ...string`` 保持为空，则从调用者的结构体中获取要查询的用户电话号码。
//
//@Return
//				查询到的信息将写入调用者的结构体中，并返回。
//				error 查询正确则返回nil，否则返回Error
func (u *User) GetUserByPhone(phone ...string) error {
	if len(phone) > 1 {
		u.Phone = phone[0]
	} else if u.Phone == "" {
		return errors.New(fmt.Sprintf("非法手机号：%v", u.Phone))
	}
	if len(phone) != 11 {
		return errors.New(fmt.Sprintf("非法手机号：%v", u.Phone))
	}
	res := Databases.DB.Model(&u).Where("user_phone = ?", u.Phone).Take(&u)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

/*
UpdatePassword

@Description	更新User用户密码

@Param
				`newPassword` string 新密码，输入的密码应该是明文密码，而非加密过的密码。

@Return
				成功：error=nil，失败：返回错误信息。
*/
func (u *User) UpdatePassword(newPassword string) error {
	if same, _ := password.PasswordVerify(newPassword, u.Password); same {
		return errors.New("新旧密码相同，pass")
	}
	pwSHA := password.PasswordWithSaltGenToSHA(newPassword)
	result := Databases.DB.Model(&u).Update("user_password", pwSHA)
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

//AddNewUser
//
//@Description	添加新的用户
//
//@Param
//				nothing
//
//@Return
//				error
func (u *User) AddNewUser() error {
	if u.ID != 0 {
		return errors.New("禁止指定UserID")
	}
	u.Password = password.PasswordWithSaltGenToSHA(u.Password)
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
