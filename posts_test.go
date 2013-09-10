package blog

import "testing"

func TestPosts(t *testing.T) {

	postsFactory := new(FolderPostFactory)
	postsFactory.PostsDir = "test_data/posts"

	posts, err := postsFactory.GetPosts()
	if err != nil {
		t.Error(err)
	}

}
