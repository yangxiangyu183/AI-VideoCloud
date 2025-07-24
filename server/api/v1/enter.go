package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/device"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/task_bor"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
<<<<<<< HEAD
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	DeviceApiGroup  device.ApiGroup
=======
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	Task_borApiGroup task_bor.ApiGroup
>>>>>>> yxy
}
