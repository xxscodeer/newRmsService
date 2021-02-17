package route

import (
	"NewRmsService/controller"
	_ "NewRmsService/docs"
	"NewRmsService/middleware"
	"NewRmsService/repository"
	"NewRmsService/service"
	"github.com/gin-gonic/gin"
)
func InitAppRouter()*gin.Engine{
	app := gin.Default()

	user := controller.UserController{IUserService: service.UserService{IUserRepository: repository.UserRepository{}}}
	admin := controller.AdminController{IUserService: service.UserService{IUserRepository: repository.UserRepository{}}}
	drive := controller.DriveController{IDriveService: service.DriveService{IDriveRepository: repository.DriveRepository{}}}
	Jwt := middleware.UserJwt{IUserService: service.UserService{IUserRepository: repository.UserRepository{}}}

	userJwt:=Jwt.InitUserJwtFunction(middleware.AllAuth)
	app.Use(middleware.Cors())
	app.POST("/login",userJwt.LoginHandler)
	app.NoRoute(userJwt.MiddlewareFunc(),middleware.NoRouteHandler)
	adminJwt := Jwt.InitUserJwtFunction(middleware.AdminAuth)
	adminApi:= app.Group("/api/v1/").Use(adminJwt.MiddlewareFunc())
	{
		adminApi.POST("/addUser",admin.AddUser)
		adminApi.PUT("/modifyUser",admin.ModifyUser)
		adminApi.DELETE("/delUser/:userId",admin.DelUser)
		adminApi.GET("/getUser/:userName",admin.GetUserInfo)
		adminApi.GET("/getAllUser",admin.GetAllUser)
		adminApi.GET("/getAdminUrl",admin.GetAdminPath)
		adminApi.POST("/addUserPath",admin.PostUserPath)
		adminApi.POST("/addAdminPath",admin.PostAdminPath)
	}
	driveApi := app.Group("/api/v1/drive").Use(adminJwt.MiddlewareFunc())
	{
		driveApi.GET("/getAllDrive",drive.GetAllDrive)
		driveApi.GET("/getUserDrive/:userName",drive.GetDriveByUserName)
		driveApi.POST("/postDrive",drive.PostDriveInfo)
		driveApi.PUT("/putDrive",drive.PutDrive)
		driveApi.DELETE("/delDrive/:driveId",drive.DelDrive)
	}
	userApi := app.Group("/user").Use(userJwt.MiddlewareFunc())
	{
		userApi.GET("/info/:userName",user.GetUserinfo)
		userApi.GET("/getUserPath",user.GetUserPath)
	}
	return app
}
