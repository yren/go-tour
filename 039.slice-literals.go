/*
    切片的文法类似没有长度数组的文法。
    数组的文法：
    [3]bool{true, flase, true}
    
    下面会创建一个和上面相同的数组，然后构建一个引用了它的切片：
    
    []bool{true, flase, true}
*/

package main

import "fmt"

func main() {
    q := []int{2, 3, 5, 7, 9, 11}
    fmt.Println(q)
    
    r := []bool{true, false, true, true, false, false}
    fmt.Println(r)
    
    s := []struct {
        i int
        b bool
    }{
        {2, true},
        {3, false},
    }
    
    fmt.Println(s)
}