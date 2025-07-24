
// 自动生成模板Task
package task_bor
import (
	"time"
)

// 监测任务 结构体  Task
type Task struct {
  TaskName  *string `json:"taskName" form:"taskName" gorm:"primarykey;comment:任务名称;column:task_name;size:20;"`  //任务名称
  TaskDescription  *string `json:"taskDescription" form:"taskDescription" gorm:"comment:任务描述;column:task_description;size:250;"`  //任务描述
  TaskPriority  *string `json:"taskPriority" form:"taskPriority" gorm:"comment:任务优先级(1.高，2.中，3.低);column:task_priority;size:10;"`  //任务优先级
  CameraInterface  *string `json:"cameraInterface" form:"cameraInterface" gorm:"comment:摄像头接口;column:camera_interface;size:20;"`  //摄像头接口
  MonitorPoints  *string `json:"monitorPoints" form:"monitorPoints" gorm:"comment:监控点位(1.模拟环境，2.模拟环境，3.模拟环境);column:monitor_points;size:10;"`  //监控点位
  ApplicationScenario  *string `json:"applicationScenario" form:"applicationScenario" gorm:"comment:应用场景(1.佩戴安全帽，2.人体跌倒识别，3.电动车监测，4.大货车监测，5.房门关开监测，6.灭火器正常，7.地面垃圾);column:application_scenario;size:10;"`  //应用场景
  TaskStatus  *string `json:"taskStatus" form:"taskStatus" gorm:"comment:任务状态(1.启用，2.停用);column:task_status;size:10;"`  //任务状态
  WarnLevel  *string `json:"warnLevel" form:"warnLevel" gorm:"comment:警告级别;column:warn_level;size:20;"`  //警告级别
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`  //创建时间
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`  //监测日期
  DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"comment:删除时间;column:deleted_at;"`  //监测轮询时长
}


// TableName 监测任务 Task自定义表名 task
func (Task) TableName() string {
    return "task"
}





