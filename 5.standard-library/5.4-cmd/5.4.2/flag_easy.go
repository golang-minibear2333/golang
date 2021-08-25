package main

import (
	"flag"
	"fmt"
)

var (
	ip          string
	port        int
)

func init() {
	flag.StringVar(&ip, "ip", "127.0.0.1", "Input bind address(default: 127.0.0.1)")
	flag.IntVar(&port, "port", 80, "Input bind port(default: 80)")
}
func main() {
	flag.Parse()
	fmt.Println(fmt.Sprintf("bind address: %s:%d successfully", ip, port))
}
