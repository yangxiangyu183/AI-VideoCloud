package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(system.CasbinRule{}, device.DeviceGroup{}, device.MonitorDevice{})
	if err != nil {
		return err
	}
	return nil
}
