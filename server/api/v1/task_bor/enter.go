package task_bor

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ TaskApi }

var task表Service = service.ServiceGroupApp.Task_borServiceGroup.TaskService
