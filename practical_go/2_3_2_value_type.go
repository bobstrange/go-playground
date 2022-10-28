package main

import "unicode/utf8"

type SKUCode string

func (c SKUCode) Valid() bool {
	// チェック処理 例: 9 桁
	if utf8.RuneCountInString(string(c)) != 9 {
		return false
	}
	return true
}

func (c SKUCode) ItemCD() string {
	return string(c[0:5])
}

func (c SKUCode) SizeCD() string {
	return string(c[5:7])
}

func (c SKUCode) ColorCD() string {
	return string(c[7:9])
}
