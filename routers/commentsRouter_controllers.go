package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["anonymousFriends/controllers:MessageController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:MessageController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"] = append(beego.GlobalControllerRouter["anonymousFriends/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
