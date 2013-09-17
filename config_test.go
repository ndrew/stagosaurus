package stagosaurus

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// simple structure diff
//
func structDiff(a interface{}, b interface{}) (err error) {
	var typeA reflect.Type = toValueType(a)
	var typeB reflect.Type = toValueType(b)

	if typeA != typeB {
		return errors.New(fmt.Sprintf("%v has different type than %v", typeA, typeB))
	}

	if typeA.Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("%v is not reflect.Struct", typeA))
	}

	structA := toValue(a)
	structB := toValue(b)

	diffs := []string{}

	// loop through the struct's fields and set the map
	for i := 0; i < typeA.NumField(); i++ {
		vA := structA.Field(i)
		vB := structB.Field(i)

		if vA.Interface() != vB.Interface() {
			fieldType := typeA.Field(i)
			diffs = append(diffs, fmt.Sprintf("'%v' != '%v', field %v.%v", vA, vB, typeA, fieldType.Name))
		}
	}
	if len(diffs) > 0 {
		return errors.New(strings.Join(diffs, "\n"))
	}
	return nil
}

// testing the diff machinery
//
func TestDiff(t *testing.T) {
	v1 := "Test"
	v2 := 1
	err := structDiff(v1, v2)

	if err == nil {
		t.Error("diff should work only on same types")
	}

	err = structDiff(v1, v1)
	if err == nil {
		t.Error("diff should work only on structs")
	}

	c1 := new(Config)
	err = structDiff(c1, c1)
	assertNoError(err, t)

	c2 := new(Config)

	err = structDiff(c1, c2)
	assertNoError(err, t)

}

// test reading of dummy config
//
func TestConfig(t *testing.T) {
	config := new(Config)
	err := config.ReadConfig("test_data/sample-config.json")

	assertNoError(err, t)

	testConfig := new(Config)
	testConfig.BaseUrl = "http://localhost"
	testConfig.Port = ":8080"

	if !reflect.DeepEqual(config, testConfig) {
		err = structDiff(config, testConfig)
		assertNoError(err, t)
	}

}
