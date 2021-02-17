package controller

import (
	"NewRmsService/i18n"
	"NewRmsService/model"
	"NewRmsService/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AdminController struct {
	IUserService service.IUserService
}
// @Summary 管理添加用户
// @Tags admin
// @Description 新增用户
// @accept json
// @Produce  json
// @Param user body model.User true "user"
// @Success 200 {object} model.Res {"code":200,"msg":"用户添加成功"}
// @Failure 400 {object} model.Res {"code":400,"msg":"用户添加失败"}
// @Router /api/v1/addUser [post]
func (ac AdminController)AddUser(ctx *gin.Context)  {
	user := model.User{}
	if err := ctx.ShouldBindJSON(&user);err !=nil{
		i18n.BindJsonFail(ctx,err)
	}
	if err:=ac.IUserService.AddUser(user);err !=nil{
		i18n.ResponseFail(ctx,"添加用户失败")
		return
	}
	i18n.ResponseSuccess(ctx,"添加用户成功")
}

// @Summary 修改用户
// @Tags admin
// @Description 管理修改用户权限
// @accept json
// @Produce  json
// @Param user body model.User true "user"
// @Success 200 {object} model.Res {"code":200,"msg":"用户更新成功"}
// @Failure 400 {object} model.Res {"code":400,"msg":"用户更新失败"}
// @Router /api/v1/modifyUser [put]
func (ac AdminController) ModifyUser(ctx *gin.Context)  {
	user := model.User{}
	if err := ctx.ShouldBindJSON(&user);err!=nil{
		i18n.BindJsonFail(ctx,err)
		return
	}
	if err :=ac.IUserService.ModifyUser(user);err !=nil{
		i18n.ResponseFail(ctx,"修改用户数据失败")
		return
	}
	i18n.ResponseSuccess(ctx,"更新用户成功")
}

// @Summary 删除用户
// @Tags admin
// @Description 管理删除用户
// @accept json
// @Produce  json
// @Param userId path int true "userID"
// @Success 200 {object} model.Res {"code":200,"msg":"用户删除成功"}
// @Failure 400 {object} model.Res {"code":400,"msg":"用户删除失败"}
// @Router /api/v1/delUser/:userId [delete]
func (ac AdminController) DelUser(ctx *gin.Context)  {
	id:=ctx.Param("userId")
	userId,_ := strconv.Atoi(id)
	if err :=ac.IUserService.DelUserInfo(int64(userId));err !=nil{
		i18n.ResponseFail(ctx,fmt.Sprintln("删除用户失败",err))
		return
	}
	i18n.ResponseSuccess(ctx,"删除成功")
}

// @Summary 查询用户信息
// @Tags admin
// @Description 根据用户名查询用户信息
// @accept json
// @Produce  json
// @Param userName path string true "userName"
// @Success 200 {object} model.Res {"code":200,"msg":"用户查询成功"}
// @Failure 400 {object} model.Res {"code":400,"msg":"用户查询失败"}
// @Router /api/v1/getUser/:userName [get]
func (ac AdminController)GetUserInfo(ctx *gin.Context){
	userName:=ctx.Param("userName")
	user,err:=ac.IUserService.FindUserByName(userName)
	if err !=nil{
		i18n.ResponseFail(ctx,fmt.Sprint("用户不存在,",err))
		return
	}
	i18n.DataResponse(ctx,user)
}
// @Summary 查询全部用户
// @Tags admin
// @Description 查询全部用户信息
// @accept json
// @Produce  json
// @Success 200 {object} model.Res {"code":200,"msg":"用户查询成功"}
// @Failure 400 {object} model.Res {"code":400,"msg":"用户查询失败"}
// @Router /api/v1/getAllUser [get]
func (ac AdminController)GetAllUser(ctx *gin.Context)  {
	users,err := ac.IUserService.FindAllUser()
	if err !=nil{
		i18n.ResponseFail(ctx,"查询失败")
		return
	}
	i18n.DataResponse(ctx,users)
}
// @Summary 新增管理员路由路径
// @Tags admin
// @Description 新增管理员路径
// @accept json
// @Produce  json
// @Param adminPath body model.AdminPathModel true "adminPath"
// @Success 200 {object} model.Res {"code":200,"msg":"用户添加成功"}
// @Failure 400 {object} model.Res {"code":400,"msg":"用户添加失败"}
// @Router /api/v1/addAdminPath [post]
func (ac AdminController)PostAdminPath(ctx *gin.Context) {
	path := model.AdminPathModel{}
	if err :=ctx.ShouldBindJSON(&path);err !=nil{
		i18n.BindJsonFail(ctx,err)
		return
	}
	if err := ac.IUserService.AddAdminPath(path);err!=nil{
		i18n.ResponseFail(ctx,"添加路径失败")
		return
	}
	i18n.ResponseSuccess(ctx,"添加路径成功")
}
// @Summary 新增用户路径
// @Tags admin
// @Description 新增用户路由路径
// @accept json
// @Produce  json
// @Param userPath body model.UserPathModel true "userPath"
// @Success 200 {object} model.Res {"code":200,"msg":"用户添加成功"}
// @Failure 400 {object} model.Res {"code":400,"msg":"用户添加失败"}
// @Router /api/v1/addUserPath [post]
func (ac AdminController)PostUserPath(ctx *gin.Context){
	path := model.UserPathModel{}
	if err :=ctx.ShouldBindJSON(&path);err !=nil{
		i18n.BindJsonFail(ctx,err)
		return
	}
	if err := ac.IUserService.AddUserPath(path);err!=nil{
		i18n.ResponseFail(ctx,"添加路径失败")
		return
	}
	i18n.ResponseSuccess(ctx,"添加路径成功")
}
// @Summary 查询管理员路径
// @Tags admin
// @Description 查询管理员后台信息
// @accept json
// @Produce  json
// @Success 200 {object} model.Res {"code":200,"msg":"用户查询成功"}
// @Failure 400 {object} model.Res {"code":400,"msg":"用户查询失败"}
// @Router /api/v1/getAdminUrl [get]
func (ac AdminController) GetAdminPath(ctx *gin.Context)  {
	path,err:=ac.IUserService.QueryAdminPath()
	if err !=nil{
		i18n.ResponseFail(ctx,"获取路径失败"+err.Error())
		return
	}
	i18n.DataResponse(ctx,path)
}