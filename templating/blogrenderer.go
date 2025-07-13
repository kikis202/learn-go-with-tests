package blogrenderer

import (
	"html/template"
	"io"

	blogposts "github.com/kikis202/learn-go-with-tests/learnGoWithTests/blog"
)

const (
	postTemplate = "<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>"
)

func Render(w io.Writer, post blogposts.Post) error {
	templ, err := template.New("blog").Parse(postTemplate)
	if err != nil {
		return err
	}

	if err := templ.Execute(w, post); err != nil {
		return err
	}

	return nil
}
