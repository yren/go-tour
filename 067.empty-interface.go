/*
    没有指定方法的接口成为 空接口
    interface{}
    
    空接口可以保存任何类型的值
*/

package main

import "fmt"

func main() {
    var i interface{}
    describe(i)
    
    i = 42
    describe(i)
    
    i = "hello"
    describe(i)
}

func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}