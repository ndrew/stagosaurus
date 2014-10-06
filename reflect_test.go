package stagosaurus

import (
	"testing"
)

//
func TestReflection_ToMap_Convertation(t *testing.T) {
	data := make(map[string]string)
	data["foo"] = "bar"

	v, _ := ToMap(data)
	cfg := new(AConfig)
	cfg.cfg = v

	res, _ := cfg.String("foo")
	if res != "bar" {
		t.Error("Map convertion doesn't work")
	}

	//val := toValue(v["foo"])
	//println(val.String())

	//println(v.(map[string]string)["foo"])

}
