package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)
var(
	showVersion = kingpin.Flag("version",
		"show version information",
	).Default("false").Bool()
	isDebug = kingpin.Flag("isDebug",
		"is debug mode(default: false)",
	).Default("false").Bool()
	ip = kingpin.Flag("ip",
		"Input bind address(default: 127.0.0.1)",
	).Default("127.0.0.1").String()
	port    = kingpin.Flag("port",
		"Input bind port(default: 80)",
		).Default("80").Int()
)
func main() {
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	if *showVersion{
		fmt.Println("1.0.0")
		os.Exit(0)
	}
	if *isDebug {
		fmt.Println("set log level: debug")
	}
	fmt.Println(fmt.Sprintf("bind address: %s:%d successfully",*ip,*port))

}

