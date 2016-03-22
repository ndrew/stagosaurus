package stagosaurus

// Renderer — describes rendering interface
//
type Renderer interface {
	Render(Config, []Post) ([]Post, error)
}
