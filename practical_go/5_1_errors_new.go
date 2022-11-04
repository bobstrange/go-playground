package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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

// 4. エラーのラップ, アンラップ
type loadConfigError struct {
	msg string
	err error
}

func (e *loadConfigError) Error() string {
	return fmt.Sprintf("cannot load config: %s (%s)", e.msg, e.err.Error())
}

// Unwrap を実装しておくのがお作法
func (e *loadConfigError) Unwrap() error {
	return e.err
}

type Config any

func LoadConfig(configFilePath string) (*Config, error) {
	var cfg *Config
	data, err := os.ReadFile(configFilePath)
	// ファイルが存在しないエラーや、ファイルの読み込みに失敗したエラーを抽象度が高い loadConfigError として返す
	if err != nil {
		return nil, &loadConfigError{msg: fmt.Sprintf("read file `%s`", configFilePath), err: err}
	}
	if err = json.Unmarshal(data, &cfg); err != nil {
		return nil, &loadConfigError{msg: fmt.Sprintf("parse config file `%s`", configFilePath), err: err}
	}
	return cfg, nil
}
