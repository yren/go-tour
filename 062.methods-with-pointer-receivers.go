/*
    使用指针接收者的原因有二：
    首先，方法能够修改其接收者指向的值。
    其次，避免每次调用方法时复制该值，若值为大型结构体，这样更加高效。
    
    通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。
*/

package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := &Vertex{3, 4}
    fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
    v.Scale(5)
    fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}