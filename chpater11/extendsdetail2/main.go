package main

import (
	"fmt"
)

type A struct{
	Name string
	Age int
}

type B struct{
	Name string
	Score float64
}

type C struct{
	A
	B
}


func main(){
	var c C
	// 如果c 没有Name字段,而A,B有name字段，则需要通过匿名结构体名字来区分
	
	c.A.Name = "tome" // error
	fmt.Print(c.A.Name)
}