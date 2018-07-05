/*
    数值常量
    数值常量是高精度的值。
    一个未指定类型的常量有上下文决定类型。
*/

package main

import "fmt"

const (
    // create a huge number by shift a 1 bit left 100 places.
    // In other words, the binary number is 1 followed by 100 zeros
    Big = 1 << 100
    // shift it right again 99 places, so we end up with 1 << 1 or 2 
    Small = Big >> 99
)

func needInt(x int) int {
    return x * 10 + 1
}

func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
}