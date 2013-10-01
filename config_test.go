package stagosaurus

import (
	"testing"
)

var WTF string = "Pure Magic just happened!" // always wanted to do this!

// custom assertion
//
func assertNoError(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}

// testing contracts from Config class
//
func testConfig(c Config, t *testing.T) {
	// Get() with no params should return config itself
	if c != c.Get() {
		t.Error("Get doesn't return all config itself")
	}

	// non-existing property, no defaults => nil
	if v := c.Get("property"); v != nil {
		t.Error(WTF)
	}

	// non-existing property, default value => default value
	if v := c.Get("property", "default-value"); v != "default-value" {
		t.Error(WTF)
	}

	// test setter return value
	if v := c.Set("property", "value"); v != "value" {
		t.Error(WTF)
	}

	// test setter
	if v := c.Get("property"); v != "value" {
		t.Error(WTF)
	}

	// test String func
	if s, err := ToString(c.Get("property")); s != "value" {
		t.Error(err)
	}

	if s, err := ToString(c.Get("non-existing-property")); s != "" || err == nil {
		t.Error(WTF)
	}

	// test Bool func
	if v := c.Set("bool", true); v != true {
		t.Error(WTF)
	}

	if b, err := ToBool(c.Get("bool")); b != true {
		t.Error(err)
	}

	if b, err := ToBool(c.Get("non-exisitng-bool")); b != false || err == nil {
		t.Error(WTF)
	}

	// test for panic in case get with awkward number of args, ensure that it's done in last statement
	defer func() {
		if r := recover(); r == nil {
			t.Error("not recovered from panic")
		}
	}()

	c.Get("property", "default", "WTF")
	// ... dead code here
}

// test case config without any defaults
//
func TestConfig_Sys_Constructor(t *testing.T) {
	c := EmptyConfig()
	testConfig(c, t)
}

// test case for config with defaults
//
func TestConfig_Conveniece_Constructor(t *testing.T) {
	defaults := EmptyConfig()
	defaults.Set("default1", "1")
	defaults.Set("default2", "2")

	c := NewConfig(defaults)
	testConfig(c, t)

	// test geting defaults
	if v := c.Get("default1"); v != "1" {
		t.Error(WTF)
	}

	// test overriding defaults
	if v := c.Set("default2", "0"); v != "0" {
		t.Error(WTF)
	}

	if v := c.Get("default2"); v != "0" {
		t.Error(WTF)
	}
}
