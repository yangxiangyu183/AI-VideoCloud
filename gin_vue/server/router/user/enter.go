package user

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	AiUserRouter
	AdminUserRouter
}

var (
	aiUserApi    = api.ApiGroupApp.UserApiGroup.AiUserApi
	adminUserApi = api.ApiGroupApp.UserApiGroup.AdminUserApi
)
