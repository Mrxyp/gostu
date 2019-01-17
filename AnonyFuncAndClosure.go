package main

import "fmt"

//匿名函数与闭包
func main() {
	//1.无参匿名函数
	num:=1
	f:=func() int{
		num++
		return num
	}
	fmt.Println(f())

	//2.有参
	x,y:=func(a,b int)(min,max int){//(a,b int) 为参数 (min,max int)为返回类型
		if a>b{
			max=a
			min=b
		}else{
			min=a
			max=b
		}
		return min,max
	}(6,3)
	fmt.Println(x,y)

	//闭包
	fn:=test1()
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
}

//闭包 函数里有函数 最里面的函数可以访问外层函数的变量
func test1() func()int{
	i:=1
	f:=func() int{
		i++
		return  i
	}
	return f
}