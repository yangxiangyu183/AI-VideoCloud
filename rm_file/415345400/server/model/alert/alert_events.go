
// 自动生成模板AlertEvents
package alert
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// alertEvents表 结构体  AlertEvents
type AlertEvents struct {
    global.GVA_MODEL
  TaskId  *int `json:"taskId" form:"taskId" gorm:"comment:关联任务id;column:task_id;size:19;" binding:"required"`  //关联任务id
  CameraId  *int `json:"cameraId" form:"cameraId" gorm:"comment:关联摄像头id;column:camera_id;size:19;" binding:"required"`  //关联摄像头id
  AlertType  *string `json:"alertType" form:"alertType" gorm:"comment:告警类型;column:alert_type;size:50;" binding:"required"`  //告警类型
  Servenrity  *string `json:"servenrity" form:"servenrity" gorm:"comment:警告级别;column:servenrity;size:50;" binding:"required"`  //警告级别
  ImagePath  *string `json:"imagePath" form:"imagePath" gorm:"comment:截图路径;column:image_path;size:50;" binding:"required"`  //截图路径
  VideoPath  *string `json:"videoPath" form:"videoPath" gorm:"comment:视频片段路径;column:video_path;size:100;" binding:"required"`  //视频片段路径
  IsRead  *string `json:"isRead" form:"isRead" gorm:"comment:是否已读;column:is_read;size:50;" binding:"required"`  //是否已读
  IsProcessed  *string `json:"isProcessed" form:"isProcessed" gorm:"comment:是否已经处理;column:is_processed;size:50;" binding:"required"`  //是否已经处理
  ProcessedBy  *int `json:"processedBy" form:"processedBy" gorm:"comment:处理用户id;column:processed_by;size:19;" binding:"required"`  //处理用户id
  ProcessedTime  *time.Time `json:"processedTime" form:"processedTime" gorm:"comment:处理时间;column:processed_time;" binding:"required"`  //处理时间
  ProcessedResult  *string `json:"processedResult" form:"processedResult" gorm:"comment:处理结果;column:processed_result;size:50;" binding:"required"`  //处理结果
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName alertEvents表 AlertEvents自定义表名 alert_events
func (AlertEvents) TableName() string {
    return "alert_events"
}





