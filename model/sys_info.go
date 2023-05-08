package model

type SysInfo struct {
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
	Disk   string `json:"disk"`
	Gpu    string `json:"gpu"`
}
