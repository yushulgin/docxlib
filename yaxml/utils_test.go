package yaxml

import "testing"

func TestGenerateTag(t *testing.T) {
	res := GenerateTag("asdasdsad>asd\n<asdasdas<sadasd>sada")
	t.Log("res", res)
	// t.Error(res)
}

func TestIsDocumentStart(t *testing.T) {
	res := IsDocumentStart("?asdasdasdasdada?>")
	t.Log(res)
	// t.Error(res)
}

func TestGenerateElement(t *testing.T) {
	
}
