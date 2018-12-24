package models

import "encoding/json"

/*	1. 默认float32和float64映射到数据库中为float,real,double这几种类型，这几种数据库类型数据库的实现一般都是非精确的
		如果一定要作为查询条件，请将数据库中的类型定义为Numeric或者Decimal  `xorm:"Numeric"`
	2. 复合主键  Id(xorm.PK{1, 2})

*/

const ZOMBIE_USER_NUMBER = 100

type User struct {
	UId       			int64			`description:"uId" json:"uId" xorm:"pk autoincr"`
	Avatar				string			`description:"头像" json:"avatar"`
	PhoneNumber			string			`description:"手机号" json:"phoneNumber"`
	NickName 			string			`description:"昵称" json:"nickName" xorm:"notnull "`		//string类型默认映射为varchar(255)
	Password 			string			`description:"密码" json:"password" xorm:"notnull"`
	Salt	 			string			`description:"密码" json:"salt" xorm:"notnull"`
	Gender        		int    			`description:"性别,1 男, 2 女" json:"gender" xorm:"notnull default 0"`
	Birthday        	string			`description:"出生年月" json:"birthday"`
	Status	        	int				`description:"在线状态，0离线，1在线，2隐身，3Q我吧" json:"status"`
	Province        	string			`description:"省" json:"province"`
	City	        	string			`description:"市" json:"city"`
	Area	        	string			`description:"区" json:"area"`
	Longitude			float64			`description:"经度" json:"longitude"`
	Latitude			float64			`description:"纬度" json:"latitude"`
	IsZombie			int				`description:"是否僵尸账户，1是0否" json:"isZombie"`
	Created           	int64  			`description:"注册时间" json:"created" xorm:"created"`
	Updated           	int64  			`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         	int64  			`description:"删除时间" json:"deleted" xorm:"deleted"`
}

//用户登录信息
type UserSignInDeviceInfo struct {
	UId           		int64  			`description:"uId" json:"uId" xorm:"pk"`
	System       		int    			`description:"设备系统类型 0 未知 1 android 2 ios 3 h5" json:"system"`
	Manufacturers 		int    			`description:"厂商 0 未知 1 华为 2 魅族 3 小米" json:"manufacturers" xorm:"notnull"`
	DeviceToken   		string 			`description:"deviceToken" json:"deviceToken" xorm:"varchar(70)"`
	DeviceModel   		string 			`description:"设备型号" json:"deviceModel" xorm:"varchar(50)"`
	SystemVersion 		string 			`description:"设备系统版本" json:"systemVersion" xorm:"varchar(30)"`
	AppVersion    		string 			`description:"app版本" json:"appVersion" xorm:"varchar(30) notnull"`
	Created       		int64  			`description:"创建时间" json:"created" xorm:"created"`
	Updated       		int64  			`description:"修改时间" json:"updated" xorm:"updated"`
}

//账户
type UserAccount struct {
	UId       			int64			`description:"注册时间" json:"uId" xorm:"pk"`
	CashBalance			int				`description:"现金余额，单位分" json:"cashBalance"`
	Created       		int64  			`description:"创建时间" json:"created" xorm:"created"`
	Updated       		int64  			`description:"修改时间" json:"updated" xorm:"updated"`
}

// 收支明细
type AccountTransactionRecord struct {
	Id	       			int64  			`description:"记录id" json:"id" xorm:"pk autoincr"`
	UId        			int64  			`description:"用户id" json:"uId"`
	Money      			int    			`description:"金额" json:"money"`
	MoneyType  			int    			`description:"金额类型，1:收入,-1:支出" json:"moneyType"`
	RecordType 			int    			`description:"记录类型，" json:"recordType"`
	RecordName 			string 			`description:"记录类型，" json:"recordName"`
	OrderId    			int64  			`description:"订单编号 " json:"orderId" xorm:"notnull default 0"`
	Created    			int64  			`description:"生成时间" json:"created" xorm:"created"`
	Updated    			int64   		`description:"更新时间" xorm:"updated" json:"updated"`
	Status     			int    			`description:"交易状态 0:交易未执行 1:交易已执行 2:交易取消" json:"status" xorm:"notnull default 1"`
}

//默认昵称
type DefaultNickName struct {
	Id	       			int64  			`description:"记录id" json:"id" xorm:"pk autoincr"`
	NickName			string 			`description:"默认昵称" json:"nickName"`
	Status     			int    			`description:"状态 0未使用，1已使用" json:"status" xorm:"notnull default 0"`
}

//-------------结构体如下---------------

type UserShort struct {
	UId       			int64			`description:"注册时间" json:"uId" xorm:"pk autoincr"`
	Avatar				string			`description:"头像" json:"avatar"`
	PhoneNumber			string			`description:"手机号" json:"phoneNumber"`
	NickName 			string			`description:"昵称" json:"nickName" xorm:"notnull "`		//string类型默认映射为varchar(255)
	Gender        		int    			`description:"性别,1 男, 2 女" json:"gender" xorm:"notnull default 0"`
	Birthday        	string			`description:"出生年月" json:"birthday"`
	Status	        	int				`description:"在线状态，0离线，1在线，2隐身" json:"status"`
	Province        	string			`description:"省" json:"province"`
	City	        	string			`description:"市" json:"city"`
	Area	        	string			`description:"区" json:"area"`
	Longitude			float64			`description:"经度" json:"longitude"`
	Latitude			float64			`description:"纬度" json:"latitude"`
	Created           	int64  			`description:"注册时间" json:"created" xorm:"created"`
	Updated           	int64  			`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         	int64  			`description:"删除时间" json:"deleted" xorm:"deleted"`
}

type SignInUser struct {
	User				UserShort		`description:"登录用户信息" json:"user"`
}

type UserInfo struct {
	User				UserShort		`description:"用户信息" json:"user"`
}

type UserList struct {
	UserList			[]UserShort		`description:"用户信息" json:"userList"`
}

type UserAccountInfo struct {
	UserAccount			UserAccount		`description:"用户账户信息" json:"userAccount"`
}

//-------------user方法如下--------------

func (u *User) UsetToUserShort() (userDTO *UserShort, error error) {
	josnByte, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	userD := UserShort{}
	err = json.Unmarshal(josnByte, &userD)
	if err != nil {
		return nil, err
	}
	return &userD, nil
}

//--------------默认参数----------------

var DefaultGender = []int{
	1,	//男
	2,	//女
}

var DefaultBirthday = []string{
	"199301",
	"199503",
	"199002",
	"198808",
	"198011",
	"198606",
	"198209",
	"198910",
	"199011",
	"199412",
}

var DefaultDirection = []float64{
	1.0,
	-1.0,
}

var DefaultAvatar = []string{
	"da1",
	"da2",
	"da3",
}

var ZombieMoveFlagRatio = []int{
	1,
	0,
}