// One of the most common uses for contexts is to share data, or use request
// scoped values. When you have multiple functions and you want to share data between them,
// you can do so using contexts

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = addValue(ctx)
	readValue(ctx)
}

func addValue(ctx context.Context) context.Context {
	return context.WithValue(ctx, "key", "test-value")
}

func readValue(ctx context.Context) {
	val := ctx.Value("key")
	fmt.Println(val)
}
