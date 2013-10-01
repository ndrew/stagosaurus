package stagosaurus

// 'Filestem' abstraction for retrieving/stroring Assets
//
type FileSystem struct {
}

func (this *FileSystem) Get(key ...interface{}) interface{} {
	return nil
}

func (this *FileSystem) Set(key interface{}, value interface{}) interface{} {
	return nil
}

func (this *FileSystem) Find(predicate func(interface{}, interface{}) bool) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	/*for k, v := range this.cfg {
	    if predicate(k, v) {
	        res[k] = v
	    }
	}*/
	return res
}

func (this *FileSystem) Validate(params ...interface{}) (bool, error) {
	return true, nil
}

//////////
// impl

// FileSytem constructor.
//
func NewFileSystem(config Config) (*FileSystem, error) {
	/*if v, ok := config.(site.Validator); ok {
	    // validate the config key 'greeting'
	    validator := map[interface{}](func(interface{}) bool){
	        "greeting": func(v interface{}) bool {
	            return v != nil && v != "Hello"
	        },
	    }

	    if original, _ := v.Validate(validator); !original {
	        return config, errors.New("You've provided too trivial value! Try again, be original!")
	    }
	}*/

	fs := new(FileSystem)

	validator := map[interface{}](func(interface{}) bool){
		"FS.HOME_DIR": func(v interface{}) bool {
			return v != nil // && v != "Hello"
		},
	}
	if ok, err := fs.Validate(validator); !ok {
		return fs, err
	}

	return fs, nil
}

/*
func (this *Config) FindByKey(predicate func(interface{}) bool) map[interface{}]interface{} {
    return this.Find(func(k interface{}, v interface{}) bool {
        return predicate(k)
    })
}

func (this *Config) FindByValue(predicate func(interface{}) bool) map[interface{}]interface{} {
    return this.Find(func(k interface{}, v interface{}) bool {
        return predicate(v)
    })
}

func (this *Config) Find(predicate func(interface{}, interface{}) bool) map[interface{}]interface{} {
    res := make(map[interface{}]interface{})
    for k, v := range this.cfg {
        if predicate(k, v) {
            res[k] = v
        }
    }
    return res
}

*/
