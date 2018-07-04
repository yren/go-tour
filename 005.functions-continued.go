/*
    当函数的连续多个参数类型相同，可以只保留最后一个，省略前面的类型声明。
    如：
    x int, y int
    被缩写为:
    x, y int
*/

package main

import "fmt"

func add(x, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(10, 31))
}