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

// File extends os.FileInfo
//
type File struct {
	os.FileInfo
}

// Contents — return contents of the file
//
func (f *File) Contents(sourceDir string) (*[]byte, error) {
	data, err := ioutil.ReadFile(filepath.Join(sourceDir, f.Name()))
	return &data, err
}

// GetExt — returns file extension
//
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

// SubstituteExt — returns filename with new extension
//
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

// FileSystem — 'Filestem' abstraction for retrieving/stroring Assets
//
type FileSystem struct {
	Root string
}

// Contents — return file contents
//
func (fs *FileSystem) Contents(f *File) (*[]byte, error) {
	return f.Contents(fs.Root)
}

// Get — tbd
//
func (fs *FileSystem) Get(key ...interface{}) interface{} {
	return nil
}

// Set — tbd
//
func (fs *FileSystem) Set(key interface{}, value interface{}) interface{} {
	return nil
}

// Find — finds files)
//
func (fs *FileSystem) Find(predicate func(interface{}, interface{}) bool) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	posts, _ := ioutil.ReadDir(fs.Root)

	for _, f := range posts {
		file := newFile(f)

		fname := filepath.Join(fs.Root, f.Name())
		if predicate(fname, file) {
			res[fname] = file
		}
	}
	return res
}

// Validate — tbd
//
func (fs *FileSystem) Validate(params ...interface{}) (bool, error) {
	return true, nil
}

//////////
// impl

// NewFileSystem — FileSytem constructor.
//
func NewFileSystem(root string) (*FileSystem, error) {
	fs := new(FileSystem)
	fs.Root = root
	return fs, nil
}
