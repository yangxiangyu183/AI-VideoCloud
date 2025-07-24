package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
<<<<<<< HEAD
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
=======
	"github.com/flipped-aurora/gin-vue-admin/server/model/task_bor"
>>>>>>> hjy
)

func bizModel() error {
	db := global.GVA_DB
<<<<<<< HEAD
	err := db.AutoMigrate(system.CasbinRule{}, user.AiUser{}, user.AiUser{}, user.AdminUser{})
=======
<<<<<<< HEAD
	err := db.AutoMigrate(system.CasbinRule{}, device.DeviceGroup{}, device.MonitorDevice{})
=======
	err := db.AutoMigrate(system.CasbinRule{}, task_bor.Task{})
>>>>>>> yxy
>>>>>>> hjy
	if err != nil {
		return err
	}
	return nil
}
