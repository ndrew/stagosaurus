/*
   The blog package is a library for building static generated sites, usually blogs.
*/
package stagosaurus

//import "reflect"

// Core of the blog
//
type Engine struct {
	cfg Configurable
	// posts []*Post

	posts    PostFactory
	renderer Renderer
	deployer Deployer
}

// generic constructor
//
func Create(args ...interface{}) *Engine {
	var (
		config   Configurable = nil
		posts    PostFactory  = nil
		renderer Renderer     = nil
		deployer Deployer     = nil
	)
	for _, arg := range args {
		switch v := arg.(type) {
		case Configurable:
			{
				println("configurable")
				println(arg)
				config = v
			}
		/*case int32, int64:
			fmt.Println(v)
		case SomeCustomType:
			fmt.Println(v)*/
		default:
			println("unknown")
		}
	}

	return New(config, posts, renderer, deployer)
}

// Full constructor
//
func New(cfg Configurable, posts PostFactory, renderer Renderer, deployer Deployer) *Engine {
	return &Engine{
		cfg:      cfg,
		renderer: renderer,
		deployer: deployer,
		posts:    posts,
	}
}

// regenerates website files
//
func (self Engine) Publish() (err error) { // TODO: add err handling
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
