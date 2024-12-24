package services

import (
	"reflect"
)

func HasEmptyFields(v interface{}) bool {
	val := reflect.ValueOf(v)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String && field.String() == "" {
			return true
		}
	}
	return false
}
