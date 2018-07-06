/*
    defer 推迟的函数会被压入一个栈中，当外层函数返回时，被推迟的函数会按照 LIFO 顺序调用。
    关于更多 defer 内容参考 https://blog.go-zh.org/defer-panic-and-recover
*/

package main

import "fmt"

func main() {
    fmt.Println("counting")
    
    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }
    
    fmt.Println("done")
}