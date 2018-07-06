/*
    映射的文法与结构体相似，不过必须有键名。
*/

package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex {
    "Bell Labs": Vertex {
        40.68, -74.39,
    },
    "Google": Vertex {
        37.42, -122.08,
    },
}

func main() {
    fmt.Println(m)
}