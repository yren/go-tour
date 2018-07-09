/*
    接口类型 是由一组方法签名定义的集合
    接口类型的变量可以保持任何实现了这些方法的值
*/

package main

import (
    "fmt"
    "math"
)

type Abser interface {
    Abs() float64
}

func main() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}
    
    a = f // MyFloat 实现了 Abser
    a = &v // *Vertex 实现了 Abser
    
    //a = v
    
    fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}