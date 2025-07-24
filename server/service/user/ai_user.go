package user

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	"gorm.io/gorm"
)

type AiUserService struct{}

// CreateAiUser 创建aiUser表记录
// Author [yourname](https://github.com/yourname)
func (aiUserService *AiUserService) CreateAiUser(ctx context.Context, aiUser *user.AiUser) (err error) {
	err = global.GVA_DB.Create(aiUser).Error
	return err
}

// DeleteAiUser 删除aiUser表记录
// Author [yourname](https://github.com/yourname)
func (aiUserService *AiUserService) DeleteAiUser(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&user.AiUser{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&user.AiUser{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteAiUserByIds 批量删除aiUser表记录
// Author [yourname](https://github.com/yourname)
func (aiUserService *AiUserService) DeleteAiUserByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&user.AiUser{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&user.AiUser{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateAiUser 更新aiUser表记录
// Author [yourname](https://github.com/yourname)
func (aiUserService *AiUserService) UpdateAiUser(ctx context.Context, aiUser user.AiUser) (err error) {
	err = global.GVA_DB.Model(&user.AiUser{}).Where("id = ?", aiUser.ID).Updates(&aiUser).Error
	return err
}

// GetAiUser 根据ID获取aiUser表记录
// Author [yourname](https://github.com/yourname)
func (aiUserService *AiUserService) GetAiUser(ctx context.Context, ID string) (aiUser user.AiUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&aiUser).Error
	return
}

// GetAiUserInfoList 分页获取aiUser表记录
// Author [yourname](https://github.com/yourname)
func (aiUserService *AiUserService) GetAiUserInfoList(ctx context.Context, info userReq.AiUserSearch) (list []user.AiUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&user.AiUser{})
	var aiUsers []user.AiUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&aiUsers).Error
	return aiUsers, total, err
}
func (aiUserService *AiUserService) GetAiUserPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
