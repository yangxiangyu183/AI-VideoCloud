package alert

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ AlertEventsRouter }

var alertEventsApi = api.ApiGroupApp.AlertApiGroup.AlertEventsApi
