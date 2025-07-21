package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AiUserRouter struct {}

// InitAiUserRouter 初始化 aiUser表 路由信息
func (s *AiUserRouter) InitAiUserRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	aiUserRouter := Router.Group("aiUser").Use(middleware.OperationRecord())
	aiUserRouterWithoutRecord := Router.Group("aiUser")
	aiUserRouterWithoutAuth := PublicRouter.Group("aiUser")
	{
		aiUserRouter.POST("createAiUser", aiUserApi.CreateAiUser)   // 新建aiUser表
		aiUserRouter.DELETE("deleteAiUser", aiUserApi.DeleteAiUser) // 删除aiUser表
		aiUserRouter.DELETE("deleteAiUserByIds", aiUserApi.DeleteAiUserByIds) // 批量删除aiUser表
		aiUserRouter.PUT("updateAiUser", aiUserApi.UpdateAiUser)    // 更新aiUser表
	}
	{
		aiUserRouterWithoutRecord.GET("findAiUser", aiUserApi.FindAiUser)        // 根据ID获取aiUser表
		aiUserRouterWithoutRecord.GET("getAiUserList", aiUserApi.GetAiUserList)  // 获取aiUser表列表
	}
	{
	    aiUserRouterWithoutAuth.GET("getAiUserPublic", aiUserApi.GetAiUserPublic)  // aiUser表开放接口
	}
}
