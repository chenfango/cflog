package main

import (
	"fmt"
)

type AInterface interface{
	 Say()
}

type BInterface interface{
	Hello()
}

type Student struct{
	Name string

}

type integer int

func (i integer) Say(){
	fmt.Println("interge Say i=", i)
}

func (s Student) Say(){
	fmt.Println("Stu say()")
}

type Monster struct{

}

func (m Monster) Hello(){
	fmt.Println("Hello")
}

func (m Monster) Say(){
	fmt.Println("Say")
}

func main(){
	var stu Student  // 结构体变量，实现了Say方法
	var a AInterface = stu
	a.Say()

	var i integer = 10
	var b AInterface = i
	b.Say()

	// Monster实现了AInterface和BInterface

	var monster Monster

	var a2 AInterface = monster
	var b2 BInterface = monster
	
	a2.Say()
	b2.Hello()


}