package task_bor

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ TaskRouter }

var task表Api = api.ApiGroupApp.Task_borApiGroup.TaskApi
