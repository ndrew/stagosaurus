package stagosaurus

import (
	"errors"
	"reflect"
)

//////////////////////////////////////////
// Convenience convertation functions
//

// interface{} -> string
//
func ToString(v interface{}) (string, error) {
	if nil != v {
		if ret, ok := v.(string); ok {
			return ret, nil
		}
	}
	iv := "" // identity value
	return iv, errors.New("Value is not String")
}

// interface{} -> bool
//
func ToBool(v interface{}) (bool, error) {
	if nil != v {
		if ret, ok := v.(bool); ok {
			return ret, nil
		}
	}
	iv := false // identity value
	return iv, errors.New("Value is not Bool")
}

// interface{} -> Asset
//
func ToAsset(v interface{}) (Asset, error) {
	if nil != v {
		if ret, ok := v.(Asset); ok {
			return ret, nil
		}
	}
	return nil, errors.New("Value is not Asset")
}

/////////////////////
// reflection stuff

var typeOfBytes = reflect.TypeOf([]byte(nil))

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
