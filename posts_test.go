package blog

import "testing"

// convenience function for testing
func getTestPosts(t *testing.T) []*Post {
	postsFactory := new(FileSystem)
	postsFactory.PostsDir = "test_data/posts"

	posts, err := postsFactory.GetPosts()
	if err != nil {
		t.Error(err)
	}
	if len(posts) < 1 {
		t.Errorf("No test posts have been found in %s", postsFactory.PostsDir)
	}
	return posts
}

func TestPostNew(t *testing.T) {
	postsFactory := new(FileSystem)
	post := postsFactory.New("testo")

	if post.Content != "testo" {
		t.Errorf("%s <> %s", post.Content, "testo")
	}
}

func TestPosts(t *testing.T) {

	for _, p := range getTestPosts(t) {
		if len(p.Content) == 0 {
			t.Error("test post file shouldn't be empty")
		}
	}

}
