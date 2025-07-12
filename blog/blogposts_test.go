package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
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
		const (
			firstBody = `Title: Post1
Description: Description1
Tags: tdd, go`
			secondBody = `Title: Post2
Description: Description2
Tags: tag1, tag2`
		)

		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFs(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}

		want := []blogposts.Post{
			{
				Title:       "Post1",
				Description: "Description1",
				Tags:        []string{"tdd", "go"},
			},
			{
				Title:       "Post2",
				Description: "Description2",
				Tags:        []string{"tag1", "tag2"},
			},
		}

		assertAllPosts(t, posts, want)
	})

	t.Run("failing to open FS should return error", func(t *testing.T) {
		fs := StubFailingFS{}

		_, err := blogposts.NewPostsFromFs(fs)

		if err == nil {
			t.Error("Expected to get an error when file system fails")
		}
	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertContainsPost(t *testing.T, got []blogposts.Post, want blogposts.Post) {
	t.Helper()
	for i := range got {
		if reflect.DeepEqual(got[i], want) {
			return
		}
	}
	t.Errorf("result %+v didn't contain this %+v post", got, want)
}

func assertAllPosts(t *testing.T, got, want []blogposts.Post) {
	t.Helper()
	for i := range want {
		assertContainsPost(t, got, want[i])
	}
}
