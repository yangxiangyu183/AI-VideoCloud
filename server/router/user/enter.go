package user

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ AiUserRouter }

var aiUserApi = api.ApiGroupApp.UserApiGroup.AiUserApi
