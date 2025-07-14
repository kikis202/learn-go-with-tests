package blogrenderer_test

import (
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	blogposts "github.com/kikis202/learn-go-with-tests/learnGoWithTests/blog"
	blogrenderer "github.com/kikis202/learn-go-with-tests/learnGoWithTests/templating"
)

func TestRender(t *testing.T) {
	aPost := blogposts.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatalf("Didnt expect an error, but got %q", err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
