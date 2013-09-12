package stagosaurus

import (
	"fmt"
	"testing"
	"text/template"
	"time"
)

func TestDefaultRenderingStrategy(t *testing.T) {
	renderer := new(RenderingStrategy)

	renderer.cfg = new(Config)
	renderer.cfg.BaseUrl = "http://localhost"
	renderer.cfg.Port = ":8080"

	var err error

	// more about go template syntax - http://golang.org/pkg/text/template/

	renderer.indexTemplate, err = template.New("index").Parse("Foo")
	//template.ParseFiles("test_data/templates/index.template")
	assertNoError(err, t)

	renderer.postTemplate, err = template.New("post").Parse("{{.cfg.BaseUrl}}\n{{.Name}}:{{.Content}}")
	// template.ParseFiles("test_data/templates/post.template")
	assertNoError(err, t)

	err = renderer.RenderStarted()
	assertNoError(err, t)

	var post *Post
	for i := 1; i < 5; i++ {
		post = new(Post)
		post.Content = fmt.Sprintf("Content %v", i)
		post.Name = fmt.Sprintf("Post#%v", i)

		post.Meta = new(Meta)
		post.Meta.Ready = true
		post.Meta.Date = time.Now().Add(time.Duration(i*60) * time.Second)

		renderer.Render(post)
	}

	renderer.RenderEnded()

	// renderer.Posts
	// t.Error(renderer.Posts)
}
