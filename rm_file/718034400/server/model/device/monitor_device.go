
// 自动生成模板MonitorDevice
package device
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// monitorDevice表 结构体  MonitorDevice
type MonitorDevice struct {
    global.GVA_MODEL
  DeviceName  *string `json:"deviceName" form:"deviceName" gorm:"comment:摄像头点位;column:device_name;size:100;" binding:"required"`  //摄像头点位
  GroupId  *int `json:"groupId" form:"groupId" gorm:"comment:关联分组id;column:group_id;" binding:"required"`  //关联分组id
  StreamUrl  *string `json:"streamUrl" form:"streamUrl" gorm:"comment:视频流地址;column:stream_url;size:191;" binding:"required"`  //视频流地址
  Status  *string `json:"status" form:"status" gorm:"default:1;comment:设备状态：1-在线 2-离线 3-故障;column:status;size:10;"`  //设备状态：1-在线 2-离线 3-故障
  Resolution  *string `json:"resolution" form:"resolution" gorm:"comment:分辨率;column:resolution;size:50;"`  //分辨率
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName monitorDevice表 MonitorDevice自定义表名 monitor_device
func (MonitorDevice) TableName() string {
    return "monitor_device"
}





