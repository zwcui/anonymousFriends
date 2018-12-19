package controllers

import (
	"anonymousFriends/util"
	"net/http"
	"strings"
	"io/ioutil"
)

//公共模块
type PublicController struct {
	apiController
}

//当前api请求之前调用，用于配置哪些接口需要进行head身份验证
func (this *PublicController) Prepare(){
	this.NeedBaseAuthList = []RequestPathAndMethod{}
	this.bathAuth()
	util.Logger.Info("PublicController beforeRequest ")
}

// @Title 测试http请求
// @Description 测试http请求
// @Param	url				formData		string  		true		"url"
// @Param	method			formData		string  		true		"method"
// @Param	headerKey1		formData		string  		false		"headerKey1"
// @Param	headerValue1	formData		string  		false		"headerValue1"
// @Param	headerKey2		formData		string  		false		"headerKey2"
// @Param	headerValue2	formData		string  		false		"headerValue2"
// @Param	headerKey3		formData		string  		false		"headerKey3"
// @Param	headerValue3	formData		string  		false		"headerValue3"
// @Success 200 {string} success
// @Failure 403 create users failed
// @router /testHttpRequest [post]
func (this *PublicController) TestHttpRequest() {
	url := this.MustString("url")
	method := this.MustString("method")
	headerKey1 := this.GetString("headerKey1", "")
	headerValue1 := this.GetString("headerValue1", "")
	headerKey2 := this.GetString("headerKey2", "")
	headerValue2 := this.GetString("headerValue2", "")
	headerKey3 := this.GetString("headerKey3", "")
	headerValue3 := this.GetString("headerValue3", "")


	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader("name=cjb"))
	if err != nil {
		util.Logger.Info("--http.NewRequest---err:"+err.Error())
		this.ReturnData = "--http.NewRequest---err:"+err.Error()
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if headerKey1 != "" {
		req.Header.Set(headerKey1, headerValue1)
	}
	if headerKey2 != "" {
		req.Header.Set(headerKey2, headerValue2)
	}
	if headerKey3 != "" {
		req.Header.Set(headerKey3, headerValue3)
	}


	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		util.Logger.Info("--ioutil.ReadAll---err:"+err.Error())
		this.ReturnData = "--ioutil.ReadAll---err:"+err.Error()
		return
	}

	this.ReturnData = string(body)
}
