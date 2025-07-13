package blogrenderer

import (
	"fmt"
	"io"

	blogposts "github.com/kikis202/learn-go-with-tests/learnGoWithTests/blog"
)

func Render(w io.Writer, post blogposts.Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>\n<p>%s</p>\n", post.Title, post.Description)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(w, "Tags: <ul>")
	for _, tag := range post.Tags {
		fmt.Fprintf(w, "<li>%s</li>", tag)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprint(w, "</ul>")
	if err != nil {
		return err
	}

	return nil
}
