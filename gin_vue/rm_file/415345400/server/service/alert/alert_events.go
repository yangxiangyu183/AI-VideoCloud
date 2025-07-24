
package alert

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alert"
    alertReq "github.com/flipped-aurora/gin-vue-admin/server/model/alert/request"
    "gorm.io/gorm"
)

type AlertEventsService struct {}
// CreateAlertEvents 创建alertEvents表记录
// Author [yourname](https://github.com/yourname)
func (alertEventsService *AlertEventsService) CreateAlertEvents(ctx context.Context, alertEvents *alert.AlertEvents) (err error) {
	err = global.GVA_DB.Create(alertEvents).Error
	return err
}

// DeleteAlertEvents 删除alertEvents表记录
// Author [yourname](https://github.com/yourname)
func (alertEventsService *AlertEventsService)DeleteAlertEvents(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&alert.AlertEvents{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&alert.AlertEvents{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteAlertEventsByIds 批量删除alertEvents表记录
// Author [yourname](https://github.com/yourname)
func (alertEventsService *AlertEventsService)DeleteAlertEventsByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&alert.AlertEvents{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&alert.AlertEvents{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateAlertEvents 更新alertEvents表记录
// Author [yourname](https://github.com/yourname)
func (alertEventsService *AlertEventsService)UpdateAlertEvents(ctx context.Context, alertEvents alert.AlertEvents) (err error) {
	err = global.GVA_DB.Model(&alert.AlertEvents{}).Where("id = ?",alertEvents.ID).Updates(&alertEvents).Error
	return err
}

// GetAlertEvents 根据ID获取alertEvents表记录
// Author [yourname](https://github.com/yourname)
func (alertEventsService *AlertEventsService)GetAlertEvents(ctx context.Context, ID string) (alertEvents alert.AlertEvents, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&alertEvents).Error
	return
}
// GetAlertEventsInfoList 分页获取alertEvents表记录
// Author [yourname](https://github.com/yourname)
func (alertEventsService *AlertEventsService)GetAlertEventsInfoList(ctx context.Context, info alertReq.AlertEventsSearch) (list []alert.AlertEvents, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&alert.AlertEvents{})
    var alertEventss []alert.AlertEvents
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.TaskId != nil {
        db = db.Where("task_id = ?", *info.TaskId)
    }
    if info.CameraId != nil {
        db = db.Where("camera_id = ?", *info.CameraId)
    }
    if info.AlertType != nil && *info.AlertType != "" {
        db = db.Where("alert_type = ?", *info.AlertType)
    }
    if info.Servenrity != nil && *info.Servenrity != "" {
        db = db.Where("servenrity = ?", *info.Servenrity)
    }
    if info.ImagePath != nil && *info.ImagePath != "" {
        db = db.Where("image_path = ?", *info.ImagePath)
    }
    if info.VideoPath != nil && *info.VideoPath != "" {
        db = db.Where("video_path = ?", *info.VideoPath)
    }
    if info.IsRead != nil && *info.IsRead != "" {
        db = db.Where("is_read = ?", *info.IsRead)
    }
    if info.IsProcessed != nil && *info.IsProcessed != "" {
        db = db.Where("is_processed = ?", *info.IsProcessed)
    }
    if info.ProcessedBy != nil {
        db = db.Where("processed_by = ?", *info.ProcessedBy)
    }
    if info.ProcessedTime != nil {
        db = db.Where("processed_time = ?", *info.ProcessedTime)
    }
    if info.ProcessedResult != nil && *info.ProcessedResult != "" {
        db = db.Where("processed_result = ?", *info.ProcessedResult)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&alertEventss).Error
	return  alertEventss, total, err
}
func (alertEventsService *AlertEventsService)GetAlertEventsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
