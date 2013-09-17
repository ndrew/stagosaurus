package stagosaurus

import (
	"reflect"
)

// if a pointer to a struct is passed, get the type of the dereferenced object
//
func toValueType(i interface{}) reflect.Type {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

func toValue(i interface{}) reflect.Value {
	t := reflect.ValueOf(i)
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}
