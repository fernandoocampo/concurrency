package confinement

import "strconv"

func GenerateData() <-chan string {
	dataStream := make(chan string)
	go func() {
		defer close(dataStream)
		for i := 0; i < 11; i++ {
			dataStream <- strconv.Itoa(i)
		}
	}()
	return dataStream
}
