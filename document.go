package docxlib

import (
	"strings"

	"github.com/haiyux/docxlib/yaxml"
)

type Paragraph struct {
	Dom *yaxml.Element
}

func GenerateParagraph(dom *yaxml.Element) *Paragraph {
	par := &Paragraph{}
	par.Dom = dom
	return par
}

func (p *Paragraph) GetContent() (text []string) {
	text = make([]string, 0)
	for _, t := range p.Dom.FindAll("t") {
		text = append(text, t.Text()...)
	}
	return
}

func (p *Paragraph) GetDom() *yaxml.Element {
	return p.Dom
}

type Document struct {
	Docu *yaxml.Element
}

func GenerateDocument(document *yaxml.Element) *Document {
	doc := &Document{}
	doc.Docu = document
	doc.GetContent()
	return doc
}

func (d *Document) GetContent() (par []*Paragraph) {
	for _, child := range d.Docu.Find("body").Chilren {
		if !strings.HasPrefix(child.Name, "sectPr") {
			par = append(par, GenerateParagraph(child))
		}
	}
	return
}

func (d *Document) ClearContent() {
	d.Docu.Find("body").Clear()
}

func (d *Document) ClearContentWithoutHeader() {
	sect := d.Docu.Find("sectPr")
	if sect == nil {
		return
	}
	sect.Extract()
	d.Docu.Find("body").Clear()
	d.Docu.Find("body").Append(sect)
}

func (d *Document) AppendContent(dom *yaxml.Element) {
	section := d.Docu.Find("sectPr")
	if section != nil {
		section.InsertBefore(dom)
	} else {
		d.Docu.Find("body").Append(dom)
	}
}

func (d *Document) GetDom(dom *yaxml.Element) *yaxml.Element {
	return d.Docu
}

func (d *Document) AppendParagraph(text string, align string) {
	dom := GetTextDom(text, align)
	d.AppendContent(dom)
}

func (d *Document) AppendPicture(rid string, width int, height int, align string) {
	dom := GetImageDom(rid, width, height, align)
	d.AppendContent(dom)
}

func (d *Document) AppendChunk(rid string) {
	dom := GetChunkDom(rid)
	d.AppendContent(dom)
}

func (d *Document) AppendPageBreak() {
	dom := GetBreakPage()
	d.AppendContent(dom)
}

func (d *Document) ToString() string {
	return d.Docu.ToString()
}
