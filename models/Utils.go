package models

type AllInfo struct {
	CPU struct {
		CPUAll
	} `json:"cpu"`
	GPU struct {
		GPUAll
	} `json:"gpu"`
	Disk struct {
		DiskInfo
	} `json:"disk"`
	Mem struct {
		MemInfo
	} `json:"mem"`
	Net struct {
		NetSpeed
	} `json:"net"`
}
