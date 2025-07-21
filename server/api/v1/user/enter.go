package user

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ AiUserApi }

var aiUserService = service.ServiceGroupApp.UserServiceGroup.AiUserService
