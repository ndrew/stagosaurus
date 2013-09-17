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
type FakeGenerator struct {
	t *testing.T
	c *Config

	posts Posts

	log string
}

// dummy implementation of Renderer

func (self *FakeGenerator) Render(post *Post) error {
	self.log += "<render>"
	return nil
}

func (self *FakeGenerator) RenderStarted() error {
	self.log += "<start>"
	return nil
}

func (self *FakeGenerator) RenderEnded() error {
	self.log += "<end>"
	return nil
}

func (self *FakeGenerator) GetRenderedPosts() ([]*Post, error) {
	dummy := new(Post)
	dummy.Name = "DUMMY"
	return []*Post{dummy}, nil
}

// dummy implementation of Deployer

func (self *FakeGenerator) Deploy(posts []*Post) error {
	if len(posts) != 1 {
		self.t.Error("incorrect posts passed to deployer")
	}
	return nil
}

// dummy implementation of Configurable

func (self *FakeGenerator) GetConfig() *Config {
	return self.c
}

// dummy implementation of Posts

func (self *FakeGenerator) New(data string, name string) (*Post, error) {
	return self.posts.New(data, name)
}

func (self *FakeGenerator) GetPosts() (posts []*Post, err error) {
	return self.posts.GetPosts()
}

//

func getFakeImpl(t *testing.T) *FakeGenerator {
	cfg := new(Config)
	// maybe use config created by hand?
	cfg.ReadConfig("test_data/sample-config.json")

	postsFactory := new(FileSystem)
	postsFactory.PostsDir = "test_data/posts"

	dummy := new(FakeGenerator)
	dummy.c = cfg
	dummy.t = t
	dummy.posts = postsFactory

	return dummy
}

func testGeneratorFunctionality(blog *Generator, dummy *FakeGenerator, t *testing.T) {
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

// test basic operation
//
func TestGenerator(t *testing.T) {
	dummy := getFakeImpl(t)
	blog := NewGenerator(dummy, dummy, dummy, dummy)
	testGeneratorFunctionality(blog, dummy, t)
}

// test basic operation
//
func TestGeneratorCreation(t *testing.T) {
	dummy := getFakeImpl(t)
	blog, err := New(dummy)
	assertNoError(err, t)

	testGeneratorFunctionality(blog, dummy, t)

	blog, err = New("foo")
	if nil != blog || nil == err {
		t.Error("error should be thrown on site creation")
	}

	blog, err = New("foo", "bar", "buzz")
	if nil != blog || nil == err {
		t.Error("error should be thrown on site creation")
	}

	blog, err = New("foo", "bar", "buzz")
	if nil != blog || nil == err {
		t.Error("error should be thrown on site creation")
	}

	blog, err = New("foo", "bar", new(FileSystem), "buzz")
	if nil != blog || nil == err {
		t.Error("error should be thrown on site creation")
	}

	blog, err = New("foo", "bar", new(FileSystem), "buzz", new(RenderingStrategy))
	if nil != blog || nil == err {
		t.Error("error should be thrown on site creation")
	}

	blog, err = New("foo", "bar", new(FileSystem), "buzz", new(RenderingStrategy), dummy)
	if nil == blog || nil != err {
		t.Error("no error should be thrown on site creation")
	}

}
