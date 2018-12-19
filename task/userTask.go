package task

import (
	"strconv"
	"anonymousFriends/util"
)

func TestTimedTask(){
	util.Logger.Info(strconv.FormatInt(util.UnixOfBeijingTime(), 10)+"-->"+util.FormatTimestamp(util.UnixOfBeijingTime()))
}