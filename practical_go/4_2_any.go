package main

var slices = []any{
	"関ヶ原",
	1600,
}

// Go 1.17 以前は any は使えない
var slicesOld = []interface{}{
	"大政奉還",
	1867,
}

var ieyasu = map[string]any{
	"名前":  "徳川家康",
	"生まれ": 1543,
}

var nobunaga = map[string]interface{}{
	"名前": "織田信長",
}
