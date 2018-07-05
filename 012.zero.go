/*
    没有明确初始化的变量会被赋予零值
    零值是:
    数值类型为 0,
    布尔类型为 false,
    字符串为 ""
*/

package main

import "fmt"

func main() {
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)
}