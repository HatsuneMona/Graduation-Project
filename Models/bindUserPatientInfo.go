package Models

type BindUserPatientInfo struct {
	BindID     int `json:"bind_id" gorm:"primary_key;column:bind_id;type:int(11) auto_increment;comment:'绑定信息id'"`                                                                       // 绑定信息id
	UserID     int `json:"user_id" gorm:"column:user_id;type:int(11);not null;index:public_bind_user_patient_public_user_user_id_fk;comment:'注册用户id'"`                                   // 注册用户id
	PatientAID int `json:"patient_a_id" gorm:"column:patient_a_id;type:int(11);not null;index:public_bind_user_patient_public_patient_patient_id_fk;default:'0';comment:'第一位患者的id信息'"`   // 第一位患者的id信息
	PatientBID int `json:"patient_b_id" gorm:"column:patient_b_id;type:int(11);not null;index:public_bind_user_patient_public_patient_patient_id_fk_2;default:'0';comment:'第二位患者的id信息'"` // 第二位患者的id信息
	PatientCID int `json:"patient_c_id" gorm:"column:patient_c_id;type:int(11);not null;index:public_bind_user_patient_public_patient_patient_id_fk_3;default:'0';comment:'第三位患者的id信息'"` // 第三位患者的id信息
	PatientDID int `json:"patient_d_id" gorm:"column:patient_d_id;type:int(11);not null;index:public_bind_user_patient_public_patient_patient_id_fk_4;default:'0';comment:'第四位患者的id信息'"` // 第四位患者的id信息
	PatientEID int `json:"patient_e_id" gorm:"column:patient_e_id;type:int(11);not null;index:public_bind_user_patient_public_patient_patient_id_fk_5;default:'0';comment:'第五位患者的id信息'"` // 第五位患者的id信息
}

// TableName returns the table name of the BindUserPatientInfo model
func (b *BindUserPatientInfo) TableName() string {
	return "public_bind_user_patient"
}