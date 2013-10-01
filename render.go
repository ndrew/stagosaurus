package stagosaurus

type Renderer interface {
	Render(Config, []Post) ([]Post, error)
}
