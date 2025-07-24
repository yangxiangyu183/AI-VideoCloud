package task_bor

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/task_bor"
    task_borReq "github.com/flipped-aurora/gin-vue-admin/server/model/task_bor/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TaskApi struct {}



// CreateTask 创建监测任务
// @Tags Task
// @Summary 创建监测任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body task_bor.Task true "创建监测任务"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /task表/createTask [post]
func (task表Api *TaskApi) CreateTask(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var task表 task_bor.Task
	err := c.ShouldBindJSON(&task表)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = task表Service.CreateTask(ctx,&task表)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteTask 删除监测任务
// @Tags Task
// @Summary 删除监测任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body task_bor.Task true "删除监测任务"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /task表/deleteTask [delete]
func (task表Api *TaskApi) DeleteTask(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	taskName := c.Query("taskName")
	err := task表Service.DeleteTask(ctx,taskName)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteTaskByIds 批量删除监测任务
// @Tags Task
// @Summary 批量删除监测任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /task表/deleteTaskByIds [delete]
func (task表Api *TaskApi) DeleteTaskByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	taskNames := c.QueryArray("taskNames[]")
	err := task表Service.DeleteTaskByIds(ctx,taskNames)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTask 更新监测任务
// @Tags Task
// @Summary 更新监测任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body task_bor.Task true "更新监测任务"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /task表/updateTask [put]
func (task表Api *TaskApi) UpdateTask(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var task表 task_bor.Task
	err := c.ShouldBindJSON(&task表)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = task表Service.UpdateTask(ctx,task表)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTask 用id查询监测任务
// @Tags Task
// @Summary 用id查询监测任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param taskName query string true "用id查询监测任务"
// @Success 200 {object} response.Response{data=task_bor.Task,msg=string} "查询成功"
// @Router /task表/findTask [get]
func (task表Api *TaskApi) FindTask(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	taskName := c.Query("taskName")
	retask表, err := task表Service.GetTask(ctx,taskName)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(retask表, c)
}
// GetTaskList 分页获取监测任务列表
// @Tags Task
// @Summary 分页获取监测任务列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query task_borReq.TaskSearch true "分页获取监测任务列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /task表/getTaskList [get]
func (task表Api *TaskApi) GetTaskList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo task_borReq.TaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := task表Service.GetTaskInfoList(ctx,pageInfo)
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

// GetTaskPublic 不需要鉴权的监测任务接口
// @Tags Task
// @Summary 不需要鉴权的监测任务接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /task表/getTaskPublic [get]
func (task表Api *TaskApi) GetTaskPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    task表Service.GetTaskPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的监测任务接口信息",
    }, "获取成功", c)
}
