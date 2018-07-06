/*
    结构体指针
    结构体字段可以通过指针来访问
    如果 p 是指向结构体的指针，那么可以通过 (*p).X 访问字段 X, 但这样写太啰嗦，所以语言允许我们
    使用隐式间接引用，直接写 p.X 即可
*/

package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    p := &v
    p.X = 100
    fmt.Println(v)
}