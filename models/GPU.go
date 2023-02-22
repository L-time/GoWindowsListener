package models

import "encoding/xml"

type GPU_XML struct {
	XMLName      xml.Name `xml:"nvidia_smi_log"`
	CudaVersion  string   `xml:"cuda_version"`
	AttachedGpus uint     `xml:"attached_gpus"`
	Gpu          struct {
		ProductionName   string `xml:"product_name"`
		FanSpeed         string `xml:"fan_speed"`         // 0 %
		PerformanceState string `xml:"performance_state"` // P8
		Utilization      struct {
			GpuUtil     string `xml:"gpu_util"`
			MemoryUtil  string `xml:"memory_util"`
			EncoderUtil string `xml:"encoder_util"`
			DecoderUtil string `xml:"decoder_util"`
		} `xml:"utilization"`
		Temperature struct {
			GpuTemp string `xml:"gpu_temp"`
		} `xml:"temperature"`
		PowerReadings struct {
			PowerState string `xml:"power_state"`
			PowerDraw  string `xml:"power_draw"`
			PowerLimit string `xml:"power_limit"`
		} `xml:"power_readings"`
	} `xml:"gpu"`
}

type GPUInfo struct {
	Name string `json:"name"`
}

type GPUStar struct {
	Fanspeed    int     `json:"fanspeed"`
	Temperature int     `json:"temperature"`
	Percent     int     `json:"percent"`
	Power       float64 `json:"power"`
	Memory      int     `json:"memory"`
}

type GPUAll struct {
	GPUInfo
	GPUStar
}
