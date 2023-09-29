package queuing

import (
	"context"
	"time"
)

func Sleep(ctx context.Context, timeout time.Duration, values <-chan any) <-chan any {
	stream := make(chan any)
	go func() {
		defer close(stream)
		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-values:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					continue
				case <-time.After(timeout):
					select {
					case <-ctx.Done():
						continue
					case stream <- value:
					}
				}
			}
		}
	}()
	return stream
}

func Buffer(ctx context.Context, buffer int, values <-chan any) <-chan any {
	stream := make(chan any, buffer)
	go func() {
		defer close(stream)
		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-values:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case stream <- value.(int):
				}
			}
		}
	}()
	return stream
}
