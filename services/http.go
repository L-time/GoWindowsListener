package services

import (
	"WindowsListener/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpSent(context *gin.Context) {
	result := util.GetAllHardWareInfo()
	context.JSON(http.StatusOK, result)
}
