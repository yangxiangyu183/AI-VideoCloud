package task_bor

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TaskRouter struct {}

// InitTaskRouter 初始化 监测任务 路由信息
func (s *TaskRouter) InitTaskRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	task表Router := Router.Group("task表").Use(middleware.OperationRecord())
	task表RouterWithoutRecord := Router.Group("task表")
	task表RouterWithoutAuth := PublicRouter.Group("task表")
	{
		task表Router.POST("createTask", task表Api.CreateTask)   // 新建监测任务
		task表Router.DELETE("deleteTask", task表Api.DeleteTask) // 删除监测任务
		task表Router.DELETE("deleteTaskByIds", task表Api.DeleteTaskByIds) // 批量删除监测任务
		task表Router.PUT("updateTask", task表Api.UpdateTask)    // 更新监测任务
	}
	{
		task表RouterWithoutRecord.GET("findTask", task表Api.FindTask)        // 根据ID获取监测任务
		task表RouterWithoutRecord.GET("getTaskList", task表Api.GetTaskList)  // 获取监测任务列表
	}
	{
	    task表RouterWithoutAuth.GET("getTaskPublic", task表Api.GetTaskPublic)  // 监测任务开放接口
	}
}
