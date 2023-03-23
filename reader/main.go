package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var _ io.ReadCloser = mockReadCloser{}

type mockReadCloser struct {
	readFn  func(p []byte) (n int, err error)
	closeFn func() error
}

// Read implements io.ReadCloser
func (m mockReadCloser) Read(p []byte) (n int, err error) {
	return m.readFn(p)
}

// Close implements io.ReadCloser
func (m mockReadCloser) Close() error {
	return m.closeFn()
}

func main() {
	err := run()
	if err != nil {
		os.Exit(1)
	}
}

func run() error {
	data := bytes.NewBuffer(bytes.Repeat([]byte(`0`), 512*2+1)).Bytes()
	dataIndex := 0 // index of data copied on read
	readCloser := mockReadCloser{
		readFn: func(p []byte) (n int, err error) {
			pEnd := cap(p)
			if pEnd > len(data)-dataIndex {
				pEnd = len(data) - dataIndex
			}

			for n < pEnd {
				p[n] = data[dataIndex]
				n++
				dataIndex++
			}

			if dataIndex >= len(data) {
				return n, io.EOF
			}

			return n, nil
		},
		closeFn: func() error {
			dataIndex = 0
			return nil
		},
	}

	readData, err := io.ReadAll(readCloser)
	if err != nil {
		return err
	}

	if !bytes.Equal(data, readData) {
		return fmt.Errorf("data not read correctly")
	}

	return nil
}
