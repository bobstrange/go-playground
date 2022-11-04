package main

import (
	"errors"
	"fmt"
)

// エラーを変数に登録して公開しておく
// 1. errors.New()
var ErrNotFound = errors.New("not found")

type Book any

func findBook(isbn string) (*Book, error) {
	// ...
	return nil, ErrNotFound
}

func validate(length int) error {
	if length <= 0 {
		// 2. fmt.Errorf()
		return fmt.Errorf("format strinlength must be greater than 0, length = %d", length)
	}
	return nil
}

// 3. 独自エラー
type HTTPError struct {
	StatusCode int
	URL        string
}

// Error() string を実装すれば Error interface を満たす
func (e *HTTPError) Error() string {
	return fmt.Sprintf("http status code = %d, url = %s", e.StatusCode, e.URL)
}
