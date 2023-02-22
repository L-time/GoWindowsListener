package models

type NetSpeed struct {
	Upload   uint64 `json:"upload"`
	Download uint64 `json:"download"`
}
