/*
    类型选择 是一种按顺序从几个类型断言中选择分支的结构。
    类型选择与一般 switch 语句相似，但 case 为类型， 它针对给定接口值锁存储值进行比较
    
    switch v := i.(type) {
    case T:
        // v 的类型为 T
    case S:
        // v 的类型为 S
    default:
        // 没有匹配
}

类型选择与类型断言 i.(T) 语法相同，只是将具体类型 T 替换成 type 关键字
*/

package main

import "fmt"

func do(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer\n")
    case string:
        fmt.Printf("String\n")
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}

func main() {
    do(21)
    do("hello")
    do(true)
}