/*
    switch 是一连串 if else 的简便方法。 它运行的哥值等于表达式的 case 语句。
    Go 的 swtich 与 java, c, javascript 不同， case 之后无需 break. 实际上， Go 自动
    在每个 case 后加入了 break 语句。
    除非以 fallthrough 语句结束， 否则分支自动结束。
    Go 的 case 无需为常量， 且取值不必为整数
*/

package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {
    case "drwin":
        fmt.Println("OS X.\n")
    case "linux":
        fmt.Println("Linux.\n")
    default:
        fmt.Printf("%s.\n", os)
    }
}