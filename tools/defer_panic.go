package tools

import (
	"fmt"
	"runtime"
)

func main() {
	defer func() {
		if reason := recover(); reason != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			ss := runtime.Stack(buf, false)
			if ss > size {
				ss = size
			}
			buf = buf[:ss]
			fmt.Println("doBatchReport panic", reason, "detail:", buf)
		}
	}()
}
