package pipelines

import (
	"context"
)

func ProductData(ctx context.Context, intStream <-chan int) <-chan int {
	productStream := make(chan int)
	go func() {
		defer close(productStream)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-intStream:
				if !ok {
					return
				}
				productStream <- v * v
			}
		}
	}()
	return productStream
}

func SumData(ctx context.Context, intStream <-chan int) <-chan int {
	sumStream := make(chan int)
	go func() {
		defer close(sumStream)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-intStream:
				if !ok {
					return
				}
				sumStream <- v + v
			}
		}
	}()
	return sumStream
}

func IntGenerator(ctx context.Context, data []int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for _, v := range data {
			select {
			case <-ctx.Done():
				return
			case intStream <- v:
			}
		}
	}()
	return intStream
}

func IntCollector(ctx context.Context, intStream <-chan int) []int {
	var result []int

	for {
		select {
		case <-ctx.Done():
			return nil
		case value, ok := <-intStream:
			if !ok {
				return result
			}
			result = append(result, value)
		}
	}
}

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
