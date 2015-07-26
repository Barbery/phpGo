package phpGo

import (
	"reflect"
)

func Isset(args ...interface{}) bool {
	arrVal := reflect.ValueOf(args[0])
	if !arrVal.IsValid() {
		return false
	}

	if len(args) < 2 || args[1] == nil {
		return issetValue(args[0])
	}

	switch arrVal.Kind() {
	case reflect.Array, reflect.Slice:
		if index, ok := args[1].(int); ok {
			return arrVal.Index(index).IsValid()
		}

		return false
	case reflect.Map:
		indexVal := arrVal.MapIndex(reflect.ValueOf(args[1]))
		if indexVal.IsValid() == false {
			return false
		}

		return true
	case reflect.Struct:
		key := reflect.ValueOf(args[1]).String()
		if !arrVal.FieldByName(key).IsValid() {
			return false
		} else {
			return true
		}
	}

	return false
}

func issetValue(value interface{}) bool {
	if value == nil {
		return false
	}

	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Bool:
		return val.Bool()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() != 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() != 0

	case reflect.Float32, reflect.Float64:
		return val.Float() != 0.00

	case reflect.Complex64, reflect.Complex128:
		return val.Complex() != 0+0i

	case reflect.String:
		return val.String() != ""

	}

	return val.IsValid()
}
