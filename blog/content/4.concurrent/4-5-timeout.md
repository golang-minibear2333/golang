# 4.4 定时器

## **...本节正在编写，未完待续，催更请留言，我会收到邮件**

## 超时关闭


[完整代码](timeout.go)

```go
package main

import "time"

func main() {
	t := make(chan bool)
	ch := make(chan int)
	defer func() {
		close(ch)
		close(t)
	}()
	go func() {
		time.Sleep(1e9) //等待1秒
		t <- true
	}()
	go func() {
		time.Sleep(time.Second * 2)
		ch <- 123
	}()
	select {
	case <-ch: //从ch中读取数据

	case <-t: //如果1秒后没有从ch中读取到数据，那么从t中读取，并进行下一步操作
	}
}
```

# 可热更新的定时器

废话不多说，直接上代码
```go
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

```

...未完待续，催更请留言，我会收到邮件
