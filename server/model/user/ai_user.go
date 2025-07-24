// 自动生成模板AiUser
package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// aiUser表 结构体  AiUser
type AiUser struct {
	global.GVA_MODEL
	UserName    *string `json:"userName" form:"userName" gorm:"column:user_name;size:30;" binding:"required"`     //userName字段
	Password    *string `json:"password" form:"password" gorm:"column:password;size:191;" binding:"required"`     //password字段
	UserEmail   *string `json:"userEmail" form:"userEmail" gorm:"column:user_email;size:100;" binding:"required"` //userEmail字段
	Permissions *string `json:"permissions" form:"permissions" gorm:"comment:用户权限;column:permissions;size:100;"`  //用户权限
	Status      *string `json:"status" form:"status" gorm:"comment:用户状态;column:status;size:1;"`                   //用户状态
	CreatedBy   uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName aiUser表 AiUser自定义表名 ai_user
func (AiUser) TableName() string {
	return "ai_user"
}
