package main

/*
    parse command argument example
*/

import (
    "fmt"
    "flag"
)

var name string

func init() {
    flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
    //flag.Parse 解析参数，必须在参数声明和设置之后，在读取命令参数之前进行
    flag.Parse()
    fmt.Printf("Hello, %s!\n", name)
}