package device

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DeviceGroupApi
	MonitorDeviceApi
}

var (
	deviceGroupService   = service.ServiceGroupApp.DeviceServiceGroup.DeviceGroupService
	monitorDeviceService = service.ServiceGroupApp.DeviceServiceGroup.MonitorDeviceService
)
