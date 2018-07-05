/*
    常量
    常量声明使用 const 关键字
    常量可以是 字符，字符串，布尔值，数值
    常量不能用 := 语法
*/

package main

import "fmt"

const Pi = 3.14

func main() {
    const World = "dragon world"
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")
    
    const Truth = true
    fmt.Println("Go rules?", Truth)
}