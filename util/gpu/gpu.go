package gpu

import (
	"WindowsListener/models"
	"WindowsListener/util/powershell"
	"encoding/xml"
	"strconv"
	"strings"
)

func GetGPUAll() models.GPUAll {
	models.Logger.Info("尝试获取GPU信息")
	xmls := powershell.Run("nvidia-smi -q -x")
	var AllInfo models.GPU_XML
	err := xml.Unmarshal([]byte(xmls), &AllInfo)
	if err != nil {
		models.Logger.Fatalf("解析GPU信息失败：%s", err)
	}
	var result models.GPUAll

	result.Name = AllInfo.Gpu.ProductionName

	result.Temperature, err = strconv.Atoi(strings.TrimSuffix(AllInfo.Gpu.Temperature.GpuTemp, " C"))
	if err != nil {
		models.Logger.Errorf("无法获取GPU温度数据：%s", err)
	}

	result.Percent, err = strconv.Atoi(strings.TrimSuffix(AllInfo.Gpu.Utilization.GpuUtil, " %"))
	if err != nil {
		models.Logger.Errorf("无法获取GPU占用率数据：%s", err)
	}

	result.Memory, err = strconv.Atoi(strings.TrimSuffix(AllInfo.Gpu.Utilization.MemoryUtil, " %"))
	if err != nil {
		models.Logger.Errorf("无法获取GPU内存占用率数据：%s", err)
	}

	result.Fanspeed, err = strconv.Atoi(strings.TrimSuffix(AllInfo.Gpu.FanSpeed, " %"))
	if err != nil {
		models.Logger.Errorf("无法获取GPU风扇数据：%s", err)
	}

	result.Power, err = strconv.ParseFloat(strings.TrimSuffix(AllInfo.Gpu.PowerReadings.PowerDraw, " W"), 64)
	if err != nil {
		models.Logger.Errorf("无法获取GPU电源信息：%s", err)
	}
	models.Logger.Info("获取可获得的GPU信息")
	return result
}
