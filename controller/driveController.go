package controller

import (
	"NewRmsService/i18n"
	"NewRmsService/model"
	"NewRmsService/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DriveController struct {
	IDriveService service.IDriveService
}
// @Summary 添加设备
// @Tags adminDrive
// @Description 添加设备信息
// @accept json
// @Produce  json
// @Param user body model.UserDrive true "userDrive"
// @Success 200 {object} model.Res {"code":200,"msg":"添加设备成功"}
// @Failure 400 {object} model.Res {"code":400,"msg":"添加设备失败"}
// @Router /api/v1/drive/postDrive [post]
func (dc DriveController)PostDriveInfo(ctx *gin.Context)  {
	drive := model.UserDrive{}
	if err := ctx.ShouldBindJSON(&drive);err !=nil{
		i18n.BindJsonFail(ctx,err)
		return
	}
	if err := dc.IDriveService.AddDrive(drive);err !=nil{
		i18n.ResponseFail(ctx,fmt.Sprintln("设备添加失败,",err))
		return
	}
	i18n.ResponseSuccess(ctx,"设备添加成功")
}

// @Summary 修改设备
// @Tags adminDrive
// @Description 修改设备信息
// @accept json
// @Produce  json
// @Param user body model.UserDrive true "userDrive"
// @Success 200 {object} model.Res {"code":200,"msg":"修改设备成功"}
// @Failure 400 {object} model.Res {"code":400,"msg":"修改设备失败"}
// @Router /api/v1/drive/putDrive [put]
func (dc DriveController) PutDrive(ctx *gin.Context){
	 var drive model.UserDrive
	 if err := ctx.ShouldBindJSON(&drive);err !=nil{
	 	i18n.BindJsonFail(ctx,err)
		 return
	 }
	 if err := dc.IDriveService.UpdateDrive(drive);err !=nil{
	 	i18n.ResponseFail(ctx,fmt.Sprintln("更新设备信息失败,",err))
		 return
	 }
	 i18n.ResponseSuccess(ctx,"更新设备成功")
}

// @Summary 删除设备
// @Tags adminDrive
// @Description 根据设备id删除设备信息
// @accept json
// @Produce  json
// @Param user path int true "userDrive"
// @Success 200 {object} model.Res {"code":200,"msg":"删除设备成功"}
// @Failure 400 {object} model.Res {"code":400,"msg":"删除设备失败"}
// @Router /api/v1/drive/delDrive/:driveId [Delete]
func (dc DriveController)DelDrive(ctx *gin.Context){
	id := ctx.Param("driveId")
	driveId,_:=strconv.Atoi(id)
	if err := dc.IDriveService.DelDrive(int64(driveId));err !=nil{
		i18n.ResponseFail(ctx,fmt.Sprintln("删除设备失败,",err))
		return
	}
	i18n.ResponseSuccess(ctx,"删除设备成功")
}
// @Summary 查询设备信息
// @Tags adminDrive
// @Description 根据用户名查询设备信息
// @accept json
// @Produce  json
// @Param userName path string true "userName"
// @Success 200 {object} model.Res {"code":200,"msg":"用户设备查询成功"}
// @Failure 400 {object} model.Res {"code":400,"msg":"用户设备查询失败"}
// @Router /api/v1/drive/getUserDrive/:userName [get]
func (dc DriveController)GetDriveByUserName(ctx *gin.Context) {
	userName := ctx.Param("userName")
	drive,err := dc.IDriveService.QueryDriveByUserName(userName)
	if err !=nil{
		i18n.ResponseFail(ctx,"设备查询失败")
		return
	}
	i18n.DataResponse(ctx,drive)
}

// @Summary 查询全部设备信息
// @Tags adminDrive
// @Description 查询所有设备信息
// @accept json
// @Produce  json
// @Param page query string true "页数"
// @Param pageSize query string true "每页显示数量"
// @Success 200 {object} model.Res {"code":200,"msg":"用户设备查询成功"}
// @Failure 400 {object} model.Res {"code":400,"msg":"用户设备查询失败"}
// @Router /api/v1/drive/getAllDrive [get]
func (dc DriveController) GetAllDrive(ctx *gin.Context)  {
	var (
		page int
		pageSize int)
	pageStr,isExist := ctx.GetQuery("page")
	if isExist{
		page,_ = strconv.Atoi(pageStr)
	}else {
		i18n.BindJsonFail(ctx,errors.New("请传入正确的页码"))
		return
	}
	pageSizeStr,isSizeExist := ctx.GetQuery("pageSize")
	if isSizeExist{
		pageSize,_ = strconv.Atoi(pageSizeStr)
	}else {
		i18n.BindJsonFail(ctx,errors.New("请传入每页的参数"))
		return
	}

	drives,err := dc.IDriveService.QueryAllDrive(page,pageSize)
	if err !=nil{
		i18n.ResponseFail(ctx,fmt.Sprintln("查询设备信息失败"))
		return
	}
	i18n.DataResponse(ctx,drives)
}