package Models

type RegisteredInfo struct {
	RegisteredID        int   `json:"registered_id" gorm:"primary_key;column:registered_id;type:int(11) auto_increment;comment:'挂号id'"`                                                              // 挂号id
	RegisteredUserID    int   `json:"registered_user_id" gorm:"column:registered_user_id;type:int(11);not null;index:public_register_public_user_user_id_fk;comment:'挂号账户id'"`                       // 挂号账户id
	RegisteredPatientID int   `json:"registered_patient_id" gorm:"column:registered_patient_id;type:int(11);not null;index:public_registed_public_patient_patient_id_fk;comment:'挂号患者id'"`           // 挂号患者id
	RegisterInfoID      int   `json:"register_info_id" gorm:"column:register_info_id;type:int(11);not null;index:public_registed_hospital_register_register_id_fk;comment:'挂号信息id（与医疗系统内号源信息id挂钩）'"` // 挂号信息id（与医疗系统内号源信息id挂钩）
	RegisteredPaid      int32 `json:"registered_paid" gorm:"column:registered_paid;type:tinyint(1);not null;default:'0';comment:'是否支付订单'"`                                                           // 是否支付订单
	RegisteredEffective int32 `json:"registered_effective" gorm:"column:registered_effective;type:tinyint(1);not null;default:'1';comment:'是否有效'"`                                                   // 是否有效
}

// TableName returns the table name of the RegisteredInfo model
func (r *RegisteredInfo) TableName() string {
	return "public_registered"
}