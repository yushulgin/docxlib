package docxlib

import (
	"github.com/zhaohaiyu1996/docxlib/yaxml"
	"fmt"
	"strings"
)

func GetBreakPage() (tag *yaxml.Element) {
	// parent := &yaxml.Element{Name: "xml"}
	tag = yaxml.GenerateElement(`<w:p><w:r><w:br w:type="page"/></w:r></w:p>`, nil, yaxml.ORIGINAL)
	return
}

var text_template = `<w:p><w:pPr><w:jc w:val="%s"/></w:pPr><w:r><w:t>%s</w:t></w:r></w:p>`

func GetTextDom(text, align string) (dom *yaxml.Element) {
	text = strings.ReplaceAll(strings.ReplaceAll(text, "<", "&lt;"), ">", "&gt;")
	text = fmt.Sprintf(text_template, align, text)
	dom = yaxml.GenerateElement(text, nil, yaxml.ORIGINAL)
	return
}

var relationship_template = `<Relationship Id="%s" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/image" Target="media/%s"/>`

func GetImageRelationshipDom(rid, filename string) (dom *yaxml.Element) {
	tempalete := fmt.Sprintf(relationship_template, rid, filename)
	dom = yaxml.GenerateElement(tempalete, nil, yaxml.ORIGINAL)
	return
}

var image_template = `<w:p><w:pPr><w:jc w:val="%s"/><w:rPr><w:rFonts w:ascii="Xingkai TC Light" w:eastAsia="Xingkai TC Light" w:hAnsi="Xingkai TC Light" w:hint="eastAsia"/></w:rPr></w:pPr><w:r><w:rPr><w:rFonts w:ascii="Xingkai TC Light" w:eastAsia="Xingkai TC Light" w:hAnsi="Xingkai TC Light"/><w:noProof/></w:rPr><w:drawing><wp:inline><wp:extent cx="%d" cy="%d"/><wp:effectExtent l="0" t="0" r="0" b="0"/><wp:docPr id="1" name="picture 1"/><wp:cNvGraphicFramePr><a:graphicFrameLocks xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" noChangeAspect="1"/></wp:cNvGraphicFramePr><a:graphic xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main"><a:graphicData uri="http://schemas.openxmlformats.org/drawingml/2006/picture"><pic:pic xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture"><pic:nvPicPr><pic:cNvPr id="1" name=""/><pic:cNvPicPr/></pic:nvPicPr><pic:blipFill><a:blip r:embed="%s"/><a:stretch><a:fillRect/></a:stretch></pic:blipFill><pic:spPr><a:xfrm><a:off x="0" y="0"/><a:ext cx="%d" cy="%d"/></a:xfrm><a:prstGeom prst="rect"><a:avLst/></a:prstGeom></pic:spPr></pic:pic></a:graphicData></a:graphic></wp:inline></w:drawing></w:r></w:p>`

func GetImageDom(rid string, width int, height int, align string) (dom *yaxml.Element) {
	image := fmt.Sprintf(image_template, align, width, height, rid, width, height)
	dom = yaxml.GenerateElement(image, nil, yaxml.ORIGINAL)
	return
}

var content_template_png = `<Default Extension="png" ContentType="image/png"/>`

func GetPngContenttype() (dom *yaxml.Element) {
	dom = yaxml.GenerateElement(content_template_png, nil, yaxml.ORIGINAL)
	return
}

var content_template_jpg = `<Default Extension="jpg" ContentType="image/jpeg"/>`

func GetJpgContenttype() (dom *yaxml.Element) {
	dom = yaxml.GenerateElement(content_template_jpg, nil, yaxml.ORIGINAL)
	return
}

var chunk = `<w:altChunk r:id="%s"/>`

func GetChunkDom(rid string) (dom *yaxml.Element) {
	c := fmt.Sprintf(chunk, rid)
	dom = yaxml.GenerateElement(c, nil, yaxml.ORIGINAL)
	return
}

var chunk_relation = `<Relationship Target="../%s" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/aFChunk" Id="%s"/>`

func GetChunkRelationDom(filename, rid string) (dom *yaxml.Element) {
	relation := fmt.Sprintf(chunk_relation, filename, rid)
	dom = yaxml.GenerateElement(relation, nil, yaxml.ORIGINAL)
	return
}

var chunk_content_type = `<Override ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml" PartName="/%s"/>`

func GetChunkContenDom(filename string) (dom *yaxml.Element) {
	relation := fmt.Sprintf(chunk_content_type, filename)
	dom = yaxml.GenerateElement(relation, nil, yaxml.ORIGINAL)
	return
}
