/*
    数组
    类型 [n]T 表示拥有 n 个 T 类型值的数组
    
    var a [10]int
    
    数组的长度是其类型的一部分，因此数组不能改变大小。 这看起来是个限制，但 Go 更加
    便利的方式来使用数组
*/

package main

import "fmt"

func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)
    
    primes := [6]int{2, 3, 5, 7, 89, 20}
    fmt.Println(primes)
}