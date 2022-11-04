package main

import (
	"context"
	"io"
	"log"
)

func main() {
	ctx := context.WithValue(context.Background(), "favorite", "ra-men")

	// ここで ctx.Value("favorite") は string ではなく interface{} なので type assertion が必要
	if s, ok := ctx.Value("favorite").(string); ok {
		log.Printf("My favorite food is %s.\n", s)
	}

	// あるいは型スイッチを使う
	switch v := ctx.Value("favorite").(type) {
	case string:
		log.Printf("I like %s.\n", v)
	case int:
		log.Printf("My favorite number is %d.\n", v)
	default:
		log.Printf("I prefer %v.\n", v)
	}

	var r any
	// interface を満たすかどうかの assertion
	if c, ok := r.(io.Closer); ok {
		c.Close()
	}

	type Fish any
	var fishList = []Fish{"サバ", "ブリ", "マグロ"}
	// ↓ の様に一括で downcast はできない
	// var fishNameList = fishList.([]string)

	// any への一括 upcast もできない
	// var anyList []any = fishList

	fishNames := make([]string, len(fishList))
	for i, f := range fishList {
		// それぞれを型チェックして downcast する必要がある
		if fn, ok := f.(string); ok {
			fishNames[i] = fn
		}
	}

	anyList := make([]any, len(fishList))
	for i, fn := range fishList {
		// upcast は型チェック不要
		anyList[i] = fn
	}

}
