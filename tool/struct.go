package tool

import (
	"errors"
	"reflect"
)

// 复制属性
func CopyStructProperty(src interface{}, dst interface{}) error {
	srcValue := reflect.ValueOf(src)
	if srcValue.Type().Kind() == reflect.Ptr {
		srcValue = srcValue.Elem()
	}
	if srcValue.Kind() != reflect.Struct {
		return errors.New("src error")
	}

	dstValue := reflect.ValueOf(dst)
	if dstValue.Type().Kind() != reflect.Ptr || dstValue.Elem().Kind() != reflect.Struct {
		return errors.New("src error")
	}

	keys := make(map[string]interface{})
	for i := 0; i < srcValue.NumField(); i++ {
		keys[srcValue.Type().Field(i).Name] = true
	}

	for i := 0; i < dstValue.Elem().NumField(); i++ {
		name := dstValue.Elem().Type().Field(i).Name
		if _, ok := keys[name]; ok {
			value := srcValue.FieldByName(name)
			if value.CanInterface() && value.CanSet() && value.CanAddr() {
				dstValue.Elem().Field(i).Set(value)
			}
		}
	}
	return nil
}
