package stagosaurus

import (
	"errors"
)

// Generic Config interface
//
type Config interface {
	Get(key ...interface{}) interface{}
	Set(key interface{}, value interface{}) interface{}
	Find(predicate func(interface{}, interface{}) bool) map[interface{}]interface{}
	Validate(params ...interface{}) (bool, error)
}

//
//
type ExtendedConfig interface {
	Get(key ...interface{}) interface{}
	Set(key interface{}, value interface{}) interface{}
	Find(predicate func(interface{}, interface{}) bool) map[interface{}]interface{}
	Validate(params ...interface{}) (bool, error)

	Bool(key ...interface{}) (bool, error)
	String(key ...interface{}) (string, error)
	SubConfig(key interface{}) (ExtendedConfig, error)
}

// Generic validator
//
type Validator interface {
	Validate(params ...interface{}) (bool, error)
}

// KeyValue config implementation
//
type AConfig struct {
	Defaults Config
	cfg      map[interface{}]interface{}
}

// Get value by key from configuration dictionary
//
// Note: as go doesn't support polymorphic funcs, I do trick with variable arguments to have
// fluent interface for providing optional overriding default value
func (this *AConfig) Get(key ...interface{}) interface{} {
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
func (this *AConfig) Set(key interface{}, value interface{}) interface{} {
	if this.cfg == nil {
		this.cfg = make(map[interface{}]interface{})
	}

	this.cfg[key] = value
	return this.Get(key)
}

// TBD — do generic validation
// Configuration validation via key => validity-predicate map.
//  Returns result of validation and list of failed keys, if any
//
func (this *AConfig) Validate(params ...interface{}) (bool, error) {
	if 1 == len(params) {
		if predicateMap, ok := params[0].(map[interface{}]func(interface{}) bool); ok {
			res := true
			noErrors := [0]interface{}{}
			errs := noErrors[:]

			for k, predicate := range predicateMap {
				if valid := predicate(this.Get(k)); false == valid {
					res = false
					errs = append(errs, k)
				}
			}
			if len(errs) > 0 {
				// TBD: more details
				return res, errors.New("Validation failed")
			}
			return res, nil
		}
	}
	return false, errors.New("Unknown stuff to validate")
}

// Search function
//
func (this *AConfig) Find(predicate func(interface{}, interface{}) bool) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	for k, v := range this.cfg {
		if predicate(k, v) {
			res[k] = v
		}
	}
	return res
}

// convenience search
//

func (this *AConfig) FindByKey(predicate func(interface{}) bool) map[interface{}]interface{} {
	return this.Find(func(k interface{}, v interface{}) bool {
		return predicate(k)
	})
}

func (this *AConfig) FindByValue(predicate func(interface{}) bool) map[interface{}]interface{} {
	return this.Find(func(k interface{}, v interface{}) bool {
		return predicate(v)
	})
}

// subconfig
//
func (this *AConfig) SubConfig(key interface{}) (ExtendedConfig, error) {
	cfg := new(AConfig)

	v, err := ToMap(this.Get(key))
	if err != nil {
		return cfg, errors.New("couldn't get a subconfig")
	}

	cfg.cfg = v
	return cfg, nil
}

//////////////
// convenience getters, but hence these will be mostly invisible - use convertation functions like ToString/ToBool/etc.

// string
//
func (this *AConfig) String(key ...interface{}) (string, error) {
	return ToString(this.Get(key...))
}

// bool
//
func (this *AConfig) Bool(key ...interface{}) (bool, error) {
	return ToBool(this.Get(key...))
}

// Constructor for empty Config
//
func EmptyConfig() Config {
	cfg := new(AConfig)
	return cfg
}

// Constructor for Config. If defaults are not needed just do new(Config)
//
func NewConfig(defaults Config) Config {
	cfg := new(AConfig)
	cfg.Defaults = defaults
	return cfg
}

func HumanConfig(config Config) ExtendedConfig {
	cfg := new(AConfig)
	cfg.Defaults = config
	return cfg
}
