package yaxml

import (
	"fmt"
	"strings"
)

const (
	ROOT           = 0
	DOCUMENT_START = 1
	TAG            = 2
	TEXT           = 3
	ORIGINAL       = 4
)

type Element struct {
	Prefix  string
	Name    string
	Type    int
	Attrs   map[string]string
	Parent  *Element
	Chilren []*Element
}

func GenerateElement(tag string, parent *Element, mytype int) (ele *Element) {
	ele = &Element{
		Prefix:  "",
		Name:    "",
		Type:    mytype,
		Attrs:   make(map[string]string, 0),
		Parent:  parent,
		Chilren: nil,
	}
	if ele.Type == TAG {
		ele.parserTag(tag)
	} else if ele.Type == TEXT || ele.Type == DOCUMENT_START || ele.Type == ORIGINAL {
		ele.Name = tag
	}
	return
}

func (e *Element) GetItem(name string) string {
	s, ok := e.Attrs[name]
	if !ok {
		return ""
	}
	return s
}

func (e *Element) SetItem(name, value string) {
	e.Attrs[name] = value
}

func (e *Element) Text() (text []string) {
	for _, child := range e.Chilren {
		if child.Type == TEXT {
			text = append(text, child.Name)
		}
	}
	return
}

func (e *Element) parserTag(tag string) {
	end := len(tag) - 1
	if TypeTagEnd(tag) == 1 {
		end--
	}
	tag = strings.Trim(strings.Trim(tag[1:end], " "), "\n")
	e.parserDetail(tag, true)
}

func (e *Element) parserName(text string) {
	text = strings.Trim(strings.TrimSpace(text), "\n")
	names := strings.Split(text, ":")
	if len(names) > 1 {
		e.Prefix = names[0]
		e.Name = names[1]
	} else {
		e.Name = names[0]
	}
}

func (e *Element) parserAttr(text string) {
	texts := strings.Split(text, "=")
	if len(texts) < 2 {
		return
	}
	e.Attrs[texts[0]] = strings.Trim(strings.Trim(texts[1], "\""), "'")
}

func (e *Element) parserDetail(text string, name bool) {
	text = strings.Trim(strings.TrimSpace(text), "\n")
	if len(text) == 0 {
		return
	}
	var quotation = 0
	var indexs = 0
	for index := 0; index < len(text); index++ {
		if text[index] == '"' {
			quotation++
		}
		if text[index] == ' ' && quotation%2 == 0 {
			break
		}
		index++
		indexs = index
	}
	if indexs == len(text)-1 {
		if name {
			e.parserName(text)
		} else {
			e.parserAttr(text)
		}
		return
	} else {
		if name {
			e.parserName(text[:indexs])
		} else {
			e.parserAttr(text[:indexs])
		}
	}
	e.parserDetail(text[indexs+1:], false)
}

func (e *Element) Find(tagName string) (res *Element) {
	strack := make([]*Element, 0)
	strack = append(strack, e.Chilren...)
	for len(strack) > 0 {
		child := strack[0]
		strack = strack[1:]
		if child.Name == tagName {
			res = child
			return
		}
		strack = append(strack, child.Chilren...)
	}
	return
}

func (e *Element) FindAll(tagName string) (res []*Element) {
	strack := make([]*Element, 0)
	strack = append(strack, e.Chilren...)
	for len(strack) > 0 {
		child := strack[0]
		strack = strack[1:]
		if child.Name == tagName {
			res = append(res, child)
		}
		strack = append(strack, child.Chilren...)
	}
	return
}

func (e *Element) extract() (err error) {
	if e.Parent == nil {
		return
	}
	err = e.delSelfInParent()
	if err != nil {
		return
	}
	e.Parent = nil
	return
}

func (e *Element) Clear() {
	for n := len(e.Chilren); n > 0; {
		child := e.Chilren[n-1]
		e.Chilren = e.Chilren[:n]
		child.Parent = nil
	}
}

func (e *Element) Insert(position int, ele *Element) {
	ele.Parent = e
	e.Chilren = eleInsert(position, ele, e.Chilren)
}

func (e *Element) InsertBefore(ele *Element) {
	index := e.GetMyIndexByParentChrildren()
	if index < 0 {
		return
	}
	e.Parent.Insert(index, ele)
}

func (e *Element) Append(ele *Element) {
	ele.Parent = e
	e.Chilren = append(e.Chilren, ele)
}

func (e *Element) ToString() string {
	doc := ""
	if e.Type == ROOT {
		for _, chirld := range e.Chilren {
			doc += chirld.ToString()
		}
		return doc
	} else if e.Type == DOCUMENT_START {
		doc = e.Name + "\n"
		return doc
	} else if e.Type == TEXT || e.Type == ORIGINAL {
		return e.Name
	}
	var tagName string
	if e.Prefix != "" {
		tagName = e.Name
	} else {
		tagName = e.Prefix + ":" + e.Name
	}
	attrs := make([]string, 0)
	for _, key := range e.Attrs {
		attrs = append(attrs, key+"=\""+e.Attrs[key]+"\"")
	}
	attrstr := strings.Join(attrs, " ")
	var space string
	if len(attrstr) > 0 {
		space = " "
	} else {
		space = ""
	}
	tag := "<" + tagName + space + attrstr
	if len(e.Chilren) == 0 {
		tag += "/>"
		return tag
	}
	tag += ">"
	for _, child := range e.Chilren {
		tag += child.ToString()
	}
	tag += fmt.Sprintf("</%s>", tagName)
	return tag
}

func (e *Element) delSelfInParent() (err error) {
	meIndex := -1
	for i := 0; i < len(e.Parent.Chilren); i++ {
		if e.Parent.Chilren[i] == e {
			meIndex = i
			break
		}
	}
	if meIndex == -1 {
		err = fmt.Errorf("not found me in parent.children")
		return
	}
	e.Parent.Chilren = append(e.Parent.Chilren[:meIndex], e.Parent.Chilren[meIndex+1:]...)
	return
}

func eleInsert(position int, ele *Element, lst []*Element) (elist []*Element) {
	n := len(lst)
	if position <= 0 {
		elist = append(elist, ele)
		elist = append(elist, lst...)
		return
	} else if position >= n {
		elist = append(lst, ele)
		return
	}
	elist = append(lst[:position], ele)
	elist = append(elist, lst[position:]...)
	return
}

func (e *Element) GetMyIndexByParentChrildren() int {
	index := -1
	for i := 0; i < len(e.Parent.Chilren); i++ {
		if e.Parent.Chilren[i] == e {
			index = i
			return index
		}
	}
	return index
}
