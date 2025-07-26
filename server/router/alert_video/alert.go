package alert_video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AlertRouter struct {}

// InitAlertRouter 初始化 alert表 路由信息
func (s *AlertRouter) InitAlertRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	alertRouter := Router.Group("alert").Use(middleware.OperationRecord())
	alertRouterWithoutRecord := Router.Group("alert")
	alertRouterWithoutAuth := PublicRouter.Group("alert")
	{
		alertRouter.POST("createAlert", alertApi.CreateAlert)   // 新建alert表
		alertRouter.DELETE("deleteAlert", alertApi.DeleteAlert) // 删除alert表
		alertRouter.DELETE("deleteAlertByIds", alertApi.DeleteAlertByIds) // 批量删除alert表
		alertRouter.PUT("updateAlert", alertApi.UpdateAlert)    // 更新alert表
	}
	{
		alertRouterWithoutRecord.GET("findAlert", alertApi.FindAlert)        // 根据ID获取alert表
	}
	{
	    alertRouterWithoutAuth.GET("getAlertPublic", alertApi.GetAlertPublic)  // alert表开放接口
	}
}