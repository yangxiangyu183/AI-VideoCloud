package alert

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ AlertEventsApi }

var alertEventsService = service.ServiceGroupApp.AlertServiceGroup.AlertEventsService
