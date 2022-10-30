package main

import "sync"

type BigStruct struct {
	Member string
}

var pool = &sync.Pool{
	// New フィールドに、使用したい構造体を初期化する処理の関数を設定する
	New: func() interface{} {
		return &BigStruct{}
	},
}

// ファクトリの内部で sync.Pool を使う
func NewBigStruct() *BigStruct {
	b := pool.Get().(*BigStruct)
	return b
}

func (b *BigStruct) Release() {
	b.Member = ""
	pool.Put(b)
}

func main() {
	// BigStruct が初期化済みならそれを、されていない場合は New() を呼び出す
	b := pool.Get().(*BigStruct)
	pool.Put(b)
}
