package main

import (
	"fmt"
)

type BInterface interface{
	test01()
}

type CInterface interface{
	test02()
}

type AInterface interface{
	BInterface
	CInterface
	test03()

}

// 如果我们需要实现AInterface接口，我们需要将B和C接口的方法都实现
type Stu struct{

}

func (stu Stu) test01(){

}

func (stu Stu) test02(){

}

func (stu Stu) test03(){

}

func main(){
	// 创建一个Stu的结构体
	var stu Stu
	// var a AInterface = stu
	stu.test01()
	
}
