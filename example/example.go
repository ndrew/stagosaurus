package main

import (
	site "./.." // substite this to "github.com/ndrew/stagosaurus" for production use
	"errors"
	"strings"
)

func main() {
	// provide defaults for configuration
	defaults := new(site.Config)
	defaults.Set("greeting", "Hello")

	config := site.NewConfig(defaults)
	config.Set("blogName", "Stagosaurus")

	// create 'blog' example
	blog, err := New(config)
	if err != nil {
		// config validation: ["greeting"=>"Hello"] won't pass, so we fix config and recreate blog
		config.Set("greeting", "Rhoaarrrr")
		blog, err = New(config)
	}

	// we don't have any posts yet, so will be rendering data from config
	if err = blog.HelloWorld(); err == nil {
		println("\nCongratulation, you did it!")
	} else {
		println(err.Error())
	}
}

func New(config *site.Config) (*Blog, error) { // blog constructor
	// validate the config key 'greeting'
	validator := map[interface{}](func(interface{}) bool){
		"greeting": func(v interface{}) bool {
			return v != nil && v != "Hello"
		},
	}
	if original, _ := config.Validate(validator); !original {
		return nil, errors.New("You've provided too trivial value! Try again, be original!")
	}

	blog := new(Blog)
	blog.Config = config
	return blog, nil
}

// your future blog generator
type Blog struct {
	Config *site.Config
}

// test validation and 'publish' hello-world
//
func (this *Blog) HelloWorld() error {
	output := ""

	// cast to type manually
	hello := this.Config.Get("greeting")
	if hello != nil {
		if helloStr, ok := hello.(string); ok {
			output = helloStr
		} else {
			return errors.New("hello is not a string!")
		}
	}
	// or use shorthand for common types: string/bool/int
	world, err := this.Config.String("blogName")
	if err != nil {
		return err
	}
	output = output + " " + world + "!"

	return this.Render(output)
}

// rendering text in a fancy border
//
func (this *Blog) Render(output string) error {
	filler := strings.Repeat("═", len(output))

	println("╔═" + filler + "═╗")
	println("╟ " + output + " ╢")
	println("╚═" + filler + "═╝")

	return nil
}
