/*
    Go 的返回值可以被命名，它们被视为定义在函数顶部的变量。
    返回值的名称应该有意义，它可以作为文档使用。
    没有参数的 return 语句返回已命名的返回值。也就是直接返回。
    直接返回用在短函数中。在长函数中会影响代码可读性。
*/

package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}