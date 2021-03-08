package Models

import (
	"database/sql"
)

type Doctor struct {
	DoctorID         int            `json:"doctor_id" gorm:"primary_key;column:doctor_id;type:int(11) auto_increment;comment:'医生工号'"`                                                                            // 医生工号
	DoctorName       string         `json:"doctor_name" gorm:"column:doctor_name;type:varchar(8);not null;comment:'医生名字'"`                                                                                       // 医生名字
	DoctorPassword   string         `json:"doctor_password" gorm:"column:doctor_password;type:char(32);not null;comment:'医生登录密码'"`                                                                               // 医生登录密码
	DoctorDepartment int            `json:"doctor_department" gorm:"column:doctor_department;type:int(11);not null;index:hospital_doctor_hospital_department_department_id_fk;default:'0';comment:'医生所在科室（部门）'"` // 医生所在科室（部门）
	DoctorClass      int            `json:"doctor_class" gorm:"column:doctor_class;type:int(11);not null;default:'0';comment:'医生职位'"`                                                                            // 医生职位
	DoctorIntro      sql.NullString `json:"doctor_intro" gorm:"column:doctor_intro;type:text;comment:'医生简介'"`                                                                                                    // 医生简介
}

// TableName returns the table name of the Doctor model
func (d *Doctor) TableName() string {
	return "hospital_doctor"
}