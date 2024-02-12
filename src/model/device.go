package model

var Devices []DeviceList = nil

type DeviceList struct {
	DeviceId        string            `json:"deviceid"`
	DeviceName      string            `json:"devicename"`
	Applications    []ApplicationInfo `json:"appinfo"`
	AvailableMemory string            `json:"availablememory"`
	TotalMemory     string            `json:"totalmemory"`
	Status          string            `json:"status"`
}

type DeviceRegistration struct {
	DeviceName string `json:"devicename"`
}

type ApplicationInfo struct {
	ApplicationId   string `json:"applicationid"`
	ApplicationName string `json:"applicationname"`
}
