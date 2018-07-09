/*
    io 包指定了 io.Reader 接口，它表示从数据流的末尾进行读取。
    Go 包含了该接口的许多实现，包括 文件，网络连接，压缩和加密等等。
    
    io.Reader 接口有一个 Read 方法
    func (T) Read(b []byte) (n int, err error)
    
    Read 用数据填充给定字节切片并返回填充字节数和错误值。在遇到数据流结尾时，它会返回一个 io.EOF 错误。
*/

package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Hello, yufei!")
    b := make([]byte, 8)
    
    for {
        n, err := r.Read(b)
        fmt.Printf("n = %v err = %v b = %v \n", n, err, b)
        fmt.Printf("b[:n] = %q\n", b[:n])
        if err == io.EOF {
            break
        }
    }
}