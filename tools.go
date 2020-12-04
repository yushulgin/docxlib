package docxlib

import (
	"strconv"
)

var (
	UNIT = [5]string{"", "十", "百", "千", "万"}
	NUM  = [10]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
)

// 数字转汉字
func NumToHans(num int) string {
	deleteFitst := false
	if num >= 10 && num <= 20 {
		deleteFitst = true
	}
	numStr := strconv.Itoa(num)
	numStrLen := len(numStr)
	temp := ""
	zero := false
	for x := 0; x < numStrLen; x++ {
		index := x
		curNum, _ := strconv.Atoi(string(numStr[index]))
		curUint := numStrLen - index - 1
		if curNum == 0 && zero {
			continue
		}
		if curNum == 0 {
			zero = true
			temp += NUM[curNum]
			continue
		}
		temp += (NUM[curNum] + UNIT[curUint])
		zero = false
	}
	// fmt.Println()
	if zero {
		temp = temp[:len(temp)-3]
	}
	if deleteFitst {
		temp = temp[3:]
	}
	return temp
}
