package phpGo

import (
	"reflect"
)

func Empty(args ...interface{}) bool {
	arrVal := reflect.ValueOf(args[0])
	if !arrVal.IsValid() {
		return true
	}

	if len(args) < 2 || args[1] == nil {
		return isEmptyValue(args[0])
	}

	switch arrVal.Kind() {
	case reflect.Array, reflect.Slice:
		if index, ok := args[1].(int); ok {
			return isEmptyValue(arrVal.Index(index).Interface())
		}
		return true

	case reflect.Map:
		indexVal := arrVal.MapIndex(reflect.ValueOf(args[1]))
		if indexVal.IsValid() == false {
			return true
		}
		return isEmptyValue(indexVal.Interface())

	case reflect.Struct:
		key := reflect.ValueOf(args[1]).String()
		if !arrVal.FieldByName(key).IsValid() {
			return true
		} else {
			return isEmptyValue(arrVal.FieldByName(key).Interface())
		}
	}

	return false
}

func isEmptyValue(value interface{}) bool {
	if value == nil {
		return true
	}

	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Bool:
		return !val.Bool()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return val.Float() == 0.00

	case reflect.Complex64, reflect.Complex128:
		return val.Complex() == 0+0i

	case reflect.String:
		realVal := val.String()
		return realVal == "" || realVal == "0"

	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return val.Len() == 0

	case reflect.Struct:
		return val.NumField() == 0
	}

	return false
}
