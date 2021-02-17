package model

type UserDrive struct {
	DriveId      int64	`json:"DriveId" gorm:"primary_key;not_null;auto_increment"`
	UserName    string	`json:"userName" gorm:"unique;not_null" from:"userName" binding:"required"`
	DriveName string	`json:"driveName" from:"DriveName" binding:"required"`
	CpuModel string	`json:"cpuModel" from:"cpuModel" binding:"required"`
	MemorySize string	`json:"memorySize" from:"memorySize" binding:"required"`
	DisplayCardModel string	`json:"displayCardModel"`
	HardSize string	`json:"diskSize" from:"diskSize" binding:"required"`
	SsdSize string	`json:"ssdSize"`
	Display1 string	`json:"display1" from:"display1" binding:"required"`
	Display2 string	`json:"display2"`
	InTime int64	`json:"inTime" from:"inTime" binding:"-"`
	CreatedAt int64	`json:"createdAt" gorm:"autoCreateTime "`
	UpdatedAt int64 `json:"updatedAt" gorm:"autoUpdateTime:nano"`
}
