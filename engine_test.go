package blog

import "testing"

func TestEngine(t *testing.T) {
	cfg := new(Config)
	cfg.BaseUrl = "http://localhost:666/blog/"

	renderingStrategy := new(RenderingStrategy)

	postsFactory := new(FolderPostFactory)
	postsFactory.PostsDir = "/Users/ndrw/Desktop/dev/site/blog/posts"

	posts := postsFactory.GetPosts()

	//t.Fail()

	engine := New(cfg, renderingStrategy, posts)
	engine.Publish()

	return
}
