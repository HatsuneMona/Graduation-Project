package Models

import (
	"errors"
	"fmt"
	"service/Databases/DBPool"
)

type AdminPermissionsInfo struct {
	ID        int    `json:"permission_id" gorm:"primary_key;column:permission_id;type:int(11) auto_increment;comment:'权限编号（id）'"`                   // 权限编号（id）
	ShortName string `json:"permission_short_name" gorm:"column:permission_short_name;type:varchar(20);not null;comment:'权限名（简要）'"`                  // 权限名（简要）
	Introduce string `json:"permission_introduce" gorm:"column:permission_introduce;type:varchar(200);not null;comment:'权限详细介绍'"`                    // 权限详细介绍
	Choose    string `json:"permission_choose" gorm:"column:permission_choose;type:varchar(100);not null;comment:'如果是混合权限，则该值代表有哪些权限可用，例如【1,6,13】'"` // 如果是混合权限，则该值代表有哪些权限可用，例如【1,6,13】
	ChooseID  []int
	AdminPermissionsInfoHandlers
}

// TableName returns the table name of the AdminPermissionsInfo model
func (info *AdminPermissionsInfo) TableName() string {
	return "hospital_admin_permissions"
}

type AdminPermissionsInfoHandlers interface {
	adminPermissionsInfoReader
	adminPermissionsInfoSetter
}

type adminPermissionsInfoReader interface {
	GetAdminPermissionsInfoByID(id int) error
	GetAdminPermissionsInfoByShortName(ShortName string) error
	haveAdminPermissionID() (bool, error)
}

type adminPermissionsInfoSetter interface {
	UpdatePermissionShortName(newShortName string) error
	UpdatePermissionIntroduce(newIntroduce string) error
	UpdatePermissionChoose(newChooses ...int) error
	AddNewPermission() error
}

func (info *AdminPermissionsInfo) GetAdminPermissionsInfoByID(id int) error {
	info.ID = id
	res := DBPool.DB.Take(&info)
	if res.Error != nil {
		//log.Logger.Info(fmt.Sprintf("查询Admin错误，查询错误=%v", res.Error))
		return errors.New(fmt.Sprintf("查询AdminPermissionsInfo错误，查询错误=%v", res.Error))
	}
	return nil
}

func (info *AdminPermissionsInfo) GetAdminPermissionsInfoByShortName(ShortName string) error {
	res := DBPool.DB.Where("permission_short_name = ?", ShortName).Take(&info)
	if res.Error != nil {
		//log.Logger.Info(fmt.Sprintf("查询Admin错误，查询错误=%v", res.Error))
		return errors.New(fmt.Sprintf("查询AdminPermissionsInfo错误，查询错误=%v", res.Error))
	}
	return nil
}

//给定权限ID，查询是否有该权限。
func (info *AdminPermissionsInfo) haveAdminPermissionID() (bool, error) {
	res := DBPool.DB.Select("permission_id").First(&info)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func (info *AdminPermissionsInfo) UpdatePermissionShortName(newShortName string) error {
	if info.ShortName == newShortName {
		return errors.New("新旧数据相同，pass")
	}
	res := DBPool.DB.Model(&info).Update("permission_short_name", newShortName)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (info *AdminPermissionsInfo) UpdatePermissionIntroduce(newIntroduce string) error {
	if info.Introduce == newIntroduce {
		return errors.New("新旧数据相同，pass")
	}
	res := DBPool.DB.Model(&info).Update("permission_introduce", newIntroduce)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (info *AdminPermissionsInfo) UpdatePermissionChoose(newChooseInts ...int) error {
	newChoose := fmt.Sprint(newChooseInts)
	if info.Choose == newChoose {
		return errors.New("新旧数据相同，pass")
	}
	res := DBPool.DB.Model(&info).Update("permission_choose", newChoose)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

//新建权限
func (info *AdminPermissionsInfo) AddNewPermission() error {
	if info.ID != 0 {
		return errors.New("禁止指定permissionID")
	}
	if info.Choose != "" {
		return errors.New("禁止指定permissionChoose字段")
	}

	info.Choose = fmt.Sprint(info.ChooseID)

	result := DBPool.DB.Create(&info)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (info *AdminPermissionsInfo) generateChoose(reGenerate bool) error {
	//当权限ID有一个及以上开始进行检查
	if len(info.ChooseID) > 1 {
		for i, permissionID := range info.ChooseID {
			tempPermissionInfo := AdminPermissionsInfo{
				ID: permissionID,
			}
			if have, err := tempPermissionInfo.haveAdminPermissionID(); err != nil {
				if have {
					continue
				} else {
					return errors.New(fmt.Sprintf("权限多选，第%v项错误：未找到编号为%v的权限", i+1, permissionID))
				}
			} else {
				return err
			}
		}
	}
	if reGenerate == false && info.Choose != "" {

	}
	return nil
}
