
package user

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
    userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
    "gorm.io/gorm"
)

type AdminUserService struct {}
// CreateAdminUser 创建管理厅员表记录
// Author [yourname](https://github.com/yourname)
func (adminUserService *AdminUserService) CreateAdminUser(ctx context.Context, adminUser *user.AdminUser) (err error) {
	err = global.GVA_DB.Create(adminUser).Error
	return err
}

// DeleteAdminUser 删除管理厅员表记录
// Author [yourname](https://github.com/yourname)
func (adminUserService *AdminUserService)DeleteAdminUser(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&user.AdminUser{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&user.AdminUser{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteAdminUserByIds 批量删除管理厅员表记录
// Author [yourname](https://github.com/yourname)
func (adminUserService *AdminUserService)DeleteAdminUserByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&user.AdminUser{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&user.AdminUser{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateAdminUser 更新管理厅员表记录
// Author [yourname](https://github.com/yourname)
func (adminUserService *AdminUserService)UpdateAdminUser(ctx context.Context, adminUser user.AdminUser) (err error) {
	err = global.GVA_DB.Model(&user.AdminUser{}).Where("id = ?",adminUser.ID).Updates(&adminUser).Error
	return err
}

// GetAdminUser 根据ID获取管理厅员表记录
// Author [yourname](https://github.com/yourname)
func (adminUserService *AdminUserService)GetAdminUser(ctx context.Context, ID string) (adminUser user.AdminUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&adminUser).Error
	return
}
// GetAdminUserInfoList 分页获取管理厅员表记录
// Author [yourname](https://github.com/yourname)
func (adminUserService *AdminUserService)GetAdminUserInfoList(ctx context.Context, info userReq.AdminUserSearch) (list []user.AdminUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&user.AdminUser{})
    var adminUsers []user.AdminUser
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

	err = db.Find(&adminUsers).Error
	return  adminUsers, total, err
}
func (adminUserService *AdminUserService)GetAdminUserPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
