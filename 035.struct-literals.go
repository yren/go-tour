/*
    结构体文法通过直接列出字段的值来分配一个新的结构体。
    使用 Name: 语法可以仅列出部分字段 (字段名顺序无关)
    & 返回指向结构体的指针
*/

package main

import "fmt"

type Vertex struct {
    X, Y int
}

var (
    v1 = Vertex{1, 2}
    v2 = Vertex{X: 1}  //Y:0 is implicit
    v3 = Vertex{}  //X:0 and Y:0
    p = &Vertex{1, 2}
)

func main() {
    fmt.Println(v1, p, v2, v3)
}

