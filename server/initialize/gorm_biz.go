package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/task_bor"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(
		system.CasbinRule{},
		device.DeviceGroup{},
		device.MonitorDevice{},
		user.AiUser{},
		user.AdminUser{},
		task_bor.Task{},
	)
	if err != nil {
		return err
	}
	return nil
}
