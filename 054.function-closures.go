/*
    闭包是一个函数，它可以引用函数体之外的变量。该函数可以访问并赋予其引用变量的值。换
    句话说该函数被绑定在了这些变量上。
*/

package main

import "fmt"

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(pos(i), neg(-2*i))
    }
}