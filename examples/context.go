package examples

import (
	"context"
	"fmt"
	"os"
)

func ContextExample() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	c := favContextKey("color")
	ctx := context.WithValue(context.Background(), k, "Go")
	ctx = context.WithValue(ctx, c, "Green")

	f(ctx, k)
	f(ctx, c)

	fmt.Println(os.UserHomeDir())
}
