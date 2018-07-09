/*
    信道可以带缓冲, 缓冲长度作为第二个参数提供给 make
    
    ch := make(chan int, 100)
    仅当信道的缓冲区填满后，向其发送数据才会阻塞，当缓冲区为空，接收方会阻塞。
*/

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}