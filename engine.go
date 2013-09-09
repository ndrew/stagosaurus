/*
   The blog package is a library for building static generated sites, usually blogs. 
*/
package blog

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
	println(postName)
}

func (self Engine) RunServer(dir string, port string) { // "."
	//port.star
	runServer(dir, port)
}
