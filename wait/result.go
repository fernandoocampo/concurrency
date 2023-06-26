package wait

func Sum(a, b int) <-chan int {
	sumStream := make(chan int)
	go func() {
		defer close(sumStream)
		sumStream <- a + b
	}()
	return sumStream
}
