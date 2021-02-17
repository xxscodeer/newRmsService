package vaildata

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func NameValid(fl validator.FieldLevel) bool{
	if s,ok:=fl.Field().Interface().(string);ok{
		switch s {
		case "admin":
			return true
		}
	}
	return false
}

func RunBindVail() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_=v.RegisterValidation("name_valid",NameValid)
	}
}
