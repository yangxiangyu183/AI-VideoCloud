
// 自动生成模板AdminUser
package user
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 管理厅员表 结构体  AdminUser
type AdminUser struct {
    global.GVA_MODEL
  UserId  *int `json:"userId" form:"userId" gorm:"comment:用户的ID;column:user_id;size:10;" binding:"required"`  //用户的ID
  AdminPassword  *string `json:"adminPassword" form:"adminPassword" gorm:"comment:管理员的密码;column:admin_password;size:32;" binding:"required"`  //管理员的密码
  AdminName  *string `json:"adminName" form:"adminName" gorm:"comment:管理员账号;column:admin_name;size:25;" binding:"required"`  //管理员账号
  Status  *string `json:"status" form:"status" gorm:"default:1;comment:用户的账号状态;column:status;size:2;" binding:"required"`  //用户的账号状态
  Permissions  *string `json:"permissions" form:"permissions" gorm:"default:4;comment:用户的使用权限;column:permissions;size:2;" binding:"required"`  //用户的使用权限
  IdAdmin  *string `json:"idAdmin" form:"idAdmin" gorm:"default:2;comment:管理员等级（1普通,2超级）;column:id_admin;size:2;" binding:"required"`  //管理员等级（1普通,2超级）
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 管理厅员表 AdminUser自定义表名 admin_user
func (AdminUser) TableName() string {
    return "admin_user"
}





