package main

/*
	https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/
*/

import (
	"time"
	"fmt"
	"context"
)

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <- ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <- time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	//go handle(ctx, 500 * time.Millisecond)
	go handle(ctx, 1500000 * time.Millisecond)

	select {
	case <- ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}