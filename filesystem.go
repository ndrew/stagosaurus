package blog

import (
	"io/ioutil"
	"os"
	"strings"
)

// file class with some convenince functions added
//
type file struct {
	os.FileInfo
}

// hidden constructor for file
//
func _file(f os.FileInfo) *file {
	return &file{f}
}

// return file extenstion
//
func (f file) GetExt() string {
	var a = strings.Split(f.Name(), ".")
	if 0 == len(a) {
		return ""
	}
	var ext = a[len(a)-1]
	return ext
}

// checks if file is markdown file
//
func (f file) isMarkdown() bool {
	return strings.EqualFold("md", f.GetExt()) || strings.EqualFold("markdown", f.GetExt())
}

func concatFilePaths(path1 string, path2 string) string {
	separator := string(os.PathSeparator)
	return strings.TrimSuffix(path1, separator) + separator + strings.TrimPrefix(path2, separator)
}

// traverses files in directory and runs a callback 
//
func traverseFiles(dir string, callback func(*file)) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		println("can't read from directory " + dir)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			traverseFiles(concatFilePaths(dir, f.Name()), callback)
		} else {
			callback(_file(f))
		}
	}
}
