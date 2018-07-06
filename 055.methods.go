/*
    Go 没有类。不过你可以为结构体类型定义方法。
    方法是带特殊的 接受者 参数的函数
    接受者位于 func 和 方法名之间
*/

package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
}