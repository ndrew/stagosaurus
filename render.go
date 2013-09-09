package blog

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
	// Renderers []Renderer
	// indexTemplate *template.Template
	// postTemplate  *template.Template
}

func (self *RenderingStrategy) Render(post *Post) error {
	return nil
}

func (self *RenderingStrategy) RenderStarted() error {
	println("RenderStarted")
	return nil
}

func (self *RenderingStrategy) RenderEnded() error {
	println("RenderEnded")
	// flush changes here 
	return nil

}
