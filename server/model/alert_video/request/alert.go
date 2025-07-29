
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

// 设备状态常量
const (
	DeviceStatusOnline  = "1" // 在线
	DeviceStatusOffline = "2" // 离线
	DeviceStatusFault   = "3" // 故障
)

type AlertSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Video  *string `json:"video" form:"video"` 
      CameraAddress  *string `json:"cameraAddress" form:"cameraAddress"` 
      AlertType  *string `json:"alertType" form:"alertType"` 
      AlertTime  *time.Time `json:"alertTime" form:"alertTime"` 
      CameraId  *int `json:"cameraId" form:"cameraId"` 
      DeviceStatus *string `json:"deviceStatus" form:"deviceStatus"` // 设备状态筛选：1-在线 2-离线 3-故障
      DeviceName   *string `json:"deviceName" form:"deviceName"`     // 设备名称搜索，支持模糊匹配
      Keyword      *string `json:"keyword" form:"keyword"`           // 关键词搜索
    request.PageInfo
}

// ValidateDeviceStatus 验证设备状态参数
func (a *AlertSearch) ValidateDeviceStatus() bool {
	if a.DeviceStatus == nil {
		return true // 允许为空
	}
	
	status := *a.DeviceStatus
	return status == DeviceStatusOnline || status == DeviceStatusOffline || status == DeviceStatusFault
}

// IsValidDeviceStatus 检查设备状态是否有效
func IsValidDeviceStatus(status string) bool {
	return status == DeviceStatusOnline || status == DeviceStatusOffline || status == DeviceStatusFault
}
