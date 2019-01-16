package iotastu

import (
	"fmt"
)
//iota 遇到const 会被重新赋值为0
const a = iota
const b = iota

//顺序声明会自增
const(
	c =iota
	d =iota
)
//跳值使用 _相当于是一个垃圾堆
const(
	e = iota
	_ = iota
	f = iota
)
//插入使用
const(
	g = iota
	h=123
	i=iota
)
//公式继承
const(
	j = iota*2 + iota
	k
)
func main() {
	fmt.Println(a,b)
	fmt.Println(c,d)
	fmt.Println(e,f)
	fmt.Println(g,i)
	fmt.Println(j,k)
}