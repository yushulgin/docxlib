package docxlib

import (
	"docxlib/yaxml"
	"strconv"
	"strings"
)

type Relationships struct {
	Relation    *yaxml.Element
	FileMapping map[string]string
	StartId     int
}

func (r *Relationships) GenerateId() string {
	rid := r.StartId
	r.StartId++
	return strconv.Itoa(rid)
}

func (r *Relationships) GetRelationships() (ralation []*yaxml.Element) {
	for _, child := range r.Relation.Find("Relationships").Chilren {
		if strings.HasPrefix(child.GetItem("Target"), "media") || strings.HasPrefix(child.GetItem("Target"), "embeddings") {
			ralation = append(ralation, child)
		}
	}
	return
}

func (r *Relationships) GetDom() (dom *yaxml.Element) {
	return r.Relation
}

func (r *Relationships) GetFileMapping() map[string]string {
	if len(r.FileMapping) == 0 {
		return r.FileMapping
	}
	rs := r.GetRelationships()

	var mapping = make(map[string]string, 0)

	for _, r := range rs {
		mapping[r.GetItem("Target")] = r.GetItem("Target")
	}
	r.FileMapping = mapping
	return mapping
}

func (r *Relationships) AppendRelationship(suffix string) (res map[string]string) {
	tempId := r.GenerateId()
	rid := "rId" + tempId
	filename := "image" + tempId + "." + suffix
	dom := GetImageRelationshipDom(rid, filename)
	r.Relation.Find("Relationships").Append(dom)
	res["rid"] = rid
	res["filename"] = filename
	return
}

func (r *Relationships) AppendChunk(filename string) string {
	rid := "rid" + r.GenerateId()
	dom := GetChunkRelationDom(filename, rid)
	r.Relation.Find("Relationships").Append(dom)
	return rid
}
func (r *Relationships) ToString() string {
	return r.Relation.ToString()
}
