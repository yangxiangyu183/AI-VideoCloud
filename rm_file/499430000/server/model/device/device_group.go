
// 自动生成模板DeviceGroup
package device
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// deviceGroup表 结构体  DeviceGroup
type DeviceGroup struct {
    global.GVA_MODEL
  GroupName  *string `json:"groupName" form:"groupName" gorm:"comment:分组名称;column:group_name;size:100;" binding:"required"`  //分组名称
  Pid  *int `json:"pid" form:"pid" gorm:"comment:父分组ID;column:pid;size:19;"`  //父分组ID
  CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"comment:创建时间;column:create_time;"`  //创建时间
  UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"comment:更新时间;column:update_time;"`  //更新时间
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName deviceGroup表 DeviceGroup自定义表名 device_group
func (DeviceGroup) TableName() string {
    return "device_group"
}





