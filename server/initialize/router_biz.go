package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		systemRouter := router.RouterGroupApp.System
		systemRouter.InitCasbinRuleRouter(privateGroup, publicGroup)
	} // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
	{
		task_borRouter := router.RouterGroupApp.Task_bor
		task_borRouter.InitTaskRouter(privateGroup, publicGroup)
	}
	{
		deviceRouter := router.RouterGroupApp.Device
		deviceRouter.InitDeviceGroupRouter(privateGroup, publicGroup) // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
		deviceRouter.InitMonitorDeviceRouter(privateGroup, publicGroup)
	}
	{
		userRouter := router.RouterGroupApp.User
		userRouter.InitAiUserRouter(privateGroup, publicGroup)
		userRouter.InitAdminUserRouter(privateGroup, publicGroup)
	}
}
