package inits

import (
	"WindowsListener/models"
)

func init() {
	LogInit()
}

func AllInit() {
	models.Logger.Info("初始化完成")
}
