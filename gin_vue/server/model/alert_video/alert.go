
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
  CameraId  *int `json:"cameraId" form:"cameraId" gorm:"comment:摄像头点位id;column:camera_id;size:19;" binding:"required"`  //摄像头点位id
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}

// AlertWithDevice 包含设备信息的告警结构体
type AlertWithDevice struct {
    Alert
    DeviceName   *string `json:"deviceName"`   // 设备名称
    DeviceStatus *string `json:"deviceStatus"` // 设备状态
    Resolution   *string `json:"resolution"`   // 分辨率
    StreamUrl    *string `json:"streamUrl"`    // 视频流地址
    GroupId      *int    `json:"groupId"`      // 分组ID
}

// 设备状态常量
const (
    DeviceStatusOnline  = "1" // 在线
    DeviceStatusOffline = "2" // 离线
    DeviceStatusFault   = "3" // 故障
)


// TableName alert表 Alert自定义表名 alert
func (Alert) TableName() string {
    return "alert"
}





