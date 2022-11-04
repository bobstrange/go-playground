package main

import (
	"bufio"
	"io"
	"os"
	"strings"

	"golang.org/x/text/unicode/norm"
)

// ファイルから読み込んで、全角カタカナやスペースを全て正規化する

// ファイル名や、ファイル自体ではなく、io.Writer, io.Reader の interface に依存した I/F を作る
func Normalize(w io.Writer, r io.Reader) error {
	br := bufio.NewReader(r)
	for {
		s, err := br.ReadString('\n')
		if s != "" {
			io.WriteString(w, norm.NFKC.String(s))
		}
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
	}
}

// ファイルを入出力にする場合の Wrapper
func NormalizeFile(input, output string) error {
	r, err := os.Open(input)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.Create(output)
	if err != nil {
		return err
	}
	defer w.Close()

	return Normalize(w, r)
}

// 文字列を入出力にする場合の Wrapper
func NormalizeString(i string) (string, error) {
	r := strings.NewReader(i)
	var w strings.Builder
	err := Normalize(w, r)
	if err != nil {
		return "", err
	}
	return w.String(), nil
}
