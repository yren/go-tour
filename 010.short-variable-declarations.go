/*
    在函数中， 简洁赋值语句 := 可在类型明确的地方代替 var 声明。
    函数外的每个语句必须以关键字开始 (var, func 等)， 因此 := 不能在函数外使用。
*/

package main

import "fmt"

func main() {
    i, j := 1, 2
    k := 3
    c, python, java := true, false, "cool!"
    fmt.Println(i, j, k, c, python, java)
}