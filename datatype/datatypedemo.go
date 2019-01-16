package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a uint  //无符号整型 默认值为0 位数与系统相关
	var b int8  //整型 八位
	var c int16 //整型16位
	var d int32 //整型32fmt.Println(unsafe.Sizeof(a))位
	var e int64 //整型64位
	var f float32
	var g bool
	var h complex64 //复数
	fmt.Println("占内存情况：")
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(d))
	fmt.Println(unsafe.Sizeof(e))
	fmt.Println(unsafe.Sizeof(f))
	fmt.Println(unsafe.Sizeof(g))
	fmt.Println(unsafe.Sizeof(h))
	fmt.Println("默认值:")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}