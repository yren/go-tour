/*
    声明变量是不指定类型，类型根据右侧值推导而出
    如：
    i := 42
    f := 3.14
*/

package main

import "fmt"

func main() {
    v := 42
    fmt.Printf("v is of type %T\n", v)
}