package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/alert_video"
	"github.com/flipped-aurora/gin-vue-admin/server/router/device"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/task_bor"
	"github.com/flipped-aurora/gin-vue-admin/server/router/user"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System      system.RouterGroup
	Example     example.RouterGroup
	Device      device.RouterGroup
	User        user.RouterGroup
	Task_bor    task_bor.RouterGroup
	Alert_video alert_video.RouterGroup
}
