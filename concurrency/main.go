package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wclaro123/stackpath/backend-developer-tests/concurrency/concurrency"
)

func main() {

	sp := concurrency.NewSimplePool(3)

	for i := 0; i < 100; i++ {
		sp.Submit(func() {
			fmt.Println(time.Now())
		})
	}

	ap, err := concurrency.NewAdvancedPool(2, 1)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {

		if i > 50 {
			err := ap.Close(context.Background())
			if err != nil {
				fmt.Println(fmt.Printf("close error: %s", err.Error()))
				continue
			}
		}

		err = ap.Submit(context.Background(), func(ctx context.Context) {
			time.Sleep(10 * time.Millisecond)
			fmt.Println(time.Now())
		})
		if err != nil {
			fmt.Println(fmt.Printf("submit error: %s", err.Error()))
		}
	}

	time.Sleep(5 * time.Second)

}
