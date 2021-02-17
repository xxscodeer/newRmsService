package service

import (
	"NewRmsService/model"
	"NewRmsService/repository"
	"NewRmsService/tools"
)

type UserService struct {
	repository.IUserRepository
}

type IUserService interface {
	AddUser(user model.User)error
	FindUserRole(userName string)(string,error)
	FindUserByName(userName string)(model.User,error)
	FindAllUser()([]model.User,error)
	AddAdminPath(pathModel model.AdminPathModel)error
	AddUserPath(pathModel model.UserPathModel)error
	QueryUserPath()([]model.UserPathModel,error)
	QueryAdminPath()([]model.AdminPathModel,error)
	CheckLogin(loginModel model.UserLoginModel)(bool,error)
	ModifyUser(user model.User)error
	DelUserInfo(userId int64)error
}

func (s UserService) AddUser(user model.User)error {
	userPwd := user.UserPwd
	basePwd := tools.EncodePwd(userPwd)
	user.UserPwd = basePwd
	return  s.IUserRepository.CreateUser(user)
}

func (s UserService)FindUserRole(userName string)(string,error) {
	users,err:=s.IUserRepository.FindUserByName(userName)
	return users.UserRole, err
}

func (s UserService) ModifyUser(user model.User)error {
	return s.IUserRepository.UpdateUser(user)
}

func (s UserService) DelUserInfo(userId int64)error  {
	return s.IUserRepository.DelUser(userId)
}

func (s UserService)CheckLogin(loginModel model.UserLoginModel)(bool,error)  {
	userName := loginModel.UserName
	userPwd := loginModel.UserPwd
	user,err:=s.IUserRepository.FindUserByName(userName)
	if err !=nil{
		return false, err
	}
	if user.UserName == userName && user.UserPwd == userPwd{
		return true,nil
	}
	return false,nil
}

func (s UserService) FindUserByName(userName string)(model.User,error)  {
	return s.IUserRepository.FindUserByName(userName)
}

func (s UserService) FindAllUser()([]model.User,error)  {
	return s.IUserRepository.FindAllUser()
}

func (s UserService)AddAdminPath(pathModel model.AdminPathModel)error {
	return s.IUserRepository.AddAdminPath(pathModel)
}

func (s UserService) QueryAdminPath()([]model.AdminPathModel,error) {
	return s.IUserRepository.FindAdminPath()
}

func (s UserService)QueryUserPath()([]model.UserPathModel,error) {
	return s.IUserRepository.FindUserPath()
}

func (s UserService)AddUserPath(pathModel model.UserPathModel)error {
	return s.IUserRepository.AddUserPath(pathModel)
}