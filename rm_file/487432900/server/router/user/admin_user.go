package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AdminUserRouter struct {}

// InitAdminUserRouter 初始化 管理厅员表 路由信息
func (s *AdminUserRouter) InitAdminUserRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	adminUserRouter := Router.Group("adminUser").Use(middleware.OperationRecord())
	adminUserRouterWithoutRecord := Router.Group("adminUser")
	adminUserRouterWithoutAuth := PublicRouter.Group("adminUser")
	{
		adminUserRouter.POST("createAdminUser", adminUserApi.CreateAdminUser)   // 新建管理厅员表
		adminUserRouter.DELETE("deleteAdminUser", adminUserApi.DeleteAdminUser) // 删除管理厅员表
		adminUserRouter.DELETE("deleteAdminUserByIds", adminUserApi.DeleteAdminUserByIds) // 批量删除管理厅员表
		adminUserRouter.PUT("updateAdminUser", adminUserApi.UpdateAdminUser)    // 更新管理厅员表
	}
	{
		adminUserRouterWithoutRecord.GET("findAdminUser", adminUserApi.FindAdminUser)        // 根据ID获取管理厅员表
		adminUserRouterWithoutRecord.GET("getAdminUserList", adminUserApi.GetAdminUserList)  // 获取管理厅员表列表
	}
	{
	    adminUserRouterWithoutAuth.GET("getAdminUserPublic", adminUserApi.GetAdminUserPublic)  // 管理厅员表开放接口
	}
}
