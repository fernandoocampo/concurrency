package drops

import (
	"context"
	"log"
	"strconv"
)

func DropSomething(ctx context.Context, work, capacity int) <-chan string {
	result := make(chan string, capacity)

	go func() {
		defer close(result)
		for i := 0; i < work; i++ {
			select {
			case <-ctx.Done():
				log.Println("context was cancelled:", ctx.Err())
				return
			case result <- strconv.Itoa(i):
			default: // key element of the drop pattern
				log.Println("dropped data:", i)
			}
		}
	}()

	return result
}
