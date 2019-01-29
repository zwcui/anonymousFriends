/*
@Time : 2019/1/8 下午5:33 
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"anonymousFriends/util"
)

//调试模块
type TestController struct {
	apiController
}

func (this *TestController) Prepare(){
	this.NeedBaseAuthList = []RequestPathAndMethod{
	}
	this.bathAuth()
}

// @Title 测试高德逆地理编码
// @Description 测试高德逆地理编码
// @Param	longitude			query			float64	  		true		"longitude"
// @Param	latitude			query			float64	  		true		"latitude"
// @Success 200 {object} models.RegeoResponse
// @router /testAmapRegeoApi [get]
func (this *TestController) TestAmapRegeoApi() {
	longitude := this.MustFloat64("longitude")
	latitude := this.MustFloat64("latitude")

	regeoResponse, err := GetRegeocode(longitude, latitude)
	if err != nil {
		this.ReturnData = err.Error()
	} else {
		this.ReturnData = regeoResponse
	}
}

// @Title 测试高德天气查询
// @Description 测试高德天气查询
// @Param	province			query			string	  		false		"province"
// @Param	city				query			string	  		false		"city"
// @Param	area				query			string	  		false		"area"
// @Param	longitude			query			float64	  		false		"longitude"
// @Param	latitude			query			float64	  		false		"latitude"
// @Success 200 {object} models.RegeoResponse
// @router /testAmapWeatherApi [get]
func (this *TestController) TestAmapWeatherApi() {
	province := this.GetString("province", "")
	city := this.GetString("city", "")
	area := this.GetString("area", "")
	longitude, _ := this.GetFloat("longitude", 0)
	latitude, _ := this.GetFloat("latitude", 0)

	weatherResponse, err := GetCurrentWeather(province, city, area, longitude, latitude)
	if err != nil {
		this.ReturnData = err.Error()
	} else {
		this.ReturnData = weatherResponse
	}
}

// @Title 测试其他
// @Description 测试其他
// @Success 200 {string}
// @router /test [get]
func (this *TestController) Test() {
	util.Logger.Info(util.GetCurrentHour(util.UnixOfBeijingTime()))


	user := "service@mail.zwcui.cn"
	password := ""
	host := "smtpdm.aliyun.com:465"
	to := "747660511@qq.com"

	subject := "使用Golang发送邮件"

	body := `
		<html>
		<body>
		<h3>
		"Test send to email"
		</h3>
		</body>
		</html>
		`
	util.Logger.Info("send email")
	err := SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		util.Logger.Info("Send mail error!")
		util.Logger.Info(err)
	} else {
		util.Logger.Info("Send mail success!")
	}

}