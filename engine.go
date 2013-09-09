/*
   The blog package is a library for building static generated sites, usually blogs. 
*/
package blog

import (
	//	"bytes"
	//"path/filepath"
	"time"
)

// Additional information about blog entry 
// 
type Meta struct {
	Ready   bool
	Title   string
	Date    time.Time
	Url     string
	Summary string
}

// Blog entry
//
type Post struct {
	Content string
	Meta    Meta
}

type Renderer interface {
	Render(post *Post) ([]byte, error)
	RenderStarted() error
	RenderEnded() error
}

type RenderingStrategy struct {
	// Renderers []Renderer
	// indexTemplate *template.Template
	// postTemplate  *template.Template
}

func (self *RenderingStrategy) Render(post *Post) ([]byte, error) {
	return []byte("FUUUUUU"), nil
}

func (self *RenderingStrategy) RenderStarted() error {
	println("RenderStarted")
	return nil
}

func (self *RenderingStrategy) RenderEnded() error {
	println("RenderEnded")
	// flush changes here 
	return nil

}

// Configuration about blog itself
//
type Config struct {
	BaseUrl string

	PublishDir   string
	TemplatesDir string
}

// Post retrival
//
type PostFactory interface {
	GetPosts() []*Post
}

// Post iterator from filesystem
//
type FolderPostFactory struct {
	PostsDir string
}

func (self FolderPostFactory) GetPosts() []*Post {
	posts := []*Post{}

	callback := func(file *file) {
		if file.isMarkdown() {
			post := new(Post)
			post.Content = file.Name()
			posts = append(posts, post)
		}
	}
	traverseFiles(self.PostsDir, callback)

	// get posts
	return posts
}

// Core of the blog
//
type Engine struct {
	cfg      *Config
	renderer Renderer
	posts    []*Post
}

// Constructor
//
func New(cfg *Config, renderer Renderer, posts []*Post) *Engine {
	return &Engine{
		cfg:      cfg,
		renderer: renderer,
		posts:    posts,
	}
}

func (self Engine) Publish() { // TODO: add err handling
	err := self.renderer.RenderStarted()
	if err == nil {
		println("error while starting rendering")
	}

	for _, post := range self.posts {
		if post.Meta.Ready {
			self.renderer.Render(post)
		}
	}

	self.renderer.RenderEnded()
}
