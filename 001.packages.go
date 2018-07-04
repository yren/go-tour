/*
    每个 go 程序都是有包构成的， 程序从 main 包开始运行。
    本程序通过导入路径 "fmt" 和 "math/rand" 来使用这两个包。
    按照约定，包名与导入路径最后一个元素一致。例如， "math/rand" 包中源码以  package rand 开始。
*/

package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println("my number is ", rand.Intn(10))
}