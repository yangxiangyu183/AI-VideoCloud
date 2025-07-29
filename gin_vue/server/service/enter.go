package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/alert_video"
	"github.com/flipped-aurora/gin-vue-admin/server/service/device"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/task_bor"
	"github.com/flipped-aurora/gin-vue-admin/server/service/user"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup      system.ServiceGroup
	ExampleServiceGroup     example.ServiceGroup
	DeviceServiceGroup      device.ServiceGroup
	UserServiceGroup        user.ServiceGroup
	Task_borServiceGroup    task_bor.ServiceGroup
	Alert_videoServiceGroup alert_video.ServiceGroup
}
