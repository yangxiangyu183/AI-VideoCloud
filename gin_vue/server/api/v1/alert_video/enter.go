package alert_video

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ AlertApi }

var alertService = service.ServiceGroupApp.Alert_videoServiceGroup.AlertService
