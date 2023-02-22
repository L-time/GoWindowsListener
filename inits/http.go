package inits

import (
	"WindowsListener/models"
	"WindowsListener/models/Network"
	"WindowsListener/services"
	"github.com/gin-gonic/gin"
)

func HTTPInit() {
	Network.Services = gin.Default()
	Network.Services.GET("/Monitor", services.HttpSent)
	models.Logger.Info("启动HTTP服务......")
	err := Network.Services.Run(":12345")
	if err != nil {
		models.Logger.Fatalf("服务启动失败：%s", err)
	}
}
