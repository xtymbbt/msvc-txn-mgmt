package common

import (
	"math/rand"
	"time"
)

//获得定长字符串
//str 填充字符串
//length 获得定长的长度
func RandomFixedStr(baseEle string, length int) string {
	if len(baseEle) == 0 {
		return ""
	}
	if length == 0 {
		return ""
	}
	n := len(baseEle)
	result := ""
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		seed := rand.Intn(n)
		//fmt.Println("seed is: ", seed)
		//fmt.Println("element is: ", string(baseEle[seed]))
		result += string(baseEle[seed])
	}
	return result
}
