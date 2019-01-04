// @APIVersion 1.0.0
// @Title anonymousFriends 接口基础工程
// @Description 接入Redis,MongoDB,xorm,seelog等，加入部署脚本
// @Contact zwcui2017@163.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"anonymousFriends/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/public",
			beego.NSInclude(
				&controllers.PublicController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/message",
			beego.NSInclude(
				&controllers.MessageController{},
			),
		),
		beego.NSNamespace("/socialDynamics",
			beego.NSInclude(
				&controllers.SocialDynamicsController{},
			),
		),
		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),
		beego.NSNamespace("/friend",
			beego.NSInclude(
				&controllers.FriendController{},
			),
		),
		beego.NSNamespace("/push",
			beego.NSInclude(
				&controllers.PushController{},
			),
		),
		beego.NSNamespace("/driftBottle",
			beego.NSInclude(
				&controllers.DriftBottleController{},
			),
		),
		beego.NSNamespace("/tag",
			beego.NSInclude(
				&controllers.TagController{},
			),
		),
		beego.NSNamespace("/adminNotice",
			beego.NSInclude(
				&controllers.AdminNoticeController{},
			),
		),
		beego.NSNamespace("/sharePosition",
			beego.NSInclude(
				&controllers.SharePositionController{},
			),
		),
		beego.NSNamespace("/report",
			beego.NSInclude(
				&controllers.ReportController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
