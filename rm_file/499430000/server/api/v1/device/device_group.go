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

type DeviceGroupApi struct {}



// CreateDeviceGroup 创建deviceGroup表
// @Tags DeviceGroup
// @Summary 创建deviceGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body device.DeviceGroup true "创建deviceGroup表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /deviceGroup/createDeviceGroup [post]
func (deviceGroupApi *DeviceGroupApi) CreateDeviceGroup(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var deviceGroup device.DeviceGroup
	err := c.ShouldBindJSON(&deviceGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deviceGroup.CreatedBy = utils.GetUserID(c)
	err = deviceGroupService.CreateDeviceGroup(ctx,&deviceGroup)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteDeviceGroup 删除deviceGroup表
// @Tags DeviceGroup
// @Summary 删除deviceGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body device.DeviceGroup true "删除deviceGroup表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /deviceGroup/deleteDeviceGroup [delete]
func (deviceGroupApi *DeviceGroupApi) DeleteDeviceGroup(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := deviceGroupService.DeleteDeviceGroup(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteDeviceGroupByIds 批量删除deviceGroup表
// @Tags DeviceGroup
// @Summary 批量删除deviceGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /deviceGroup/deleteDeviceGroupByIds [delete]
func (deviceGroupApi *DeviceGroupApi) DeleteDeviceGroupByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := deviceGroupService.DeleteDeviceGroupByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateDeviceGroup 更新deviceGroup表
// @Tags DeviceGroup
// @Summary 更新deviceGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body device.DeviceGroup true "更新deviceGroup表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /deviceGroup/updateDeviceGroup [put]
func (deviceGroupApi *DeviceGroupApi) UpdateDeviceGroup(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var deviceGroup device.DeviceGroup
	err := c.ShouldBindJSON(&deviceGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deviceGroup.UpdatedBy = utils.GetUserID(c)
	err = deviceGroupService.UpdateDeviceGroup(ctx,deviceGroup)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindDeviceGroup 用id查询deviceGroup表
// @Tags DeviceGroup
// @Summary 用id查询deviceGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询deviceGroup表"
// @Success 200 {object} response.Response{data=device.DeviceGroup,msg=string} "查询成功"
// @Router /deviceGroup/findDeviceGroup [get]
func (deviceGroupApi *DeviceGroupApi) FindDeviceGroup(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	redeviceGroup, err := deviceGroupService.GetDeviceGroup(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(redeviceGroup, c)
}
// GetDeviceGroupList 分页获取deviceGroup表列表
// @Tags DeviceGroup
// @Summary 分页获取deviceGroup表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query deviceReq.DeviceGroupSearch true "分页获取deviceGroup表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /deviceGroup/getDeviceGroupList [get]
func (deviceGroupApi *DeviceGroupApi) GetDeviceGroupList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo deviceReq.DeviceGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := deviceGroupService.GetDeviceGroupInfoList(ctx,pageInfo)
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

// GetDeviceGroupPublic 不需要鉴权的deviceGroup表接口
// @Tags DeviceGroup
// @Summary 不需要鉴权的deviceGroup表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /deviceGroup/getDeviceGroupPublic [get]
func (deviceGroupApi *DeviceGroupApi) GetDeviceGroupPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    deviceGroupService.GetDeviceGroupPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的deviceGroup表接口信息",
    }, "获取成功", c)
}
