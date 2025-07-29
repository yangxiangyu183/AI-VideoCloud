package alert_video

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/alert_video"
    alert_videoReq "github.com/flipped-aurora/gin-vue-admin/server/model/alert_video/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AlertApi struct {}



// CreateAlert 创建alert表
// @Tags Alert
// @Summary 创建alert表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body alert_video.Alert true "创建alert表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /alert/createAlert [post]
func (alertApi *AlertApi) CreateAlert(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var alert alert_video.Alert
	err := c.ShouldBindJSON(&alert)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = alertService.CreateAlert(ctx,&alert)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteAlert 删除alert表
// @Tags Alert
// @Summary 删除alert表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body alert_video.Alert true "删除alert表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /alert/deleteAlert [delete]
func (alertApi *AlertApi) DeleteAlert(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := alertService.DeleteAlert(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAlertByIds 批量删除alert表
// @Tags Alert
// @Summary 批量删除alert表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /alert/deleteAlertByIds [delete]
func (alertApi *AlertApi) DeleteAlertByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := alertService.DeleteAlertByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAlert 更新alert表
// @Tags Alert
// @Summary 更新alert表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body alert_video.Alert true "更新alert表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /alert/updateAlert [put]
func (alertApi *AlertApi) UpdateAlert(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var alert alert_video.Alert
	err := c.ShouldBindJSON(&alert)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = alertService.UpdateAlert(ctx,alert)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAlert 用id查询alert表
// @Tags Alert
// @Summary 用id查询alert表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询alert表"
// @Success 200 {object} response.Response{data=alert_video.Alert,msg=string} "查询成功"
// @Router /alert/findAlert [get]
func (alertApi *AlertApi) FindAlert(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	realert, err := alertService.GetAlert(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(realert, c)
}
// GetAlertList 分页获取alert表列表
// @Tags Alert
// @Summary 分页获取alert表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query alert_videoReq.AlertSearch true "分页获取alert表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /alert/getAlertList [get]
func (alertApi *AlertApi) GetAlertList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo alert_videoReq.AlertSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := alertService.GetAlertInfoList(ctx,pageInfo)
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

// GetAlertPublic 不需要鉴权的alert表接口
// @Tags Alert
// @Summary 不需要鉴权的alert表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /alert/getAlertPublic [get]
func (alertApi *AlertApi) GetAlertPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    alertService.GetAlertPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的alert表接口信息",
    }, "获取成功", c)
}
