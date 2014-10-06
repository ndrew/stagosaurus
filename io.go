package stagosaurus

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

/******************************
 * Playing with go io system
 */
type File struct {
	os.FileInfo
}

/**
 * return contents of the file
 */
func (f *File) Contents(sourceDir string) *[]byte {
	data, _ := ioutil.ReadFile(filepath.Join(sourceDir, f.Name()))
	return &data
}

/**
 * returns file extension
 */
func (f File) GetExt() string {
	var a = strings.Split(f.Name(), ".")
	if 0 == len(a) {
		return ""
	}
	var ext = a[len(a)-1]
	return ext
}

/**
 * checks if file is markdown file
 */
func (f File) isMarkdownFile() bool {
	return strings.EqualFold("md", f.GetExt()) || strings.EqualFold("markdown", f.GetExt())
}

// hidden constructor :)
func newFile(f os.FileInfo) *File {
	return &File{f}
}

/**
 * returns filename with new extension
 */
func (f File) SubstituteExt(with string) string {
	// TODO: handing with == ""
	var a = strings.Split(f.Name(), ".")
	if 0 == len(a) {
		a[len(a)] = with
	} else {
		a[len(a)-1] = with
	}
	return strings.Join(a, ".")
}

// 'Filestem' abstraction for retrieving/stroring Assets
//
type FileSystem struct {
	root string
}

func (this *FileSystem) Get(key ...interface{}) interface{} {
	return nil
}

func (this *FileSystem) Set(key interface{}, value interface{}) interface{} {
	return nil
}

func (this *FileSystem) Find(predicate func(interface{}, interface{}) bool) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	posts, _ := ioutil.ReadDir(this.root)

	for _, f := range posts {
		file := newFile(f)

		fname := filepath.Join(this.root, f.Name())
		if predicate(fname, file) {
			res[fname] = file
		}
	}
	return res
}

func (this *FileSystem) Validate(params ...interface{}) (bool, error) {
	return true, nil
}

//////////
// impl

// FileSytem constructor.
//
func NewFileSystem(cfg Config) (*FileSystem, error) {
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

	config := HumanConfig(cfg)

	validator := map[interface{}](func(interface{}) bool){
		"source-dir": func(v interface{}) bool {
			return v != nil // && v != "Hello"
		},
	}
	if ok, err := config.Validate(validator); !ok {
		return nil, err
	}

	root, _ := config.String("source-dir")

	fs := new(FileSystem)
	fs.root = root
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
