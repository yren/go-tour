/*
    切片的切片
    切片可以包含任何类型，甚至其它的切片。
*/

package main

import (
    "fmt"
    "strings"
)

func main() {
    board := [][]string{
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
    }
    
    board[0][0] = "X"
    board[2][2] = "O"
    
    for i := 0; i < len(board); i++ {
        fmt.Printf("%s\n", strings.Join(board[i], " "))
    }
}