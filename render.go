package blog

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
}

// 'Composite' rendering strategy, renders each ready post + table of contents(i.e. index.html) and rss stream  
//
type RenderingStrategy struct {
	cfg *Config
	// Renderers []Renderer
	indexTemplate *template.Template
	postTemplate  *template.Template
	// results
	index string
	posts []string

	_meta []*Meta // for sorting

}

////////////////////////////
// sort.Interface impl

func (s *RenderingStrategy) Len() int {
	return len(s.posts)
}

func (s *RenderingStrategy) Swap(i, j int) {
	s.posts[i], s.posts[j] = s.posts[j], s.posts[i]
}

func (s *RenderingStrategy) Less(i, j int) bool {
	return s.sortMethod(i, j)
}

// custom sortMethod
//
func (s *RenderingStrategy) sortMethod(i, j int) bool {
	// todo: add nil check
	return s._meta[i].Date.Before(s._meta[j].Date)
}

////////////////////////////
// Renderer intefrace impl

//
//
func (self *RenderingStrategy) Render(post *Post) error {
	var html bytes.Buffer

	self.postTemplate.Execute(&html, post)

	// store meta for further sorting
	self._meta = append(self._meta, post.Meta)
	self.posts = append(self.posts, html.String())
	return nil
}

//
//
func (self *RenderingStrategy) RenderStarted() error {
	self.posts = []string{}
	return nil
}

//
//
func (self *RenderingStrategy) RenderEnded() error {
	// sort posts by date	
	sort.Sort(self)

	// make index 
	self.index = "test"
	return nil

}
