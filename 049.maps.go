/*
    映射的零值为 nil . nil 映射没有键，也不能添加键
    make 函数返回戈丁类型映射，并将其初始化备用
*/

package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.68, -74.39,
    }
    fmt.Println(m)
}