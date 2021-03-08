package Models

type Admin struct {
	AdminID           int    `json:"admin_id" gorm:"primary_key;column:admin_id;type:int(11) auto_increment;comment:'管理员id'"`                                                                      // 管理员id
	AdminUsername     string `json:"admin_username" gorm:"column:admin_username;type:char(20);not null;unique_index:hospital_admin_admin_username_uindex;comment:'管理员用户名'"`                        // 管理员用户名
	AdminPassword     string `json:"admin_password" gorm:"column:admin_password;type:char(30);not null;comment:'管理员密码'"`                                                                           // 管理员密码
	AdminName         string `json:"admin_name" gorm:"column:admin_name;type:varchar(10);not null;comment:'管理员名字'"`                                                                                // 管理员名字
	AdminPhonenum     string `json:"admin_phoneNum" gorm:"column:admin_phoneNum;type:char(15);not null;comment:'管理员电话'"`                                                                           // 管理员电话
	AdminPermissionID int    `json:"admin_permission_id" gorm:"column:admin_permission_id;type:int(11);not null;index:hospital_admin_hospital_admin_permissions_permission_id_fk;comment:'管理员权限'"` // 管理员权限
}

// TableName returns the table name of the Admin model
func (a *Admin) TableName() string {
	return "hospital_admin"
}