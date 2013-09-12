/*
   The blog package is a library for building static generated sites, usually blogs. 
*/
package blog

// Core of the blog
//
type Engine struct {
	cfg *Config
	// posts []*Post

	posts    PostFactory
	renderer Renderer
	deployer Deployer
}

// Constructor
//
func New(cfg *Config, posts PostFactory, renderer Renderer, deployer Deployer) *Engine {
	return &Engine{
		cfg:      cfg,
		renderer: renderer,
		deployer: deployer,
		posts:    posts,
	}
}

func (self Engine) Publish() (err error) { // TODO: add err handling
	e := self.renderer.RenderStarted()
	if e != nil {
		return e
	}

	posts, e := self.posts.GetPosts()
	if e != nil {
		return e
	}

	for _, post := range posts {
		// don't use post.Meta.Ready for more generic behaviour 
		e = self.renderer.Render(post)
		if e != nil {
			return e
		}
	}

	e = self.renderer.RenderEnded()
	if e != nil {
		return e
	}

	// todo: use
	return nil
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
