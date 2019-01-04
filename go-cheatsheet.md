# go 基本类型
```
bool

string

int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8

rune //uint32

float32 float64

complex64 complex128

```

# 常量
```
const 关键字
不能用 :=
```

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

## slice (切片)
数组大小固定，切片为数组提供了动态大小，灵活的视角

`[]T` 表示元素 T 类型的 slice

切片通过上界，下界来界定，用 `:` 分隔 `a[low : high]`, 是一个半开区间，包括第一个元素，排除最后一个元素。

a[1, 4] 切片包含 1 到 3 的元素。

```
package main

import "fmt"

func main() {
    primes := [6]int{1, 2, 3, 4, 5, 6}
    s := primes[1: 4]
    fmt.Println(s)
}
```

## slice 像数组的引用
slice 不存储任何数据，它描述底层数组的一段

更改切边会修改底层数组的元素

共享底层数组的其他 slice 会看到修改

```
package main

import "fmt"

func main() {
    names := [4]string{
        "a",
        "b",
        "c",
        "d",
    }
    fmt.Println(names)
    a := names[0, 2]
    b := names[1, 3]
    fmt.Println(a, b)

    b[0] = aaa
    fmt.Println(a, b)
    fmt.Println(names)
}
```

## slice 文法
切片类似没有长度数组的文法

```
// 数组
[3]bool{true, false, false}

// 切片
[]bool{true, false, false}
```

```
package main

import "fmt"

func main() {
    q := []int{1, 2, 3}
    r := []bool{true, false}
    s := []struct {
        i int
        b bool
    }{
        {1, true},
        {2, false}
        }

    fmt.Println(q, r, s)
}
```

## slice 的默认边界
slice 下界的默认是 `0`, 上界默认是切片长度

```
//对于数组

var a [10]int

// 以下切片是等价的

a[0: 10]
a[: 10]
a[0:]
a[:] 
```

## slice 的长度(length)和容量 (capacity)

切片有 length 和 capacity

length 是切片包含元素的个数

capacity 是第一个元素开始，到底层元素末尾的个数

可以用表达式 `len(s)`, `cap(s)` 获取长度和容量

## nil 切片
切片的零值是 `nil`

`nil` 切片长度和容量为 0, 且没有底层数组

```
package main

import "fmt"

func main() {
    var s []int
    fmt.Println(s, len(s), cap(s))
    if s == nil {
        fmt.Println("nil")
    }
}
```

## make 创建 slice
使用内建函数 make 创建 slice, 也可以创建动态数组

make 函数分配一个元素为零值得数组并返回一个引用了它的切片:

```
a := make([]int, 5) // len(a) = 5
```

指定容量需要传入第三个参数:

```
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b) = 5, cap(b) = 5
b = b[1:]  // len(b) = 4, cap(b) = 4
```

## 切片的切片
切片可以包含切片

```
package main 

import (
    "fmt"
    "strings"
)

func main() {
    board := [][]string{
        []string{"a", "b", "c"},
        []string{"e", "f", "g"},
        []string{"h", "i", "j"},
    }
    board[0][0] = "x"
}
```

## 向切片追加元素
内建的 append 函数

```
func append(s []T, ...T) []T
```
append 结果是一个添加新元素的切片。当底层数组太小，会分配更大数组，返回切片指向新的数组。

```
package main

import "fmt"

func main() {
    var s []int
    s = append(s, 0) // append works on nil slice
    s = append(s, 1, 2, 3)
    printSlice(s)
}

func printSlice(s []int) {
    fmt.Printf("len=%d, cap=%d, %v", len(s), cap(s), s)
}
```

# range
for 循环的 range 形式遍历切片或映射

使用 for 遍历 slice 时，每次迭代返回两个值，第一个是当前元素的下标，第二个是元素对于元素的一个副本。

```
package main

import "fmt"

var pow = []int{1, 2, 3, 4, 5}

func main() {
    for i, v := range pow {
        fmt.Printf("%d, %d", i, v)
    }
}
```

使用 `_` 忽略索引或值

```
func main() {
    pow := make([]int, 10)
    for i := range pow {
        pow[i] = 1 << uint(i) 
    }
    for _, value := range pow {
        fmt.Println("%d", value)
    }
}
```

# map 映射
map 将键映射到值。

零值为 nil, nil 映射即没有键，也不能添加 key

make 函数返回给定类型的 map, 并将其初始化备用。

```
package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell"] = Vertex{40.662, 34.555}
}
```

## map 的文法
map 与 struct 文法类似，必须有键名

```
package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex {
    "abc": Vertex {
        40.112, 34.223,
    },
    "def": Vertex {
        20.234, 45.123,
    },
}

```

如果顶级类型只是一个类型名，可以在文法中省略它。

```
package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell": {29.34, 19.23},
    "Mike": {34.12, 34.56},
}
```

## 使用 map
插入或修改 map, 
```
m[key] = elem
```

获取元素
```
elem = m[key]
```
删除元素 `delete(m, key)`, 

双赋值检测某个键是否存在 `elem, ok = m[key]` , 如果 key 在 m 中，ok=true, 否则 ok=false

# 函数值
函数可以像其他值一样传递，作为参数，返回值。

```
package main

import (
    "fmt"
    "math"
)

func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(compute(hypot))
}
```

# 函数闭包 closure
闭包是一个函数值，它引用了函数体之外的变量。该函数可以访问并赋予其引用变量的值。

```
package main

import "fmt"

func addr() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := addr(), addr()
    for i:=0, i<10; i++ {
        fmt.Println(
            pos(i),
            neg(-2 * i)
        )
    }
}
```

# 方法 method
Go 没有类，不过你可以为 struct 定义方法。 方法是带`接收者`参数的函数.

接收者位于 `func` 和方法名之间，在方法接收者自己的参数列表内。

```
package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
}
```

## method 只是带有接收者的 func
方法和正常函数功能没有变化

```
package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func Abs(v Vertex) float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(Abs(v))
}
```

* 可以为非 struct 声明方法，如 MyFloat 数值类型。只能为同一包中的类型接收者声明方法，不能为内建类型 (如: int) 声明方法。

```
package main

import (
    "fmt"
    "math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f<0 {
        return float64(-f)
    } else {
        return float64(f)
    }
}

func main() {
    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs())
}

```

## 指针接收者
可以为指针声明方法

对于某类型 T , 方法接收者可以是 `*T` (但不能是内建类型指针， 如: `*int`)

## 方法与指针重定向
带指针的函数，必须接受一个指针

指针接收者的方法调用时，接收者可以是指针，或值。 Go 会将值重定向为指针 (&v).Scale(5)

值为接收者的方法调用，接收者可以为值或指针。 Go 会将指针重定向为值 `(*v).Scale(5)`

使用指针作为接收者，1 方法可以修改指针指向的值， 2 避免每次调用复制该值，对大型结构体更高效


# 接口 interface
一组方法签名定义的集合

接口类型的变量可以保存任何实现了这些方法的值。

```
package main

import (
    "fmt"
    "math"
)

type Abser interface {
    Abs() float64
}
```

## 接口与隐式实现
类型通过实现接口所有方法来实现接口，无需显示声明，不需要 `implements` 关键字

隐式接口从接口实现中解耦了定义，这样的接口实现可出现在任何包中，无需提前准备。

```
package main

import "fmt"

type I interface {
    M()
}

type T struct {
    S string
}

func (t T) M() {
    fmt.Println(t.S)
}

func main() {
    var i I = T{"hello"}
    i.M()
}
```

## 接口值 interface value

## 空接口 interface{}
空接口可以保存任何类型的值。

# 类型断言

```
t, ok := i.(T)
```

## 类型选择

```
package main 

import "fmt"

func do(i interface{}) {
    switch v := i.(type) {
        case int:
            
        case string:

        default:
    }
}

```

# Stringer
fmt 包中的 Stringer

```
type Stringer interface {
    String() string
}
```

# Error
GO 使用 error 表示错误状态， error 类型是一个内建接口

type error interface {
    Error() string
}

通常函数会返回一个 error 值，判断 error 是否为 nil 

# Reader 
io.Reader 接口

```
func (T) Read(b byte[]) (n int, err error)
```

# Go routine
goroutine 是由 Go 运行时管理的轻量级线程

```
go f(x,y,z)
```

# channel
信道是带类型的管道 `<-`

```
ch <- v // 将 v 发送至信道 ch
v := <-ch

ch := make(chan int)
```

```
pacakge main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum
}

func main()
```

## buffer channel

```
ch := make(chan int, 100)
```

## range 和 close

发送者可以 close 关闭一个信道来表示没有需要发送的值

## select 语句


