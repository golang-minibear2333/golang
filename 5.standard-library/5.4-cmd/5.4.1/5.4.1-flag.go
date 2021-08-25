package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	version = "1.0.0"
	showVersion = flag.Bool("version", false, "show version information")
	isDebug = flag.Bool("debug", false, "is debug")
	ip      = flag.String("ip", "127.0.0.1", "Input bind address")
	port    = flag.Int("port", 80, "Input bind port")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Println(version)
		os.Exit(0)
	}
	if *isDebug {
		fmt.Println("set log level: debug")
	}
	fmt.Println(fmt.Sprintf("bind address: %s:%d successfully",*ip,*port))
}
