package repository

import (
	"NewRmsService/model"
	"NewRmsService/tools"
)

type UserRepository struct {
}

type IUserRepository interface {
	CreateUser(user model.User)error
	FindUserByName(userName string)(model.User,error)
	FindAllUser()([]model.User,error)
	UpdateUser(user model.User)error
	DelUser(userId int64)error
	AddUserPath(pathModel model.UserPathModel)error
	AddAdminPath(pathModel model.AdminPathModel)error
	FindUserPath()([]model.UserPathModel,error)
	FindAdminPath()([]model.AdminPathModel,error)
}

func (ur UserRepository)CreateUser(user model.User)error {
	return tools.DbEngine.Create(&user).Error
}

func (ur UserRepository)FindUserByName(userName string)(user model.User,err error) {
	return user,tools.DbEngine.Where("user_name=?",userName).First(&user).Error
}

func (ur UserRepository) UpdateUser(user model.User)error {
	return tools.DbEngine.Model(&user).Updates(user).Error
}

func (ur UserRepository)DelUser(userId int64)error {
	return tools.DbEngine.Delete(&model.User{},userId).Error
}

func (ur UserRepository) FindAllUser()(users []model.User,err error)  {
	return users,tools.DbEngine.Find(&users).Error
}

func (ur UserRepository) AddUserPath(pathModel model.UserPathModel)error  {
	return tools.DbEngine.Create(&pathModel).Error
}

func (ur UserRepository) AddAdminPath(pathModel model.AdminPathModel)error  {
	return tools.DbEngine.Create(&pathModel).Error
}

func (ur UserRepository)FindUserPath()(paths []model.UserPathModel,err error) {
	return paths,tools.DbEngine.Preload("OneUrlList").Find(&paths).Error
}
func (ur UserRepository)FindAdminPath()(paths []model.AdminPathModel,err error)  {
	return paths,tools.DbEngine.Preload("AdminOne").Find(&paths).Error
}
