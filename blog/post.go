package blogposts

import (
	"bufio"
	"io"
	"strings"
)

const (
	titleSeperator       = "Title: "
	descriptionSeperator = "Description: "
)

type Post struct {
	Title       string
	Description string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titleSeperator)
	description := readMetaLine(descriptionSeperator)

	return Post{
		Title:       title,
		Description: description,
	}, nil
}
