package util

import (
	"WindowsListener/models"
	"WindowsListener/util/cpu"
	"WindowsListener/util/disc"
	"WindowsListener/util/gpu"
	"WindowsListener/util/mem"
	"WindowsListener/util/net"
)

func GetAllHardWareInfo() models.AllInfo {
	var result models.AllInfo
	result.CPU.CPUAll = cpu.GetCPUAll()
	result.GPU.GPUAll = gpu.GetGPUAll()
	result.Mem.MemInfo = mem.GetMemInfo()
	result.Disk.DiskInfo = disc.GetDiskInfo()
	result.Net.NetSpeed = net.GetSpeed()
	return result
}
