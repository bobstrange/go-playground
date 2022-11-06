package main

import (
	"io"
	"os"
)

func writeExample() error {
	var fd os.File
	var p1, p2, p3 []byte
	var a, b, c, d, e, f int
	_, err := fd.Write(p1[a:b])
	if err != nil {
		return err
	}
	_, err = fd.Write(p2[c:d])
	if err != nil {
		return err
	}
	_, err = fd.Write(p3[e:f])
	if err != nil {
		return err
	}
	// ...
	return nil
}

type errorWriter struct {
	w   io.Writer
	err error
}

func (ew *errorWriter) write(buf []byte) {
	if ew.err != nil {
		return
	}
	_, ew.err = ew.w.Write(buf)
}

// io.Writer を Wrap した errorWriter を利用することで、下のようにエラー処理を簡潔に書ける
func writeExample2() error {
	var fd io.Writer
	var p1, p2, p3 []byte
	var a, b, c, d, e, f int
	ew := &errorWriter{w: fd}
	ew.write(p1[a:b])
	ew.write(p2[c:d])
	ew.write(p3[e:f])
	if ew.err != nil {
		return ew.err
	}
	return nil
}
