package blogrenderer

import (
	"fmt"
	"io"

	blogposts "github.com/kikis202/learn-go-with-tests/learnGoWithTests/blog"
)

func Render(w io.Writer, post blogposts.Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>", post.Title)
	return err
}
