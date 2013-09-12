package stagosaurus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Configuration about blog itself
//
type Config struct {
	BaseUrl string

	PublishDir   string
	TemplatesDir string

	Port string
}

func (cfg *Config) ReadConfig(path string) (err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("can't read file "+path, err)
		return err
	}
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		fmt.Println("error while unmarshaling json ", err)
	}
	return err
}
