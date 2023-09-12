package pipelines

import (
	"context"
)

func ReverseSlice(ctx context.Context, data []string) []string {
	return dataCollector(ctx, reverse(ctx, data))
}

func dataCollector(ctx context.Context, data <-chan string) []string {
	var result []string

	for {
		select {
		case <-ctx.Done():
			return nil
		case value, ok := <-data:
			if !ok {
				return result
			}
			result = append(result, value)
		}
	}
}

func reverse(ctx context.Context, data []string) <-chan string {
	dataStream := make(chan string)
	go func() {
		defer close(dataStream)

		for i := len(data) - 1; i >= 0; i-- {
			select {
			case <-ctx.Done():
				return
			case dataStream <- data[i]:
			}
		}
	}()

	return dataStream
}
