package alert

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/alert"
    alertReq "github.com/flipped-aurora/gin-vue-admin/server/model/alert/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type AlertEventsApi struct {}



// CreateAlertEvents 创建alertEvents表
// @Tags AlertEvents
// @Summary 创建alertEvents表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body alert.AlertEvents true "创建alertEvents表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /alertEvents/createAlertEvents [post]
func (alertEventsApi *AlertEventsApi) CreateAlertEvents(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var alertEvents alert.AlertEvents
	err := c.ShouldBindJSON(&alertEvents)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    alertEvents.CreatedBy = utils.GetUserID(c)
	err = alertEventsService.CreateAlertEvents(ctx,&alertEvents)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteAlertEvents 删除alertEvents表
// @Tags AlertEvents
// @Summary 删除alertEvents表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body alert.AlertEvents true "删除alertEvents表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /alertEvents/deleteAlertEvents [delete]
func (alertEventsApi *AlertEventsApi) DeleteAlertEvents(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := alertEventsService.DeleteAlertEvents(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAlertEventsByIds 批量删除alertEvents表
// @Tags AlertEvents
// @Summary 批量删除alertEvents表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /alertEvents/deleteAlertEventsByIds [delete]
func (alertEventsApi *AlertEventsApi) DeleteAlertEventsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := alertEventsService.DeleteAlertEventsByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAlertEvents 更新alertEvents表
// @Tags AlertEvents
// @Summary 更新alertEvents表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body alert.AlertEvents true "更新alertEvents表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /alertEvents/updateAlertEvents [put]
func (alertEventsApi *AlertEventsApi) UpdateAlertEvents(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var alertEvents alert.AlertEvents
	err := c.ShouldBindJSON(&alertEvents)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    alertEvents.UpdatedBy = utils.GetUserID(c)
	err = alertEventsService.UpdateAlertEvents(ctx,alertEvents)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAlertEvents 用id查询alertEvents表
// @Tags AlertEvents
// @Summary 用id查询alertEvents表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询alertEvents表"
// @Success 200 {object} response.Response{data=alert.AlertEvents,msg=string} "查询成功"
// @Router /alertEvents/findAlertEvents [get]
func (alertEventsApi *AlertEventsApi) FindAlertEvents(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	realertEvents, err := alertEventsService.GetAlertEvents(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(realertEvents, c)
}
// GetAlertEventsList 分页获取alertEvents表列表
// @Tags AlertEvents
// @Summary 分页获取alertEvents表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query alertReq.AlertEventsSearch true "分页获取alertEvents表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /alertEvents/getAlertEventsList [get]
func (alertEventsApi *AlertEventsApi) GetAlertEventsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo alertReq.AlertEventsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := alertEventsService.GetAlertEventsInfoList(ctx,pageInfo)
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

// GetAlertEventsPublic 不需要鉴权的alertEvents表接口
// @Tags AlertEvents
// @Summary 不需要鉴权的alertEvents表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /alertEvents/getAlertEventsPublic [get]
func (alertEventsApi *AlertEventsApi) GetAlertEventsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    alertEventsService.GetAlertEventsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的alertEvents表接口信息",
    }, "获取成功", c)
}
