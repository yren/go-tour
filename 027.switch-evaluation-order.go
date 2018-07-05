/*
    switch 的 case 语句从上到下顺次执行，直到匹配成功为止
*/

package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("When's Saturday?")
    today := time.Now().Weekday()
}