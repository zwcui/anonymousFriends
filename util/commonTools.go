package util

import (
	"math/rand"
)

//隐藏昵称后几位，用*替代
func FormatNickname(nickname string, length int) (formatName string) {
	rs := []rune(nickname)
	rl := len(rs)
	end := length

	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	formatName = string(rs[0:end]) + "*"
	return formatName
}

//手机号中间打*号
func FormatPhoneNo(phoneNo string) (formatString string) {
	if len(phoneNo) < 11{
		return phoneNo
	}
	rs := []rune(phoneNo)

	formatString = string(rs[0:3]) + "****" +  string(rs[7:11])
	return formatString
}

//范围之内生成随机数
func GenerateRangeNum(min, max int) int {
	randNum := rand.Intn(max - min) + min
	return randNum
}

//过滤不良内容
func FilterContent(content string) (find bool, filterContent string) {
	find, _ = ContentFilter.FindIn(content)
	if find {
		filterContent = ContentFilter.Replace(content, 42)
	}
	return find, filterContent
}