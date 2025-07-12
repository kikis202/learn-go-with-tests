package blogposts

import (
	"bufio"
	"io"
	"strings"
)

const (
	titleSeperator       = "Title: "
	descriptionSeperator = "Description: "
	tagSeperator         = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titleSeperator)
	description := readMetaLine(descriptionSeperator)
	tags := strings.Split(readMetaLine(tagSeperator), ", ")

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
	}, nil
}
