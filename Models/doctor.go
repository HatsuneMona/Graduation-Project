package Models

import (
	"errors"
	"fmt"
	"service/Databases"
	"service/pkg/password"
)

type Doctor struct {
	ID         int    `json:"doctor_id" gorm:"primary_key;column:doctor_id;type:int(11) auto_increment;comment:'医生工号'"`                                                                            // 医生工号
	Name       string `json:"doctor_name" gorm:"column:doctor_name;type:varchar(8);not null;comment:'医生名字'"`                                                                                       // 医生名字
	Password   string `json:"doctor_password" gorm:"column:doctor_password;type:char(32);not null;comment:'医生登录密码'"`                                                                               // 医生登录密码
	Department int    `json:"doctor_department" gorm:"column:doctor_department;type:int(11);not null;index:hospital_doctor_hospital_department_department_id_fk;default:'0';comment:'医生所在科室（部门）'"` // 医生所在科室（部门）
	Class      int    `json:"doctor_class" gorm:"column:doctor_class;type:int(11);not null;default:'0';comment:'医生职位'"`                                                                            // 医生职位
	Intro      string `json:"doctor_intro" gorm:"column:doctor_intro;type:text;not null;comment:'医生简介'"`                                                                                           // 医生简介
	DoctorHandlers
}

// TableName returns the table name of the Doctor model
func (d *Doctor) TableName() string {
	return "hospital_doctor"
}

type DoctorHandlers struct {
	doctorReader
	DoctorSetter
}

type doctorReader interface {
	GetDoctorByID(id ...int) error
	GetDoctorByName(name ...string) error
}
type DoctorSetter interface {
	UpdatePassword(newPassword string) error
	UpdateName(newName string) error
	UpdateDepartment(newDepartment int) error
	UpdateClass(newClass int) error
	UpdateIntro(newIntro string) error
	AddNewDoctor() error
	Delete() error
}

//GetDoctorByID
//
//@Description	通过给定Doctor.ID来查询一名医生的信息
//
//@Param
//				`id ...int` 需要查询的ID，可选参数。（优先）（若提供多个值，则仅查询第一个值）
//				当该参数为空时，会从调用者结构体中寻找Doctor.ID并进行查询。
//
//@Return
//				查询到的信息将写入调用者的结构体中，并返回。
//				若出错，则返回error，当error!=nil则结果不可用。
func (d *Doctor) GetDoctorByID(id ...int) error {
	if len(id) > 1 {
		d.ID = id[0]
	} else if d.ID == 0 {
		return errors.New("待查询id错误，id = 0")
	}
	result := Databases.DB.Take(&d)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//GetDoctorByName
//
//@Description	通过给定Doctor.ID来查询一名医生的信息
//
//@Param
//				`id ...int` 需要查询的ID，可选参数。（优先查询）
//				当该参数为空时，会从调用者结构体中寻找Doctor.ID并进行查询。
//
//@Return
//				查询到的信息将写入调用者的结构体中，并返回。
//				若出错，则返回error，当error!=nil则结果不可用。
func (d *Doctor) GetDoctorByName(name ...string) error {
	if len(name) > 1 {
		d.Name = name[0]
	} else if d.Name == "" {
		return errors.New(fmt.Sprint("查询非法用户名，doctor.Name=", d.Name))
	}
	result := Databases.DB.Take(&d, "doctor_name = ?", d.Name)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//UpdatePassword 更新密码
//
//@Description	更新医生密码
//
//@Param
//				`newPassword string` 新的密码（名文），
//
//@Return
//				修改成功后，密文密码会更新到调用结构体Doctor中。
//				若修改失败，通过`error`返回错误信息。
func (d *Doctor) UpdatePassword(newPassword string) error {
	if same, _ := password.PasswordVerify(newPassword, d.Password); same {
		return errors.New("新旧密码相同，pass")
	}
	pwSHA := password.PasswordWithSaltGenToSHA(newPassword)
	result := Databases.DB.Model(&d).Update("user_password", pwSHA)
	if result.Error != nil {
		return result.Error
	}
	d.Password = pwSHA
	return nil
}

//UpdateName	更新名字
//
//@Description	更新Doctor的名字
//
//@Param
//				`newName string` 新的名字
//
//@Return
//				修改成功后，新的名字会更新到调用者的结构体中
//				若修改失败，通过`error`返回错误信息。
func (d *Doctor) UpdateName(newName string) error {
	if newName == d.Name {
		return errors.New("新旧值相同")
	}
	result := Databases.DB.Model(&d).Update("doctor_name", newName)
	if result.Error != nil {
		return result.Error
	}
	d.Name = newName
	return nil
}

//UpdateDepartment 更新医生所在的科室
//
//@Description	更新医生所在的科室
//
//@Param
//				`newDepartment int` 新的科室的科室ID
//
//@Return
//				修改陈宫后，新的科室ID将更新到调用者的结构体中
//				若修改失败，则通过`error`返回错误信息。
func (d *Doctor) UpdateDepartment(newDepartment int) error {
	if newDepartment == d.Department {
		return errors.New("新旧值相同")
	}
	result := Databases.DB.Model(&d).Update("doctor_department", newDepartment)
	if result != nil {
		return result.Error
	}
	d.Department = newDepartment
	return nil
}

func (d *Doctor) UpdateClass(newClass int) error {
	return nil
}

func (d *Doctor) UpdateIntro(newIntro string) error {
	return nil
}

func (d *Doctor) AddNewDoctor() error {
	return nil
}

func (d *Doctor) Delete() error {
	return nil
}
