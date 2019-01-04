/*
@Time : 2019/1/4 下午4:37 
@Author : zwcui
@Software: GoLand
*/
package models

type Report struct {
	Id       					int64			`description:"id" json:"id" xorm:"pk autoincr"`
	Type       					int 			`description:"类型，1:对app的建议 2: " json:"id" xorm:"pk autoincr"`


}

