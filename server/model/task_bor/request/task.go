
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TaskSearch struct{
    request.PageInfo
    TaskName         *string    `json:"taskName" form:"taskName"`                   // 任务名称
    SearchContent    *string    `json:"searchContent" form:"searchContent"`         // 搜索内容
    WarnLevel        *string    `json:"warnLevel" form:"warnLevel"`                 // 告警级别
    TaskStatus       *string    `json:"taskStatus" form:"taskStatus"`               // 任务状态
    StartDate        *time.Time `json:"startDate" form:"startDate"`                 // 开始日期
    EndDate          *time.Time `json:"endDate" form:"endDate"`                     // 结束日期
    ActiveTab        *string    `json:"activeTab" form:"activeTab"`                 // 当前选中的tab
}
