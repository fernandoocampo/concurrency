package tees

import (
	"context"

	"github.com/fernandoocampo/concurrency/ors"
)

func Tee(ctx context.Context, input <-chan any) (<-chan any, <-chan any) {
	outStream1 := make(chan any)
	outStream2 := make(chan any)

	go func() {
		defer close(outStream1)
		defer close(outStream2)

		for val := range ors.OrDone(ctx.Done(), input) {
			out1, out2 := outStream1, outStream2
			for i := 0; i < 2; i++ {
				select {
				case <-ctx.Done():
				case out1 <- val:
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}

	}()

	return outStream1, outStream2
}
