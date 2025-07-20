
package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type CasbinRuleService struct {}
// CreateCasbinRule 创建casbinRule表记录
// Author [yourname](https://github.com/yourname)
func (casbinRuleService *CasbinRuleService) CreateCasbinRule(ctx context.Context, casbinRule *system.CasbinRule) (err error) {
	err = global.GVA_DB.Create(casbinRule).Error
	return err
}

// DeleteCasbinRule 删除casbinRule表记录
// Author [yourname](https://github.com/yourname)
func (casbinRuleService *CasbinRuleService)DeleteCasbinRule(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&system.CasbinRule{},"id = ?",ID).Error
	return err
}

// DeleteCasbinRuleByIds 批量删除casbinRule表记录
// Author [yourname](https://github.com/yourname)
func (casbinRuleService *CasbinRuleService)DeleteCasbinRuleByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.CasbinRule{},"id in ?",IDs).Error
	return err
}

// UpdateCasbinRule 更新casbinRule表记录
// Author [yourname](https://github.com/yourname)
func (casbinRuleService *CasbinRuleService)UpdateCasbinRule(ctx context.Context, casbinRule system.CasbinRule) (err error) {
	err = global.GVA_DB.Model(&system.CasbinRule{}).Where("id = ?",casbinRule.ID).Updates(&casbinRule).Error
	return err
}

// GetCasbinRule 根据ID获取casbinRule表记录
// Author [yourname](https://github.com/yourname)
func (casbinRuleService *CasbinRuleService)GetCasbinRule(ctx context.Context, ID string) (casbinRule system.CasbinRule, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&casbinRule).Error
	return
}
// GetCasbinRuleInfoList 分页获取casbinRule表记录
// Author [yourname](https://github.com/yourname)
func (casbinRuleService *CasbinRuleService)GetCasbinRuleInfoList(ctx context.Context, info systemReq.CasbinRuleSearch) (list []system.CasbinRule, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&system.CasbinRule{})
    var casbinRules []system.CasbinRule
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&casbinRules).Error
	return  casbinRules, total, err
}
func (casbinRuleService *CasbinRuleService)GetCasbinRulePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
