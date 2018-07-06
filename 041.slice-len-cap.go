/*
    切片拥有 长度 和 容量
    切片的长度是包含元素的个数。
    切片的容量是从它的抵押给元素开始数，到其底层数组元素末尾的个数。
    len(s) 和 cap(s)
    可以通过重新切片来扩展一个切片。
*/

package main

import "fmt"

func main() {
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s)
    // give the slice 0 length
    s = s[:0]
    printSlice(s)
    
    s = s[:6]
    printSlice(s)
    
    s = s[2:4]
    printSlice(s)

}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}