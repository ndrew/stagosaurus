package stagosaurus

import (
	"testing"
)

func TestFileSystemImpl(t *testing.T) {
	fs, err := NewFileSystem()
	if err != nil {
		t.Error(err)
	}

	var config Config = fs
	if nil == config {
		t.Error(WTF)
	}
}
