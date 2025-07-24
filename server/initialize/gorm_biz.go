package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/task_bor"
)

func bizModel() error {
	db := global.GVA_DB
<<<<<<< HEAD
	err := db.AutoMigrate(system.CasbinRule{}, device.DeviceGroup{}, device.MonitorDevice{})
=======
	err := db.AutoMigrate(system.CasbinRule{}, task_bor.Task{})
>>>>>>> yxy
	if err != nil {
		return err
	}
	return nil
}
