package stagosaurus

import (
	"errors"
	"reflect"
)

//////////////////////////////////////////
// Convenience convertation functions
//

// ToString — interface{} -> string
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

// ToBool – interface{} -> bool
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

// ToAsset — interface{} -> Asset
//
func ToAsset(v interface{}) (Asset, error) {
	if nil != v {
		if ret, ok := v.(Asset); ok {
			return ret, nil
		}
	}
	return nil, errors.New("Value is not Asset")
}

// ToMap — interface{} -> map[interface{}]interface{}
//
func ToMap(v interface{}) (map[interface{}]interface{}, error) {
	if nil != v {
		ret := make(map[interface{}]interface{})
		tv := toValue(v)
		for _, k := range tv.MapKeys() {
			val := tv.MapIndex(k)
			ret[k.Interface()] = val.Interface()
		}

		return ret, nil
	}
	return nil, errors.New("Value is not a map")
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

// AsConfig — config->ExtendedConfig
//
func AsConfig(config interface{}) ExtendedConfig {
	t := reflect.ValueOf(config)
	i := t.Interface()

	r, _ := i.(ExtendedConfig)
	return r
}
