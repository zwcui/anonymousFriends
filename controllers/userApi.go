package controllers

import (
	"anonymousFriends/base"
	"anonymousFriends/models"
	"anonymousFriends/util"
	"strconv"
	"math/rand"
	"github.com/satori/go.uuid"
	"time"
)

//用户模块
type UserController struct {
	apiController
}

//当前api请求之前调用，用于配置哪些接口需要进行head身份验证
func (this *UserController) Prepare(){
	//this.NeedBaseAuthList = []RequestPathAndMethod{{".+", "post"}, {".+", "patch"}, {".+", "delete"}}
	this.NeedBaseAuthList = []RequestPathAndMethod{{"/updateUserInfo","patch"}, {"/updateUserPassword","patch"}, {"/updateUserPosition","patch"}, {"getUserAccountInfo", "get"}}
	this.bathAuth()
	util.Logger.Info("UserController beforeRequest ")
}

// @Title 注册
// @Description 注册
// @Param	nickName		formData		string  		true		"昵称"
// @Param	avatar			formData		string  		false		"头像"
// @Param	phoneNumber		formData		string  		false		"手机号"
// @Param	password		formData		string  		true		"密码"
// @Param   veriCode        formData        string  		false       "veriCode"
// @Param   needVeriCode    formData        string  		false       "是否需要验证码，需要为1，不需要为0"
// @Param	gender			formData		int		  		false		"性别,1 男, 2 女"
// @Param	birthday		formData		string  		false		"出生年月"
// @Param   system         	formData        int     		false       "系统类型 1:android 2:ios 3:h5"
// @Param   deviceToken     formData    	string  		false       "deviceToken"
// @Param   deviceModel     formData    	string  		false       "设备型号"
// @Param   systemVersion   formData    	string  		false       "systemVersion"
// @Param   appVersion      formData    	string  		false       "app版本号"
// @Param   manufacturers   formData    	int     		false       "厂商 1:华为 2:魅族 3:小米"
// @Success 200 {string} success
// @Failure 403 create users failed
// @router /signUp [post]
func (this *UserController) SignUp() {
	nickName := this.MustString("nickName")
	avatar := this.GetString("avatar", "")
	phoneNumber := this.GetString("phoneNumber", "")
	password := this.MustString("password")
	gender, _ := this.GetInt("gender", 2)
	birthday := this.GetString("birthday", "")
	system, _ := this.GetInt("system", 0)
	deviceToken := this.GetString("deviceToken", "")
	deviceModel := this.GetString("deviceModel", "")
	systemVersion := this.GetString("systemVersion", "")
	appVersion := this.GetString("appVersion", "")
	manufacturers, _ := this.GetInt("manufacturers", 0)

	//昵称唯一
	isSensitive, hasSameNickNameUser := checkSameNickName(nickName)
	if isSensitive {
		this.ReturnData = util.GenerateAlertMessage(models.UserError105)
		return
	}
	if hasSameNickNameUser {
		this.ReturnData = util.GenerateAlertMessage(models.UserError100)
		return
	}

	if len(password) < 6 || len(password) > 30 {
		this.ReturnData = util.GenerateAlertMessage(models.UserError104)
		return
	}

	var user models.User
	user.Avatar = avatar
	user.NickName = nickName
	user.PhoneNumber = phoneNumber
	user.Gender = gender
	user.Birthday = birthday
	user.Status = 1
	hashedPassword, salt, err := util.EncryptPassword(password)
	if err != nil {
		this.ReturnData = util.GenerateAlertMessage(models.UserError101)
		return
	}
	user.Password = hashedPassword
	user.Salt = salt
	base.DBEngine.Table("user").InsertOne(&user)

	//推送用表
	var userSignInDeviceInfo models.UserSignInDeviceInfo
	userSignInDeviceInfo.UId = user.UId
	userSignInDeviceInfo.System = system
	userSignInDeviceInfo.Manufacturers = manufacturers
	userSignInDeviceInfo.DeviceToken = deviceToken
	userSignInDeviceInfo.DeviceModel = deviceModel
	userSignInDeviceInfo.SystemVersion = systemVersion
	userSignInDeviceInfo.AppVersion = appVersion
	base.DBEngine.Table("user_sign_in_device_info").InsertOne(&userSignInDeviceInfo)

	//账户表
	var userAccount models.UserAccount
	userAccount.CashBalance = 0
	base.DBEngine.Table("user_account").InsertOne(&userAccount)

	//用户注册通知管理员
	var adminNotice models.AdminNotice
	adminNotice.Type = 2
	adminNotice.TypeId = user.UId
	adminNotice.Content = "新用户注册："+user.NickName+"(uId:"+strconv.FormatInt(user.UId, 10)+")"
	adminNotice.Status = 0
	base.DBEngine.Table("admin_notice").InsertOne(&adminNotice)

	userShort, _ := user.UsetToUserShort()
	this.ReturnData = models.UserInfo{*userShort}
}


// @Title 登录
// @Description 登录
// @Param	nickName		formData		string  		true		"昵称"
// @Param	phoneNumber		formData		string  		false		"手机号"
// @Param	password		formData		string  		true		"密码"
// @Param   system         	formData        int     		false       "系统类型 1:android 2:ios 3:h5"
// @Param   deviceToken     formData    	string  		false       "deviceToken"
// @Param   deviceModel     formData    	string  		false       "设备型号"
// @Param   systemVersion   formData    	string  		false       "systemVersion"
// @Param   appVersion      formData    	string  		false       "app版本号"
// @Param   manufacturers   formData    	int     		false       "厂商 1:华为 2:魅族 3:小米"
// @Success 200 {object} models.SignInUser
// @router /signIn [post]
func (this *UserController) SignIn() {
	nickName := this.MustString("nickName")
	//phoneNumber := this.GetString("phoneNumber", "")
	password := this.MustString("password")
	system, _ := this.GetInt("system", 0)
	deviceToken := this.GetString("deviceToken", "")
	deviceModel := this.GetString("deviceModel", "")
	systemVersion := this.GetString("systemVersion", "")
	appVersion := this.GetString("appVersion", "")
	manufacturers, _ := this.GetInt("manufacturers", 0)

	var storedUser models.User
	hasStoredUser, _ := base.DBEngine.Table("user").Where("nick_name=?", nickName).Get(&storedUser)
	if !hasStoredUser {
		this.ReturnData = util.GenerateAlertMessage(models.UserError102)
		return
	}

	hashedPassword, _ := util.EncryptPasswordWithSalt(password, storedUser.Salt)
	if hashedPassword != storedUser.Password {
		this.ReturnData = util.GenerateAlertMessage(models.UserError103)
		return
	}

	storedUser.Status = 1
	base.DBEngine.Table("user").Where("u_id=?", storedUser.UId).Cols("status").Update(&storedUser)

	//推送用表
	var userSignInDeviceInfo models.UserSignInDeviceInfo
	base.DBEngine.Table("user_sign_in_device_info").Where("u_id=?", storedUser.UId).Get(&userSignInDeviceInfo)
	userSignInDeviceInfo.System = system
	userSignInDeviceInfo.Manufacturers = manufacturers
	userSignInDeviceInfo.DeviceToken = deviceToken
	userSignInDeviceInfo.DeviceModel = deviceModel
	userSignInDeviceInfo.SystemVersion = systemVersion
	userSignInDeviceInfo.AppVersion = appVersion
	base.DBEngine.Table("user_sign_in_device_info").Where("u_id=?", storedUser.UId).AllCols().Update(&userSignInDeviceInfo)

	userShort, _ := storedUser.UsetToUserShort()
	this.ReturnData = models.UserInfo{*userShort}
}

// @Title 获取用户详情
// @Description 获取用户详情
// @Param	uId				query			int64	  		true		"uId"
// @Success 200 {object} models.UserInfo
// @router /getUserInfo [get]
func (this *UserController) GetUserInfo() {
	uId := this.MustInt64("uId")

	var user models.UserShort
	base.DBEngine.Table("user").Where("u_id=?", uId).Get(&user)

	this.ReturnData = models.UserInfo{user}
}

// @Title 获取用户账户详情
// @Description 获取用户账户详情
// @Param	uId				query			int64	  		true		"uId"
// @Success 200 {object} models.UserAccountInfo
// @router /getUserAccountInfo [get]
func (this *UserController) GetUserAccountInfo() {
	uId := this.MustInt64("uId")

	var userAccount models.UserAccount
	base.DBEngine.Table("user_account").Where("u_id=?", uId).Get(&userAccount)

	this.ReturnData = models.UserAccountInfo{userAccount}
}

// @Title 更新用户信息
// @Description 更新用户信息
// @Param	uId				formData		int64	  		true		"uId"
// @Param	avatar			formData		string  		false		"头像"
// @Param	nickName		formData		string  		false		"昵称"
// @Param	phoneNumber		formData		string  		false		"手机号"
// @Param	gender			formData		int		  		false		"性别,1 男, 2 女"
// @Param	birthday		formData		string  		false		"出生年月"
// @Param	tagNames		formData		string  		false		"标签名称，多个,分隔"
// @Param	tagIds			formData		string  		false		"标签id，多个,分隔"
// @Param	status			formData		string  		true		"在线状态，0离线，1在线，2隐身，3Q我吧"
// @Success 200 {object} models.UserInfo
// @router /updateUserInfo [patch]
func (this *UserController) UpdateUserInfo() {
	uId := this.MustInt64("uId")
	avatar := this.GetString("avatar", "")
	nickName := this.GetString("nickName", "")
	phoneNumber := this.GetString("phoneNumber", "")
	birthday := this.GetString("birthday", "")
	tagNames := this.GetString("tagNames", "")
	tagIds := this.GetString("tagIds", "")
	gender, _ := this.GetInt("gender", 0)
	status := this.MustInt("status")

	//昵称唯一
	if nickName != "" {
		isSensitive, hasSameNickNameUser := checkSameNickName(nickName)
		if isSensitive {
			this.ReturnData = util.GenerateAlertMessage(models.UserError105)
			return
		}
		if hasSameNickNameUser {
			this.ReturnData = util.GenerateAlertMessage(models.UserError100)
			return
		}
	}

	var user models.User
	base.DBEngine.Table("user").Where("u_id=?", uId).Get(&user)
	if nickName != "" {
		oldNickName := user.NickName
		user.NickName = nickName

		base.RedisCache.Delete(REDIS_BATHAUTH + oldNickName)
		base.RedisCache.Put(REDIS_BATHAUTH + user.NickName, user.Password, 60*60*2*time.Second)
	}
	if avatar != "" {
		user.Avatar = avatar
	}
	if phoneNumber != "" {
		user.PhoneNumber = phoneNumber
	}
	if birthday != "" {
		user.Birthday = birthday
	}
	if gender != 0 {
		user.Gender = gender
	}
	if tagNames != "" {
		user.TagNames = tagNames
	}
	if tagIds != "" {
		user.TagIds = tagIds
	}

	user.Status = status
	base.DBEngine.Table("user").Where("u_id=?", uId).Cols("avatar", "nick_name", "phone_number", "birthday", "gender", "status", "tag_names", "tag_ids").Update(&user)

	userShort, _ := user.UsetToUserShort()
	this.ReturnData = models.UserInfo{*userShort}
}

// @Title 更新用户密码
// @Description 更新用户密码
// @Param	uId				formData		int64	  		true		"uId"
// @Param	oldPassword		formData		string  		true		"旧密码"
// @Param	newPassword		formData		string  		true		"新密码"
// @Success 200 {string} success
// @router /updateUserPassword [patch]
func (this *UserController) UpdateUserPassword() {
	uId := this.MustInt64("uId")
	oldPassword := this.MustString("oldPassword")
	newPassword := this.MustString("newPassword")

	var user models.User
	base.DBEngine.Table("user").Where("u_id=?", uId).Get(&user)

	hashedPassword, _ := util.EncryptPasswordWithSalt(oldPassword, user.Salt)
	if hashedPassword != oldPassword {
		this.ReturnData = util.GenerateAlertMessage(models.UserError103)
		return
	}

	hashedPassword, salt, err := util.EncryptPassword(newPassword)
	if err != nil {
		this.ReturnData = util.GenerateAlertMessage(models.UserError101)
		return
	}
	user.Password = hashedPassword
	user.Salt = salt
	base.DBEngine.Table("user").Where("u_id=?", uId).Cols("password", "salt").Update(&user)

	base.RedisCache.Put(REDIS_BATHAUTH + user.NickName, user.Password, 60*60*2*time.Second)

	this.ReturnData = "success"
}

// @Title 接口调用更新用户位置，心跳实现相同功能
// @Description 接口调用更新用户位置，心跳实现相同功能
// @Param	uId				formData		int64	  		true		"uId"
// @Param	province		formData		string  		false		"省"
// @Param	city			formData		string  		false		"市"
// @Param	area			formData		string  		false		"区"
// @Param	longitude		formData		float64  		false		"经度"
// @Param	latitude		formData		float64  		false		"纬度"
// @Success 200 {string} success
// @router /updateUserPosition [patch]
func (this *UserController) UpdateUserPosition() {
	uId := this.MustInt64("uId")
	province := this.GetString("province", "")
	city := this.GetString("city", "")
	area := this.GetString("area", "")
	longitude , _ := this.GetFloat("longitude", 0)
	latitude , _ := this.GetFloat("latitude", 0)

	var user models.User
	base.DBEngine.Table("user").Where("u_id=?", uId).Get(&user)

	if province != "" {
		user.Province = province
	}
	if city != "" {
		user.City = city
	}
	if area != "" {
		user.Area = area
	}
	if longitude != 0 {
		user.Longitude = longitude
	}
	if latitude != 0 {
		user.Latitude = latitude
	}
	base.DBEngine.Table("user").Where("u_id=?", uId).Cols("province", "city", "area", "longitude", "latitude").Update(&user)

	this.ReturnData = "success"
}

// @Title 根据当前位置要求获取用户列表
// @Description 根据当前位置要求获取用户列表
// @Param	uId				query			int64	  		true		"uId"
// @Param	province		query			string  		false		"省"
// @Param	city			query			string  		false		"市"
// @Param	area			query			string  		false		"区"
// @Param	longitudeMax	query			float64  		false		"经度最大"
// @Param	longitudeMin	query			float64  		false		"经度最小"
// @Param	latitudeMax		query			float64  		false		"纬度最大"
// @Param	latitudeMin		query			float64  		false		"纬度最小"
// @Success 200 {object} models.UserList
// @router /getUserListByPosition [get]
func (this *UserController) GetUserListByPosition() {
	uId := this.MustInt64("uId")
	province := this.GetString("province", "")
	city := this.GetString("city", "")
	area := this.GetString("area", "")
	longitudeMax, _ := this.GetFloat("longitudeMax", 0)
	longitudeMin, _ := this.GetFloat("longitudeMin", 0)
	latitudeMax, _ := this.GetFloat("latitudeMax", 0)
	latitudeMin, _ := this.GetFloat("latitudeMin", 0)

	var user models.User
	base.DBEngine.Table("user").Where("u_id=?", uId).Get(&user)

	whereSql := " and u_id !='"+strconv.FormatInt(uId, 10)+"' "
	if province != "" {
		whereSql += " and province='" + province + "' "
	}
	if city != "" {
		whereSql += " and city='" + city + "' "
	}
	if area != "" {
		whereSql += " and area='" + area + "' "
	}
	if longitudeMax != 0 {
		whereSql += " and longitude<=" + strconv.FormatFloat(longitudeMax, 'f', 6, 64) + " "
	}
	if longitudeMin != 0 {
		whereSql += " and longitude>=" + strconv.FormatFloat(longitudeMin, 'f', 6, 64) + " "
	}
	if latitudeMax != 0 {
		whereSql += " and latitude<=" + strconv.FormatFloat(latitudeMax, 'f', 6, 64) + " "
	}
	if latitudeMin != 0 {
		whereSql += " and latitude>=" + strconv.FormatFloat(latitudeMin, 'f', 6, 64) + " "
	}

	var userList []models.UserShort
	base.DBEngine.Table("user").Where("1=1 "+whereSql).Find(&userList)

	if userList == nil {
		userList = make([]models.UserShort, 0)
	}

	//如果周围的真实用户少于10个，则创建僵尸用户直到10个
	if len(userList) < 10 {
		zombieList := createZombieUser(user, 10 - len(userList))
		for _, zombie := range zombieList {
			userList = append(userList, zombie)
		}
	}

	this.ReturnData = models.UserList{userList}
}

// @Title 登出
// @Description 登出
// @Param	uId				formData		int64	  		true		"uId"
// @Success 200 {string} success
// @router /signOut [post]
func (this *UserController) SignOut() {
	uId := this.MustInt64("uId")

	//推送用表
	var userSignInDeviceInfo models.UserSignInDeviceInfo
	base.DBEngine.Table("user_sign_in_device_info").Where("u_id=?", uId).Get(&userSignInDeviceInfo)
	userSignInDeviceInfo.DeviceToken = ""
	userSignInDeviceInfo.DeviceModel = ""
	userSignInDeviceInfo.SystemVersion = ""
	userSignInDeviceInfo.AppVersion = ""
	base.DBEngine.Table("user_sign_in_device_info").Where("u_id=?", uId).AllCols().Update(&userSignInDeviceInfo)


	this.ReturnData = "success"
}


































//通过手机号获取用户信息
func UserWithPhoneNumber(phoneNumber string) (user models.User, err error) {
	hasUser, err := base.DBEngine.Table("user").Where("phone_number=?", phoneNumber).Get(&user)
	if err != nil {
		return user, err
	}
	if !hasUser {
		return user, nil
	}
	return user, nil
}

//通过昵称获取用户信息
func UserWithNickName(nickName string) (user models.User, err error) {
	hasUser, err := base.DBEngine.Table("user").Where("nick_name=?", nickName).Get(&user)
	if err != nil {
		return user, err
	}
	if !hasUser {
		return user, nil
	}
	return user, nil
}

//检查昵称是否唯一
func checkSameNickName(nickName string) (isSensitive bool, hasSameNickNameUser bool) {
	isSensitive, _ = util.FilterContent(nickName)
	hasSameNickNameUser, _ = base.DBEngine.Table("user").Where("nick_name=?", nickName).Get(new(models.User))
	return isSensitive, hasSameNickNameUser
}

//创建僵尸用户，如果用户退出则还原僵尸账户位置
func createZombieUser(user models.User, number int) []models.UserShort {
	var zombieList []models.User
	var userList []models.UserShort
	base.DBEngine.Table("user").Where("is_zombie=1").And("longitude=0 and latitude=0").Find(&zombieList)
	if zombieList == nil {
		zombieList = make([]models.User, 0)
	}
	storedUnuseZombieNumber := len(zombieList)
	//累加僵尸用户
	if number > len(zombieList) {
		for i:=0; i< (number-storedUnuseZombieNumber);i++ {
			var zombie models.User
			zombie.NickName = GetDefaultNickName()
			zombie.Gender, zombie.Avatar = GetRandomGenderAndAvatar()
			hashedPassword, salt, _ := util.EncryptPassword("iamzombie")
			zombie.Password = hashedPassword
			zombie.Salt = salt
			zombie.Birthday = GetRandomBirthday()
			zombie.Status = 1
			zombie.IsZombie = 1
			base.DBEngine.Table("user").InsertOne(&zombie)
			zombieList = append(zombieList, zombie)
		}
	}
	//提供僵尸用户定位
	for _, zombie := range zombieList {
		zombie.Province = user.Province
		zombie.City = user.City
		zombie.Area = user.Area
		zombie.Longitude, zombie.Latitude = CalcZombiePositionByRangeMeter(user.Longitude, user.Latitude, 300)
		base.DBEngine.Table("user").Where("u_id=?", zombie.UId).Cols("province", "city", "area", "longitude", "latitude").Update(&zombie)

		user, _ := zombie.UsetToUserShort()
		userList = append(userList, *user)
	}
	return userList
}

//根据用户定位获得僵尸定位
//纬度每差1度，实际距离为111千米
//在纬线上，经度每差1度，实际距离为111×cos(角)千米
//300米范围随机加减
//考虑不能移动的地方，如河海
func CalcZombiePositionByRangeMeter(longitude float64, latitude float64, rangeMeter int) (float64, float64) {
	zombieLongitudeChange := float64(util.GenerateRangeNum(0, rangeMeter))/1000000.0 * GetRandomChange()
	zombieLatitudeChange := float64(util.GenerateRangeNum(0, rangeMeter))/1000000.0 * GetRandomChange()
	return longitude + zombieLongitudeChange, latitude + zombieLatitudeChange
}

//获得默认昵称
func GetDefaultNickName() string {
	var defaultNickName models.DefaultNickName
	hasDefaultNickName, _ := base.DBEngine.Table("default_nick_name").Where("status=0").Asc("id").Limit(1, 0).Get(&defaultNickName)
	if hasDefaultNickName {
		defaultNickName.Status = 1
		base.DBEngine.Table("default_nick_name").Where("id=?", defaultNickName.Id).Cols("status").Update(&defaultNickName)
		return defaultNickName.NickName
	} else {
		randomUUId, _ := uuid.NewV4()
		nickName := "匿名"+randomUUId.String()
		return nickName
	}
}

//获得随机性别
func GetRandomBirthday() (birthday string) {
	year := util.GenerateRangeNum(models.DefaultBirthdayMin, models.DefaultBirthdayMax)
	month := util.GenerateRangeNum(1, 12)
	if month < 10 {
		birthday = strconv.Itoa(year)+"0"+strconv.Itoa(month)
	} else {
		birthday = strconv.Itoa(year)+strconv.Itoa(month)
	}
	return birthday
}

//获得随机性别和头像
func GetRandomGenderAndAvatar() (gender int, avatar string) {
	sIndex := rand.Intn(len(models.DefaultGenderAndAvatar))
	return models.DefaultGenderAndAvatar[sIndex].Gender, models.DefaultGenderAndAvatar[sIndex].Avatar
}

//获得随机经纬度加减
func GetRandomChange() float64 {
	sIndex := rand.Intn(len(models.DefaultDirection))
	return models.DefaultDirection[sIndex]
}