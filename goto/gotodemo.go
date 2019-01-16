package main

import "fmt"

func main() {
	goto One
	fmt.Println("中间代码块")
	One:
		fmt.Println("这是代码块One")
		goto One //这里将产生无限循环
}
