/*
@Time : 2018/12/26 上午9:50 
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"strconv"
	"strings"
	"github.com/astaxie/beego"
	//"wenshixiongServer/models"
	"anonymousFriends/util"
	"encoding/json"
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"net/http"
	"bytes"
	"io/ioutil"
	"errors"
	"anonymousFriends/models"
	"anonymousFriends/base"
	"golang.org/x/net/websocket"
)

const (
	UMENG_PUSH_URL = "http://msg.umeng.com/api/send"
	UMENG_API_URL  = "http://api.umeng.com/"
)

const DEFAULT_SOUND = ""


var umengiOSAppkey, umengiOSAppMasterSecret, umengAndroidAppkey, umengAndroidAppMasterSecret, umengAuthToken string

func init(){
	umengAuthToken = ""
	if beego.BConfig.RunMode == base.RUN_MODE_DEV {
		umengiOSAppkey = "5c2319fdf1f55602c0000296"
		umengiOSAppMasterSecret = "u7dslpj3uffgeyxrgrbrptwbnybrhebk"
		umengAndroidAppkey = ""
		umengAndroidAppMasterSecret = ""
	} else if beego.BConfig.RunMode == base.RUN_MODE_TEST {
		umengiOSAppkey = "5c2319fdf1f55602c0000296"
		umengiOSAppMasterSecret = "u7dslpj3uffgeyxrgrbrptwbnybrhebk"
		umengAndroidAppkey = ""
		umengAndroidAppMasterSecret = ""
	} else {
		umengiOSAppkey = "5c2319fdf1f55602c0000296"
		umengiOSAppMasterSecret = "u7dslpj3uffgeyxrgrbrptwbnybrhebk"
		umengAndroidAppkey = ""
		umengAndroidAppMasterSecret = ""
	}
}

func PushCommonMessageToUser(uId int64, message *models.Message, sound string, badge int, cmd string) (ok bool, error error) {
	if len(sound) == 0 {
		sound = DEFAULT_SOUND
	}

	var userSignInDeviceInfo models.UserSignInDeviceInfo
	hasSignIn, _ := base.DBEngine.Table("user_sign_in_device_info").Where("u_id=?", uId).Get(&userSignInDeviceInfo)
	if !hasSignIn {
		return false, errors.New("用户未登录")
	}

	deviceToken := userSignInDeviceInfo.DeviceToken

	if userSignInDeviceInfo.System == models.SYSTEM_ANDROID {
		error = umengPushToUsersFoAndroid(deviceToken, "", message.Content, message.ActionUrl, sound)
	} else {
		error = umengPushToUsersForiOS(deviceToken, message.Content, message.ActionUrl, sound, badge, cmd)
	}


	ok = true
	if error != nil {
		ok = false
	}
	return
}

//APP在前台socket消息，后台及离线则推送消息
func PushSocketMessageToUser(uId int64, message *models.Message, sound string, badge int, cmd string, socketMessageType int) (ok bool, error error) {
	pushSocketSuccess := false

	var userSignInDeviceInfo models.UserSignInDeviceInfo
	hasSignIn, _ := base.DBEngine.Table("user_sign_in_device_info").Where("u_id=?", uId).Get(&userSignInDeviceInfo)
	if !hasSignIn {
		return false, errors.New("用户未登录")
	}

	deviceToken := userSignInDeviceInfo.DeviceToken

	socketConnection, ok := UserSocketConnections[uId]
	if !ok || socketConnection.ExpireTime < util.UnixOfBeijingTime() {
		pushSocketSuccess = false
	} else {
		var pushSocketMessage models.PushSocketMessage
		pushSocketMessage.Message = *message
		pushSocketMessage.Sound = sound
		messageJsonByte, err := json.Marshal(pushSocketMessage)

		var socketMessage models.SocketMessage
		socketMessage.MessageType = socketMessageType
		socketMessage.MessageSendTime = util.UnixOfBeijingTime()
		socketMessage.MessageSenderUid = message.SenderUid
		socketMessage.MessageExpireTime = util.UnixOfBeijingTime()
		socketMessage.MessageContent = string(messageJsonByte)
		socketMessage.MessageToken = deviceToken
		socketMessage.MessageSign = SignMessage(socketMessage)

		socketMessageJsonByte, err := json.Marshal(socketMessage)
		if err != nil {
			pushSocketSuccess = false
			util.Logger.Info("---json to string---pushSocketMessage----err:"+err.Error())
		} else {
			if err := websocket.Message.Send(socketConnection.Conn, string(socketMessageJsonByte)); err != nil {
				util.Logger.Info("----PushSocketMessageToUser  err:", err.Error())
				//移除出错的链接
				delete(UserSocketConnections, uId)
			} else {
				pushSocketSuccess = true
			}
		}
	}

	//cmd不为空表示就socket推一次，即使失败也不推推送
	if cmd != "" {
		return pushSocketSuccess, error
	}

	if !pushSocketSuccess {
		ok, err := PushCommonMessageToUser(uId, message, sound, badge, cmd)
		if ok {
			return true, nil
		} else {
			return false, err
		}
	} else {
		return true, nil
	}
}

func postSign(method, url, masterSecret string, params map[string]interface{}) (string, []byte, error) {
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return "", nil, err
	}

	signSource := strings.ToUpper(method) + url + string(jsonBytes) + masterSecret

	hasher := md5.New()
	hasher.Write([]byte(signSource))

	return hex.EncodeToString(hasher.Sum(nil)), jsonBytes, nil
}

func iOSPostSign(params map[string]interface{}) (string, []byte, error) {
	return postSign("POST", UMENG_PUSH_URL, umengiOSAppMasterSecret, params)
}

func androidPostSign(params map[string]interface{}) (string, []byte, error) {
	return postSign("POST", UMENG_PUSH_URL, umengAndroidAppMasterSecret, params)
}

type umengPushResponseData struct {
	ErrorCode string `json:"error_code"`
}

type umengPushResponse struct {
	Ret  string                `json:"ret"`
	Data umengPushResponseData `json:"data"`
}

type UmengDTO struct {
	Date          string `json:"date"`
	NewUsers      int64 `json:"new_users"`
	ActiveUsers   int64 `json:"active_users"`
	Launches      int64 `json:"launches"`
	Installations int64 `json:"installations"`
}

func UmengPush(isAndroid bool, params map[string]interface{}) (body []byte, err error) {
	signResult := ""
	var jsonBytes []byte
	if isAndroid {
		signResult, jsonBytes, err = androidPostSign(params)
	} else {
		signResult, jsonBytes, err = iOSPostSign(params)
	}

	if err != nil {
		return
	}

	var Url *url.URL
	Url, err = url.Parse(UMENG_PUSH_URL)
	if err != nil {
		return
	}
	Url.RawQuery = "sign=" + signResult

	req, err := http.NewRequest("POST", Url.String(), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	//util.Logger.Debug("UmengPush response Status:", resp.Status)

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		response := umengPushResponse{}
		err1 := json.Unmarshal(body, &response)
		if err1 == nil {
			util.Logger.Warnf("umeng push err code %s", response.Data.ErrorCode)
			util.Logger.Info("UmengPush params:", params)
			util.Logger.Info("UmengPush response Body:", string(body))
		} else {
			util.Logger.Warnf("umeng push Unmarshal err %s", err1.Error())
		}
		err = errors.New("umeng push err code " + response.Data.ErrorCode)
		return
	}
	return
}

func umengPushToUsersForiOS(deviceTokens, msg, actionUrl, sound string, badge int, cmd string) error {
	params := make(map[string]interface{})
	params["appkey"] = umengiOSAppkey
	params["timestamp"] = strconv.FormatInt(util.UnixOfBeijingTime(), 10)
	if strings.Contains(deviceTokens, ",") {
		params["type"] = "listcast"
	} else {
		params["type"] = "unicast"
	}
	params["device_tokens"] = deviceTokens

	payload := make(map[string]interface{})
	aps := make(map[string]interface{})
	aps["alert"] = msg
	aps["sound"] = sound
	if badge == 1 {
		aps["badge"] = 0
	} else {
		aps["badge"] = 1
	}
	payload["aps"] = aps
	payload["actionUrl"] = actionUrl

	payload["time"] = strconv.FormatInt(util.UnixOfBeijingTime(), 10)

	if len(cmd) > 0 {
		payload["cmd"] = cmd
	}
	params["payload"] = payload
	if beego.BConfig.RunMode == base.RUN_MODE_DEV {
		params["production_mode"] = "false"
	} else {
		params["production_mode"] = "true"
	}

	_, err := UmengPush(false, params)
	return err
}

func umengPushToUsersFoAndroid(deviceTokens, title, msg, actionUrl, sound string) error {
	msgMap := make(map[string]string)
	msgMap["msg"] = msg
	msgMap["title"] = title
	msgMap["actionUrl"] = actionUrl
	msgMap["time"] = strconv.FormatInt(util.UnixOfBeijingTime(), 10)

	jsonBytes, err := json.Marshal(msgMap)
	if err != nil {
		return err
	}

	params := make(map[string]interface{})
	params["appkey"] = umengAndroidAppkey
	params["timestamp"] = strconv.FormatInt(util.UnixOfBeijingTime(), 10)
	if strings.Contains(deviceTokens, ",") {
		params["type"] = "listcast"
	} else {
		params["type"] = "unicast"
	}
	params["device_tokens"] = deviceTokens

	payload := make(map[string]interface{})
	body := make(map[string]string)

	body["sound"] = sound

	body["custom"] = string(jsonBytes)
	payload["body"] = body
	payload["display_type"] = "message"

	params["payload"] = payload
	if beego.BConfig.RunMode == base.RUN_MODE_DEV {
		params["production_mode"] = "false"
	} else {
		params["production_mode"] = "true"
	}

	_, err = UmengPush(true, params)
	return err
}

//func UmengPushToAllForiOS(msg, actionUrl string) error {
//	params := make(map[string]interface{})
//	params["appkey"] = umengiOSAppkey
//	params["timestamp"] = strconv.FormatInt(util.UnixOfBeijingTime(), 10)
//	params["type"] = "broadcast"
//
//	payload := make(map[string]interface{})
//	aps := make(map[string]string)
//	aps["alert"] = msg
//	aps["sound"] = "default"
//	//	aps["badge"] = "1"
//	payload["aps"] = aps
//	payload["actionUrl"] = actionUrl
//	payload["time"] = strconv.FormatInt(util.UnixOfBeijingTime(), 10)
//
//	params["payload"] = payload
//	if beego.BConfig.RunMode == RUN_MODE_DEV {
//		params["production_mode"] = "false"
//	} else {
//		params["production_mode"] = "true"
//	}
//
//	//	fmt.Println("params = ", params)
//
//	_, err := UmengPush(false, params, 1)
//	return err
//}
//
//func UmengPushToAllForAndroid(title, msg, actionUrl string) error {
//	msgMap := make(map[string]string)
//	msgMap["msg"] = msg
//	msgMap["title"] = title
//	msgMap["actionUrl"] = actionUrl
//	msgMap["time"] = strconv.FormatInt(util.UnixOfBeijingTime(), 10)
//
//	jsonBytes, err := json.Marshal(msgMap)
//	if err != nil {
//		return err
//	}
//
//	params := make(map[string]interface{})
//	params["appkey"] = umengAndroidAppkey
//	params["timestamp"] = strconv.FormatInt(util.UnixOfBeijingTime(), 10)
//
//	params["type"] = "broadcast"
//	payload := make(map[string]interface{})
//	body := make(map[string]string)
//
//	body["custom"] = string(jsonBytes)
//	payload["body"] = body
//	payload["display_type"] = "message"
//
//	params["payload"] = payload
//	if beego.BConfig.RunMode == RUN_MODE_DEV {
//		params["production_mode"] = "false"
//	} else {
//		params["production_mode"] = "true"
//	}
//
//	_, err = UmengPush(true, params, 1)
//	return err
//}
