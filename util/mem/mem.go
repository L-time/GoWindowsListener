package mem

import (
	"WindowsListener/models"
	"github.com/shirou/gopsutil/mem"
)

func GetMemInfo() models.MemInfo {
	models.Logger.Info("尝试获取内存信息")
	var result models.MemInfo
	AllInfo, err := mem.VirtualMemory()
	if err != nil {
		models.Logger.Errorf("获取内存数据错误：%s", err)
	}
	result.Total = AllInfo.Total
	result.Used = AllInfo.Used
	result.Free = AllInfo.Free
	result.Available = AllInfo.Available
	result.UsedPercent = AllInfo.UsedPercent
	models.Logger.Debugf("内存信息：Total: %d , Free: %d, Used: %d, Available: %d, Percent: %f %%",
		result.Total,
		result.Free,
		result.Used,
		result.Available,
		result.UsedPercent)
	models.Logger.Info("获取内存数据成功")
	return result
}
