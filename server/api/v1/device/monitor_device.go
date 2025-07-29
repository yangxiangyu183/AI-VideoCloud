package device

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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

// downloadAndUploadVideoStream 从视频流URL下载视频内容并上传到MinIO
func (monitorDeviceApi *MonitorDeviceApi) downloadAndUploadVideoStream(ctx context.Context, streamUrl string) (string, error) {
	global.GVA_LOG.Info("开始从视频流下载内容", zap.String("streamUrl", streamUrl))

	var videoData []byte
	var contentType string
	var err error

	// 检查是否为HTTP/HTTPS URL
	if parsedURL, err := url.Parse(streamUrl); err == nil && (parsedURL.Scheme == "http" || parsedURL.Scheme == "https") {
		// 处理HTTP/HTTPS URL
		global.GVA_LOG.Info("检测到HTTP/HTTPS URL，开始下载")

		// 创建HTTP客户端，设置超时时间
		client := &http.Client{
			Timeout: 30 * time.Second,
		}

		// 发起HTTP请求获取视频流
		req, err := http.NewRequestWithContext(ctx, "GET", streamUrl, nil)
		if err != nil {
			return "", fmt.Errorf("创建请求失败: %v", err)
		}

		// 设置请求头，模拟视频播放器
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
		req.Header.Set("Accept", "video/mp4,video/*,*/*")

		resp, err := client.Do(req)
		if err != nil {
			return "", fmt.Errorf("请求视频流失败: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("视频流响应状态码错误: %d", resp.StatusCode)
		}

		// 检查内容类型
		contentType = resp.Header.Get("Content-Type")
		global.GVA_LOG.Info("视频流内容类型", zap.String("contentType", contentType))

		// 限制下载大小（例如最大100MB）
		const maxSize = 100 * 1024 * 1024 // 100MB
		limitedReader := io.LimitReader(resp.Body, maxSize)

		// 读取视频内容到内存
		videoData, err = io.ReadAll(limitedReader)
		if err != nil {
			return "", fmt.Errorf("读取视频流数据失败: %v", err)
		}
	} else {
		// 处理本地文件路径
		global.GVA_LOG.Info("检测到本地文件路径，开始读取文件")

		// 检查文件是否存在
		if _, err := os.Stat(streamUrl); os.IsNotExist(err) {
			return "", fmt.Errorf("文件不存在: %s", streamUrl)
		}

		// 读取本地文件
		videoData, err = os.ReadFile(streamUrl)
		if err != nil {
			return "", fmt.Errorf("读取本地文件失败: %v", err)
		}

		// 根据文件扩展名确定内容类型
		ext := strings.ToLower(filepath.Ext(streamUrl))
		switch ext {
		case ".mp4":
			contentType = "video/mp4"
		case ".avi":
			contentType = "video/avi"
		case ".mov":
			contentType = "video/quicktime"
		case ".wmv":
			contentType = "video/x-ms-wmv"
		case ".flv":
			contentType = "video/x-flv"
		case ".webm":
			contentType = "video/webm"
		case ".jpg", ".jpeg":
			contentType = "image/jpeg"
		case ".png":
			contentType = "image/png"
		case ".webp":
			contentType = "image/webp"
		default:
			contentType = "application/octet-stream"
		}

		global.GVA_LOG.Info("本地文件内容类型", zap.String("contentType", contentType))
	}

	if len(videoData) == 0 {
		return "", fmt.Errorf("视频数据为空")
	}

	global.GVA_LOG.Info("视频内容读取完成", zap.Int("size", len(videoData)))

	// 生成文件名
	timestamp := time.Now().Unix()
	var fileName string

	// 根据内容类型确定文件扩展名
	if strings.Contains(contentType, "video/mp4") {
		fileName = fmt.Sprintf("monitor_video_%d.mp4", timestamp)
	} else if strings.Contains(contentType, "video/") {
		fileName = fmt.Sprintf("monitor_video_%d.mp4", timestamp) // 默认使用mp4
	} else if strings.Contains(contentType, "image/") {
		// 处理图片文件
		if strings.Contains(contentType, "jpeg") {
			fileName = fmt.Sprintf("monitor_image_%d.jpg", timestamp)
		} else if strings.Contains(contentType, "png") {
			fileName = fmt.Sprintf("monitor_image_%d.png", timestamp)
		} else if strings.Contains(contentType, "webp") {
			fileName = fmt.Sprintf("monitor_image_%d.webp", timestamp)
		} else {
			fileName = fmt.Sprintf("monitor_image_%d.jpg", timestamp)
		}
	} else {
		// 如果不是视频或图片类型
		fileName = fmt.Sprintf("monitor_content_%d.bin", timestamp)
	}

	// 上传到MinIO
	reader := bytes.NewReader(videoData)
	uploadInfo, err := middleware.Upload(fileName, reader, int64(len(videoData)))
	if err != nil {
		return "", fmt.Errorf("上传到MinIO失败: %v", err)
	}

	// 构建MinIO访问URL
	minioURL := fmt.Sprintf("http://14.103.149.194:9000/test/%s", fileName)

	global.GVA_LOG.Info("视频流上传到MinIO成功",
		zap.String("fileName", fileName),
		zap.String("location", uploadInfo.Location),
		zap.String("minioURL", minioURL))

	return minioURL, nil
}

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

		// 保存原始视频流URL用于下载
		originalStreamUrl := *monitorDevice.StreamUrl

		// 1. 先将监控设备视频流上传到MinIO
		minioURL, err := monitorDeviceApi.downloadAndUploadVideoStream(ctx, originalStreamUrl)
		if err != nil {
			global.GVA_LOG.Error("上传视频流到MinIO失败", zap.Error(err))
			// 不中断流程，继续创建设备，但记录错误
			global.GVA_LOG.Warn("视频流上传失败，继续创建设备", zap.Error(err))
		} else {
			global.GVA_LOG.Info("视频流上传到MinIO成功", zap.String("minioURL", minioURL))

			// 将MinIO URL存储到数据库的stream_url字段中
			monitorDevice.StreamUrl = &minioURL
			global.GVA_LOG.Info("已将MinIO URL设置为设备的stream_url",
				zap.String("originalUrl", originalStreamUrl),
				zap.String("minioUrl", minioURL))

			// 2. 提交视频审核任务（暂时跳过，因为腾讯云无法访问MinIO URL）

			taskId, err := middleware.SubmitVideoModerationTask(minioURL)
			if err != nil {
				global.GVA_LOG.Error("提交审核任务失败",
					zap.Error(err),
					zap.String("minioURL", minioURL))
				// 不中断流程，继续创建设备
				global.GVA_LOG.Warn("视频审核失败，继续创建设备", zap.Error(err))
			} else {
				fmt.Printf("审核任务已提交，任务ID: %s\n", taskId)

				// 3. 等待审核完成（视频审核是异步过程，设置超时时间）
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

								// 4. 处理审核结果
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
