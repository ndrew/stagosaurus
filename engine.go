/*
   The blog package is a library for building static generated sites, usually blogs. 
*/
package blog

import (
	//	"bytes"
	//"path/filepath"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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
	Render(post *Post) error
	RenderStarted() error
	RenderEnded() error
}

type RenderingStrategy struct {
	// Renderers []Renderer
	// indexTemplate *template.Template
	// postTemplate  *template.Template
}

func (self *RenderingStrategy) Render(post *Post) error {
	return nil
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

	Port string
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

func (self Engine) NewPost(postName string) {
	println(postName)
}

func (self Engine) EditPost(postName string) {

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
	// defer r.Body.Close()
}

func stopServer(w http.ResponseWriter, req *http.Request) {
	responseString := "Bye-bye"

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(responseString)))
	io.WriteString(w, responseString)

	f, canFlush := w.(http.Flusher)
	if canFlush {
		f.Flush()
	}

	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		//fmt.Printf("error while shutting down: %v", err)
	}

	conn.Close()

	println("Shutting down")
	os.Exit(0)
}

func (self Engine) RunServer(dir string, port string) { // "."
	//port.star

	http.HandleFunc("/exit", stopServer)
	http.HandleFunc("/preview", handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)

	//http.ListenAndServe(":8080", http.FileServer(http.Dir(dir)))
}
