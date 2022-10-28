package main

var i int

type ErrorCode int

var e ErrorCode

func main() {
	// ErrorCode を int にキャスト
	i = int(e)
	// int を ErrorCode にキャスト
	e = ErrorCode(i)
}
