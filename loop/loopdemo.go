package main

import (
	"fmt"
	"time"
)

func main() {

	//方式1
	for i:=1;i<=10;i++{
		fmt.Println(i)
	}
	//方式2
	arr := []int{1,2,3,4,5,6}
	for key,value:=range arr{
		fmt.Println(key,value)
	}
    //不要下标
	for _,value:=range arr{
		fmt.Println(value)
	}

	//无限循环
	for{
		fmt.Println("balabala")
		time.Sleep(1*time.Second)
	}

}
