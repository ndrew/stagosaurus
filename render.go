package blog

import (
	"bytes"
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
}

func (self *RenderingStrategy) Render(post *Post) error {
	var html bytes.Buffer
	self.postTemplate.Execute(&html, post)

	self.posts = append(self.posts, html.String())
	return nil
}

func (self *RenderingStrategy) RenderStarted() error {
	self.posts = []string{}
	return nil
}

func (self *RenderingStrategy) RenderEnded() error {
	// sort posts by date	

	// make index 
	self.index = "test"
	return nil

}
