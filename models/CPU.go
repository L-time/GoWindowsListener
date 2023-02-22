package models

import (
	"github.com/shirou/gopsutil/cpu"
)

type CPUInfo struct {
	PhysicalCore int            `json:"physicalCore"`
	LogicalCore  int            `json:"logicalCore"`
	More         []cpu.InfoStat `json:"more"`
}

type CPUStar struct {
	Temp         float64   `json:"temp"`
	TotalPercent float64   `json:"totalPercent"`
	PerPercent   []float64 `json:"perPercent"`
}

type CPUAll struct {
	CPUInfo
	CPUStar
}
