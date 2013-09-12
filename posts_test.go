package blog

import "testing"

// convenience function for testing
/*func getTestPosts(t *testing.T) []*Post {
	postsFactory := new(FileSystem)
	postsFactory.PostsDir = "test_data/posts"

	posts, err := postsFactory.GetPosts()
	assertNoError(err, t)
	if len(posts) < 1 {
		t.Errorf("No test posts have been found in %s", postsFactory.PostsDir)
	}
	return posts
} */

// todo: instead of real test system use mock

func TestPostNew(t *testing.T) {
	postsFactory := new(FileSystem)
	post, err := postsFactory.New("testo", "Untitled")
	assertNoError(err, t)

	if post.Content != "testo" {
		t.Errorf("%s <> %s", post.Content, "testo")
	}
	if post.Name != "Untitled" {
		t.Errorf("%s <> %s", post.Name, "Untitled")
	}

}

func TestPosts(t *testing.T) {
	postsFactory := new(FileSystem)
	posts, err := postsFactory.GetPosts()
	assertNoError(err, t)

	for _, p := range posts {
		if len(p.Content) == 0 {
			t.Error("test post file shouldn't be empty")
		}
	}

	_, err = postsFactory.New("testo", "Untitled")
	assertNoError(err, t)
}
