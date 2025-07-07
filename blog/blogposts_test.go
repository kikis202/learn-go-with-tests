package blogposts_test

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"

	blogposts "github.com/kikis202/learn-go-with-tests/learnGoWithTests/blog"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I always fail")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("reading FS with files should return corresponding posts", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("hi")},
			"hello-world2.md": {Data: []byte("hola")},
		}

		posts, err := blogposts.NewPostsFromFs(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})

	t.Run("failing to open FS should return error", func(t *testing.T) {
		fs := StubFailingFS{}

		_, err := blogposts.NewPostsFromFs(fs)

		if err == nil {
			t.Error("Expected to get an error when file system fails")
		}
	})
}
