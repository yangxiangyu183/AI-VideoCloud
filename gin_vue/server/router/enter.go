package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/alert"
	"github.com/flipped-aurora/gin-vue-admin/server/router/alert_video"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System      system.RouterGroup
	Example     example.RouterGroup
	Alert       alert.RouterGroup
	Alert_video alert_video.RouterGroup
}
