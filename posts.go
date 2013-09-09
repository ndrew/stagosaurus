package blog

// Blog entry
//
type Post struct {
	Content string
	Meta    Meta
}

// Post retrival
//
type PostFactory interface {
	GetPosts() []*Post
}

/////////////////////////
// I M P L 

// Post iterator from filesystem
//
type FolderPostFactory struct {
	PostsDir string
}

func (self FolderPostFactory) GetPosts() []*Post {
	posts := []*Post{}

	callback := func(file *file) {
		if file.isMarkdown() {
			post := new(Post)
			post.Content = file.Name()
			posts = append(posts, post)
		}
	}
	traverseFiles(self.PostsDir, callback)

	// get posts
	return posts
}
