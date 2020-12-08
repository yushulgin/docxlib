package docxlib

import (
	"docxlib/yaxml"
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

func Parser(text string) *yaxml.Element {
	g := yaxml.GenerateTag(text)
	root := yaxml.GenerateElement("", nil, yaxml.ROOT)
	var stack = make([]*yaxml.Element, 0)
	stack = append(stack, root)
	for _, tag := range g {
		if yaxml.IsDocumentStart(tag) {
			root.Append(yaxml.GenerateElement(tag, nil, yaxml.DOCUMENT_START))
			continue
		}
		if yaxml.IsTag(tag) {
			if yaxml.TypeTagEnd(tag) == 0 {
				e := yaxml.GenerateElement(tag, nil, yaxml.TAG)
				stack[len(stack)-1].Append(e)
				stack = append(stack, e)
			} else if yaxml.TypeTagEnd(tag) == 1 {
				stack[len(stack)-1].Append(yaxml.GenerateElement(tag, nil, yaxml.TAG))
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack[len(stack)-1].Append(yaxml.GenerateElement(tag, nil, yaxml.TEXT))
		}
	}
	return root
}
