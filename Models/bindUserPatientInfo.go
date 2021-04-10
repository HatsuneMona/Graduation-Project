package Models

//废弃
type BindUserPatientInfo struct {
	BindID     int `json:"bind_id" gorm:"primary_key;column:bind_id;type:int(11) auto_increment;comment:'绑定信息id'"`                            // 绑定信息id
	UserID     int `json:"user_id" gorm:"column:user_id;type:int(11);not null;index:public_bind_user_patient_user_id_index;comment:'注册用户id'"` // 注册用户id
	PatientsID int `json:"patients_id" gorm:"column:patients_id;type:char(200);not null;default:'0';comment:'第一位患者的id信息'"`                    // 所有患者的id信息
	BindCount  int `json:"bind_count" gorm:"column:bind_count;type:int(11);not null;default:'0';comment:'第二位患者的id信息'"`                        // 绑定患者的数量
	Patients   []int
}

// TableName returns the table name of the BindUserPatientInfo model
func (b *BindUserPatientInfo) TableName() string {
	return "public_bind_user_patient"
}

type bindUserPatientHandler interface {
	GetBindInfoByBindID(id ...int)
	GetBindInfoByUserID(id ...int)
}
