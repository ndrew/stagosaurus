package stagosaurus

import (
	"bytes"
	"sort"
	"text/template"
)

// 'Renderer' interface. Renderer in wide sence: markdown posts -> render -> html files.
//
type Renderer interface {
	Render(post *Post) error
	RenderStarted() error
	RenderEnded() error

	GetPosts() []*Post
}

// 'Composite' rendering strategy, renders each ready post + table of contents(i.e. index.html) and rss stream
//
type RenderingStrategy struct {
	cfg *Config
	// Renderers []Renderer
	indexTemplate *template.Template
	postTemplate  *template.Template
	// results
	IndexPage *Post
	Posts     []*Post
}

////////////////////////////
// sort.Interface impl

func (s *RenderingStrategy) Len() int {
	return len(s.Posts)
}

func (s *RenderingStrategy) Swap(i, j int) {
	s.Posts[i], s.Posts[j] = s.Posts[j], s.Posts[i]
}

func (s *RenderingStrategy) Less(i, j int) bool {
	return s.sortMethod(i, j)
}

// custom sortMethod
//
func (s *RenderingStrategy) sortMethod(i, j int) bool {
	// todo: add nil check
	return s.Posts[i].Meta.Date.After(s.Posts[j].Meta.Date)
}

////////////////////////////
// Renderer intefrace impl

//
//
func (self *RenderingStrategy) Render(post *Post) error {
	var html bytes.Buffer

	self.postTemplate.Execute(&html, post)

	// store meta for further sorting
	newPost := new(Post)
	newPost.Meta = post.Meta
	newPost.Content = html.String()

	self.Posts = append(self.Posts, newPost)
	return nil
}

//
//
func (self *RenderingStrategy) RenderStarted() error {
	self.Posts = []*Post{}
	return nil
}

//
//
func (self *RenderingStrategy) RenderEnded() error {
	// sort posts by date
	sort.Sort(self)

	// make index
	//self.Index = "test"
	return nil
}

//
//
func (self *RenderingStrategy) GetPosts() []*Post {
	return self.Posts
}
