package repository

import (
	"NewRmsService/model"
	"NewRmsService/tools"
	"gorm.io/gorm"
)

type DriveRepository struct {

}

type IDriveRepository interface {
	CreateDrive(drive model.UserDrive)error
	FindDriveByName(userName string)(model.UserDrive,error)
	FindAllDrive(page,pageSize int)([]model.UserDrive,error)
	UpdateDrive(drive model.UserDrive)error
	DelDrive(DriveId int64)error
}

func (dr DriveRepository)CreateDrive(drive model.UserDrive)error  {
	return tools.DbEngine.Create(&drive).Error
}

func (dr DriveRepository) FindDriveByName(userName string)(drive model.UserDrive,err error) {
	return drive,tools.DbEngine.Where("user_name=?",userName).First(&drive).Error
}

func (dr DriveRepository) FindAllDrive(page,pageSize int)(drives []model.UserDrive,err error) {
	return drives,tools.DbEngine.Scopes(Paginate(page,pageSize)).Find(&drives).Error
}

func (dr DriveRepository)UpdateDrive(drive model.UserDrive)error {
	return tools.DbEngine.Model(&drive).Updates(&drive).Error
}

func (dr DriveRepository) DelDrive(DriveId int64)error {
	return tools.DbEngine.Delete(&model.UserDrive{},DriveId).Error
}

//分页封装
func Paginate(page int,pageSize int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}