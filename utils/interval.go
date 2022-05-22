package utils

import "time"

func SetInterval(fn func(), delay time.Duration) chan bool {
	stop := make(chan bool)
	go func() {
		for {
			select {
				case <- time.After(delay):
					fn()
				case <- stop:
					return
			}
		}
	}()
	return stop
}

func ClearInterval(stop chan bool) {
	stop <- true
}
