package Models

type User struct {
	UserID       int    `json:"user_id" gorm:"primary_key;column:user_id;type:int(11) auto_increment;comment:'注册用户id'"`                                                   // 注册用户id
	UserNickname string `json:"user_nickname" gorm:"column:user_nickname;type:varchar(16);not null;index:public_user_user_nickname_user_phoneNum_index;comment:'注册用户昵称'"` // 注册用户昵称
	UserPassword string `json:"user_password" gorm:"column:user_password;type:char(32);not null;comment:'注册用户密码'"`                                                        // 注册用户密码
	UserPhonenum string `json:"user_phoneNum" gorm:"column:user_phoneNum;type:char(16);not null;index:public_user_user_nickname_user_phoneNum_index;comment:'注册用户手机号'"`   // 注册用户手机号
}

// TableName returns the table name of the User model
func (u *User) TableName() string {
	return "public_user"
}