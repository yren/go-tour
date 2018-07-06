/*
    在进行切片时，你可以利用它的默认行为忽略上下界。
    下界的默认值是 0， 上界默认是切片长度。
    
    对于数组
    var a [10]int
    
    来说，以下切片时等价的:
    a[0:10]
    a[:]
    a[0:]
    a[:10]
*/

package main

import "fmt"

func main() {
    s := []int{2, 3, 5, 7, 11, 13}
    s = s[1:4]
    fmt.Println(s)
    
    s = s[:2]
    fmt.Println(s)
    
    s = s[1:]
    fmt.Println(s)
}