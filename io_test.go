package stagosaurus

import (
	"testing"
)

func TestFileSystemImpl(t *testing.T) {
	config := EmptyConfig()
	config.Set("FS.HOME_DIR", "FOO")

	fs, err := NewFileSystem(config)
	if err != nil {
		t.Error(err)
	}

	var configTest Config = fs
	if nil == configTest {
		t.Error(WTF)
	}
}
