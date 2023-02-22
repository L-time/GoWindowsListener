package disc

import (
	"WindowsListener/models"
	"github.com/shirou/gopsutil/disk"
)

type disks struct {
	disk.PartitionStat
	disk.UsageStat
}

func GetDiskInfo() models.DiskInfo {
	models.Logger.Info("尝试查询磁盘信息")
	var result models.DiskInfo
	diskList, err := disk.Partitions(false)
	if err != nil {
		models.Logger.Errorf("查询磁盘物理信息失败：%s", err)
	}
	for _, eachDisk := range diskList {
		var newDisk disks
		newDisk.PartitionStat = eachDisk
		tmpDisk, err := disk.Usage(eachDisk.Device)
		if err != nil {
			models.Logger.Errorf("查询 %s 磁盘分区失败：%s", eachDisk.Device, err)
		}
		newDisk.UsageStat = *tmpDisk
		models.Logger.Debugf("%s 分区磁盘使用情况：Mem: %d, Used: %d, Free: %d", newDisk.Device, newDisk.Total, newDisk.Used, newDisk.Free)
		result.TotalMem += tmpDisk.Total
		result.TotalFree += tmpDisk.Free
		result.TotalUsed += tmpDisk.Used
		result.DiskList = append(result.DiskList, newDisk)
	}
	models.Logger.Info("查询磁盘信息成功")
	return result
}
