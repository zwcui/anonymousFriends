package controllers

import "github.com/astaxie/beego"

//基础controller
type baseController struct {
	beego.Controller
	Err        error
	ErrCode    int
	ReturnData interface{}
}

type apiController struct {
	baseController
	NeedBaseAuthList []RequestPathAndMethod
}

//请求路径
type RequestPathAndMethod struct {
	PathRegexp string
	Method     string
}