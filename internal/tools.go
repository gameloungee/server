package internal

import (
	"encoding/json"
	"io"
	"reflect"
)

func Unmarshal[T any](r io.Reader) (*T, error) {
	jsonByte, err := io.ReadAll(r)

	if err != nil {
		return nil, err
	}

	var obj T
	err = json.Unmarshal(jsonByte, &obj)

	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func HasNilFileds[T any](obj *T) bool {
	v := reflect.ValueOf(obj).Elem()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.IsNil() {
			return true
		}
	}

	return false
}
