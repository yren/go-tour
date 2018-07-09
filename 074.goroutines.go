/*
    goroutine 是 Go 运行时管理的轻量级线程。
    go f(x, y, z)
    
    会启动一个新的 goroutine 并执行
    f(x, y, z)
    
    f, x, y, z 的求值发生在当前 Go 程中，而 f 的执行发生在新的 Go 程中。
    Go 程在相同的地址空间中运行，因此访问共享的内存时必须同步， sync 包提供了这种能力。
*/

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}