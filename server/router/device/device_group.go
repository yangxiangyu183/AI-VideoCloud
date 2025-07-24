package device

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DeviceGroupRouter struct {}

// InitDeviceGroupRouter 初始化 deviceGroup表 路由信息
func (s *DeviceGroupRouter) InitDeviceGroupRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	deviceGroupRouter := Router.Group("deviceGroup").Use(middleware.OperationRecord())
	deviceGroupRouterWithoutRecord := Router.Group("deviceGroup")
	deviceGroupRouterWithoutAuth := PublicRouter.Group("deviceGroup")
	{
		deviceGroupRouter.POST("createDeviceGroup", deviceGroupApi.CreateDeviceGroup)   // 新建deviceGroup表
		deviceGroupRouter.DELETE("deleteDeviceGroup", deviceGroupApi.DeleteDeviceGroup) // 删除deviceGroup表
		deviceGroupRouter.DELETE("deleteDeviceGroupByIds", deviceGroupApi.DeleteDeviceGroupByIds) // 批量删除deviceGroup表
		deviceGroupRouter.PUT("updateDeviceGroup", deviceGroupApi.UpdateDeviceGroup)    // 更新deviceGroup表
	}
	{
		deviceGroupRouterWithoutRecord.GET("findDeviceGroup", deviceGroupApi.FindDeviceGroup)        // 根据ID获取deviceGroup表
		deviceGroupRouterWithoutRecord.GET("getDeviceGroupList", deviceGroupApi.GetDeviceGroupList)  // 获取deviceGroup表列表
	}
	{
	    deviceGroupRouterWithoutAuth.GET("getDeviceGroupPublic", deviceGroupApi.GetDeviceGroupPublic)  // deviceGroup表开放接口
	}
}
