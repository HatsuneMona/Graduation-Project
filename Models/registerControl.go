package Models

import (
	"database/sql"
	"time"
)

type RegisterControlInfo struct {
	ControlID          int          `json:"control_id" gorm:"primary_key;column:control_id;type:int(11) auto_increment;comment:'控制器id'"`            // 控制器id
	ControlRegisterID  int          `json:"control_register_id" gorm:"column:control_register_id;type:int(11);not null;comment:'控制器_号源id'"`         // 控制器_号源id
	ControlWeak        string       `json:"control_weak" gorm:"column:control_weak;type:char(8);not null;default:'1234567';comment:'控制器_放号日期（星期）'"` // 控制器_放号日期（星期）
	ControlTime        time.Time    `json:"control_time" gorm:"column:control_time;type:time;not null;comment:'控制器_放号时间'"`                          // 控制器_放号时间
	ControlDestroyTime sql.NullTime `json:"control_destroy_time" gorm:"column:control_destroy_time;type:date;comment:'控制器_截止时间'"`                   // 控制器_截止时间
}

// TableName returns the table name of the RegisterControlInfo model
func (r *RegisterControlInfo) TableName() string {
	return "hospital_register_control"
}