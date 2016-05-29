package stagosaurus

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// Config — generic Config interface
//
type Config interface {
	Get(key ...interface{}) interface{}
	Set(key interface{}, value interface{}) interface{}
	Find(predicate func(interface{}, interface{}) bool) map[interface{}]interface{}
	Validate(params ...interface{}) (bool, error)
}

// ExtendedConfig — convenience wrapper for Config -
//
type ExtendedConfig interface {
	Get(key ...interface{}) interface{}
	Set(key interface{}, value interface{}) interface{}
	Validate(params ...interface{}) (bool, error)

	Bool(key ...interface{}) (bool, error)
	String(key ...interface{}) (string, error)
	SubConfig(key interface{}) (ExtendedConfig, error)

	Find(predicate func(interface{}, interface{}) bool) map[interface{}]interface{}
	//	FindByKey(predicate func(interface{}) bool) map[interface{}]interface{}
	//	FindByValue(predicate func(interface{}) bool) map[interface{}]interface{}

}

// Validator — generic validator
//
type Validator interface {
	Validate(params ...interface{}) (bool, error)
}

// AConfig - KeyValue config implementation
//
type AConfig struct {
	Defaults Config
	cfg      map[interface{}]interface{}
}

// Get value by key from configuration dictionary
//
// Note: as go doesn't support polymorphic funcs, I do trick with variable arguments to have
// fluent interface for providing optional overriding default value
func (config *AConfig) Get(key ...interface{}) interface{} {
	if config.cfg == nil {
		config.cfg = make(map[interface{}]interface{})
	}
	switch {
	case 0 == len(key):
		return config
	case 1 == len(key):
		if nil != config.Defaults && config != config.Defaults && nil == config.cfg[key[0]] {
			return config.Defaults.Get(key[0])
		}
		return config.cfg[key[0]]
	case 2 == len(key):
		v := config.cfg[key[0]]
		if nil == v {
			return key[1]
		}
		return v
	default:
		panic("Incorrect number of params passed to Config.Get")
	}
}

// Set - sets value for key in configuration dictionary
//
func (config *AConfig) Set(key interface{}, value interface{}) interface{} {
	if config.cfg == nil {
		config.cfg = make(map[interface{}]interface{})
	}

	config.cfg[key] = value
	return config.Get(key)
}

// Validate — does validation
//
// Configuration validation via key => validity-predicate map.
// Returns result of validation and list of failed keys, if any
// TBD — do generic validation
func (config *AConfig) Validate(params ...interface{}) (bool, error) {
	if 1 == len(params) {
		if predicateMap, ok := params[0].(map[interface{}]func(interface{}) bool); ok {
			res := true
			noErrors := [0]interface{}{}
			errs := noErrors[:]

			for k, predicate := range predicateMap {
				if valid := predicate(config.Get(k)); false == valid {
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

// Find — search function
//
func (config *AConfig) Find(predicate func(interface{}, interface{}) bool) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	for k, v := range config.cfg {
		if predicate(k, v) {
			res[k] = v
		}
	}
	return res
}

// FindByKey — convenience search by key
//
func (config *AConfig) FindByKey(predicate func(interface{}) bool) map[interface{}]interface{} {
	return config.Find(func(k interface{}, v interface{}) bool {
		return predicate(k)
	})
}

// FindByValue — convenience search by value
//
func (config *AConfig) FindByValue(predicate func(interface{}) bool) map[interface{}]interface{} {
	return config.Find(func(k interface{}, v interface{}) bool {
		return predicate(v)
	})
}

// SubConfig - returns config for key
//
func (config *AConfig) SubConfig(key interface{}) (ExtendedConfig, error) {
	val := config.Get(key)

	if existingConfig, ok := val.(ExtendedConfig); ok {
		return existingConfig, nil
	}
	cfg := new(AConfig)

	v, err := ToMap(val)
	if err != nil {
		return cfg, errors.New("couldn't get a subconfig")
	}

	cfg.cfg = v
	return cfg, nil
}

//////////////
// convenience getters, but hence these will be mostly invisible - use convertation functions like ToString/ToBool/etc.

// String - return configuration for key as string
//
func (config *AConfig) String(key ...interface{}) (string, error) {
	return ToString(config.Get(key...))
}

// Bool - return configuration for key as bool
//
func (config *AConfig) Bool(key ...interface{}) (bool, error) {
	return ToBool(config.Get(key...))
}

// EmptyConfig — constructor for empty Config
//
func EmptyConfig() Config {
	cfg := new(AConfig)
	return cfg
}

// NewConfig — Constructor for Config. If defaults are not needed just do new(Config)
//
func NewConfig(defaults Config) Config {
	cfg := new(AConfig)
	cfg.Defaults = defaults
	return cfg
}

// HumanConfig — creates ExtendedConfig from Config
//
func HumanConfig(config Config) ExtendedConfig {
	cfg := new(AConfig)
	cfg.Defaults = config
	return cfg
}

// ConfigFromMap — creates config from map
//
func ConfigFromMap(m map[interface{}]interface{}) ExtendedConfig {
	config := new(AConfig)
	config.cfg = m
	return config
}

// ReadJSONConfig — reads json files and creates Config from it
//
func ReadJSONConfig(path string, defaults Config) (Config, error) {
	cfg := NewConfig(defaults)

	var realpath, err = filepath.Abs(path)
	if err != nil {
		return cfg, err
	}

	source, err := ioutil.ReadFile(realpath)
	if err != nil {
		return cfg, fmt.Errorf("ERR: Config file '%v' is not found.\n", realpath)
	}

	//
	// JSON stuff
	//
	// TODO: move config reading to stago lib
	var data map[string]*json.RawMessage
	err = json.Unmarshal(source, &data)

	if err != nil || len(data) == 0 {
		return cfg, fmt.Errorf("ERR: can't parse JSON from '%v'\n", realpath)
	}

	for k, v := range data {
		var value interface{}
		err = json.Unmarshal(*v, &value)

		if err == nil {
			cfg.Set(k, value)
		} else {
			// does this really occurs?
			return cfg, fmt.Errorf("ERR: couldn't interpret json, '%v':%v \n", k, *v)
		}
	}

	return cfg, nil
}
