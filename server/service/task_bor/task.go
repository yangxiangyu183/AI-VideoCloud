
package task_bor

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/task_bor"
    task_borReq "github.com/flipped-aurora/gin-vue-admin/server/model/task_bor/request"
)

type TaskService struct {}
// CreateTask 创建监测任务记录
// Author [yourname](https://github.com/yourname)
func (task表Service *TaskService) CreateTask(ctx context.Context, task表 *task_bor.Task) (err error) {
	err = global.GVA_DB.Create(task表).Error
	return err
}

// DeleteTask 删除监测任务记录
// Author [yourname](https://github.com/yourname)
func (task表Service *TaskService)DeleteTask(ctx context.Context, taskName string) (err error) {
	err = global.GVA_DB.Delete(&task_bor.Task{},"task_name = ?",taskName).Error
	return err
}

// DeleteTaskByIds 批量删除监测任务记录
// Author [yourname](https://github.com/yourname)
func (task表Service *TaskService)DeleteTaskByIds(ctx context.Context, taskNames []string) (err error) {
	err = global.GVA_DB.Delete(&[]task_bor.Task{},"task_name in ?",taskNames).Error
	return err
}

// UpdateTask 更新监测任务记录
// Author [yourname](https://github.com/yourname)
func (task表Service *TaskService)UpdateTask(ctx context.Context, task表 task_bor.Task) (err error) {
	err = global.GVA_DB.Model(&task_bor.Task{}).Where("task_name = ?",task表.TaskName).Updates(&task表).Error
	return err
}

// GetTask 根据taskName获取监测任务记录
// Author [yourname](https://github.com/yourname)
func (task表Service *TaskService)GetTask(ctx context.Context, taskName string) (task表 task_bor.Task, err error) {
	err = global.GVA_DB.Where("task_name = ?", taskName).First(&task表).Error
	return
}
// GetTaskInfoList 分页获取监测任务记录
// Author [yourname](https://github.com/yourname)
func (task表Service *TaskService)GetTaskInfoList(ctx context.Context, info task_borReq.TaskSearch) (list []task_bor.Task, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&task_bor.Task{})
    var task表s []task_bor.Task
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&task表s).Error
	return  task表s, total, err
}
func (task表Service *TaskService)GetTaskPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
