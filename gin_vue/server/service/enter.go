package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/alert"
	"github.com/flipped-aurora/gin-vue-admin/server/service/alert_video"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup      system.ServiceGroup
	ExampleServiceGroup     example.ServiceGroup
	AlertServiceGroup       alert.ServiceGroup
	Alert_videoServiceGroup alert_video.ServiceGroup
}
