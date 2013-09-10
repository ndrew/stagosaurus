package blog

import "testing"

func TestEngine(t *testing.T) {
	cfg := new(Config)
	err := cfg.ReadConfig("test_data/sample-config.json")
	if err != nil {
		t.Error(err)
	}

	renderingStrategy := new(RenderingStrategy)

	postsFactory := new(FolderPostFactory)
	postsFactory.PostsDir = "/Users/ndrw/Desktop/dev/site/blog/posts"

	posts := postsFactory.GetPosts()

	engine := New(cfg, renderingStrategy, posts)
	err = engine.Publish()
	if err != nil {
		t.Error(err)
	}

}
