package Models

import (
	"database/sql"
)

type AdminPermissionsInfo struct {
	PermissionID          int            `json:"permission_id" gorm:"primary_key;column:permission_id;type:int(11) auto_increment;comment:'权限编号（id）'"`                      // 权限编号（id）
	PermissionMulti       int32          `json:"permission_multi" gorm:"column:permission_multi;type:tinyint(1);not null;default:'0';comment:'是否为混合权限'"`                    // 是否为混合权限
	PermissionShortName   string         `json:"permission_short_name" gorm:"column:permission_short_name;type:varchar(20);not null;comment:'权限名（简要）'"`                     // 权限名（简要）
	PermissionIntroduce   string         `json:"permission_introduce" gorm:"column:permission_introduce;type:varchar(200);not null;comment:'权限详细介绍'"`                       // 权限详细介绍
	PermissionMultiChoose sql.NullString `json:"permission_multi_choose" gorm:"column:permission_multi_choose;type:varchar(100);comment:'如果是混合权限，则该值代表有哪些权限可用，例如【1,6,13】'"` // 如果是混合权限，则该值代表有哪些权限可用，例如【1,6,13】
}

// TableName returns the table name of the AdminPermissionsInfo model
func (a *AdminPermissionsInfo) TableName() string {
	return "hospital_admin_permissions"
}