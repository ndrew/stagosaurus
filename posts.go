package blog

import (
	"fmt"
)

// Blog entry
//
type Post struct {
	Content string
	Meta    Meta
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
type FolderPostFactory struct {
	PostsDir string
}

func (self FolderPostFactory) New(data string) *Post {
	post := new(Post)
	// TODO: init meta
	return post
}

func (self FolderPostFactory) GetPosts() ([]*Post, error) {
	posts := []*Post{}
	var err error = nil

	callback := func(file *file) {
		if file.isMarkdown() {
			contents, err := file.Contents()
			if err != nil {
				fmt.Println(err)
				return
			}
			post := self.New(string(*contents))
			posts = append(posts, post)
		}
	}
	traverseFiles(self.PostsDir, callback)
	// get posts
	return posts, err
}
