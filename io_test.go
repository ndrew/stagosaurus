package stagosaurus

import (
	"testing"
)

func TestFileSystemImpl(t *testing.T) {
	config := EmptyConfig()
	config.Set("source-dir", ".")

	fs, err := NewFileSystem(config)
	if err != nil {
		t.Error(err)
	}

	var configTest Config = fs
	if nil == configTest {
		t.Error(WTF)
	}

	res := fs.Find(func(k interface{}, v interface{}) bool {
		// fmt.Println(v.(*File).Name())
		return v.(*File).Name() == "io_test.go"
	})

	cfg := ConfigFromMap(res)
	f := cfg.Get("io_test.go")
	if f == nil {
		t.Error("filtering by filename had been broken")
	}

	file := f.(*File)
	content := string(*file.Contents("."))
	if content == "" {
		t.Error("file hasn't been read")
	}
}
