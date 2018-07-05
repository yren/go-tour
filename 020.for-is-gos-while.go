/*
    for 是 Go 中的 while
    去掉分号 for 和 while 的作用相同
*/

package main

import "fmt"

func main() {
    sum := 1
    for sum < 100 {
        sum += sum
    }
    fmt.Println(sum)
}