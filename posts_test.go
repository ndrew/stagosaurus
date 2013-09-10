package blog

import "testing"

func TestPostNew(t *testing.T) {
	postsFactory := new(FileSystem)
	post := postsFactory.New("testo")

	if post.Content != "testo" {
		t.Errorf("%s <> %s", post.Content, "testo")
	}
}

func TestPosts(t *testing.T) {

	postsFactory := new(FileSystem)
	postsFactory.PostsDir = "test_data/posts"

	posts, err := postsFactory.GetPosts()
	if err != nil {
		t.Error(err)
	}

	for _, p := range posts {
		println(p.Content)
	}

}
