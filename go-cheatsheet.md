# 指针
Go 有指针，保存值的内存地址

类型 `*T` 是指向 `T` 类型的指针，零值为 `nil`

`var p *int`

`&` 操作符会生成一个指向其操作数的指针

```
var p *int
i := 42
p = &i
```

`*` 操作符表示指针指向的底层值

```
fmt.Println(*p)
*p = 21 // 这就是间接引用或重定向
```

Go 没有指针运算

```
package main 

import "fmt"

func main() {
    i, j := 42, 27
    p := &i 
    fmt.Println(*p)
    *p = 21
    fmt.Println(i)

    p = &j
    *p = *p / 37
    fmt.Println(j)
}
```

# struct (结构体)
struct (结构体)就是一个字段的集合

```
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    fmt.Println(Vertex{1, 2})
}
```

## struct field (结构体字段)
结构体字段使用 `.` 号访问

```
package main 

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)
}
```