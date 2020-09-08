package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}){
	 //if i,ok:=p.(int);ok{
	 //	fmt.Println("Integer ",i)
	 //	return
	 //}
	 //if s,ok:=p.(string);ok{
	 //	fmt.Println("string",s)
	 //	return
	 //}
	 //fmt.Println("Unknow Type")

	switch v:=p.(type) {
	case int:
		fmt.Println("Inter",v)

	case string:
		fmt.Println("String",v)
	default:
		fmt.Println("Unknow Type",v)
	}
}


func TestEmptyInterfaceAssertion(t *testing.T){
	DoSomething(10)

}