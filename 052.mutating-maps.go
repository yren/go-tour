/*
    插入或修改元素
    m[key] = elem
    
    获取元素
    elem = m[key]
    
    删除元素
    delete(m, key)
    
    双赋值检测某个键是否存在:
    elem, ok = m[key]
    若 key 在，则 ok 为 true, 否则 ok 为 false
    若 key 不在， elem 为类型的零值
*/

package main

import "fmt"

func main() {
    m := make(map[string]int)
    
    m["Answer"] = 10
    fmt.Println("The value: ", m["Answer"])
    
    m["Answer"] = 20
    fmt.Println("The value: ", m["Answer"])
    
    delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}