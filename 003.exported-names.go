/*
    在 Go 中，如果名字以大写字母开头，那么它就是已导出的。如: Pizza 是个已导出名，Pi 也是， 它导出自 math 包。
    pizza 和 pi 没有以大写字母开头，它们是未导出的。
    在导入一个包时，你只能引用已导出的名字。任何 "未导出" 的名字在该包外无法访问。
*/

package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Pi)
}