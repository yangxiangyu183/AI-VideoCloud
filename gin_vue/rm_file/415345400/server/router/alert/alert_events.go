package alert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AlertEventsRouter struct {}

// InitAlertEventsRouter 初始化 alertEvents表 路由信息
func (s *AlertEventsRouter) InitAlertEventsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	alertEventsRouter := Router.Group("alertEvents").Use(middleware.OperationRecord())
	alertEventsRouterWithoutRecord := Router.Group("alertEvents")
	alertEventsRouterWithoutAuth := PublicRouter.Group("alertEvents")
	{
		alertEventsRouter.POST("createAlertEvents", alertEventsApi.CreateAlertEvents)   // 新建alertEvents表
		alertEventsRouter.DELETE("deleteAlertEvents", alertEventsApi.DeleteAlertEvents) // 删除alertEvents表
		alertEventsRouter.DELETE("deleteAlertEventsByIds", alertEventsApi.DeleteAlertEventsByIds) // 批量删除alertEvents表
		alertEventsRouter.PUT("updateAlertEvents", alertEventsApi.UpdateAlertEvents)    // 更新alertEvents表
	}
	{
		alertEventsRouterWithoutRecord.GET("findAlertEvents", alertEventsApi.FindAlertEvents)        // 根据ID获取alertEvents表
		alertEventsRouterWithoutRecord.GET("getAlertEventsList", alertEventsApi.GetAlertEventsList)  // 获取alertEvents表列表
	}
	{
	    alertEventsRouterWithoutAuth.GET("getAlertEventsPublic", alertEventsApi.GetAlertEventsPublic)  // alertEvents表开放接口
	}
}
