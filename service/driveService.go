package service

import (
	"NewRmsService/model"
	"NewRmsService/repository"
)

type DriveService struct {
	IDriveRepository repository.IDriveRepository
}

type IDriveService interface {
	AddDrive(drive model.UserDrive)error
	QueryDriveByUserName(userName string)(model.UserDrive,error)
	QueryAllDrive(page,pageSize int)([]model.UserDrive,error)
	UpdateDrive(drive model.UserDrive)error
	DelDrive(driveId int64)error
}

func (ds DriveService)AddDrive(drive model.UserDrive)error{
	return ds.IDriveRepository.CreateDrive(drive)
}

func (ds DriveService) QueryDriveByUserName(userName string)(model.UserDrive,error)  {
	return ds.IDriveRepository.FindDriveByName(userName)
}

func (ds DriveService)QueryAllDrive(page,pageSize int)([]model.UserDrive,error)  {
	return ds.IDriveRepository.FindAllDrive(page,pageSize)
}
func (ds DriveService)UpdateDrive(drive model.UserDrive)error  {
	return ds.IDriveRepository.UpdateDrive(drive)
}

func (ds DriveService)DelDrive(driveId int64)error {
	return ds.IDriveRepository.DelDrive(driveId)
}