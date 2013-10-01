package stagosaurus

import (
	"errors"
	"reflect"
)

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

func ToString(v interface{}) (string, error) {
	if nil != v {
		if ret, ok := v.(string); ok {
			return ret, nil
		}
	}
	iv := "" // identity value
	return iv, errors.New("Value in config is not String")
}

/*func ToAsset(v interface{}) (string, error) {
    if nil != v {
        if ret, ok := v.(string); ok {
            return ret, nil
        }
    }
    iv := "" // identity value
    return iv, errors.New("Value in config is not String")
}*/
