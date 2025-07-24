package user

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	AiUserApi
	AdminUserApi
}

var (
	aiUserService    = service.ServiceGroupApp.UserServiceGroup.AiUserService
	adminUserService = service.ServiceGroupApp.UserServiceGroup.AdminUserService
)
