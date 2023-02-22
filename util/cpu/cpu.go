package cpu

import (
	"WindowsListener/models"
	"WindowsListener/util/powershell"
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"strconv"
	"time"
)

func getCPUMessage() models.CPUInfo {
	message := new(models.CPUInfo)
	models.Logger.Info("尝试获取CPU信息")
	var err error
	message.More, err = cpu.Info()
	if err != nil {
		models.Logger.Errorf("获取CPU信息失败：%s", err)
	}
	models.Logger.Info("获取CPU信息成功")
	tmp, _ := json.Marshal(message)
	str := string(tmp)
	models.Logger.Debugf("已获取到的CPU信息:%s", str)
	return *message
}

func getTemp() float64 {
	models.Logger.Info("CPU温度查询中")
	strTemp := powershell.Run("(Get-WmiObject MSAcpi_ThermalZoneTemperature -Namespace \"root/wmi\").CurrentTemperature")
	Temp, err := strconv.ParseFloat(strTemp, 64)
	if err != nil {
		models.Logger.Errorf("获取CPU温度错误：%s", err)
	}
	Temp = Temp/10 - 273.15
	models.Logger.Info("获取CPU温度成功")
	return Temp
}

// TODO:这里也得重新改动
func getPercent(isPer bool) []float64 {
	var str string
	if isPer {
		str = "逻辑核心查询"
	} else {
		str = "总使用率查询"
	}
	models.Logger.Infof("尝试查询实时使用率，使用模式：%s", str)
	cpuPercent, err := cpu.Percent(time.Second, isPer)
	if err != nil {
		models.Logger.Errorf("查询失败：%s", err)
	}
	models.Logger.Info("查询CPU使用率成功")
	return cpuPercent
}

// GetCPUAll TODO: 重写GetAll实现重复查询
func GetCPUAll() models.CPUAll {
	var result models.CPUAll
	result.CPUInfo = getCPUMessage()
	result.Temp = getTemp()
	result.TotalPercent = getPercent(false)[0]
	return result
}
