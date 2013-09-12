package blog

import "testing"

func assertNoError(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}

// rendering strategy that doesn't render anything
type DummyRenderingStrategy struct {
	posts *Post
	log   string
}

func (self *DummyRenderingStrategy) Render(post *Post) error {
	self.log += "<render>"
	return nil
}

func (self *DummyRenderingStrategy) RenderStarted() error {
	self.log += "<start>"
	return nil
}

func (self *DummyRenderingStrategy) RenderEnded() error {
	self.log += "<end>"
	return nil
}

func (self *DummyRenderingStrategy) GetPosts() []*Post {
	return []*Post{}
}

//
type DummyDeployer struct {
}

func (self *DummyDeployer) Deploy(posts []*Post) error {
	//
	return nil
}

func TestEngine(t *testing.T) {
	cfg := new(Config)
	err := cfg.ReadConfig("test_data/sample-config.json")
	assertNoError(err, t)

	postsFactory := new(FileSystem)
	postsFactory.PostsDir = "test_data/posts"
	posts, err := postsFactory.GetPosts()
	assertNoError(err, t)

	renderingStrategy := new(DummyRenderingStrategy)

	deployer := new(DummyDeployer)

	blog := New(cfg, postsFactory, renderingStrategy, deployer)
	err = blog.Publish()
	assertNoError(err, t)

	log := "<start>"
	for i := 0; i < len(posts); i++ {
		log += "<render>"
	}
	log += "<end>"

	if renderingStrategy.log != log {
		t.Errorf("renderering was done not as expected '%v' vs '%v' ", log, renderingStrategy.log)
	}
}
