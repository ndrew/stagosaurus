package stagosaurus

import (
	"errors"
)

type Config interface {
	Get(key ...interface{}) interface{}
	Set(key interface{}, value interface{}) interface{}
	Find(predicate func(interface{}, interface{}) bool) map[interface{}]interface{}

	// convenience funcs
	//	String(key ...interface{}) (string, error)
}

type Validator interface {
	Validate(params ...interface{}) (bool, error)
}

// Generic configuration class
//
//type Config struct {
//    Defaults *Config
//    cfg      map[interface{}]interface{}
//}

type MapConfig struct {
	Defaults Config
	cfg      map[interface{}]interface{}
}

// Constructor for Config. If defaults are not needed just do new(Config)
//
func NewConfig(defaults Config) Config {
	cfg := new(MapConfig)
	cfg.Defaults = defaults
	return cfg
}

// Get value by key from configuration dictionary
//
// Note: as go doesn't support polymorphic funcs, I do trick with variable arguments to have
// fluent interface for providing optional overriding default value
func (this *MapConfig) Get(key ...interface{}) interface{} {
	if this.cfg == nil {
		this.cfg = make(map[interface{}]interface{})
	}
	switch {
	case 0 == len(key):
		return this
	case 1 == len(key):
		if nil != this.Defaults && this != this.Defaults && nil == this.cfg[key[0]] {
			return this.Defaults.Get(key[0])
		}
		return this.cfg[key[0]]
	case 2 == len(key):
		v := this.cfg[key[0]]
		if nil == v {
			return key[1]
		}
		return v
	default:
		panic("Incorrect number of params passed to Config.Get")
	}
}

// Sets value for key in configuration dictionary
//
func (this *MapConfig) Set(key interface{}, value interface{}) interface{} {
	if this.cfg == nil {
		this.cfg = make(map[interface{}]interface{})
	}

	this.cfg[key] = value
	return this.Get(key)
}

// Configuration validation via key => validity-predicate map. Returns result of validation and list of failed keys, if any
//
func (this *MapConfig) Validate(predicateMap map[interface{}]func(interface{}) bool) (bool, []interface{}) {
	res := true
	noErrors := [0]interface{}{}
	errors := noErrors[:]

	for k, predicate := range predicateMap {
		if valid := predicate(this.Get(k)); false == valid {
			res = false
			errors = append(errors, k)
		}
	}

	return res, errors
}

func (this *MapConfig) FindByKey(predicate func(interface{}) bool) map[interface{}]interface{} {
	return this.Find(func(k interface{}, v interface{}) bool {
		return predicate(k)
	})
}

func (this *MapConfig) FindByValue(predicate func(interface{}) bool) map[interface{}]interface{} {
	return this.Find(func(k interface{}, v interface{}) bool {
		return predicate(v)
	})
}

func (this *MapConfig) Find(predicate func(interface{}, interface{}) bool) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	for k, v := range this.cfg {
		if predicate(k, v) {
			res[k] = v
		}
	}
	return res
}

// convenience getters

// string
//
func (this *MapConfig) String(key ...interface{}) (string, error) {
	v := this.Get(key...)
	if nil != v {
		if ret, ok := v.(string); ok {
			return ret, nil
		}
	}
	iv := "" // identity value
	return iv, errors.New("Value in config is not String")
}

//   func (v Value) Bytes() []byte

func (this *MapConfig) Bool(key ...interface{}) (bool, error) {
	v := this.Get(key...)
	if nil != v {
		if ret, ok := v.(bool); ok {
			return ret, nil
		}
	}
	iv := false // identity value
	return iv, errors.New("Value in config is not Bool")
}
