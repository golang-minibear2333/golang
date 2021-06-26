package main

import (
	"fmt"
	"time"
)

type Server struct {
	tk     *time.Ticker
	reset  chan struct{}
	Close  chan struct{}
	Period int64
}


func main() {
	s := CreateServer(1)
	go s.Start()
	time.Sleep(time.Duration(10) * time.Second)
	s.Update(3)
	time.Sleep(time.Duration(10) * time.Second)
	s.Stop()
	fmt.Println("good bye")
}

func CreateServer(Period int64) *Server {
	return &Server{
		tk:     nil,
		reset:  make(chan struct{}),
		Close:  make(chan struct{}),
		Period: Period,
	}
}

// 程序启动
func (s *Server) Start() {
	// 定时
	s.tk = time.NewTicker(time.Duration(s.Period) * time.Second)
	defer s.tk.Stop()
	for {
		select {
		case <-s.Close:
			return
		case <-s.tk.C:
			fmt.Println("定时唤醒:", time.Now().Format("2006-01-02 15:04:05"))
		case <-s.reset:
			s.tk.Stop()
			s.tk = time.NewTicker(time.Duration(s.Period) * time.Second)
		}
	}
}

func (s *Server) Stop() {
	close(s.Close)
	close(s.reset)
}

func (s *Server) Update(p int64) {
	s.Period = p
	s.reset <- struct{}{}
}
