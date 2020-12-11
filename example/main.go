package main

import (
	"docxlib"
	"fmt"
)

func main() {
	// // docxlib.Printl()
	// doc, err := docxlib.MakeNewDocument()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// // doc.AppendParagraph("21312312312", "right")
	// err = doc.AppendPicture("/tmp/1.jpg", "left")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// doc.Save("example/t1q1222.docx")
	// // fmt.Println(doc)
	// d, err := doc.GetDocument()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// c := d.GetContent()
	// for i, v := range c {
	// 	d.ClearContent()
	// 	d.AppendContent(v.GetDom())
	// 	doc.Save(fmt.Sprintf("./example/a/%d.docx", i))
	// }
	// doc.Close()
	// doc,err := docxlib.MakeNewDocument()
	// if err != nil {
	// 	return
	// }
	// fmt.Println(doc)
	// doc, err := docxlib.MakeNewDocument()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// _ = doc
	// fmt.Println(doc.Document.ToString())
	// fmt.Printf("%#v\n", doc)
	// doc.Save("./example/1111.docx")

	// a := "赵海宇"
	// fmt.Println(len([]rune(a)))
	// fmt.Println(a[1:])

	// 	s := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	// <w:document xmlns:wpc="http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas" xmlns:cx="http://schemas.microsoft.com/office/drawing/2014/chartex" xmlns:cx1="http://schemas.microsoft.com/office/drawing/2015/9/8/chartex" xmlns:cx2="http://schemas.microsoft.com/office/drawing/2015/10/21/chartex" xmlns:cx3="http://schemas.microsoft.com/office/drawing/2016/5/9/chartex" xmlns:cx4="http://schemas.microsoft.com/office/drawing/2016/5/10/chartex" xmlns:cx5="http://schemas.microsoft.com/office/drawing/2016/5/11/chartex" xmlns:cx6="http://schemas.microsoft.com/office/drawing/2016/5/12/chartex" xmlns:cx7="http://schemas.microsoft.com/office/drawing/2016/5/13/chartex" xmlns:cx8="http://schemas.microsoft.com/office/drawing/2016/5/14/chartex" xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" xmlns:aink="http://schemas.microsoft.com/office/drawing/2016/ink" xmlns:am3d="http://schemas.microsoft.com/office/drawing/2017/model3d" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:m="http://schemas.openxmlformats.org/officeDocument/2006/math" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:wp14="http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing" xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing" xmlns:w10="urn:schemas-microsoft-com:office:word" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml" xmlns:w15="http://schemas.microsoft.com/office/word/2012/wordml" xmlns:w16cid="http://schemas.microsoft.com/office/word/2016/wordml/cid" xmlns:w16se="http://schemas.microsoft.com/office/word/2015/wordml/symex" xmlns:wpg="http://schemas.microsoft.com/office/word/2010/wordprocessingGroup" xmlns:wpi="http://schemas.microsoft.com/office/word/2010/wordprocessingInk" xmlns:wne="http://schemas.microsoft.com/office/word/2006/wordml" xmlns:wps="http://schemas.microsoft.com/office/word/2010/wordprocessingShape" mc:Ignorable="w14 w15 w16se w16cid wp14"><w:body><w:p w:rsidR="001E4F60" w:rsidRPr="001C6ACF" w:rsidRDefault="001E4F60" w:rsidP="00E4139B"><w:pPr><w:rPr><w:rFonts w:ascii="Xingkai TC Light" w:eastAsia="Xingkai TC Light" w:hAnsi="Xingkai TC Light" w:hint="eastAsia"/></w:rPr></w:pPr><w:bookmarkStart w:id="0" w:name="_GoBack"/><w:bookmarkEnd w:id="0"/></w:p><w:sectPr w:rsidR="001E4F60" w:rsidRPr="001C6ACF" w:rsidSect="00E5722F"><w:pgSz w:w="11900" w:h="16840"/><w:pgMar w:top="1440" w:right="1800" w:bottom="1440" w:left="1800" w:header="851" w:footer="992" w:gutter="0"/><w:cols w:space="425"/><w:docGrid w:type="lines" w:linePitch="312"/></w:sectPr></w:body></w:document>`
	// 	// print(len(s))
	// 	s,_ := docxlib.ReadAll("/private/tmp/docxs_temp/8cb91942392911eba20dacde48001122/word/document.xml")
	// 	fmt.Println(len(strings.ReplaceAll(s,"\r\n","\n")))
	// os.Mkdir("/tmp/docx_temp/_rel", 0777)
	// files := []string{"example/a/1.docx", "example/a/2.docx", "example/a/3.docx"}
	// err := docxlib.MergeFiles(files, "example/he2.docx", false)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// doc,err := docxlib.NewDocx("example/a/1.docx",false)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// lst := []string{"example/a/2.docx","example/a/3.docx"}

	// err = doc.Merge(lst,false,false)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// doc.Save("example/merge.docx")
	// doc, err := docxlib.NewDocx("a/1.docx",false)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// a := []string{"a/2.docx","a/4.docx","a/45.docx"}
	// err = doc.Merge(a, true, false)
	// if err != nil {
	// 	return
	// }

	// doc.Save("mmf.docx")
	replace()
}

func replace() {
	doc,err := docxlib.NewDocx("./replace.docx",false)
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	doc.Replace("赵海宇","yzzzzzz")
	doc.Save("./replace_res.docx")
}

