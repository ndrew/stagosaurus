package stagosaurus

import "testing"

func assertNoError(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}

// mocked engine
//
type FakeEngine struct {
	log string

	t *testing.T
	c *Config

	posts PostFactory
}

func (self *FakeEngine) Render(post *Post) error {
	self.log += "<render>"
	return nil
}

func (self *FakeEngine) RenderStarted() error {
	self.log += "<start>"
	return nil
}

func (self *FakeEngine) RenderEnded() error {
	self.log += "<end>"
	return nil
}

// renderer
func (self *FakeEngine) GetRenderedPosts() ([]*Post, error) {
	dummy := new(Post)
	dummy.Name = "DUMMY"
	return []*Post{dummy}, nil
}

func (self *FakeEngine) Deploy(posts []*Post) error {
	if len(posts) != 1 {
		self.t.Error("incorrect posts passed to deployer")
	}
	return nil
}

func (self *FakeEngine) GetConfig() *Config {
	return self.c
}

func (self *FakeEngine) New(data string, name string) (*Post, error) {
	return self.posts.New(data, name)
}

// posts
func (self *FakeEngine) GetPosts() (posts []*Post, err error) {
	return self.posts.GetPosts()
}

func TestEngine(t *testing.T) {
	cfg := new(Config)
	err := cfg.ReadConfig("test_data/sample-config.json")
	assertNoError(err, t)

	postsFactory := new(FileSystem)
	postsFactory.PostsDir = "test_data/posts"
	posts, err := postsFactory.GetPosts()
	assertNoError(err, t)

	dummy := new(FakeEngine)
	dummy.c = cfg
	dummy.t = t
	dummy.posts = postsFactory

	blog := New(dummy, dummy, dummy, dummy)

	blog1 := Create(dummy)
	println(blog1)

	err = blog.Publish()
	assertNoError(err, t)

	log := "<start>"
	for i := 0; i < len(posts); i++ {
		log += "<render>"
	}
	log += "<end>"

	if dummy.log != log {
		t.Errorf("renderering was done not as expected '%v' vs '%v' ", log, dummy.log)
	}
}
