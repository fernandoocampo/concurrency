package wait

import (
	"context"
	"fmt"
	"math"
)

func PowerOfThree(ctx context.Context, valueStream <-chan int) (int, error) {
	var base float64
	select {
	case <-ctx.Done():
		return 0, fmt.Errorf("getting value to calculate the power of three: %w", ctx.Err())
	case value := <-valueStream:
		base = float64(value)
	}

	result := math.Pow(base, 3)

	return int(result), nil
}
