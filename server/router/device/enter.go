package device

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	DeviceGroupRouter
	MonitorDeviceRouter
}

var (
	deviceGroupApi   = api.ApiGroupApp.DeviceApiGroup.DeviceGroupApi
	monitorDeviceApi = api.ApiGroupApp.DeviceApiGroup.MonitorDeviceApi
)
