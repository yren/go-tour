/*
    此代码用圆括号组合了导入， 这是 "分组" 形式的导入。
    当然你也可以使用多个导入语句，如:
    import "fmt"
    import "math"
    
    但使用分组导入语句是更好的形式
*/

package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Printf("Now you have %g points.\n", math.Sqrt(9))
}