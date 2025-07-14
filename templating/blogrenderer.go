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

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, post blogposts.Post) error {

	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", post); err != nil {
		return err
	}

	return nil
}
