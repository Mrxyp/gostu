//程序所属包
package main

//导入fmt包
import (
	"fmt"
)

//常量  string 代表 NAME 为string 类型 可要可不要 go自带类型推断机制
const NAME string = "hello go!"

//全局变量
var globalVar = 1

//一般类型声明  声明myType为int 型
type myType int

//结构体声明
type Student struct {
}

//声明接口
type IInterface interface {
}

//main函数
func main() {
	fmt.Println(NAME)
	add()
	fmt.Println(globalVar)
}

//自定义函数
func add() {
	globalVar++
}
