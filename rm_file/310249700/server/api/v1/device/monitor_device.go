package device

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/device"
    deviceReq "github.com/flipped-aurora/gin-vue-admin/server/model/device/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type MonitorDeviceApi struct {}



// CreateMonitorDevice 创建monitorDevice表
// @Tags MonitorDevice
// @Summary 创建monitorDevice表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body device.MonitorDevice true "创建monitorDevice表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /monitorDevice/createMonitorDevice [post]
func (monitorDeviceApi *MonitorDeviceApi) CreateMonitorDevice(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var monitorDevice device.MonitorDevice
	err := c.ShouldBindJSON(&monitorDevice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    monitorDevice.CreatedBy = utils.GetUserID(c)
	err = monitorDeviceService.CreateMonitorDevice(ctx,&monitorDevice)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMonitorDevice 删除monitorDevice表
// @Tags MonitorDevice
// @Summary 删除monitorDevice表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body device.MonitorDevice true "删除monitorDevice表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /monitorDevice/deleteMonitorDevice [delete]
func (monitorDeviceApi *MonitorDeviceApi) DeleteMonitorDevice(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := monitorDeviceService.DeleteMonitorDevice(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMonitorDeviceByIds 批量删除monitorDevice表
// @Tags MonitorDevice
// @Summary 批量删除monitorDevice表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /monitorDevice/deleteMonitorDeviceByIds [delete]
func (monitorDeviceApi *MonitorDeviceApi) DeleteMonitorDeviceByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := monitorDeviceService.DeleteMonitorDeviceByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMonitorDevice 更新monitorDevice表
// @Tags MonitorDevice
// @Summary 更新monitorDevice表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body device.MonitorDevice true "更新monitorDevice表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /monitorDevice/updateMonitorDevice [put]
func (monitorDeviceApi *MonitorDeviceApi) UpdateMonitorDevice(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var monitorDevice device.MonitorDevice
	err := c.ShouldBindJSON(&monitorDevice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    monitorDevice.UpdatedBy = utils.GetUserID(c)
	err = monitorDeviceService.UpdateMonitorDevice(ctx,monitorDevice)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMonitorDevice 用id查询monitorDevice表
// @Tags MonitorDevice
// @Summary 用id查询monitorDevice表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询monitorDevice表"
// @Success 200 {object} response.Response{data=device.MonitorDevice,msg=string} "查询成功"
// @Router /monitorDevice/findMonitorDevice [get]
func (monitorDeviceApi *MonitorDeviceApi) FindMonitorDevice(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	remonitorDevice, err := monitorDeviceService.GetMonitorDevice(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(remonitorDevice, c)
}
// GetMonitorDeviceList 分页获取monitorDevice表列表
// @Tags MonitorDevice
// @Summary 分页获取monitorDevice表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query deviceReq.MonitorDeviceSearch true "分页获取monitorDevice表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /monitorDevice/getMonitorDeviceList [get]
func (monitorDeviceApi *MonitorDeviceApi) GetMonitorDeviceList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo deviceReq.MonitorDeviceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := monitorDeviceService.GetMonitorDeviceInfoList(ctx,pageInfo)
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

// GetMonitorDevicePublic 不需要鉴权的monitorDevice表接口
// @Tags MonitorDevice
// @Summary 不需要鉴权的monitorDevice表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /monitorDevice/getMonitorDevicePublic [get]
func (monitorDeviceApi *MonitorDeviceApi) GetMonitorDevicePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    monitorDeviceService.GetMonitorDevicePublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的monitorDevice表接口信息",
    }, "获取成功", c)
}
