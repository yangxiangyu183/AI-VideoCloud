
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type MonitorDeviceSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      DeviceName  *string `json:"deviceName" form:"deviceName"` 
      GroupId  *int `json:"groupId" form:"groupId"` 
      StreamUrl  *string `json:"streamUrl" form:"streamUrl"` 
      Status  *string `json:"status" form:"status"` 
    request.PageInfo
}
