package blog

import (
	"fmt"
)

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
	New(string) (*Post, error)
}

/////////////////////////
// I M P L 

// Post iterator from filesystem
//
type FileSystem struct {
	PostsDir string
}

func (self FileSystem) New(data string, name string) *Post {
	post := new(Post)
	post.Content = data
	post.Name = name
	// TODO: init meta
	return post
}

func (self FileSystem) GetPosts() ([]*Post, error) {
	posts := []*Post{}
	var err error = nil

	callback := func(file *file) {
		if file.isMarkdown() {
			contents, err := file.Contents()
			if err != nil {
				fmt.Println(err)
				return
			}
			post := self.New(string(*contents), "Untitled")
			posts = append(posts, post)
		}
	}
	traverseFiles(self.PostsDir, callback)
	// get posts
	return posts, err
}
