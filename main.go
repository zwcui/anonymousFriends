package main

import (
	_ "baseApi/routers"
	"github.com/astaxie/beego"
	_ "baseApi/models"
)


/*
	1. govendor add +e 将项目使用到但为加入vendor的包加入工程，参考vendor.json
	2. xorm 同步结构体与表结构，默认驼峰  其他参考： https://www.kancloud.cn/kancloud/xorm-manual-zh-cn/56013
	3. 使用swagger 第一次执行  bee run -gendoc=true -downdoc=true
	4. 如果启动工程时想执行其他包中的init()方法，则引入进来，前面加 "_ "
 */
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
