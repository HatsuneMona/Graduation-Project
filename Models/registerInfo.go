package Models

type RegisterInfo struct {
	RegisterID            int `json:"register_id" gorm:"primary_key;column:register_id;type:int(11) auto_increment;comment:'挂号信息id'"`                                                                  // 挂号信息id
	RegisterDepartmentID  int `json:"register_department_id" gorm:"column:register_department_id;type:int(11);not null;index:hospital_register_hospital_department_department_id_fk;comment:'挂号科室id'"` // 挂号科室id
	RegisterDoctorID      int `json:"register_doctor_id" gorm:"column:register_doctor_id;type:int(11);not null;index:hospital_register_hospital_doctor_doctor_id_fk;comment:'挂号医师id'"`                 // 挂号医师id
	RegisterQuantity      int `json:"register_quantity" gorm:"column:register_quantity;type:int(11);not null;comment:'每次预约的总数量'"`                                                                      // 每次预约的总数量
	RegisterOrderQuantity int `json:"register_order_quantity" gorm:"column:register_order_quantity;type:int(11);not null;default:'0';comment:'已预约数量'"`                                                 // 已预约数量
	RegisterPrise         int `json:"register_prise" gorm:"column:register_prise;type:int(11);not null;comment:'挂号价格'"`                                                                                // 挂号价格
}

// TableName returns the table name of the RegisterInfo model
func (r *RegisterInfo) TableName() string {
	return "hospital_register_info"
}