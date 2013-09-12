package blog

import (
	"testing"
	"text/template"
)

func TestDefaultRenderingStrategy(t *testing.T) {
	renderer := new(RenderingStrategy)

	var err error
	renderer.indexTemplate, err = template.ParseFiles("test_data/templates/index.template")
	assertNoError(err, t)

	renderer.postTemplate, err = template.ParseFiles("test_data/templates/post.template")
	assertNoError(err, t)

	err = renderer.RenderStarted()
	assertNoError(err, t)

	post := new(Post)
	post.Content = "I am content"
	post.Name = "Dummy"

	//post.Meta := new(Meta)
	//post.Meta.Ready = true

	renderer.Render(post)

	renderer.RenderEnded()

	// todo: template loading

	//t.Error("todo: template loading")
}
