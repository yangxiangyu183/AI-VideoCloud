package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/alert"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/alert_video"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup      system.ApiGroup
	ExampleApiGroup     example.ApiGroup
	AlertApiGroup       alert.ApiGroup
	Alert_videoApiGroup alert_video.ApiGroup
}
