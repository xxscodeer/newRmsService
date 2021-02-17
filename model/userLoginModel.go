package model

type UserLoginModel struct {
	UserName string	`json:"userName" from:"userName" binding:"required"`
	UserPwd string	`json:"userPwd" from:"userPwd" binding:"required,max=16,min=5"`
}


type Res struct {
	Code int
	Msg string
}