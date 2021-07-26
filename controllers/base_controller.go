package controllers

import "acp10/models/base"

func BaseResponse(code int, message string, data interface{}) interface{} {
	return base.BaseResponseData{code, message, data}
}
