/*
@Time : 2018/12/18 上午10:58 
@Author : zwcui
@Software: GoLand
*/
package util

import (
	"strings"
	"anonymousFriends/models"
)

//根据错误码获取
func getErrorCodeAndDescription(error string) (errorCode string, errorDescription string) {
	errorCode = strings.Split(error, models.ErrorSpliter)[0]
	errorDescription = strings.Split(error, models.ErrorSpliter)[1]
	return errorCode, errorDescription
}

//生成错误返回结构体
func GenerateAlertMessage(error string) models.AlertMessage{
	var alert models.AlertMessage
	alert.AlertCode, alert.AlertMessage = getErrorCodeAndDescription(models.UserError101)
	return alert
}