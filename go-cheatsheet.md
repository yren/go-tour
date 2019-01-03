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

## 结构体指针 struct pointer
结构体字段可以通过结构体指针访问。

`p` 是指向 struct 的指针，可以 `(*p).X` 访问字段 `X`, 不用这么啰嗦，语言允许隐式间接引用，直接用 `p.X`

```
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    p := &v
    p.X = 20
    fmt.Println(v)
}
```

## 结构体文法 (struct literal)
* 直接列出字段值分配一个新的 struct
* 使用 `Name:` 列出部分字段
* 用 `&` 返回 struct pointer

```
package main

import "fmt"

type Vertex struct {
    X, Y int
}

var (
    v1 = Vertex{1, 2}
    v2 = Vertex{X: 5} // Y:0 
    v3 = Vertex{} // X:0, Y:0
    p = &Vertex{1, 2}
)

func main() {
    fmt.Println(v1, p, v2, v3)
}
```

# 数组 array
类型 `[n]T` n 个 T 类型的数组

```
var a [10]T
```

数组的长度是其类型的一部分，数组不能改变大小。Go 通过 `slice` 使用数组

```
package main

import "fmt"

func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{1, 2, 3, 4, 5, 6}
    fmt.Println(primes)
}
```