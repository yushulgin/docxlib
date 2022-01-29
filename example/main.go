package main

import (
	"fmt"

	"github.com/haiyux/docxlib"
)

func main() {
	// appendthing()
	// split()
	// replace() // 替换测试
	// merge()
}

func replace() {
	doc, err := docxlib.NewDocx("./test.docx", false)
	defer doc.Close()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	doc.Replace("兮", "不兮可惜")
	doc.Save("./replace_res.docx")
}

func split() {
	doc, err := docxlib.NewDocx("./test.docx", false)
	defer doc.Close()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	document, err := doc.GetDocument()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	contents := document.GetContent()
	fmt.Println(len(contents))
	for index, value := range contents {
		document.ClearContent()
		document.AppendContent(value.GetDom())
		err := doc.Save(fmt.Sprintf("./a/%d.docx", index))
		if err != nil {
			fmt.Println("err:", err)
			return
		}
	}

}

func appendthing() {
	doc, err := docxlib.MakeNewDocument()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	doc.AppendParagraph("https://zhaohaiyu.com", docxlib.LEFT)
	doc.AppendPicture("./1.png", "left")
	doc.AppendParagraph("https://zhaohaiyu.com/", "right")
	doc.AppendParagraph("https://zhaohaiyu.com/index.html", "left")
	doc.AppendPicture("./2.jpg", "left")
	doc.Save("appendthing.docx")
}

func merge() {
	doc, err := docxlib.NewDocx("a/1.docx", false)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// err = doc.Merge([]string{"a/5.docx", "a/2.docx", "7/34.docx"}, false, false)
	err = doc.Merge([]string{"a/5.docx", "a/2.docx", "7/34.docx"}, true, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = doc.Save("./merge.docx")
	if err != nil {
		fmt.Println(err)
		return
	}
}
