package blog

import (
//    "io/ioutil"
//    "encoding/json"
    "fmt"
)

func HelloFromLib() {
    fmt.Println("hello from blog lib")
}

// TODO: fix this mess
/*type AppCfg struct{
        Foo string
        //AparentAddr string
        //Aaddr string
        //Aip string        idmap map[string] string
}

func (l *AppCfg) ReadConfig(path string) (err error) {
        b, err := ioutil.ReadFile(path)
        if err != nil {
                return
        }
        err = json.Unmarshal(b, &l)
        if err != nil {
                fmt.Print("error while reading json ", err)
        }
        return
}*/