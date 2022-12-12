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
	// fake in memory file system
	fs := fstest.MapFS{
		"hello_world.md":   {Data: []byte("Title: Post 1")},
		"hello_world_2.md": {Data: []byte("Title: Post 2")},
	}

	posts, err := blog.NewPostsFromFS(fs)

	got := posts[0]
	want := blog.Post{Title: "Post 1"}

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
