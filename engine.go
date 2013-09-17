/*
   The blog package is a library for building static generated sites, usually blogs.
*/
package stagosaurus

// Core of the blog
//
type Engine struct {
	cfg Congigurable
	// posts []*Post

	posts    PostFactory
	renderer Renderer
	deployer Deployer
}

// generic constructor
//
func Create(args ...interface{}) *Engine {
	var (
		config   Congigurable = nil
		posts    PostFactory  = nil
		renderer Renderer     = nil
		deployer Deployer     = nil
	)
	for _, arg := range args {
		println(arg)
	}

	return New(config, posts, renderer, deployer)
}

// Full constructor
//
func New(cfg Congigurable, posts PostFactory, renderer Renderer, deployer Deployer) *Engine {
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

	if e = self.renderer.RenderEnded(); e != nil {
		return e
	}

	posts, e = self.renderer.GetRenderedPosts()
	if e != nil {
		return e
	}

	if e = self.deployer.Deploy(posts); e != nil {
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
