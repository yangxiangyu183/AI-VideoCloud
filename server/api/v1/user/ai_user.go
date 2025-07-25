package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AiUserApi struct{}

// CreateAiUser 创建aiUser表
// @Tags AiUser
// @Summary 创建aiUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body user.AiUser true "创建aiUser表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /aiUser/createAiUser [post]
func (aiUserApi *AiUserApi) CreateAiUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var aiUser user.AiUser
	err := c.ShouldBindJSON(&aiUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	aiUser.CreatedBy = utils.GetUserID(c)
	err = aiUserService.CreateAiUser(ctx, &aiUser)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAiUser 删除aiUser表
// @Tags AiUser
// @Summary 删除aiUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body user.AiUser true "删除aiUser表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /aiUser/deleteAiUser [delete]
func (aiUserApi *AiUserApi) DeleteAiUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := aiUserService.DeleteAiUser(ctx, ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAiUserByIds 批量删除aiUser表
// @Tags AiUser
// @Summary 批量删除aiUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /aiUser/deleteAiUserByIds [delete]
func (aiUserApi *AiUserApi) DeleteAiUserByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := aiUserService.DeleteAiUserByIds(ctx, IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAiUser 更新aiUser表
// @Tags AiUser
// @Summary 更新aiUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body user.AiUser true "更新aiUser表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /aiUser/updateAiUser [put]
func (aiUserApi *AiUserApi) UpdateAiUser(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var aiUser user.AiUser
	err := c.ShouldBindJSON(&aiUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	aiUser.UpdatedBy = utils.GetUserID(c)
	err = aiUserService.UpdateAiUser(ctx, aiUser)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAiUser 用id查询aiUser表
// @Tags AiUser
// @Summary 用id查询aiUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询aiUser表"
// @Success 200 {object} response.Response{data=user.AiUser,msg=string} "查询成功"
// @Router /aiUser/findAiUser [get]
func (aiUserApi *AiUserApi) FindAiUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reaiUser, err := aiUserService.GetAiUser(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reaiUser, c)
}

// GetAiUserList 分页获取aiUser表列表
// @Tags AiUser
// @Summary 分页获取aiUser表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query userReq.AiUserSearch true "分页获取aiUser表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /aiUser/getAiUserList [get]
func (aiUserApi *AiUserApi) GetAiUserList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo userReq.AiUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := aiUserService.GetAiUserInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetAiUserPublic 不需要鉴权的aiUser表接口
// @Tags AiUser
// @Summary 不需要鉴权的aiUser表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /aiUser/getAiUserPublic [get]
func (aiUserApi *AiUserApi) GetAiUserPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	aiUserService.GetAiUserPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的aiUser表接口信息",
	}, "获取成功", c)
}
