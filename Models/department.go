package Models

import (
	"errors"
	"service/Databases"
)

type DepartmentInfo struct {
	ID     int    `json:"department_id" gorm:"primary_key;column:department_id;type:int(11) auto_increment;comment:'科室id'"`                                                    // 科室id
	Name   string `json:"department_name" gorm:"column:department_name;type:varchar(30);not null;comment:'科室名'"`                                                               // 科室名
	Header int    `json:"department_header" gorm:"column:department_header;type:int(11);not null;index:hospital_department_hospital_doctor_doctor_id_fk;comment:'科室主任（工号id）'"` // 科室主任（工号id）
	Intro  string `json:"department_intro" gorm:"column:department_intro;not null;type:text;comment:'科室简介'"`                                                                   // 科室简介
}

// TableName returns the table name of the DepartmentInfo model
func (d *DepartmentInfo) TableName() string {
	return "hospital_department"
}

type DepartmentHandler interface {
	GetDepartmentByID(id ...int) error

	UpdateName(newName string) error
	UpdateHeader(newHeaderID int) error
	UpdateIntro(newIntro string) error
}

//GetDepartmentByID
//
//@Description	通货给定科室ID，返回科室所有信息
//
//@Param
//				`id ...int` 需要查询的ID，可选参数。
//				若 `id ...int` 保持为空，则从调用者的结构体中获取要查询的科室ID。
//
//@Return
//				查询到的信息将写入调用者的结构体中，并返回。
//				error 查询正确则返回nil，否则返回Error
func (d *DepartmentInfo) GetDepartmentByID(id ...int) error {
	if len(id) > 0 {
		d.ID = id[0]
	}
	ret := Databases.DB.Take(&d)
	if ret.Error != nil {
		return ret.Error
	} else {
		return nil
	}
}

/*
UpdateName

@Description	更新科室名称

@Param
				`newName` string 新名称

@Return
				成功：error=nil，失败：返回错误信息。
*/
func (d *DepartmentInfo) UpdateName(newName string) error {
	if d.Name == newName {
		return errors.New("新旧密码相同，pass")
	}
	ret := Databases.DB.Model(&d).Update("department_name", newName)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

/*
UpdateHeader

@Description	更新科室领导

@Param
				`newHeaderId` id 新领导的id

@Return
				成功：error=nil，失败：返回错误信息。
*/
func (d *DepartmentInfo) UpdateHeader(newHeaderID int) error {
	if d.Header == newHeaderID {
		return errors.New("新旧数据相同，pass")
	}
	ret := Databases.DB.Model(&d).Update("department_header", newHeaderID)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

/*
UpdateIntro

@Description	更新科室介绍

@Param
				`newIntro` string 新科室介绍

@Return
				成功：error=nil，失败：返回错误信息。
*/
func (d *DepartmentInfo) UpdateIntro(newIntro string) error {
	ret := Databases.DB.Model(&d).Update("department_intro", newIntro)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
