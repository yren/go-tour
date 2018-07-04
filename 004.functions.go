/*
    函数可以有零到多个参数。本例子中 add 接受两个 int 类型参数。
    注意类型在变量名之后。
*/

package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(40, 15))
}