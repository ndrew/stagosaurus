package stagosaurus

import (
	"errors"
)

// Constructor for Config. If defaults are not needed just do new(Config)
//
func NewConfig(defaults *Config) *Config {
	cfg := new(Config)
	cfg.Defaults = defaults
	return cfg
}

// Generic configuration class
//
type Config struct {
	Defaults *Config
	cfg      map[interface{}]interface{}
}

// Get value by key from configuration dictionary
//
// Note: as go doesn't support polymorphic funcs, I do trick with variable arguments to have
// fluent interface for providing optional overriding default value
func (this *Config) Get(key ...interface{}) interface{} {
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
func (this *Config) Set(key interface{}, value interface{}) interface{} {
	if this.cfg == nil {
		this.cfg = make(map[interface{}]interface{})
	}

	this.cfg[key] = value
	return this.Get(key)
}

// Configuration validation via key => validity-predicate map. Returns result of validation and list of failed keys, if any
//
func (this *Config) Validate(predicateMap map[interface{}]func(interface{}) bool) (bool, []interface{}) {
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

// convenience getters

// string
//
func (this *Config) String(key ...interface{}) (string, error) {
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

func (this *Config) Bool(key ...interface{}) (bool, error) {
	v := this.Get(key...)
	if nil != v {
		if ret, ok := v.(bool); ok {
			return ret, nil
		}
	}
	iv := false // identity value
	return iv, errors.New("Value in config is not Bool")
}
