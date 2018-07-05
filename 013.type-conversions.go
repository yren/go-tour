/*
    类型转换
    Go 需要显示的类型转换。 T(v) 将值 v 转换为类型 T
    例如：
    var i int = 42
    var f float64 = float64(i)
    var u unit = unit(f)
    
    或者
    
    i := 42
    f := float64(i)
    u := unit(f)
*/

package main

import (
    "fmt"
    "math"
)

func main() {
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x * x + y * y))
    var z uint = uint(f)
    fmt.Println(x, y, z)
}