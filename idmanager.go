package docxlib

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

type IdAble struct {
	Id   string
	Num  int
	Part int
}

func GenerateIdAble() (ia *IdAble, err error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return
	}
	ia.Id = strings.ReplaceAll(uid.String(), "-", "")
	rand.Seed(time.Now().UnixNano())
	ia.Num = rand.Intn(60000-100) + 100
	ia.Part = 0
	return
}

func (id *IdAble) GetIdAndInc() int {
	part := id.Part
	id.Part += 1
	return part
}

func GenerateId(sid, suffix, separation string) (toId string) {
	if sid == "" {
		return
	}
	toId = sid + separation + suffix
	return
}

func GenerateFileName(fName, suffix string, replace bool) (tempName string) {
	if fName == "" {
		return
	}
	fnames := strings.Split(fName, ".")
	//
	if replace {
		tempName = suffix + "." + fnames[len(fnames)-1]
	} else {
		tempName = fnames[0] + "_" + suffix + "." + fnames[len(fnames)-1]
	}
	return
}
