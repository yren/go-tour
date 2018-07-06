/*
    向切片追加元素
    append 函数
    func append(s []T, vs ...T) []T
    
    append 的第一个参数 s 是一个元素乐行为 T 的切片， 其余类型为 T 的值会追加的切片的末尾。
    append 的结果返回一个包含原切片元素加上新添加元素的切片。
    当 s 的底层数组太小，不足以容纳给定值。它就分配一个更大的数组，返回的切片指向新数组。
*/

package main

import "fmt"

func main() {
    var s []int
    printSlice(s)
    
    // append works on nil slices.
    s = append(s, 0)
    printSlice(s)
    
    s = append(s, 1)
    printSlice(s)
    
    s = append(s, 2, 3, 4)
    printSlice(s) 
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}