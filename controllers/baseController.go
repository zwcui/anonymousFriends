package controllers

import (
	"github.com/astaxie/beego"
	"regexp"
	"strings"
	_ "github.com/astaxie/beego/cache/redis"
	"baseApi/util"
	"time"
	"baseApi/base"
)

//数据返回结构体
type baseController struct {
	beego.Controller
	Err        error
	ErrCode    int
	ReturnData interface{}
}

//api接口统一controller
type apiController struct {
	baseController
	NeedBaseAuthList []RequestPathAndMethod
}

//需要验证的请求路径
type RequestPathAndMethod struct {
	PathRegexp string
	Method     string
}

const (
	REDIS_BATHAUTH = "BaseAuth_"
)

//默认需要验证的请求路径
func (this *apiController) prepare(){
	this.NeedBaseAuthList = []RequestPathAndMethod{{".+", "post"}, {".+", "patch"}, {".+", "delete"}, {".+", "put"}}
	this.bathAuth()
}

//对路径进行校验
func (this *apiController) bathAuth(){
	pathNeedAuth := false
	for _, value := range this.NeedBaseAuthList {
		if ok, _ := regexp.MatchString(value.PathRegexp, this.Ctx.Request.URL.Path); ok && strings.ToUpper(this.Ctx.Request.Method) == strings.ToUpper(value.Method) {
			pathNeedAuth = true
			break
		}
	}

	//要求head中放Authorization，内容格式为 "Basic 18800000000:123456"  密码为加密后的密文，加密方式为base64
	if pathNeedAuth {
		phoneNumber, encryptedPassword, ok := this.Ctx.Request.BasicAuth()
		if !ok {
			w := this.Ctx.ResponseWriter
			w.Header().Set("WWW-Authenticate", `Base Auth failed : "`+"empty auth"+`"`)
			w.WriteHeader(401)
			w.Write([]byte("401 Unauthorized\n"))
			this.ServeJSON()
			this.StopRun()
		}
		redisTemp := base.RedisCache.Get(REDIS_BATHAUTH+phoneNumber)
		if redisTemp == nil {
			user, err := UserWithPhoneNumber(phoneNumber)
			if err != nil || user == nil {
				w := this.Ctx.ResponseWriter
				w.Header().Set("WWW-Authenticate", `Base Auth failed : "`+err.Error()+`"`)
				w.WriteHeader(401)
				w.Write([]byte("401 Unauthorized\n"))
				this.ServeJSON()
				this.StopRun()
			}
			//校验密码
			passwordByte := util.Base64Encode([]byte(encryptedPassword))
			password := string(passwordByte)
			hashedPwd, _ := util.EncryptPasswordWithSalt(password, user.Salt)
			if hashedPwd != user.Password {
				w := this.Ctx.ResponseWriter
				w.Header().Set("WWW-Authenticate", `Base Auth failed : "`+"password error"+`"`)
				w.WriteHeader(401)
				w.Write([]byte("401 Unauthorized\n"))
				this.ServeJSON()
				this.StopRun()
			}

			//存入redis
			if user != nil {
				base.RedisCache.Put(REDIS_BATHAUTH+phoneNumber, encryptedPassword, 60*60*2*time.Second)
			}
		} else {
			if encryptedPassword != redisTemp {
				w := this.Ctx.ResponseWriter
				w.Header().Set("WWW-Authenticate", `Base Auth failed : "`+"password redis error"+`"`)
				w.WriteHeader(401)
				w.Write([]byte("401 Unauthorized\n"))
				this.ServeJSON()
				this.StopRun()
			}
		}

	}



}