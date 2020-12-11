package docxlib

import (
	"archive/zip"
	"bufio"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"

	"strings"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/google/uuid"
)

var TEMP_BASE_DIR = path.Join("/tmp", "/docx_temp")

const (
	LEFT  = "left"
	RIGHt = "right"
)

type Docx struct {
	*IdAble
	Document        *Document
	ContentTypes    *ContentTypes
	Relationships   *Relationships
	Numbering       string
	Styles          string
	BaseDir         string
	FilePath        string
	ReplaceMap      map[string]string
	RelsDict        map[string]string
	RelsFiles       map[string]bool
	CurrentIdTarget map[string]bool
}

func NewDocx(zipPath string, clearContent bool) (docx *Docx, err error) {
	docx = &Docx{}
	docx.IdAble, err = GenerateIdAble()
	if err != nil {
		return
	}
	exist, err := PathExists(TEMP_BASE_DIR)
	if err != nil {
		return
	}
	if !exist {
		err = os.Mkdir(TEMP_BASE_DIR, 0777)
		if err != nil {
			return
		}
	}
	// TODO
	docx.Document = nil
	docx.ContentTypes = nil
	docx.Relationships = nil
	docx.Numbering = ""
	docx.Styles = ""
	uid, err := uuid.NewUUID()
	if err != nil {
		return
	}
	docx.BaseDir = strings.ReplaceAll(uid.String(), "-", "")
	docx.FilePath = path.Join(TEMP_BASE_DIR, docx.BaseDir)
	err = os.Mkdir(docx.FilePath, 0777)
	if err != nil {
		return
	}
	err = Unzip(zipPath, docx.FilePath)
	if err != nil {
		return
	}

	// TODO
	_, err = docx.GetDocument()
	_, err = docx.getContentTypes()
	_, err = docx.GetRelationships(false)
	docx.ReplaceMap = make(map[string]string, 0)
	if clearContent {
		docx.Document.ClearContentWithoutHeader()
	}
	docx.RelsDict = make(map[string]string, 0)
	docx.RelsFiles = make(map[string]bool, 0)
	docx.CurrentIdTarget = make(map[string]bool, 0)
	return
}

func (d *Docx) GetDocument() (*Document, error) {
	if d.Document != nil {
		return d.Document, nil
	}
	docPath := path.Join(d.FilePath, "word/document.xml")
	docStr, err := ReadAll(docPath)
	if err != nil {
		return nil, err
	}
	document := Parser(docStr)
	d.Document = GenerateDocument(document)

	return d.Document, nil
}

func (d *Docx) GetRelationships(refresh bool) (*Relationships, error) {
	if d.Relationships != nil && !refresh {
		return d.Relationships, nil
	}
	docPath := path.Join(d.FilePath, "word/_rels/document.xml.rels")
	docStr, err := ReadAll(docPath)
	if err != nil {
		return nil, err
	}
	doc := Parser(docStr)
	d.Relationships = GenerateRelationships(doc)
	return d.Relationships, nil
}

func (d *Docx) getAllIdTarget(refresh bool) (dic map[string]string, err error) {
	//
	dic = make(map[string]string, 0)
	r, err := d.GetRelationships(true)
	if err != nil {
		dic = d.RelsDict
		return
	}
	rels := r.GetRelationships()
	for _, rel := range rels {
		_ = rel
		dic[rel.GetItem("Id")] = filepath.Base(rel.GetItem("Target"))
	}
	d.RelsDict = dic

	return
}

func (d *Docx) getRelsAllFiles(refresh bool) (files map[string]bool, err error) {
	if d.RelsFiles != nil && !refresh {
		files = d.RelsFiles
		return
	}
	ids, err := d.getAllIdTarget(refresh)
	if err != nil {
		return
	}
	for _, rid := range ids {
		files[ids[rid]] = true
	}
	d.RelsFiles = files
	return
}

func (d *Docx) getCurrentIdTYarget(refresh bool) (ret map[string]bool, err error) {
	ret = make(map[string]bool, 0)
	if d.CurrentIdTarget != nil && !refresh {
		ret = d.CurrentIdTarget
		return
	}
	allRels, err := d.getAllIdTarget(refresh)
	if err != nil {
		return
	}
	data, err := ReadAll(path.Join(d.FilePath, "word/document.xml"))
	if err != nil {
		return
	}
	idPatten := regexp.MustCompile(`rId\d+`)
	ids := idPatten.FindAllString(data, -1)
	for _, i := range ids {
		if s, ok := allRels[i]; ok {
			ret[s] = true
		}
	}
	d.CurrentIdTarget = ret
	return
}

func (d *Docx) canSave(filename string) (can bool, err error) {
	rels, err := d.getRelsAllFiles(false)
	if err != nil {
		return
	}
	if _, ok := rels[filename]; !ok {
		can = true
		return
	}
	cids, err := d.getCurrentIdTYarget(false)
	if err != nil {
		return
	}
	if _, ok := cids[filename]; ok {
		can = true
		return
	}
	can = false
	return
}

func (d *Docx) getContentTypes() (*ContentTypes, error) {
	if d.ContentTypes != nil {
		return d.ContentTypes, nil
	}
	contentPath := path.Join(d.FilePath, "[Content_Types].xml")
	data, err := ReadAll(contentPath)
	if err != nil {
		return nil, err
	}
	contentTypes := Parser(data)
	d.ContentTypes = GenerateContentTypes(contentTypes)
	return d.ContentTypes, nil
}

func (d *Docx) Merge(filelist []string, page bool, remove bool) (err error) {
	for _, file := range filelist {
		exist, err := PathExists(file)
		if !exist || err != nil {
			err = nil
			continue
		}
		curFile, err := d.cpFile(file)
		if err != nil {
			return err
		}
		d.ContentTypes.AppendDocx(curFile)
		rid := d.Relationships.AppendChunk(curFile)
		if page {
			d.Document.AppendPageBreak()
		}
		d.Document.AppendChunk(rid)
		if remove {
			_ = os.Remove(file)
		}
	}
	return
}

func (d *Docx) Save(name string) (err error) {
	d.saveDocument()
	d.saveContentTypes()

	d.saveRelationships()

	d.getCurrentIdTYarget(true)

	err = d.Zip(name)
	if err != nil {
		return
	}
	return
}

func (d *Docx) saveDocument() (err error) {
	document := d.Document.ToString()
	for key, value := range d.ReplaceMap {
		document = strings.ReplaceAll(document, key, value)
	}
	err = writeStringToFile(path.Join(d.FilePath, "word/document.xml"), document)
	return
}

func (d *Docx) saveContentTypes() (err error) {
	err = writeStringToFile(path.Join(d.FilePath, "[Content_Types].xml"), d.ContentTypes.ToString())
	return
}

func (d *Docx) saveRelationships() (err error) {
	err = writeStringToFile(path.Join(d.FilePath, "word/_rels/document.xml.rels"), d.Relationships.ToString())
	return
}

func (d *Docx) AppendParagraph(text, align string) {
	d.Document.AppendParagraph(text, align)
}

func (d *Docx) AppendPicture(srcfilepath, align string) (err error) {
	if exists, err := PathExists(srcfilepath); !exists || err != nil {
		return fmt.Errorf("filepath not exists!")
	}

	mediaDir := path.Join(d.FilePath, "word/media")
	exists, err := PathExists(mediaDir)
	if err != nil {
		return
	}
	if !exists {
		os.Mkdir(mediaDir, 0777)
	}
	suffix := path.Ext(srcfilepath)
	d.ContentTypes.AppendExtension(suffix)
	idFile := d.Relationships.AppendRelationship(suffix)
	filePath := path.Join(d.FilePath, fmt.Sprintf("word/media/%s", idFile["filename"]))
	_, err = CopyFile(filePath, srcfilepath)
	if err != nil {
		return
	}
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	c, _, err := image.DecodeConfig(file)
	if err != nil {
		return
	}
	width := c.Width * 6350
	height := c.Height * 6350
	d.Document.AppendPicture(idFile["rid"], width, height, align)
	return
}

func (d *Docx) Replace(src, des string) {
	d.ReplaceMap[src] = des
}

func (d *Docx) Close() {
	_ = exec.Command(fmt.Sprintf("rm -rf %s", d.FilePath))
}

func (d *Docx) cpFile(filename string) (string, error) {
	partId := d.GetIdAndInc()
	basename := filepath.Base(filename)
	tFileBase := GenerateFileName(basename, fmt.Sprintf("part%d", partId), false)
	tFile := filepath.Join(d.FilePath, tFileBase)
	_, err := CopyFile(tFile, filename)
	if err != nil {
		return tFileBase, err
	}
	return tFileBase, nil
}

func Unzip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, 0777)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), 0777); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadAll(filePth string) (string, error) {
	f, err := os.Open(filePth)
	defer f.Close()
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	// _ = b
	return string(b), nil
}

//写入文件
func writeStringToFile(filepath, content string) (err error) {
	//打开文件，没有则创建，有则append内容
	w1, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return
	}
	_, err = w1.Write([]byte(content))
	if err != nil {
		return
	}

	err = w1.Close()
	if err != nil {
		return
	}
	return
}

// srcFile could be a single file or a directory
func (d *Docx) Zip(destZip string) error {
	srcFile := d.FilePath + string(filepath.Separator)
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		t, err := os.Open(path)
		defer t.Close()
		name := strings.TrimPrefix(path, srcFile)
		var flag = true
		if strings.HasSuffix(name, "embeddings") || strings.HasSuffix(name, "media") {
			if ok, err := d.canSave(info.Name()); !ok || err != nil {
				flag = false
			}
		}

		if flag {
			tf, err := archive.Create(name)
			if err != nil {
				return err
			}
			data, err := ioutil.ReadAll(t)
			if err != nil {
				return err
			}
			_, err = tf.Write(data)
			if err != nil {
				return err
			}
		}
		return err
	})

	return err
}

func MergeFiles(filelist []string, filename string, page bool) (err error) {
	if filelist == nil || len(filelist) == 0 {
		return fmt.Errorf("filelist is nil")
	}
	doc, err := NewDocx(filelist[0], false)
	if err != nil {
		return
	}
	doc.Merge(filelist[1:], page, false)
	doc.Save(filename)
	doc.Close()
	return
}

func MakeNewDocument() (doc *Docx, err error) {
	str, _ := os.Getwd()
	template := path.Join(str, "templates/template.docx")
	doc, err = NewDocx(template, false)
	if err != nil {
		return
	}
	d, err := doc.GetDocument()
	if err != nil {
		return
	}
	d.ClearContentWithoutHeader()
	return
}

func (d *Docx) printline() {
	p := path.Join(TEMP_BASE_DIR, d.BaseDir, "word/_rels/document.xml.rels")
	s, err := ReadAll(p)
	if err != nil {
		return
	}
	fmt.Println("lines:", len(s))
}

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		return
	}
	defer srcFile.Close()
	//通过srcfile ,获取到 Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}

	//通过dstFile, 获取到 Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)
}
