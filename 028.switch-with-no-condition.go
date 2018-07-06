/*
    没有条件的 swtich, 等同于 swtich true.
    这种形式可以将一长串 if then else 写的更清晰
*/

package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good Morning.")
    case t.Hour() < 18:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }
}