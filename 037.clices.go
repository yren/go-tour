/*
    切片 slice
    每个数组的大小是固定的。切片为数组元素提供动态大小的，灵活的视角。实践中，切片比数组常用。
    类型 []T 表示一个类型为 T 的切片
    
    切片通过两个下标来界定，一个上界和一个下界，用冒号分隔 :
    a[low: high]
    
    是一个半开区间，包含第一个，排除最后一个
*/

package main

import "fmt"

func main() {
    primes := [6]int{1, 4, 7, 9, 10, 45}
    var s []int = primes[1: 4]
    fmt.Println(s)
}