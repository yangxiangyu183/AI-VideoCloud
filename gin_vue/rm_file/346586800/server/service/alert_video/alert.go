
package alert_video

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alert_video"
    alert_videoReq "github.com/flipped-aurora/gin-vue-admin/server/model/alert_video/request"
    "gorm.io/gorm"
)

type AlertService struct {}
// CreateAlert 创建alert表记录
// Author [yourname](https://github.com/yourname)
func (alertService *AlertService) CreateAlert(ctx context.Context, alert *alert_video.Alert) (err error) {
	err = global.GVA_DB.Create(alert).Error
	return err
}

// DeleteAlert 删除alert表记录
// Author [yourname](https://github.com/yourname)
func (alertService *AlertService)DeleteAlert(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&alert_video.Alert{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&alert_video.Alert{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteAlertByIds 批量删除alert表记录
// Author [yourname](https://github.com/yourname)
func (alertService *AlertService)DeleteAlertByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&alert_video.Alert{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&alert_video.Alert{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateAlert 更新alert表记录
// Author [yourname](https://github.com/yourname)
func (alertService *AlertService)UpdateAlert(ctx context.Context, alert alert_video.Alert) (err error) {
	err = global.GVA_DB.Model(&alert_video.Alert{}).Where("id = ?",alert.ID).Updates(&alert).Error
	return err
}

// GetAlert 根据ID获取alert表记录
// Author [yourname](https://github.com/yourname)
func (alertService *AlertService)GetAlert(ctx context.Context, ID string) (alert alert_video.Alert, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&alert).Error
	return
}
// GetAlertInfoList 分页获取alert表记录
// Author [yourname](https://github.com/yourname)
func (alertService *AlertService)GetAlertInfoList(ctx context.Context, info alert_videoReq.AlertSearch) (list []alert_video.Alert, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&alert_video.Alert{})
    var alerts []alert_video.Alert
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Video != nil && *info.Video != "" {
        db = db.Where("video = ?", *info.Video)
    }
    if info.CameraAddress != nil && *info.CameraAddress != "" {
        db = db.Where("camera_address = ?", *info.CameraAddress)
    }
    if info.AlertType != nil && *info.AlertType != "" {
        db = db.Where("alert_type = ?", *info.AlertType)
    }
    if info.AlertTime != nil {
        db = db.Where("alert_time = ?", *info.AlertTime)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&alerts).Error
	return  alerts, total, err
}
func (alertService *AlertService)GetAlertPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
