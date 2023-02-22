package models

type MemInfo struct {
	Total uint64 `json:"total"`

	Available uint64 `json:"available"`

	Free uint64 `json:"free"`

	Used uint64 `json:"used"`

	UsedPercent float64 `json:"usedPercent"`
}
