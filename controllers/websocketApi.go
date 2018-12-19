package controllers

import (
	"net/http"
	"golang.org/x/net/websocket"
	"strconv"
	"encoding/json"
	"sort"
	"crypto/md5"
	"encoding/hex"
	"strings"
	"anonymousFriends/models"
	"anonymousFriends/util"
)

//socket连接池
var UserSocketConnections map[int64]models.SocketConnection

type WSServer struct {
	ListenAddr string
}

func (this *WSServer) Handler (conn *websocket.Conn) {
	util.Logger.Info("-------------socket--Handler--start------------"+strconv.FormatInt(util.UnixOfBeijingTime(), 10))
	if UserSocketConnections == nil {
		UserSocketConnections = make(map[int64]models.SocketConnection)
	}

	util.Logger.Info("a new ws conn: conn.RemoteAddr()="+conn.RemoteAddr().String()+"  conn.LocalAddr()="+conn.LocalAddr().String())
	var err error
	for {
		var reply string
		var socketMessage models.SocketMessage
		err = websocket.Message.Receive(conn, &reply)
		//连接出错则移除
		if err != nil {
			for k, v := range UserSocketConnections {
				if v.Conn == conn {
					delete(UserSocketConnections, k)
				}
			}
			util.Logger.Info("receive conn err:",err.Error())
			break
		}

		//util.Logger.Info("-----reply----  "+reply)
		err = json.Unmarshal([]byte(reply), &socketMessage)
		if err != nil {
			util.Logger.Info("----socketMessage--json.Unmarshal--err---- "+err.Error())
			continue
		}

		//验签
		if socketMessage.MessageSign != SignMessage(socketMessage) {
			util.Logger.Info("----socketMessage--签名验证失败")
			break
		}

		//建立心跳，区分前台后台，并当后台转至前台时补发离线消息
		if socketMessage.MessageType == -1 || socketMessage.MessageType == 0 {
			handleHeartbeat(&socketMessage, conn)
			if socketMessage.MessageType == 0 {
				handleUnsentSocketMessage(&socketMessage, conn)
			}
		}

		//普通聊天/直连聊天
		if socketMessage.MessageType == 1 {
			handleUserMessage(&socketMessage, conn)
		}


	}
	util.Logger.Info("-------------socket--Handler--end------------"+strconv.FormatInt(util.UnixOfBeijingTime(), 10))
}

func (this *WSServer) Start() (error) {
	http.Handle("/ws", websocket.Handler(this.Handler))
	util.Logger.Info("websocket----begin to listen")
	err := http.ListenAndServe(this.ListenAddr, nil)
	if err != nil {
		util.Logger.Info("ListenAndServe:", err)
		return err
	}
	util.Logger.Info("websocket----start end")
	return nil
}

//处理心跳
func handleHeartbeat(socketMessage *models.SocketMessage, conn *websocket.Conn){
	if _, ok := UserSocketConnections[socketMessage.MessageSenderUid]; !ok {
		var socketConnection models.SocketConnection
		socketConnection.Conn = conn
		if socketMessage.MessageType == 0 {
			socketConnection.ConnType = 1
		} else if socketMessage.MessageType == -1 {
			socketConnection.ConnType = 2
		}
		socketConnection.ExpireTime = socketMessage.MessageExpireTime
		socketConnection.Token = socketMessage.MessageToken
		UserSocketConnections[socketMessage.MessageSenderUid] = socketConnection
		util.Logger.Info("-----socketConnection.heartbeat.ExpireTime----start-"+strconv.FormatInt(socketMessage.MessageSenderUid, 10)+"--"+strconv.FormatInt(UserSocketConnections[socketMessage.MessageSenderUid].ExpireTime, 10))
	} else {
		//以token和appNo作为唯一标示
		if UserSocketConnections[socketMessage.MessageSenderUid].Token != socketMessage.MessageToken {
			//账户被挤下线
			util.Logger.Info("-------您的账户已在其他地方登陆-------"+strconv.FormatInt(socketMessage.MessageSenderUid, 10))
			alert := util.GenerateAlertMessage(models.WebsocketError100)
			alertJsonByte, err := json.Marshal(alert)
			if err != nil {
				util.Logger.Info("---json to string---您的账户已在其他地方登陆----err:"+err.Error())
				//return
			}

			var replySocketMessage models.SocketMessage
			replySocketMessage.MessageType = 2
			replySocketMessage.MessageSendTime = util.UnixOfBeijingTime()
			replySocketMessage.MessageSenderUid = socketMessage.MessageSenderUid
			replySocketMessage.MessageExpireTime = util.UnixOfBeijingTime()+3
			replySocketMessage.MessageContent = string(alertJsonByte)
			replySocketMessage.MessageToken = socketMessage.MessageToken
			replySocketMessage.MessageSign = SignMessage(replySocketMessage)

			replySocketMessageJsonByte, err := json.Marshal(replySocketMessage)
			if err != nil {
				util.Logger.Info("---json to string---replySocketMessage----err:"+err.Error())
				//return
			}

			if err := websocket.Message.Send(UserSocketConnections[socketMessage.MessageSenderUid].Conn, string(replySocketMessageJsonByte)); err != nil {
				util.Logger.Info("----userMessage--websocket.Message.Send 您的账户已在其他地方登陆 err:", err.Error())
				//移除出错的链接
				delete(UserSocketConnections, socketMessage.MessageSenderUid)
			}
			//新登陆的账户
			var socketConnection models.SocketConnection
			socketConnection.Conn = conn
			if socketMessage.MessageType == 0 {
				socketConnection.ConnType = 1
			} else if socketMessage.MessageType == -1 {
				socketConnection.ConnType = 2
			}
			socketConnection.ExpireTime = socketMessage.MessageExpireTime
			socketConnection.Token = socketMessage.MessageToken
			UserSocketConnections[socketMessage.MessageSenderUid] = socketConnection
		} else {
			//查看是否后台切换至前台，如果是则补发离线消息
			storedSocketConnection := UserSocketConnections[socketMessage.MessageSenderUid]
			if storedSocketConnection.ConnType == 2 && socketMessage.MessageType == 0 {
				//查看redis缓存消息，仅后台跳至前台补发直连消息socket，如果被取消，则只发取消的socket
				handleUnsentSocketMessage(socketMessage, conn)
			}

			util.Logger.Info("-----socketConnection.heartbeat.ExpireTime-----"+strconv.FormatInt(socketMessage.MessageSenderUid, 10)+"--"+strconv.FormatInt(UserSocketConnections[socketMessage.MessageSenderUid].ExpireTime, 10))
			socketConnection := UserSocketConnections[socketMessage.MessageSenderUid]
			if socketMessage.MessageType == 0 {
				socketConnection.ConnType = 1
			} else if socketMessage.MessageType == -1 {
				socketConnection.ConnType = 2
			}
			socketConnection.ExpireTime = socketMessage.MessageExpireTime
			UserSocketConnections[socketMessage.MessageSenderUid] = socketConnection
		}
	}

}

//处理聊天
func handleUserMessage(socketMessage *models.SocketMessage, conn *websocket.Conn) {

}



//redis缓存消息转发
func handleUnsentSocketMessage(socketMessage *models.SocketMessage, conn *websocket.Conn){
	//billUnsentSocketMessageKey := models.SOCKET_UNSENT_BILL_MESSAGE+"_"+strconv.FormatInt(socketMessage.MessageSenderUid, 10)
	//billUnsentSocketMessage := baseServer.RedisCache.Get(billUnsentSocketMessageKey)
	//if billUnsentSocketMessage != nil {
	//	if err := websocket.Message.Send(conn, billUnsentSocketMessage); err != nil {
	//		util.Logger.Info("---- redis缓存消息转发 err:", err.Error())
	//	}
	//}
	//baseServer.RedisCache.Delete(billUnsentSocketMessageKey)
}

//签名
func SignMessage(socketMessage models.SocketMessage) string {
	params := make(map[string]string)
	params["messageType"] = strconv.Itoa(socketMessage.MessageType)
	params["messageSendTime"] = strconv.FormatInt(socketMessage.MessageSendTime, 10)
	params["messageSenderUid"] = strconv.FormatInt(socketMessage.MessageSenderUid, 10)
	params["messageReceiverUid"] = strconv.FormatInt(socketMessage.MessageReceiverUid, 10)
	params["messageExpireTime"] = strconv.FormatInt(socketMessage.MessageExpireTime, 10)
	params["messageContent"] = socketMessage.MessageContent
	params["messageToken"] = socketMessage.MessageToken

	keys := make([]string, len(params))

	i := 0
	for k := range params {
		keys[i] = k
		i++
	}

	//util.Logger.Info("keys", keys)
	sort.Strings(keys)
	//util.Logger.Info("sorted keys", keys)

	strTemp := ""
	for _, key := range keys {
		strTemp = strTemp + key + "=" + params[key] + "&"
	}
	strTemp += "key=" + models.SOCKET_MESSAGE_SIGN_KEY
	util.Logger.Info("strTemp = ", strTemp)

	hasher := md5.New()
	hasher.Write([]byte(strTemp))
	md5Str := hex.EncodeToString(hasher.Sum(nil))

	util.Logger.Info("md5 = ", md5Str)
	return strings.ToUpper(md5Str)
}

//每分钟检查失效的socket连接
func checkSocketHeartbeat(){
	util.Logger.Info("定时任务，每分钟检查失效的socket连接")
	for uId, socketConnection := range UserSocketConnections {
		util.Logger.Info("-----定时任务  遍历users-----  util.UnixOfBeijingTime()="+strconv.FormatInt(util.UnixOfBeijingTime(), 10)+"   uid="+strconv.FormatInt(uId, 10)+"   ExpireTime="+strconv.FormatInt(socketConnection.ExpireTime, 10))
		if (socketConnection.ExpireTime + 15) <= util.UnixOfBeijingTime() {
			delete(UserSocketConnections, uId)
		}
	}
}


