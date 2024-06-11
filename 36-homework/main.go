package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctxTimeOut, cancel := context.WithTimeout(ctx, time.Second)

	defer cancel()

	go FuncContext(&ctxTimeOut, 7)

	select {
	case <-ctxTimeOut.Done():
		fmt.Println("Function finished!!!")
	}

}

func FuncContext(ctx *context.Context, i int) {
	n := 0
	for {
		n++
		fmt.Println("func is runnig ...")
	}
}
