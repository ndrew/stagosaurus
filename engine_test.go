package blog

import "testing"

func TestEngine(t *testing.T) {
	cfg := new(Config)
	err := cfg.ReadConfig("test_data/sample-config.json")
	if err != nil {
		t.Error(err)
	}

	renderingStrategy := new(RenderingStrategy)

	postsFactory := new(FileSystem)
	postsFactory.PostsDir = "test_data/posts"

	posts, err := postsFactory.GetPosts()
	if err != nil {
		t.Error(err)
	}

	if len(posts) < 1 {
		t.Errorf("No test posts have been found in %s", postsFactory.PostsDir)
	}

	engine := New(cfg, renderingStrategy, posts)
	err = engine.Publish()
	if err != nil {
		t.Error(err)
	}

}
