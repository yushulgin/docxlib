package yaxml

import (
	"strings"
)

// generate tag from text
func GenerateTag(text string) (res []string) {
	var startIndex = 0
	var endIndex = 0
	text = strings.ReplaceAll(text, "\r\n", "")
	for index := 0; index < len(text); index++ {
		if text[index] == '<' {
			startIndex = index
			if index-endIndex > 1 {
				res = append(res, text[endIndex+1:startIndex])
			}
		} else if text[index] == '>' {
			endIndex = index
			res = append(res, text[startIndex:endIndex+1])
		}
	}
	return
}

// is start of document?
func IsDocumentStart(text string) (isStart bool) {
	if strings.HasPrefix(text, "<?") && strings.HasSuffix(text, "?>") {
		isStart = true
	}
	return
}

// is tag?
func IsTag(text string) (res bool) {
	if strings.HasPrefix(text, "<") && strings.HasSuffix(text, ">") {
		res = true
	}
	return
}

// close state
// 0 is not close
// 1 is self close
// 2 is normal close
func TypeTagEnd(text string) (res int) {
	if strings.HasSuffix(text, "/>") {
		res = 1
	} else if strings.HasPrefix(text, "</") {
		res = 2
	}
	return
}
