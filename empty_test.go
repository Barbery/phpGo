package phpGo

import (
	"testing"
)

type tmpStruct struct {
	A string
	B int
}

func TestEmpty(t *testing.T) {
	if !Empty(nil) {
		t.Error()
	}

	if !Empty("") {
		t.Error()
	}

	if !Empty("0") {
		t.Error()
	}

	if !Empty(false) {
		t.Error()
	}

	tmpArr := [2]interface{}{0, 1}
	if !Empty(tmpArr, 0) || Empty(tmpArr, 1) {
		t.Error()
	}

	if !Empty([]int{}) || !Empty([]string{}) {
		t.Error()
	}

	tmpMap := map[interface{}]interface{}{"asd": "1", 123: "2", "123": "3"}
	if !Empty(tmpMap, "aa") || Empty(tmpMap, 123) || Empty(tmpMap, "asd") {
		t.Error()
	}

	tmpStruct := tmpStruct{A: "123", B: 0}
	if Empty(tmpStruct) || Empty(tmpStruct, "A") || !Empty(tmpStruct, "B") {
		t.Error()
	}
}
