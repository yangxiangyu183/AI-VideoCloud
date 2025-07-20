
// 自动生成模板CasbinRule
package system
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// casbinRule表 结构体  CasbinRule
type CasbinRule struct {
    global.GVA_MODEL
  Ptype  *string `json:"ptype" form:"ptype" gorm:"column:ptype;size:100;"`  //ptype字段
  V0  *string `json:"v0" form:"v0" gorm:"column:v0;size:100;"`  //v0字段
  V1  *string `json:"v1" form:"v1" gorm:"column:v1;size:100;"`  //v1字段
  V2  *string `json:"v2" form:"v2" gorm:"column:v2;size:100;"`  //v2字段
  V3  *string `json:"v3" form:"v3" gorm:"column:v3;size:100;"`  //v3字段
  V4  *string `json:"v4" form:"v4" gorm:"column:v4;size:100;"`  //v4字段
  V5  *string `json:"v5" form:"v5" gorm:"column:v5;size:100;"`  //v5字段
}


// TableName casbinRule表 CasbinRule自定义表名 casbin_rule
func (CasbinRule) TableName() string {
    return "casbin_rule"
}





