package heartbeats

import "time"

type Pulse string

const APulse Pulse = "lup-dum-dum-lup"

func DoBasicProcess(done <-chan any, input []string, pulseInterval time.Duration) ([]string, []Pulse) {
	resultStream, heartbeat := doBasicWork(done, Generate(done, input), pulseInterval)
	var result []string
	var pulses []Pulse
	for {
		select {
		case <-done:
			return result, pulses
		case aPulse, ok := <-heartbeat:
			if !ok {
				return result, pulses
			}
			pulses = append(pulses, aPulse)
		case r, ok := <-resultStream:
			if !ok {
				return result, pulses
			}
			result = append(result, r)
		}
	}
}

func doBasicWork(done <-chan any, data <-chan string, pulseInterval time.Duration) (<-chan string, <-chan Pulse) {
	heartbeat := make(chan Pulse)
	results := make(chan string)
	go func() {
		defer close(heartbeat)
		defer close(results)

		pulse := time.NewTicker(pulseInterval)

		sendPulse := func(done <-chan any) {
			select {
			case <-done:
				return
			case heartbeat <- APulse:
			default:
			}
		}

		sendResult := func(r string) {
			for {
				select {
				case <-done:
					return
				case <-pulse.C:
					sendPulse(done)
				case results <- r:
					return
				}
			}
		}

		for {
			select {
			case <-done:
				return
			case <-pulse.C:
				sendPulse(done)
			case r, ok := <-data:
				if !ok {
					return
				}
				time.Sleep(5 * time.Millisecond)
				sendResult(r)
			}
		}
	}()
	return results, heartbeat
}

func Generate(done <-chan any, input []string) <-chan string {
	dataStream := make(chan string)
	go func() {
		defer close(dataStream)
		for _, v := range input {
			select {
			case <-done:
				return
			case dataStream <- v:
			}
		}
	}()
	return dataStream
}
