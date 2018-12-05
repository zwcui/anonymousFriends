package main

import (
	_ "baseApi/routers"
	"github.com/astaxie/beego"
	_ "baseApi/models"
	_ "baseApi/base"
	"baseApi/util"
)


/*
	1. govendor add +e 将项目使用到但为加入vendor的包加入工程，参考vendor.json
	2. xorm 同步结构体与表结构，默认驼峰  其他参考： https://www.kancloud.cn/kancloud/xorm-manual-zh-cn/56013
	3. 使用swagger 第一次执行  bee run -gendoc=true -downdoc=true
	4. 如果启动工程时想执行其他包中的init()方法，则引入进来，前面加 "_ "
	5. 为了对老版本的兼容，docker部署多个容器，由请求head中的api-version区分，通过nginx进行不同端口的跳转；可能同一套数据库多个服务会产生新版本更新时老版本仍需更新，配置主从数据库？
	6. 域名解析时，注意接口前缀如api，网页前缀如www
	7. 考虑加入gRPC，远程调用，方便创建分布式应用
	8. controller.go中，初始化后请求前调用Prepare()，请求后调用Finish()等等
	9. 集成了seelog日志
	10.deploy_dev.sh 用于部署开发服务器docker，取git最新程序同步至服务器；
	   deploy_test.sh用于部署测试服务器docker，取git最新程序同步至服务器；
       deploy_prod.sh用于部署正式服务器docker，取git最新程序同步至服务器；
 */
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	defer util.Logger.Flush()

	beego.Run()
}
