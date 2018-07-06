/*
    range
    for 循环的 range 形式可以遍历切片或映射
    当 for 循环遍历 切片时，每次迭代都返回两个值，第一个值当前元素的下标，第二个值下标对应元素的副本
*/

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }
}