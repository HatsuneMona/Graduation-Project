package Models

import (
	"database/sql"
	"errors"
	"fmt"
	"service/Databases"
)

type Patient struct {
	ID           int    `json:"patient_id" gorm:"primary_key;column:patient_id;type:int(11) auto_increment;comment:'患者id'"`                                                                 // 患者id
	Name         string `json:"patient_name" gorm:"column:patient_name;type:varchar(10);not null;comment:'患者真实姓名'"`                                                                         // 患者真实姓名
	IdentityCard string `json:"patient_identity_card" gorm:"column:patient_identity_card;type:char(18);not null;unique_index:public_patient_patient_identity_card_uindex;comment:'患者身份证号'"` // 患者身份证号
	// TODO PatientPhoneNum 需要换成string not null。
	Phonenum sql.NullString `json:"patient_phoneNum" gorm:"column:patient_phoneNum;type:char(16);comment:'患者电话号'"` // 患者电话号
	PatientHandlers
}

// TableName returns the table name of the Patient model
func (p *Patient) TableName() string {
	return "public_patient"
}

type PatientHandlers struct {
	patientReader
	patientSetter
}

type patientReader interface {
	GetPatientByID(id ...int) error
	GetPatientByName(name ...string) error
	GetPatientByIdentityCard(identicityCard ...string) error
	GetPatientByPhone(phone ...string) error
}

type patientSetter interface {
	SetPatientPhone(newPhone string) error
	AddNewPatirnt() error
	Delete() error
	// SetPatientByName(name ...string) error
	// SetPatientByIdentityCard(identicityCard ...string) error
	// SetPatientByPhone(phone ...string) error
}

//GetPatientByID
//
//@Description	通过给定患者的ID，来返回患者的所有信息。
//
//@Param
//				`id ...int` 需要查询的ID，可选参数。
//				若 `id ...int` 保持为空，则从调用者的结构体中获取要查询的患者ID。
//
//@Return
//				查询到的信息将写入调用者的结构体中，并返回。
//				error 查询正确则返回nil，否则返回Error
func (p *Patient) GetPatientByID(id ...int) error {
	if len(id) > 0 {
		p.ID = id[0]
	}
	res := Databases.DB.Take(&p)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

//GetPatientByIdentityCard
//
//@Description	通过给定患者的身份证号，来返回用户的所有信息。
//
//@Param
//				`idCard ...string` 需要查询的ID，可选参数。
//				若 `idCard ...string` 保持为空，则从调用者的结构体中获取要查询的用户ID。
//
//@Return
//				查询到的信息将写入调用者的结构体中，并返回。
//				error 查询正确则返回nil，否则返回Error
func (p *Patient) GetPatientByIdentityCard(idCard ...string) error {
	if len(idCard) > 1 {
		p.IdentityCard = idCard[0]
	}
	res := Databases.DB.Model(&p).Where("patient_identity_card = ?", p.IdentityCard).Take(&p)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

//GetPatientByPhone
//
//@Description	通过给定患者的身份证号，来返回用户的所有信息。
//
//@Param
//				`phone ...string` 需要查询的ID，可选参数。
//				若 `phone ...string` 保持为空，则从调用者的结构体中获取要查询的用户ID。
//
//@Return
//				查询到的信息将写入调用者的结构体中，并返回。
//				error 查询正确则返回nil，否则返回Error
func (p *Patient) GetPatientByPhone(phone ...string) error {
	if len(phone) > 1 {
		p.IdentityCard = phone[0]
	}
	res := Databases.DB.Model(&p).Where("patient_phoneNum = ?", p.IdentityCard).Take(&p)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

//AddNewPatient
//
//@Description	添加新的患者
//
//@Param
//				添加的患者信息需要在调用者结构体中给出。
//
//@Return
//				error
func (p *Patient) AddNewPatirnt() error {
	if p.ID != 0 {
		return errors.New("禁止指定新PatientID")
	}
	result := Databases.DB.Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Delete
//
//@Description	删除该患者
//
//@Param
//				添加的患者信息需要在调用者结构体中给出。
//
//@Return
//				error
func (p *Patient) Delete() error {
	if p.ID == 0 {
		return errors.New("请提供正确的userID")
	}
	result := Databases.DB.Model(&p).Update("patient_id", -p.ID)
	//TODO 检查Update后内存中的数据是否发生变化
	if result.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("删除失败，查无此ID（DeleteID=%v）", p.ID))
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
