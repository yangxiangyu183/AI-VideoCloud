
package device

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
    deviceReq "github.com/flipped-aurora/gin-vue-admin/server/model/device/request"
    "gorm.io/gorm"
)

type DeviceGroupService struct {}
// CreateDeviceGroup 创建deviceGroup表记录
// Author [yourname](https://github.com/yourname)
func (deviceGroupService *DeviceGroupService) CreateDeviceGroup(ctx context.Context, deviceGroup *device.DeviceGroup) (err error) {
	err = global.GVA_DB.Create(deviceGroup).Error
	return err
}

// DeleteDeviceGroup 删除deviceGroup表记录
// Author [yourname](https://github.com/yourname)
func (deviceGroupService *DeviceGroupService)DeleteDeviceGroup(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&device.DeviceGroup{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&device.DeviceGroup{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteDeviceGroupByIds 批量删除deviceGroup表记录
// Author [yourname](https://github.com/yourname)
func (deviceGroupService *DeviceGroupService)DeleteDeviceGroupByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&device.DeviceGroup{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&device.DeviceGroup{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateDeviceGroup 更新deviceGroup表记录
// Author [yourname](https://github.com/yourname)
func (deviceGroupService *DeviceGroupService)UpdateDeviceGroup(ctx context.Context, deviceGroup device.DeviceGroup) (err error) {
	err = global.GVA_DB.Model(&device.DeviceGroup{}).Where("id = ?",deviceGroup.ID).Updates(&deviceGroup).Error
	return err
}

// GetDeviceGroup 根据ID获取deviceGroup表记录
// Author [yourname](https://github.com/yourname)
func (deviceGroupService *DeviceGroupService)GetDeviceGroup(ctx context.Context, ID string) (deviceGroup device.DeviceGroup, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&deviceGroup).Error
	return
}
// GetDeviceGroupInfoList 分页获取deviceGroup表记录
// Author [yourname](https://github.com/yourname)
func (deviceGroupService *DeviceGroupService)GetDeviceGroupInfoList(ctx context.Context, info deviceReq.DeviceGroupSearch) (list []device.DeviceGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&device.DeviceGroup{})
    var deviceGroups []device.DeviceGroup
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.GroupName != nil && *info.GroupName != "" {
        db = db.Where("group_name LIKE ?", "%"+ *info.GroupName+"%")
    }
    if info.Pid != nil {
        db = db.Where("pid = ?", *info.Pid)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&deviceGroups).Error
	return  deviceGroups, total, err
}
func (deviceGroupService *DeviceGroupService)GetDeviceGroupPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
