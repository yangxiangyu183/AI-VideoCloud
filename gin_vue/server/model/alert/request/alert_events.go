
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AlertEventsSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      TaskId  *int `json:"taskId" form:"taskId"` 
      CameraId  *int `json:"cameraId" form:"cameraId"` 
      AlertType  *string `json:"alertType" form:"alertType"` 
      Servenrity  *string `json:"servenrity" form:"servenrity"` 
      ImagePath  *string `json:"imagePath" form:"imagePath"` 
      VideoPath  *string `json:"videoPath" form:"videoPath"` 
      IsRead  *string `json:"isRead" form:"isRead"` 
      IsProcessed  *string `json:"isProcessed" form:"isProcessed"` 
      ProcessedBy  *int `json:"processedBy" form:"processedBy"` 
      ProcessedTime  *time.Time `json:"processedTime" form:"processedTime"` 
      ProcessedResult  *string `json:"processedResult" form:"processedResult"` 
    request.PageInfo
}
