package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/device"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/task_bor"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
<<<<<<< HEAD
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	DeviceServiceGroup  device.ServiceGroup
=======
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	Task_borServiceGroup task_bor.ServiceGroup
>>>>>>> yxy
}
