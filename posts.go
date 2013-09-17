package stagosaurus

// Blog entry
//
type Post struct {
	Content string
	Meta    *Meta
	Name    string
}

// Post retrival
//
type PostFactory interface {
	GetPosts() ([]*Post, error)
	New(string, string) (*Post, error)
}

/////////////////////////
// I M P L

// Post iterator from filesystem
//
type FileSystem struct {
	PostsDir string
}

func (self FileSystem) New(data string, name string) (*Post, error) {
	post := new(Post)
	post.Content = data
	post.Name = name
	// TODO: init meta
	return post, nil
}

func (self FileSystem) GetPosts() (posts []*Post, err error) {

	callback := func(file *file) {
		if file.isMarkdown() {
			contents, err := file.Contents()
			if err != nil {
				return
			}
			// TODO:
			post, err := self.New(string(*contents), "Untitled")
			posts = append(posts, post)
		}
	}
	traverseFiles(self.PostsDir, callback)
	return
}
