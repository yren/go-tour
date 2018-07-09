/*
    Stringer
    fmt 包中定义的 Stringer 是最普遍的接口之一
    type Stringer interface {
        String() string
    }
    
    Stringer 是一个可以用字符串描述自己的类型。 fmt 包通过此接口来打印值

*/

package main

import "fmt"

type Person struct {
    Name string
    Age int
}

func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
    a := Person{"Mike", 23}
    z := Person{"Jerry", 19}
    fmt.Println(a, z)
}