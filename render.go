package stagosaurus

type Renderer interface {
	Render([]Post) ([]Post, error)
}
