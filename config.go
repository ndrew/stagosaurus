package blog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// TODO: fix this mess
type AppCfg struct {
	Foo string

	// foo string - not exported
	//AparentAddr string
	//Aaddr string
	//Aip string        idmap map[string] string
}

//func NewFile(fd int, name string) *File {
//	if fd < 0 {
//		return nil
//	}
//	return &File{fd: fd, name: name} // ← Create a new File 
//}

//func NewConfig()

func HelloFromLib() {
	fmt.Println("hello from blog lib")
}

func (l *AppCfg) ReadConfig(path string) (err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("can't read file "+path, err)
		return
	}
	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Println("error while reading json ", err)
	}
	return
}
