/*
    可以将下标 或 值 赋予 _ 来忽略他。
    如果只需要索引，去掉 ,value 部分即可
*/

package main

import "fmt"

func main() {
    pow := make([]int, 10)
    
    for i := range pow {
        pow[i] = 1 << uint(i)
    }
    for _, value := range pow {
        fmt.Printf("%d\n", value)
    }
}