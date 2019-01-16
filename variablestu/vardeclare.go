package variablestu

import "fmt"

//变量声明与类型转换

//全局变量必须带var
var a int //方式1

//方式2
var(
	b int
	c string
)
//方式3
var d,e,f int = 1,2,3

func main() {
	//方式4  局部变量可以不带var
	g,h:=4,5
	fmt.Println(g,h)
	fmt.Println(a,b,c)
	fmt.Println(d,e,f)

	//相似类型可转换 但可能会损失精度 不同类型不可以转换 如int与bool
	var a float32 = 3.02
	var b = int(a)
	fmt.Println(b)
}