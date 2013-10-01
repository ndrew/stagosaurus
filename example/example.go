package main

import (
	site "./.." // substite this to "github.com/ndrew/stagosaurus" for production use
	"errors"
	"math/rand"
	"strings"
	"time"
)

var BlogHeader string

// Blog constructor
//
func New() (*Blog, error) { // blog constructor
	blog := new(Blog)

	return blog, nil
}

// your future blog generator
//
type Blog struct {
	Config site.Config
	Assets map[string]site.Asset
}

// ConfigSource
//
func (this *Blog) GetConfig() (site.Config, error) {
	// provide defaults for configuration
	defaults := new(site.MapConfig)
	defaults.Set("greeting", "Rhoaarrrr")

	config := site.NewConfig(defaults)
	config.Set("blogName", "Stagosaurus")

	if v, ok := config.(site.Validator); ok {
		// validate the config key 'greeting'
		validator := map[interface{}](func(interface{}) bool){
			"greeting": func(v interface{}) bool {
				return v != nil && v != "Hello"
			},
		}

		if original, _ := v.Validate(validator); !original {
			return config, errors.New("You've provided too trivial value! Try again, be original!")
		}
	}

	return config, nil
}

// PostsSource
//
func (this *Blog) GetPosts(meta site.Config) ([]site.Post, error) {
	assets := []site.Asset{}
	posts := []site.Post{}

	// generate random posts' contents
	rand.Seed(time.Now().UTC().UnixNano())
	buzzwords := []string{"Super", "Stoned", "Stealed", "Autogenerated", "Awesome", "Hipster", "Hilarious", "Nice"}

	for i := 1; i < 4; i++ {
		ri := rand.Intn(len(buzzwords))

		content := buzzwords[ri] + " post, bro!"

		ri = rand.Intn(len(buzzwords))
		p, err := site.NewPost(buzzwords[ri]+" Post", content, new(site.MapConfig), []site.Asset{})
		if err != nil {
			return []site.Post{}, err
		}

		posts = append(posts, p)
		// and append post as an asset for index posts for demonstration of hierarchical posts
		assets = append(assets, p)

	}

	post, err := site.NewPost("INDEX", "{HEADER}\nHere are my posts: \n{POSTS}", meta, assets)
	if err != nil {
		return []site.Post{}, err
	}

	posts = append(posts, post)
	return posts, nil
}

// rendering text in a fancy border
//
func (this *Blog) Render(cfg site.Config, posts []site.Post) ([]site.Post, error) {
	var results = []site.Post{}

	var p site.Post = nil
	var err error = nil

	for _, post := range posts {
		if "INDEX" == post.GetName() {
			p, err = this.renderIndex(post)
		} else {
			p, err = this.renderPost(post)
		}

		if err != nil {
			return results, err
		}
		results = append(results, p)

	}
	return results, nil
}

// renders index post
//
func (this *Blog) renderIndex(post site.Post) (site.Post, error) {
	hello := ""
	world := ""

	// cast to type manually
	helloProperty := post.GetConfig().Get("greeting")
	if helloProperty != nil {
		var ok bool = true
		if hello, ok = helloProperty.(string); !ok {
			return nil, errors.New("hello is not a string!")
		}
	}

	// or use shorthand for common types: string/bool/int
	world, err := site.ToString(post.GetConfig().Get("blogName"))
	if err != nil {
		return nil, err
	}

	header := hello + " " + world + "!"

	indexContent, err := post.GetContents()
	if err != nil {
		return nil, err
	}

	postsListing := ""

	// usually here you have to sort posts on some criteria (i.e. post date from meta-data), but I'll ommit it here
	for _, asset := range post.GetAssets() {
		if p, ok := asset.(site.Post); ok {
			postsListing += "\t - " + p.GetName() + "\n"
		}
	}

	content := strings.Replace(string(*indexContent), "{HEADER}", BlogHeader+"\n"+header, 1)
	content = strings.Replace(content, "{POSTS}", postsListing, 1)

	return site.NewPost("index.html", content, new(site.MapConfig), []site.Asset{})
}

// renders post
//
func (this *Blog) renderPost(post site.Post) (site.Post, error) {
	data, err := post.GetContents()
	if err != nil {
		return nil, err
	}

	content := strings.Replace(string(*data), "{HEADER}", BlogHeader, 1)
	return site.NewPost(strings.Replace(post.GetName(), " ", "_", 10)+".htm", content, new(site.MapConfig), []site.Asset{})
}

func (this *Blog) Deploy(config site.Config, posts []site.Post) ([]site.Post, error) {
	// here usually posts are being saved to filesystems, but for simplicity we will 'deploy' posts to screen
	for _, post := range posts {
		println(post.GetName())
		contents, err := post.GetContents()
		if err != nil {
			return []site.Post{}, err
		}
		println(string(*contents))
		println(strings.Repeat("=", 80))
	}

	return []site.Post{}, nil
}

//
// Entry Point
//
//
//
func main() {
	BlogHeader = "                                   \n" +
		"                                    `..-------..`    \n" +
		"                                   ./``       ```\\`  \n" +
		"                                  +``           ``\\  \n" +
		"                                  +` RHOAARRRR!!!  |` \n" +
		"                                   +.            ./` \n" +
		"                   `               `--``      `..:`  \n" +
		"             -```.: `...`            .-::.---..``    \n" +
		"            `-..``-.``/ `-``         -:``            \n" +
		"          `-``:`       `.```         ``              \n" +
		"           .-:`          --         \n" +
		"          ``--`          `/.   ``*. \n" +
		"            `:.    `.      ..  ./   \n" +
		"    `.   `/ `-``.   `.       `.`    \n" +
		"  `-+/----.--..-// ::--`/:...``     \n" +
		"          ``           `            \n"

		// most compact way to handle errors yet
	err := func() error {
		blog, err := New()
		if err != nil {
			return err
		}
		if config, err := blog.GetConfig(); err != nil {
			return err
		} else if posts, err := blog.GetPosts(config); err != nil {
			return err
		} else if renderedPosts, err := blog.Render(config, posts); err != nil {
			return err
		} else if _, err = blog.Deploy(config, renderedPosts); err != nil {
			return err
		}
		return nil
	}()

	if err != nil {
		println("Error while generation ", err.Error())
	}
}
