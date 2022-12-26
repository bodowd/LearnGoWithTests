package main

import "io"

type tape struct {
	file io.ReadWriteSeeker
}

// when we write, we go from the beginning
func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Seek(0, 0)
	return t.file.Write(p)
}
