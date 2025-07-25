package device

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MonitorDeviceRouter struct {}

// InitMonitorDeviceRouter 初始化 monitorDevice表 路由信息
func (s *MonitorDeviceRouter) InitMonitorDeviceRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	monitorDeviceRouter := Router.Group("monitorDevice").Use(middleware.OperationRecord())
	monitorDeviceRouterWithoutRecord := Router.Group("monitorDevice")
	monitorDeviceRouterWithoutAuth := PublicRouter.Group("monitorDevice")
	{
		monitorDeviceRouter.POST("createMonitorDevice", monitorDeviceApi.CreateMonitorDevice)   // 新建monitorDevice表
		monitorDeviceRouter.DELETE("deleteMonitorDevice", monitorDeviceApi.DeleteMonitorDevice) // 删除monitorDevice表
		monitorDeviceRouter.DELETE("deleteMonitorDeviceByIds", monitorDeviceApi.DeleteMonitorDeviceByIds) // 批量删除monitorDevice表
		monitorDeviceRouter.PUT("updateMonitorDevice", monitorDeviceApi.UpdateMonitorDevice)    // 更新monitorDevice表
	}
	{
		monitorDeviceRouterWithoutRecord.GET("findMonitorDevice", monitorDeviceApi.FindMonitorDevice)        // 根据ID获取monitorDevice表
		monitorDeviceRouterWithoutRecord.GET("getMonitorDeviceList", monitorDeviceApi.GetMonitorDeviceList)  // 获取monitorDevice表列表
	}
	{
	    monitorDeviceRouterWithoutAuth.GET("getMonitorDevicePublic", monitorDeviceApi.GetMonitorDevicePublic)  // monitorDevice表开放接口
	}
}
