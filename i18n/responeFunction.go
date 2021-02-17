package i18n

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func DataResponse(c *gin.Context,data interface{})  {
	c.JSON(200,gin.H{"code":200,"path":c.FullPath(),"msg":"response success","data":data})
}

func ResponseFail(c *gin.Context,msg string)  {
	c.JSON(402,gin.H{"code":402,"msg":msg +",request fail"})
}

func ResponseSuccess(c *gin.Context,msg string)  {
	c.JSON(200,gin.H{"code":200,"path":c.FullPath(),"msg":msg + ",request success"})
}

func BindJsonFail(c *gin.Context,err error)  {
	c.JSON(403,gin.H{"code":403,"msg":fmt.Sprintln("参数解析失败",err)})
}