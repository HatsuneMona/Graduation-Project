package Models

import (
	"database/sql"
)

type Patient struct {
	PatientID           int            `json:"patient_id" gorm:"primary_key;column:patient_id;type:int(11) auto_increment;comment:'患者id'"`                                                                 // 患者id
	PatientName         string         `json:"patient_name" gorm:"column:patient_name;type:varchar(10);not null;comment:'患者真实姓名'"`                                                                         // 患者真实姓名
	PatientIdentityCard string         `json:"patient_identity_card" gorm:"column:patient_identity_card;type:char(18);not null;unique_index:public_patient_patient_identity_card_uindex;comment:'患者身份证号'"` // 患者身份证号
	PatientPhonenum     sql.NullString `json:"patient_phoneNum" gorm:"column:patient_phoneNum;type:char(16);comment:'患者电话号'"`                                                                              // 患者电话号

}

// TableName returns the table name of the Patient model
func (p *Patient) TableName() string {
	return "public_patient"
}

type PatientHandlers struct {
	patientReader
	patientSetter
}
