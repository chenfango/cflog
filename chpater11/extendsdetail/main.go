package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

func (a *A) SayOk(){
	fmt.Println("A Sayok", a.Name)

}

func (a *A) hello(){
	fmt.Println("A Hello", a.Name)
}

type B struct{
	A
}

func main(){
	var b B 
	b.A.Name="tome"
	b.A.age=19
	b.A.SayOk()
	b.A.hello()
}
