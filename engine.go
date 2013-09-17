/*
   The blog package is a library for building static generated sites, usually blogs.
*/
package stagosaurus

import (
	"errors"
	"fmt"
)

// Core of the blog
//
type Generator struct {
	cfg Configurable
	// posts []*Post

	posts    Posts
	renderer Renderer
	deployer Deployer
}

// generic constructor
//
func New(args ...interface{}) (*Generator, error) {
	var (
		config   Configurable = nil
		posts    Posts        = nil
		renderer Renderer     = nil
		deployer Deployer     = nil
	)
	for _, arg := range args {
		if v, ok := arg.(Configurable); ok {
			config = v
		}
		if v, ok := arg.(Posts); ok {
			posts = v
		}
		if v, ok := arg.(Renderer); ok {
			renderer = v
		}
		if v, ok := arg.(Deployer); ok {
			deployer = v
		}
	}

	if nil == config || nil == posts || nil == renderer || nil == deployer {
		return nil, errors.New(fmt.Sprintf("stagosaurus.Create(...) error: Some of needed interfaces hadn't been provided\nconfig\t%v\nposts\t%v\nrenderer\t%v\ndeployer\t%v", config, posts, renderer, deployer))
	}

	return NewGenerator(config, posts, renderer, deployer), nil
}

// Full constructor
//
func NewGenerator(cfg Configurable, posts Posts, renderer Renderer, deployer Deployer) *Generator {
	return &Generator{
		cfg:      cfg,
		renderer: renderer,
		deployer: deployer,
		posts:    posts,
	}
}

// regenerates website files
//
func (self Generator) Publish() (err error) { // TODO: add err handling
	// 1) get posts
	posts, e := self.posts.GetPosts()
	if e != nil {
		return e
	}

	// TODO: do we need a notifier?

	// 2) notify rendering started
	if e = self.renderer.RenderStarted(); e != nil {
		return e
	}

	// 3) render each post
	for _, post := range posts {
		if e = self.renderer.Render(post); e != nil {
			return e
		}
	}

	// 4) notify rendering done
	if e = self.renderer.RenderEnded(); e != nil {
		return e
	}

	// 5) return results
	posts, e = self.renderer.GetRenderedPosts()
	if e != nil {
		return e
	}

	if e = self.deployer.Deploy(posts); e != nil {
		return e
	}

	return nil
}

func (self Generator) NewPost(postName string) {
	println(postName)
}

func (self Generator) EditPost(postName string) {
	println(postName)
}

func (self Generator) RunServer(dir string, port string) { // "."
	//port.star
	runServer(dir, port)
}
