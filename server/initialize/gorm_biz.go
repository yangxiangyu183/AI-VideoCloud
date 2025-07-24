package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/task_bor"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(system.CasbinRule{}, task_bor.Task{})
	if err != nil {
		return err
	}
	return nil
}
