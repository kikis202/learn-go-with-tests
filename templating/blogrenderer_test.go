package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	blogposts "github.com/kikis202/learn-go-with-tests/learnGoWithTests/blog"
	blogrenderer "github.com/kikis202/learn-go-with-tests/learnGoWithTests/templating"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogposts.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatalf("Didnt expect an error, but got %q", err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {

	var (
		aPost = blogposts.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		postRenderer.Render(io.Discard, aPost)
	}

}
