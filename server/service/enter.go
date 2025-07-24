package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/device"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
<<<<<<< HEAD
	"github.com/flipped-aurora/gin-vue-admin/server/service/user"
=======
	"github.com/flipped-aurora/gin-vue-admin/server/service/task_bor"
>>>>>>> hjy
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
<<<<<<< HEAD
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
<<<<<<< HEAD
	UserServiceGroup    user.ServiceGroup
=======
	DeviceServiceGroup  device.ServiceGroup
=======
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	Task_borServiceGroup task_bor.ServiceGroup
>>>>>>> yxy
>>>>>>> hjy
}
