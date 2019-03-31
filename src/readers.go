package main

import (
	"io"
)

type MyReader struct{}

func (t MyReader) Read(b []byte) (int, error) {
	for key := range b {
		b[key] = 'A'
	}
	return len(b), nil
}

type rot13Reader struct {
	r io.Reader
}

func (t rot13Reader) Read(b []byte) (int, error) {
	len, error := t.r.Read(b)
	for key, value := range b {
		if key <= len {
			b[key] = root13(value)
		}
	}
	return len, error
}

func root13(b byte) byte {
	const delta = 13
	switch {
	case b >= 'A' && b < 'A'+delta:
		return b + delta
	case b >= 'A'+delta && b < 'Z':
		return b - delta
	case b >= 'a' && b < 'a'+delta:
		return b + delta
	case b >= 'a'+delta && b < 'z':
		return b - delta
	default:
		return b
	}
}
