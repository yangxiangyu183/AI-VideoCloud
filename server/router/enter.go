package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/device"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
<<<<<<< HEAD
	"github.com/flipped-aurora/gin-vue-admin/server/router/user"
=======
	"github.com/flipped-aurora/gin-vue-admin/server/router/task_bor"
>>>>>>> hjy
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
<<<<<<< HEAD
	System  system.RouterGroup
	Example example.RouterGroup
<<<<<<< HEAD
	User    user.RouterGroup
=======
	Device  device.RouterGroup
=======
	System   system.RouterGroup
	Example  example.RouterGroup
	Task_bor task_bor.RouterGroup
>>>>>>> yxy
>>>>>>> hjy
}
