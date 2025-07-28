package device

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
	deviceReq "github.com/flipped-aurora/gin-vue-admin/server/model/device/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MonitorDeviceApi struct{}

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

	// 如果有视频流URL，进行视频审核
	if monitorDevice.StreamUrl != nil && *monitorDevice.StreamUrl != "" {
		global.GVA_LOG.Info("开始视频审核流程", zap.String("streamUrl", *monitorDevice.StreamUrl))

		//上传本地文件到腾讯云cos
		/*cos, err := middleware.UploadVideoToCOS(*monitorDevice.StreamUrl)
		if err != nil {
			return
		}*/
		// 1. 提交视频审核任务
		fmt.Println("提交视频审核任务...")
		taskId, err := middleware.SubmitVideoModerationTask(*monitorDevice.StreamUrl)
		if err != nil {
			global.GVA_LOG.Error("提交审核任务失败", zap.Error(err))
			// 不中断流程，继续创建设备
			global.GVA_LOG.Warn("视频审核失败，继续创建设备", zap.Error(err))
		} else {
			fmt.Printf("审核任务已提交，任务ID: %s\n", taskId)

			// 2. 等待审核完成（视频审核是异步过程，设置超时时间）
			fmt.Println("等待审核完成...")
			var result map[string]interface{}
			timeout := time.After(2 * time.Minute)     // 2分钟超时
			ticker := time.NewTicker(10 * time.Second) // 每10秒查询一次
			defer ticker.Stop()

			auditCompleted := false
			for !auditCompleted {
				select {
				case <-timeout:
					global.GVA_LOG.Warn("视频审核超时，继续创建设备")
					auditCompleted = true
				case <-ticker.C:
					result, err = middleware.QueryModerationResult(taskId)
					if err != nil {
						global.GVA_LOG.Error("查询审核结果失败", zap.Error(err))
						auditCompleted = true
						break
					}

					resp, ok := result["Response"].(map[string]interface{})
					if !ok {
						global.GVA_LOG.Error("无效的响应格式")
						auditCompleted = true
						break
					}

					// 任务状态：PENDING(处理中)、SUCCESS(成功)、FAILED(失败)
					if status, ok := resp["Status"].(string); ok {
						if status == "SUCCESS" {
							// 审核完成，处理结果
							global.GVA_LOG.Info("视频审核完成", zap.String("taskId", taskId))

							// 3. 处理审核结果
							fmt.Println("处理审核结果...")
							middleware.PrintModerationResult(result)

							// 输出完整JSON结果
							if resultJson, err := json.MarshalIndent(result, "", "  "); err == nil {
								fmt.Println("\n完整JSON结果:")
								fmt.Println(string(resultJson))

								// 将审核结果存储到设备信息中（可选）
								// monitorDevice.AuditResult = string(resultJson)
							}

							auditCompleted = true
						} else if status == "FAILED" {
							global.GVA_LOG.Error("审核任务失败", zap.Any("response", resp))
							auditCompleted = true
						} else {
							fmt.Println("审核中，10秒后再次查询...")
						}
					}
				}
			}
		}
	}

	err = monitorDeviceService.CreateMonitorDevice(ctx, &monitorDevice)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed("", "创建成功", c)
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
	err := monitorDeviceService.DeleteMonitorDevice(ctx, ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
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
	err := monitorDeviceService.DeleteMonitorDeviceByIds(ctx, IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
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
	err = monitorDeviceService.UpdateMonitorDevice(ctx, monitorDevice)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
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
	remonitorDevice, err := monitorDeviceService.GetMonitorDevice(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
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
	list, total, err := monitorDeviceService.GetMonitorDeviceInfoList(ctx, pageInfo)
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
