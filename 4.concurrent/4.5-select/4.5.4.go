package main

import (
	"fmt"
	"time"
)

type Worker struct {
	stream  <-chan int
	timeout time.Duration
	done    chan struct{}
}

func NewWorker(stream <-chan int, timeout int) *Worker {
	return &Worker{
		stream:  stream,
		timeout: time.Duration(timeout) * time.Second,
		done:    make(chan struct{}),
	}
}
func (w *Worker) Start() {
	w.afterTimeStop()
	for {
		select {
		case data, ok := <-w.stream:
			if !ok {
				return
			}
			fmt.Println(data)
		case <-w.done:
			close(w.done)
			return
		}
	}
}
func (w *Worker) afterTimeStop() {
	go func() {
		time.Sleep(w.timeout)
		w.done <- struct{}{}
	}()
}

func main() {
	stream := make(chan int)
	defer close(stream)

	w := NewWorker(stream, 3)
	w.Start()
}
