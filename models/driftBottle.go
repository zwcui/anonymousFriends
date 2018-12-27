/*
@Time : 2018/12/26 下午3:16 
@Author : zwcui
@Software: GoLand
*/
package models

//漂流瓶有效期，单位小时，为0表示无有效期
const ExpiryTime = 0

//漂流瓶
type DriftBottle struct {
	BottleId       				int64			`description:"bottleId" json:"bottleId" xorm:"pk autoincr"`
	BottleType     				int				`description:"类型，1普通瓶、2传递瓶、3同城瓶、4真话瓶、5暗号瓶、6提问瓶、7交往瓶、8祝愿瓶、9发泄瓶、10生日瓶、11表白瓶" json:"bottleType"`
	BottleName     				int				`description:"类型名称，1普通瓶、2传递瓶、3同城瓶、4真话瓶、5暗号瓶、6提问瓶、7交往瓶、8祝愿瓶、9发泄瓶、10生日瓶、11表白瓶" json:"bottleName"`
	SenderUid					int64			`description:"抛瓶人" json:"senderUid"`
	ReceiverUid					int64			`description:"拾瓶人" json:"receiverUid"`
	Content						string			`description:"内容" json:"content"`
	Picture						string			`description:"图片内容，多个英文逗号隔开" json:"picture"`
	Position        			string			`description:"位置名称，如全家创意产业园店" json:"position"`
	Weather        				string			`description:"天气，如'阵雨转多云 19~25摄氏度'" json:"weather"`
	Province        			string			`description:"省" json:"province"`
	City	        			string			`description:"市" json:"city"`
	Area	        			string			`description:"区" json:"area"`
	Longitude					float64			`description:"经度" json:"longitude"`
	Latitude					float64			`description:"纬度" json:"latitude"`
	Status						int				`description:"状态，0未抛出，1已抛出，2已接收，3已失效" json:"status"`
	ExpiryTime					int64			`description:"到期时间" json:"expiryTime"`
	Remark						string			`description:"备注" json:"remark"`
	ReplyNum					int				`description:"回复数" json:"replyNum"`
	Created           			int64  			`description:"创建时间" json:"created" xorm:"created"`
	Updated           			int64  			`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         			int64  			`description:"删除时间" json:"deleted" xorm:"deleted"`
}

//僵尸漂流瓶
type ZombieDriftBottle struct {
	BottleId       				int64			`description:"bottleId" json:"bottleId" xorm:"pk autoincr"`
	BottleType     				int				`description:"类型，1普通瓶、2传递瓶、3同城瓶、4真话瓶、5暗号瓶、6提问瓶、7交往瓶、8祝愿瓶、9发泄瓶、10生日瓶、11表白瓶" json:"bottleType"`
	BottleName     				int				`description:"类型名称，1普通瓶、2传递瓶、3同城瓶、4真话瓶、5暗号瓶、6提问瓶、7交往瓶、8祝愿瓶、9发泄瓶、10生日瓶、11表白瓶" json:"bottleName"`
	Content						string			`description:"内容" json:"content"`
	Picture						string			`description:"图片内容，多个英文逗号隔开" json:"picture"`
	Position        			string			`description:"位置名称，如全家创意产业园店" json:"position"`
	Weather        				string			`description:"天气，如'阵雨转多云 19~25摄氏度'" json:"weather"`
	Province        			string			`description:"省" json:"province"`
	City	        			string			`description:"市" json:"city"`
	Area	        			string			`description:"区" json:"area"`
	Longitude					float64			`description:"经度" json:"longitude"`
	Latitude					float64			`description:"纬度" json:"latitude"`
	Status						int				`description:"状态，0未使用，1已使用" json:"status"`
	Created           			int64  			`description:"创建时间" json:"created" xorm:"created"`
	Updated           			int64  			`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         			int64  			`description:"删除时间" json:"deleted" xorm:"deleted"`
}

//---------------------结构体如下-------------------------

type DriftBottleInfo struct {
	DriftBottle					DriftBottle		`description:"漂流瓶信息" json:"driftBottle"`
	CommentList					[]CommentInfo	`description:"评论列表" json:"commentList"`
}

type DriftBottleListContainer struct {
	BaseListContainer
	DriftBottleList 			[]DriftBottle 	`description:"漂流瓶列表" json:"driftBottleList"`
}