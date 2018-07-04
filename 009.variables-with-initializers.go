/*
    变量声明可以包含初始值，每个变量对应一个。
    如果初始值存在，可以省略类型声明，类型可以从初始值中推导出来。
*/

package main

import "fmt"

var i, j int = 1, 2

func main() {
    var c, python, java = true, false, "fine!"
    fmt.Println(i, j, c, python, java)
}