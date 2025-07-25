
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AlertSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Video  *string `json:"video" form:"video"` 
      CameraAddress  *string `json:"cameraAddress" form:"cameraAddress"` 
      AlertType  *string `json:"alertType" form:"alertType"` 
      AlertTime  *time.Time `json:"alertTime" form:"alertTime"` 
    request.PageInfo
}
