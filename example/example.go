package main

import (
	site "./.." // substite this to "github.com/ndrew/stagosaurus" for production use
	"errors"
	"math/rand"
	"strings"
	"time"
)

// Blog constructor
//
func New() (*Blog, error) { // blog constructor
	blog := new(Blog)

	return blog, nil
}

// your future blog generator
//
type Blog struct {
	Config *site.Config
	Assets map[string]site.Asset
}

// ConfigSource
//
func (this *Blog) GetConfig() (*site.Config, error) {
	// provide defaults for configuration
	defaults := new(site.Config)
	defaults.Set("greeting", "Rhoaarrrr")

	config := site.NewConfig(defaults)
	config.Set("blogName", "Stagosaurus")

	// validate the config key 'greeting'
	validator := map[interface{}](func(interface{}) bool){
		"greeting": func(v interface{}) bool {
			return v != nil && v != "Hello"
		},
	}
	if original, _ := config.Validate(validator); !original {
		return config, errors.New("You've provided too trivial value! Try again, be original!")
	}

	return config, nil
}

// PostsSource
//
func (this *Blog) GetPosts() ([]site.Post, error) {
	meta, err := this.GetConfig()
	if err != nil {
		return []site.Post{}, err
	}

	assets := []site.Asset{}

	// generate random posts' contents
	rand.Seed(time.Now().UTC().UnixNano())
	buzzwords := []string{"Super", "Awesome", "Hipster", "Hilarious", "Nice"}

	for i := 1; i < 4; i++ {
		ri := rand.Intn(len(buzzwords))

		content := buzzwords[ri] + " post, bro!"

		ri = rand.Intn(len(buzzwords))
		p, err := site.NewPost(buzzwords[ri]+" Post", content, new(site.Config), []site.Asset{})
		if err != nil {
			return []site.Post{}, err
		}
		assets = append(assets, p)
	}

	post, err := site.NewPost("Index Page", "Here are my posts:", meta, assets)
	if err != nil {
		return []site.Post{}, err
	}

	return []site.Post{post}, nil
}

// rendering text in a fancy border
//
func (this *Blog) Render(posts []site.Post) error {
	if len(posts) != 1 {
		return errors.New("I can render only one post. Yet")
	}
	post := posts[0]

	hello := ""
	world := ""

	// cast to type manually
	helloProperty := post.GetMeta().Get("greeting")
	if helloProperty != nil {
		var ok bool = true
		if hello, ok = helloProperty.(string); !ok {
			return errors.New("hello is not a string!")
		}
	}
	// or use shorthand for common types: string/bool/int
	world, err := post.GetMeta().String("blogName")
	if err != nil {
		return err
	}

	header := hello + " " + world + "!"
	filler := strings.Repeat("═", len(header))

	toc, err := post.GetContents()
	if err != nil {
		return err
	}

	postsListing := ""

	// usually here you have to sort posts on some criteria (i.e. post date from meta-data), but I'll ommit it here
	for _, asset := range post.GetAssets() {
		if p, ok := asset.(site.Post); ok {
			data, err := p.GetContents()
			if err != nil {
				return err
			}
			postsListing += "\t - " + p.GetName() + "\n\t\t\t" + string(*data) + "\n"
		}
	}

	println("╔═" + filler + "═╗")
	println("╟ " + header + " ╢")
	println("╚═" + filler + "═╝\n")

	println(string(*toc))
	println(postsListing)

	return nil
}

//
// Entry Point
//
//
//
func main() {
	// create a generator
	blog, err := New()
	if err != nil {
		println("Can't instantiate a blog: " + err.Error())
		return
	}
	// retrieve posts
	posts, err := blog.GetPosts()
	if err != nil {
		println("Error while retrieving posts: " + err.Error())
		return
	}
	// render them
	if err = blog.Render(posts); err != nil {
		println("Error while rendering: " + err.Error())
	}
}
