/*
@Time : 2019/1/8 下午5:33 
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"image"
	"image/color"
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
	//util.Logger.Info(util.GetCurrentHour(util.UnixOfBeijingTime()))
	//
	//
	//user := "service@mail.zwcui.cn"
	//password := ""
	//host := "smtpdm.aliyun.com:465"
	//to := "747660511@qq.com"
	//
	//subject := "使用Golang发送邮件"
	//
	//body := `
	//	<html>
	//	<body>
	//	<h3>
	//	"Test send to email"
	//	</h3>
	//	</body>
	//	</html>
	//	`
	//util.Logger.Info("send email")
	//err := SendToMail(user, password, host, to, subject, body, "html")
	//if err != nil {
	//	util.Logger.Info("Send mail error!")
	//	util.Logger.Info(err)
	//} else {
	//	util.Logger.Info("Send mail success!")
	//}





	//file, err := os.Create("/Users/youbie/Documents/test/test.jpeg")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer file.Close()
	//
	//file1, err := os.Open("/Users/youbie/Documents/test/1.jpg")	//背景图
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer file1.Close()
	//img, _ := jpeg.Decode(file1)
	//
	//file2, err := os.Open("/Users/youbie/Documents/test/2.png")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer file2.Close()
	//img2, _ := png.Decode(file2)
	//
	//if img2 == nil {
	//	util.Logger.Info("ssssssssss")
	//}
	//
	//jpg := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
	//
	//draw.Draw(jpg, jpg.Bounds(), img, img.Bounds().Min, draw.Over)                   //首先将一个图片信息存入jpg
	////draw.Draw(jpg, jpg.Bounds(), img2, img2.Bounds().Min.Sub(image.Pt(10, 20)), draw.Over)   //将另外一张图片信息存入jpg
	//draw.DrawMask(jpg, jpg.Bounds(), img2, image.ZP, &circle{image.Pt(100,200), 90}, image.ZP, draw.Over)


	//draw.DrawMask(jpg, jpg.Bounds(), img2, img2.Bounds().Min, img, img.Bounds().Min, draw.Src) // 利用这种方法不能够将两个图片直接合成？目前尚不知道原因。

	//jpeg.Encode(file, jpg, nil)








	//设置每个点的 RGBA (Red,Green,Blue,Alpha(设置透明度))
	//for y:=0;y<img.Bounds().Dy();y++ {
	//	for x:=0;x<img.Bounds().Dx();x++ {
	//		//设置一块 白色(255,255,255)不透明的背景
	//		jpg.Set(x,y,color.RGBA{255,255,255,255})
	//	}
	//}
	//读取字体数据
	//fontBytes,err := ioutil.ReadFile("/Library/Fonts/Arial Unicode.ttf")
	//if err != nil {
	//	util.Logger.Info(err.Error())
	//}
	////载入字体数据
	//font,err := freetype.ParseFont(fontBytes)
	//if err != nil {
	//	util.Logger.Info(err.Error())
	//}
	//f := freetype.NewContext()
	////设置分辨率
	//f.SetDPI(72)
	////设置字体
	//f.SetFont(font)
	////设置尺寸
	//f.SetFontSize(50)
	//f.SetClip(img.Bounds())
	////设置输出的图片
	//f.SetDst(jpg)
	////设置字体颜色(红色)
	//f.SetSrc(image.NewUniform(color.RGBA{255,0,0,255}))
	//
	////设置字体的位置
	//pt := freetype.Pt(40,40 + int(f.PointToFixed(26)) >> 8)
	//
	//_,err = f.DrawString("hello,世界,1234567890hajbcv今晚上课had王健林经理",pt)
	//if err != nil {
	//	util.Logger.Info(err.Error())
	//}
	//
	////以png 格式写入文件
	////err = png.Encode(imgfile,img)
	////if err != nil {
	////	util.Logger.Info(err.Error())
	////}
	//
	//
	//jpeg.Encode(file, jpg, nil)






















}

type circle struct {
     p image.Point
     r int
}

func (c *circle) ColorModel() color.Model {
     return color.AlphaModel
 }

func (c *circle) Bounds() image.Rectangle {
     return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
     xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
     if xx*xx+yy*yy < rr*rr {
         return color.Alpha{255}
     }
     return color.Alpha{0}
}