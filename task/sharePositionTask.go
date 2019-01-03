/*
@Time : 2019/1/3 下午2:51 
@Author : zwcui
@Software: GoLand
*/
package task

import (
	"anonymousFriends/util"
	"anonymousFriends/base"
	"anonymousFriends/models"
)

//定时任务：查看长时间未接的共享地理位置设置关闭
func checkUnclosedSharePositionGroup(){
	util.Logger.Info("定时任务：查看长时间未接的共享地理位置设置关闭")
	var sharePositionGroupList []models.SharePositionGroup
	base.DBEngine.Table("share_position_group").Where("status != 2").Find(&sharePositionGroupList)
	if sharePositionGroupList == nil {
		sharePositionGroupList = make([]models.SharePositionGroup, 0)
	}

	for _, sharePositionGroup := range sharePositionGroupList {
		if sharePositionGroup.Status == 0 {
			//超过10分钟无人应答则被关闭
			if util.UnixOfBeijingTime() - sharePositionGroup.Created >= 10 * 60 {
				sharePositionGroup.Status = 2
				sharePositionGroup.Remark += "定时任务超时未接关闭；"
				base.DBEngine.Table("share_position_group").Where("id=?", sharePositionGroup.Id).Cols("status", "remark").Update(&sharePositionGroup)
			}
		} else if sharePositionGroup.Status == 1 {
			//进行中只有一个人则关闭
			total, _ := base.DBEngine.Table("share_position_member").Where("share_position_group_id=?", sharePositionGroup.Id).And("status=1").Count(new(models.SharePositionMember))
			if total <= 1 {
				sharePositionGroup.Status = 2
				sharePositionGroup.Remark += "定时任务进行中只有一个人则关闭;"
				base.DBEngine.Table("share_position_group").Where("id=?", sharePositionGroup.Id).Cols("status", "remark").Update(&sharePositionGroup)
			}
		}
	}
}
