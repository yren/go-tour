/*
    类型断言 提供了访问接口值底层具体值的方式。
    t := i.(T)
    
    该语句断言接口值 i 保存了具体类型 T, 并将其底层类型为 T 的值赋予变量 t
    如果 i 没有保存 T 类型的值，该语句会触发一个恐慌。
    为了判断一个接口值是否保存了一个特定类型，断言可以返回连个值，底层值和断言是否成功的布尔值。
    t, ok := i.(T)
    
    若 i 保存了一个 T, t 是底层值， ok 为 true
    否则, ok 为 false, t 是 T 类型的零值，程序并不会产生恐慌。
*/

package main

import "fmt"

func main() {
    var i interface{} = "hello"
    
    s := i.(string)
    fmt.Println(s)
    
    s, ok := i.(string)
    fmt.Println(s, ok)
    
    f, ok := i.(float64)
    fmt.Println(f, ok)
}