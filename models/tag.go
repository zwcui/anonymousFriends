/*
@Time : 2018/12/26 下午2:36 
@Author : zwcui
@Software: GoLand
*/
package models

//标签
type Tag struct {
	TagId       				int64			`description:"tagId" json:"tagId" xorm:"pk autoincr"`
	TagName			       		string			`description:"标签名称" json:"tagName"`
	TagOrder					string			`description:"标签排序" json:"tagOrder"`


	Created           			int64  			`description:"创建时间" json:"created" xorm:"created"`
	Updated           			int64  			`description:"修改时间" json:"updated" xorm:"updated"`
	DeletedAt         			int64  			`description:"删除时间" json:"deleted" xorm:"deleted"`
}
