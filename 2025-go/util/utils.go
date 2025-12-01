package util

import (
	"bufio"
	"bytes"
	"io"
	"iter"
	"os"
)

func Lines(reader io.Reader) iter.Seq2[string, error] {
	s := bufio.NewScanner(reader)
	return func(yield func(string, error) bool) {
		for s.Scan() {
			if !yield(s.Text(), nil) {
				return
			}
		}
		if err := s.Err(); err != nil {
			yield("", err)
		}
	}
}

func MustReadFile(fileName string) io.Reader {
	content := Must(os.ReadFile(fileName))
	return bytes.NewReader(content)
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err.Error())
	}
	return v
}
