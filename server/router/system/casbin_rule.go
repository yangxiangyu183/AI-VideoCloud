package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CasbinRuleRouter struct {}

// InitCasbinRuleRouter 初始化 casbinRule表 路由信息
func (s *CasbinRuleRouter) InitCasbinRuleRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	casbinRuleRouter := Router.Group("casbinRule").Use(middleware.OperationRecord())
	casbinRuleRouterWithoutRecord := Router.Group("casbinRule")
	casbinRuleRouterWithoutAuth := PublicRouter.Group("casbinRule")
	{
		casbinRuleRouter.POST("createCasbinRule", casbinRuleApi.CreateCasbinRule)   // 新建casbinRule表
		casbinRuleRouter.DELETE("deleteCasbinRule", casbinRuleApi.DeleteCasbinRule) // 删除casbinRule表
		casbinRuleRouter.DELETE("deleteCasbinRuleByIds", casbinRuleApi.DeleteCasbinRuleByIds) // 批量删除casbinRule表
		casbinRuleRouter.PUT("updateCasbinRule", casbinRuleApi.UpdateCasbinRule)    // 更新casbinRule表
	}
	{
		casbinRuleRouterWithoutRecord.GET("findCasbinRule", casbinRuleApi.FindCasbinRule)        // 根据ID获取casbinRule表
		casbinRuleRouterWithoutRecord.GET("getCasbinRuleList", casbinRuleApi.GetCasbinRuleList)  // 获取casbinRule表列表
	}
	{
	    casbinRuleRouterWithoutAuth.GET("getCasbinRulePublic", casbinRuleApi.GetCasbinRulePublic)  // casbinRule表开放接口
	}
}
