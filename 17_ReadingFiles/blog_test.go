package blog_test

import (
	blog "blog"
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
H
I`
	)

	// fake in memory file system
	fs := fstest.MapFS{
		"hello_world.md":   {Data: []byte(firstBody)},
		"hello_world_2.md": {Data: []byte(secondBody)},
	}

	posts, err := blog.NewPostsFromFS(fs)

	got := posts[0]
	want := blog.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	}

	assertNoError(t, err)
	assertPostsLength(t, posts, fs)
	assertPost(t, got, want)
}

func assertPost(t *testing.T, got blog.Post, want blog.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func assertPostsLength(t *testing.T, posts []blog.Post, fs fstest.MapFS) {
	t.Helper()
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

}
