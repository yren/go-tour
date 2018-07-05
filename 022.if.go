/*
    Go 的 if 和 for 类似，无需小括号 (), 大括号 {} 是必须的
*/

package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) string {
    if x < 0 {
        return  sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func main() {
    fmt.Println(sqrt(2), sqrt(-4))
}