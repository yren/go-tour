/*
    指针
    Go 中拥有指针，指针保存了值的内存地址
    类型 *T 是指向 T 类型的指针。其零值为 nil
    
    var p *int
    
    & 操作符会生成一个指向其操作数的指针。
    
    i := 42
    p = &i
    
    * 操作符表示指针指向的底层值。
    
    fmt.Println(*p) // 通过指针 p 读取 i
    *p = 21 // 通过指针 p 设置 i
    
    这也就是通常所说的 "间接引用" 或 “重定向”
    与 c 不同， Go 没有指针运算
*/

package main

import "fmt"

func main() {
    i, j := 42, 2701
    p := &i //point to i
    fmt.Println(*p) // read i through pointer
    *p = 21 // set i through the pointer
    fmt.Println(i)
    
    p = &j // point to j
    *p = *p / 37 // divide j through pointer
    fmt.Println(j)
}