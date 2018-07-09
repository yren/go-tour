/*
    你可以为指针接收者声明方法。
    这意味着对于某种类型 T , 接收者类型可以用 *T 的文法 (此外， T 不能是像 *int 这样的指针)
    例如, 这里为 *Vertex 定义了 Scale 方法。
    指针接收者的方法可以修改接收者指定的值。 由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
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

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    v.Scale(10)
    fmt.Println(v.Abs())
}