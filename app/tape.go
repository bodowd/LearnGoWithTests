package main

import (
	"os"
)

type tape struct {
	file *os.File
}

// when we write, we go from the beginning
func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, 0)
	return t.file.Write(p)
}
