package middleware

import (
	"NewRmsService/model"
	"NewRmsService/service"
	"errors"
	"fmt"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var IdentityKey = "userName"

type UserJwt struct {
	IUserService service.IUserService
}

type AuthFunction func(data interface{}, c *gin.Context) bool

func (uj UserJwt)InitUserJwtFunction(function AuthFunction)(userJwt *jwt.GinJWTMiddleware){
	userJwt,err:= jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.UserName,
					"userRole":v.UserRole,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				UserName: claims[IdentityKey].(string),
				UserRole:claims["userRole"].(string),
			}
		},
		Authorizator:function,
		Authenticator: func(c *gin.Context) (interface{}, error) {

			var userLogin model.UserLoginModel
			if err:= c.ShouldBindJSON(&userLogin);err !=nil{
				fmt.Println(userLogin)
				return nil, errors.New("该用户无法登录")
			}
			isOk,err:=uj.IUserService.CheckLogin(userLogin)
			 if err !=nil{
				 return nil, jwt.ErrFailedAuthentication
			 }
			 if isOk{
			 	roleName,_ := uj.IUserService.FindUserRole(userLogin.UserName)
				 return &model.User{
					 UserName:userLogin.UserName,
					 UserRole: roleName,
				 }, nil
			 }
			return nil, jwt.ErrFailedAuthentication
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err !=nil{
		log.Fatalln("init jwt fail")
	}
	return
}

func AllAuth(data interface{}, c *gin.Context) bool{
	return true
}

func AdminAuth (data interface{}, c *gin.Context) bool {
	if v, ok := data.(*model.User); ok && v.UserRole == "admin" {
		return true
	}
	return false
}
func NoRouteHandler(c *gin.Context) {
	c.JSON(404, gin.H{"code": 404, "message": "访问的路径不存在"})
}