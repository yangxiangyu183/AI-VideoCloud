
package device

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
    deviceReq "github.com/flipped-aurora/gin-vue-admin/server/model/device/request"
    "gorm.io/gorm"
)

type MonitorDeviceService struct {}
// CreateMonitorDevice 创建monitorDevice表记录
// Author [yourname](https://github.com/yourname)
func (monitorDeviceService *MonitorDeviceService) CreateMonitorDevice(ctx context.Context, monitorDevice *device.MonitorDevice) (err error) {
	err = global.GVA_DB.Create(monitorDevice).Error
	return err
}

// DeleteMonitorDevice 删除monitorDevice表记录
// Author [yourname](https://github.com/yourname)
func (monitorDeviceService *MonitorDeviceService)DeleteMonitorDevice(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&device.MonitorDevice{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&device.MonitorDevice{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteMonitorDeviceByIds 批量删除monitorDevice表记录
// Author [yourname](https://github.com/yourname)
func (monitorDeviceService *MonitorDeviceService)DeleteMonitorDeviceByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&device.MonitorDevice{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&device.MonitorDevice{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateMonitorDevice 更新monitorDevice表记录
// Author [yourname](https://github.com/yourname)
func (monitorDeviceService *MonitorDeviceService)UpdateMonitorDevice(ctx context.Context, monitorDevice device.MonitorDevice) (err error) {
	err = global.GVA_DB.Model(&device.MonitorDevice{}).Where("id = ?",monitorDevice.ID).Updates(&monitorDevice).Error
	return err
}

// GetMonitorDevice 根据ID获取monitorDevice表记录
// Author [yourname](https://github.com/yourname)
func (monitorDeviceService *MonitorDeviceService)GetMonitorDevice(ctx context.Context, ID string) (monitorDevice device.MonitorDevice, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&monitorDevice).Error
	return
}
// GetMonitorDeviceInfoList 分页获取monitorDevice表记录
// Author [yourname](https://github.com/yourname)
func (monitorDeviceService *MonitorDeviceService)GetMonitorDeviceInfoList(ctx context.Context, info deviceReq.MonitorDeviceSearch) (list []device.MonitorDevice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&device.MonitorDevice{})
    var monitorDevices []device.MonitorDevice
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.DeviceName != nil && *info.DeviceName != "" {
        db = db.Where("device_name LIKE ?", "%"+ *info.DeviceName+"%")
    }
    if info.GroupId != nil {
        db = db.Where("group_id = ?", *info.GroupId)
    }
    if info.Status != nil && *info.Status != "" {
        db = db.Where("status = ?", *info.Status)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&monitorDevices).Error
	return  monitorDevices, total, err
}
func (monitorDeviceService *MonitorDeviceService)GetMonitorDevicePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
