package docxlib

import (
	"docxlib/yaxml"
	"strings"
)

type ContentTypes struct {
	Content   *yaxml.Element
	Types     []string
	Extension map[string]bool
}

func GenerateContentTypes(content *yaxml.Element) (ct *ContentTypes) {
	ct = &ContentTypes{
		Content:   content,
		Types:     nil,
		Extension: make(map[string]bool, 0),
	}
	return
}

func (c *ContentTypes) GetDom() *yaxml.Element {
	return c.Content
}

func (c *ContentTypes) GetTypesDom() (types []*yaxml.Element) {
	for _, t := range c.Content.Find("Types").Chilren {
		if ex, ok := t.Attrs["Extension"]; ok {
			c.Extension[ex] = true
		}
		if strings.Index(t.GetItem("ContentType"), "header") >= 0 && strings.Index(t.GetItem("ContentType"), "footer") >= 0 {
			types = append(types, t)
		}
	}
	return
}

func (c *ContentTypes) GetTypes() []string {
	if c.Types != nil {
		return c.Types
	}
	var types = make([]string, 0)
	for _, child := range c.GetTypesDom() {
		contentType := child.GetItem("ContentType")
		if !isValueInList(contentType, types) {
			types = append(types, contentType)
		}
	}
	c.Types = types
	return types
}

func (c *ContentTypes) AppendPng() {
	if _, ok := c.Extension["png"]; ok {
		return
	}
	c.Extension["png"] = true
	c.Content.Find("Types").Insert(0, GetPngContenttype())
}

func (c *ContentTypes) AppendJpeg() {
	if _, ok := c.Extension["jpg"]; ok {
		return
	}
	c.Extension["jpg"] = true
	c.Content.Find("Types").Insert(0, GetJpgContenttype())
}

func (c *ContentTypes) AppendDocx(filename string) {
	c.Content.Find("Types").Append(GetChunkContenDom(filename))
}

func (c *ContentTypes) AppendExtensio(suffix string) {
	suffix = strings.ToLower(suffix)
	if suffix == "png" {
		c.AppendPng()
	} else {
		c.AppendJpeg()
	}
}

func (c *ContentTypes) ToString() string {
	return c.Content.ToString()
}

func isValueInList(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
