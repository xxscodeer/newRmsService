package main

import (
	"NewRmsService/route"
	"NewRmsService/tools"
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
var cfg *tools.Config
func init() {
	fmt.Println("project config init..")
	cfg = tools.ParseConfig("./conf/conf.yaml")
	tools.InitOrmEngine(cfg.Mysql)
	fmt.Println("config init end...")
}

// @title 掌悦 rms资源管理系统
// @version 1.0
// @description rms资源管理系统后端

// @contact.name API Support
// @contact.url http://www.swagger.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8999
// @BasePath
func main() {
	app := route.InitAppRouter()

	// swagger 后端文档api
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	_ = app.Run(cfg.App.Host + ":" + cfg.App.Port)
}
