package main

import (
	"context"
	"log"
)

type TraceID string

const ZeroTraceID = ""

type traceIDKey struct{} // プリミティブな値はキーが衝突する可能性があるので、空の構造体をキーにする

func SetTraceID(ctx context.Context, tid TraceID) context.Context {
	return context.WithValue(ctx, traceIDKey{}, tid)
}

func GetTraceID(ctx context.Context) TraceID {
	tid, ok := ctx.Value(traceIDKey{}).(TraceID)
	if !ok {
		return ZeroTraceID
	}
	return tid
}

func main() {
	ctx := context.Background()
	log.Printf("trace id = %q\n", GetTraceID(ctx))
	// 2022/12/28 00:29:21 trace id = ""
	ctx = SetTraceID(ctx, "hogehoge")
	log.Printf("trace id = %q\n", GetTraceID(ctx))
	// 2022/12/28 00:29:21 trace id = "hogehoge"
}
