package net

import (
	"WindowsListener/models"
	"github.com/shirou/gopsutil/net"
	"time"
)

func GetSpeed() models.NetSpeed {
	var result models.NetSpeed
	models.Logger.Info("尝试查询当前网络速度")
	netInfo1, err1 := net.IOCounters(false)
	if err1 != nil {
		models.Logger.Errorf("第一次获取网卡数据失败: %s", err1)
	}
	time.Sleep(time.Second)
	netInfo2, err2 := net.IOCounters(false)
	if err2 != nil {
		models.Logger.Errorf("第二次获取网卡数据失败: %s", err2)
	}
	result.Upload = netInfo2[0].BytesSent - netInfo1[0].BytesSent
	result.Download = netInfo2[0].BytesRecv - netInfo1[0].BytesRecv
	models.Logger.Debugf("获取到的网卡数据：↑ %d ,↓ %d", result.Upload, result.Download)
	return result
}
