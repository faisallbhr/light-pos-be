package utils

import (
	"reflect"
)

func GetStructFieldNames[T any]() []string {
	var names []string
	var t T
	tType := reflect.TypeOf(t)
	if tType.Kind() == reflect.Ptr {
		tType = tType.Elem()
	}
	for i := 0; i < tType.NumField(); i++ {
		names = append(names, tType.Field(i).Tag.Get("json"))
	}
	return names
}
