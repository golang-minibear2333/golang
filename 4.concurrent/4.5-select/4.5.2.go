package main

func main() {
	chanInt, done := make(chan int), make(chan struct{})
	defer close(chanInt)
	defer close(done)
	go func() {
		for {
			select {
			case <-chanInt:
			case <-done:
				return
			}
		}
	}()
	done <- struct{}{}
}
