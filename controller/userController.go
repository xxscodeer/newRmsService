package controller

import (
	"NewRmsService/i18n"
	"NewRmsService/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	IUserService service.IUserService
}
// @Summary 查询用户信息
// @Tags user
// @Description 查询全部用户信息
// @accept json
// @Produce  json
// @Success 200 {object} model.Res {"code":200,"msg":"用户查询成功"}
// @Failure 400 {object} model.Res {"code":400,"msg":"用户查询失败"}
// @Router /user/info/:userName [get]
func (c UserController)GetUserinfo(ctx *gin.Context)  {
	userName :=ctx.Param("userName")
	user,err:=c.IUserService.FindUserByName(userName)
	if err !=nil{
		i18n.ResponseFail(ctx,"查询失败")
		return
	}
	i18n.DataResponse(ctx,user)
}
// @Summary 查询用户路径
// @Tags user
// @Description 查询用户可有路径
// @accept json
// @Produce  json
// @Success 200 {object} model.Res {"code":200,"msg":"用户查询成功"}
// @Failure 400 {object} model.Res {"code":400,"msg":"用户查询失败"}
// @Router /user/getUserPath [get]
func (c UserController)GetUserPath(ctx *gin.Context){
	path,err:=c.IUserService.QueryUserPath()
	if err !=nil{
		i18n.ResponseFail(ctx,"获取路径失败")
		return
	}
	i18n.DataResponse(ctx,path)
}
