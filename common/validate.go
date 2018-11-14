package common

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

func HasNoEmptyFields(object interface{}) (err error) {
	value := reflect.ValueOf(object)

	if object == nil {
		return errors.Errorf("Object is nil")
	}

	if value.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		name := value.Type().Field(i).Name

		if IsNil(field) {
			return errors.Errorf("Field %v is nil!", name)
		}

		if IsZero(field) {
			return errors.Errorf("Field %v is empty!", name)
		}
	}

	return nil
}

func IsNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

func IsZero(v reflect.Value) bool {
	for v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Struct && v.NumField() == 0 {
		return false
	}

	valueRep := fmt.Sprintf("%#v", v)
	zeroRep := fmt.Sprintf("%#v", reflect.Zero(v.Type()))

	return valueRep == zeroRep
}
