package stagosaurus

import "testing"

// custom assertion
//
func assertNoError(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}

// mocked blog engine
//
type FakeEngine struct {
	t *testing.T
	c *Config

	posts Posts

	log string
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

func getFakeImpl(t *testing.T) *FakeEngine {
	cfg := new(Config)
	cfg.ReadConfig("test_data/sample-config.json")

	postsFactory := new(FileSystem)
	postsFactory.PostsDir = "test_data/posts"

	dummy := new(FakeEngine)
	dummy.c = cfg
	dummy.t = t
	dummy.posts = postsFactory

	return dummy
}

func testEngine(blog *Engine, dummy *FakeEngine, t *testing.T) {
	posts, err := dummy.posts.GetPosts()
	assertNoError(err, t)

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

func TestEngine(t *testing.T) {
	dummy := getFakeImpl(t)
	blog := New(dummy, dummy, dummy, dummy)
	testEngine(blog, dummy, t)
}
