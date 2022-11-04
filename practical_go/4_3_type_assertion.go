package main

import (
	"context"
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
}
