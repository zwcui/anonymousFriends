package controllers

import (
	"anonymousFriends/util"
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
	"anonymousFriends/models"
	"golang.org/x/net/websocket"
	"anonymousFriends/base"
)

//公共模块
type PublicController struct {
	apiController
}

//当前api请求之前调用，用于配置哪些接口需要进行head身份验证
func (this *PublicController) Prepare(){
	this.NeedBaseAuthList = []RequestPathAndMethod{}
	this.bathAuth()
	util.Logger.Info("PublicController beforeRequest ")
}

// @Title 测试http请求
// @Description 测试http请求
// @Param	url				formData		string  		true		"url"
// @Param	method			formData		string  		true		"method"
// @Param	headerKey1		formData		string  		false		"headerKey1"
// @Param	headerValue1	formData		string  		false		"headerValue1"
// @Param	headerKey2		formData		string  		false		"headerKey2"
// @Param	headerValue2	formData		string  		false		"headerValue2"
// @Param	headerKey3		formData		string  		false		"headerKey3"
// @Param	headerValue3	formData		string  		false		"headerValue3"
// @Success 200 {string} success
// @Failure 403 create users failed
// @router /testHttpRequest [post]
func (this *PublicController) TestHttpRequest() {
	url := this.MustString("url")
	method := this.MustString("method")
	headerKey1 := this.GetString("headerKey1", "")
	headerValue1 := this.GetString("headerValue1", "")
	headerKey2 := this.GetString("headerKey2", "")
	headerValue2 := this.GetString("headerValue2", "")
	headerKey3 := this.GetString("headerKey3", "")
	headerValue3 := this.GetString("headerValue3", "")


	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader("name=cjb"))
	if err != nil {
		util.Logger.Info("--http.NewRequest---err:"+err.Error())
		this.ReturnData = "--http.NewRequest---err:"+err.Error()
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if headerKey1 != "" {
		req.Header.Set(headerKey1, headerValue1)
	}
	if headerKey2 != "" {
		req.Header.Set(headerKey2, headerValue2)
	}
	if headerKey3 != "" {
		req.Header.Set(headerKey3, headerValue3)
	}


	resp, err := client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		util.Logger.Info("--ioutil.ReadAll---err:"+err.Error())
		this.ReturnData = "--ioutil.ReadAll---err:"+err.Error()
		return
	}

	this.ReturnData = string(body)
}

// @Title 测试发送socket
// @Description 测试发送socket
// @Param	messageType  		formData		int   		true   			"消息类型，-1为后台建立连接，0为前台建立连接，1为普通聊天/直连聊天，2为被挤下线，3直播消息，4直连消息，5刷新标识，6客户端自定义"
// @Param	messageSenderUid  	formData		int64   	false   		"消息发送uid"
// @Param	messageReceiverUid  formData		int64   	true	   		"消息接受uid"
// @Param	messageExpireTime  	formData		int64   	false   		"心跳有效时间"
// @Param	messageContent  	formData		string   	false   		"消息内容"
// @Param	messageToken  		formData		string   	false   		"用户token"
// @Success 200 {string} success
// @router /sendSocketMessage [post]
func (this *PublicController) SendSocketMessage() {
	messageType := this.MustInt("messageType")
	messageSenderUid, _ := this.GetInt64("messageSenderUid", 0)
	messageReceiverUid := this.MustInt64("messageReceiverUid")
	messageExpireTime, _ := this.GetInt64("messageExpireTime", 0)
	messageContent := this.GetString("messageContent", "")
	messageToken := this.GetString("messageToken", "")

	var socketMessage models.SocketMessage
	socketMessage.MessageType = messageType
	socketMessage.MessageSendTime = util.UnixOfBeijingTime()
	socketMessage.MessageSenderUid = messageSenderUid
	socketMessage.MessageReceiverUid = messageReceiverUid
	socketMessage.MessageExpireTime = messageExpireTime
	socketMessage.MessageContent = messageContent
	socketMessage.MessageToken = messageToken
	socketMessage.MessageSign = SignMessage(socketMessage)

	replySocketMessageJsonByte, err := json.Marshal(socketMessage)
	util.Logger.Info("--test send socket--"+string(replySocketMessageJsonByte))
	if err != nil {
		util.Logger.Info("---json to string---replySocketMessage----err:"+err.Error())
	}

	ws, err := websocket.Dial(base.SocketUrl, "", "http://127.0.0.1:8080/")
	util.Logger.Info(base.SocketUrl)
	defer ws.Close()//关闭连接
	if err != nil {
		util.Logger.Info("----websocket.Dial----err:"+err.Error())
	}
	message := []byte(string(replySocketMessageJsonByte))
	_, err = ws.Write(message)
	if err != nil {
		util.Logger.Info("----websocket.ws.Write----err:"+err.Error())
	}

	this.ReturnData = "success"
}