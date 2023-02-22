package models

import "github.com/shirou/gopsutil/disk"

type DiskInfo struct {
	TotalMem  uint64 `json:"totalMem"`
	TotalUsed uint64 `json:"totalUsed"`
	TotalFree uint64 `json:"totalFree"`
	DiskList  []struct {
		disk.PartitionStat
		disk.UsageStat
	} `json:"diskList"`
}
