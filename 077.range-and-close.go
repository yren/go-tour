/*
    range 和 close
    发送者可以通过 close 关闭一个 channel 表示没有需要发送的值了。接收者可以通过为接收表达式
    分配第二个参数来测试信道是否被关闭: 若没有值可以接收且信道已经关闭，
    v, ok := <-ch 
    之后 ok 会被设置为 false.
    
    循环 for i := range c 会不断从信道接收值，知道它被关闭。
*/

package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
