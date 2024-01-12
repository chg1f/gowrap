package gowrap

import "io"

type writer struct {
	write func([]byte) (int, error)
}

func NewWriter(write func([]byte) (int, error)) io.Writer {
	return &writer{write: write}
}
func (x *writer) Write(p []byte) (n int, err error) {
	return x.write(p)
}

var _ io.Writer = (*writer)(nil)
