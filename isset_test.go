package phpGo

import (
	"testing"
)

type tmpStruct2 struct {
	A string
	B int
}

func TestIsset(t *testing.T) {
	var a int
	if Isset(a) {
		t.Error()
	}

	var f float32
	if Isset(f) {
		t.Error()
	}

	b := true
	if !Isset(b) {
		t.Error()
	}

	tmpArr := [2]interface{}{0, 1}
	if !Isset(tmpArr, 0) || !Isset(tmpArr, 1) {
		t.Error()
	}

	tmpMap := map[interface{}]interface{}{"asd": "1", 123: "2", "123": "3"}
	if Isset(tmpMap, "aa") || !Isset(tmpMap, 123) || !Isset(tmpMap, "asd") {
		t.Error()
	}

	tmpStruct := tmpStruct2{A: "123", B: 0}
	if !Isset(tmpStruct) || !Isset(tmpStruct, "A") || !Isset(tmpStruct, "B") {
		t.Error()
	}
}
