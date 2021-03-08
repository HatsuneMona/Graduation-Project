package Models

import (
	"database/sql"
)

type DepartmentInfo struct {
	DepartmentID     int            `json:"department_id" gorm:"primary_key;column:department_id;type:int(11) auto_increment;comment:'科室id'"`                                           // 科室id
	DepartmentName   string         `json:"department_name" gorm:"column:department_name;type:varchar(30);not null;comment:'科室名'"`                                                      // 科室名
	DepartmentHeader sql.NullInt64  `json:"department_header" gorm:"column:department_header;type:int(11);index:hospital_department_hospital_doctor_doctor_id_fk;comment:'科室主任（工号id）'"` // 科室主任（工号id）
	DepartmentIntro  sql.NullString `json:"department_intro" gorm:"column:department_intro;type:text;comment:'科室简介'"`                                                                   // 科室简介
}

// TableName returns the table name of the DepartmentInfo model
func (d *DepartmentInfo) TableName() string {
	return "hospital_department"
}