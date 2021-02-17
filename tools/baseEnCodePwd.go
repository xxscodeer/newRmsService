package tools

import (
	"encoding/base64"
)

func DecodePwd(pwd string)([]byte,error) {
	return base64.StdEncoding.DecodeString(pwd)

}

func EncodePwd(pwd string)string {
	bytePwd := []byte(pwd)
	return base64.StdEncoding.EncodeToString(bytePwd)
}
