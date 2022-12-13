package blog

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func(tagName string) string {
		// Scan() will read a line
		scanner.Scan()
		// then we extract the data using Text()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	titleLine := readLine(titleSeparator)
	descriptionLine := readLine(descriptionSeparator)
	tagLine := strings.Split(readLine(tagsSeparator), ", ")

	body := readBody(scanner)

	return Post{
		Title:       titleLine,
		Description: descriptionLine,
		Tags:        tagLine,
		Body:        body,
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore a line

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
