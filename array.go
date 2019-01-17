package main

import "fmt"
//数组
func main() {
	var arr = []int{1,2}
	//a[0:1]切片用法 左闭右开
	fmt.Println(arr[0:1])
	//数组参数
	arrparam(1,2,3,4)

}
//不定参数 后面必须跟类型
func arrparam(arr... int ){
	fmt.Println(arr[2:4])

	arrprint(arr...)//传递所有的参数

}
func arrprint(arr... int){
	fmt.Println(arr)

	arrprint1(1)//定参必须传值
}
//定参必须传值
func arrprint1(a int,arr... int){
	fmt.Println(a)
	fmt.Println(arr)
}