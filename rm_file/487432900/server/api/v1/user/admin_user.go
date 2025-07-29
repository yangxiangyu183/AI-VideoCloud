package user

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/user"
    userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type AdminUserApi struct {}



// CreateAdminUser 创建管理厅员表
// @Tags AdminUser
// @Summary 创建管理厅员表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body user.AdminUser true "创建管理厅员表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /adminUser/createAdminUser [post]
func (adminUserApi *AdminUserApi) CreateAdminUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var adminUser user.AdminUser
	err := c.ShouldBindJSON(&adminUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    adminUser.CreatedBy = utils.GetUserID(c)
	err = adminUserService.CreateAdminUser(ctx,&adminUser)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteAdminUser 删除管理厅员表
// @Tags AdminUser
// @Summary 删除管理厅员表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body user.AdminUser true "删除管理厅员表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /adminUser/deleteAdminUser [delete]
func (adminUserApi *AdminUserApi) DeleteAdminUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := adminUserService.DeleteAdminUser(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAdminUserByIds 批量删除管理厅员表
// @Tags AdminUser
// @Summary 批量删除管理厅员表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /adminUser/deleteAdminUserByIds [delete]
func (adminUserApi *AdminUserApi) DeleteAdminUserByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := adminUserService.DeleteAdminUserByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAdminUser 更新管理厅员表
// @Tags AdminUser
// @Summary 更新管理厅员表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body user.AdminUser true "更新管理厅员表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /adminUser/updateAdminUser [put]
func (adminUserApi *AdminUserApi) UpdateAdminUser(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var adminUser user.AdminUser
	err := c.ShouldBindJSON(&adminUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    adminUser.UpdatedBy = utils.GetUserID(c)
	err = adminUserService.UpdateAdminUser(ctx,adminUser)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAdminUser 用id查询管理厅员表
// @Tags AdminUser
// @Summary 用id查询管理厅员表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询管理厅员表"
// @Success 200 {object} response.Response{data=user.AdminUser,msg=string} "查询成功"
// @Router /adminUser/findAdminUser [get]
func (adminUserApi *AdminUserApi) FindAdminUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	readminUser, err := adminUserService.GetAdminUser(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(readminUser, c)
}
// GetAdminUserList 分页获取管理厅员表列表
// @Tags AdminUser
// @Summary 分页获取管理厅员表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query userReq.AdminUserSearch true "分页获取管理厅员表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /adminUser/getAdminUserList [get]
func (adminUserApi *AdminUserApi) GetAdminUserList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo userReq.AdminUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := adminUserService.GetAdminUserInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetAdminUserPublic 不需要鉴权的管理厅员表接口
// @Tags AdminUser
// @Summary 不需要鉴权的管理厅员表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /adminUser/getAdminUserPublic [get]
func (adminUserApi *AdminUserApi) GetAdminUserPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    adminUserService.GetAdminUserPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的管理厅员表接口信息",
    }, "获取成功", c)
}
