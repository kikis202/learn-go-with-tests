package blogrenderer

import (
	"embed"
	"html/template"
	"io"

	blogposts "github.com/kikis202/learn-go-with-tests/learnGoWithTests/blog"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, post blogposts.Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.ExecuteTemplate(w, "blog.gohtml", post); err != nil {
		return err
	}

	return nil
}
