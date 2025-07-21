
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type DeviceGroupSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      GroupName  *string `json:"groupName" form:"groupName"` 
      Pid  *int `json:"pid" form:"pid"` 
    request.PageInfo
}
