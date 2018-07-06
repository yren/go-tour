/*
    切片的零值是 nil
    nil 的 长度 和 容量 为 0， 且没有底层数组
*/

package main

import "fmt"

func main() {
    var s []int
    fmt.Println(s, len(s), cap(s))
    
    if s == nil {
        fmt.Println("nil!")
    }
}