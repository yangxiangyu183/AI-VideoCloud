
// 自动生成模板Alert
package alert_video
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// alert表 结构体  Alert
type Alert struct {
    global.GVA_MODEL
  Video  *string `json:"video" form:"video" gorm:"comment:监控视频;column:video;size:100;" binding:"required"`  //监控视频
  CameraAddress  *string `json:"cameraAddress" form:"cameraAddress" gorm:"comment:摄像头点位;column:camera_address;size:50;" binding:"required"`  //摄像头点位
  AlertType  *string `json:"alertType" form:"alertType" gorm:"comment:预警类型;column:alert_type;size:50;" binding:"required"`  //预警类型
  AlertTime  *time.Time `json:"alertTime" form:"alertTime" gorm:"comment:预警时间;column:alert_time;" binding:"required"`  //预警时间
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName alert表 Alert自定义表名 alert
func (Alert) TableName() string {
    return "alert"
}





