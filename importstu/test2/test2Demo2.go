package test2

import (
	"fmt"
	"gostu/importstu/test1"
)

func init() {
	fmt.Println("this is test2Demo2")
}
//模块中要导出的函数首字母要大写
func Test2()  {
	test1.Test1()
}