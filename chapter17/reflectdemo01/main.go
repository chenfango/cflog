/*
 * @Description: In User Settings Edit 
 * @Author: your name
 * @Date: 2019-08-15 10:29:25
 * @LastEditTime: 2019-08-15 11:27:13
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"
	"reflect"
)

// 专门演示反射的
func reflectTest01(b interface{}) {
	// 通过反射获取到传入的变量的type, kind, 值是什么
	// 先获取到reflect.Type

	// fmt.Printf("b=%v, type=%T\n", b,b)
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	
	// 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := 2.0 + rVal.Int()
	fmt.Println(n2)
	fmt.Printf("rVal=%v\n type=%T\n", rVal, rVal)

	// 我们将reflect.Value转成interface
	iV := rVal.Interface() 
	
	// 将interface通过断言转成需要的类型
	num2 := iV.(int)

	fmt.Println("num2=", num2)
	
}



// 专门演示反射(对结构体)
func reflectTest02(b interface{}){
	// fmt.Printf("b=%v, type=%T\n", b,b)
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	
	// 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	// 获取变量对应的kind
	// (1) rVal.Kind() 
	Kind1 := rVal.Kind()
	// (2) rTyp.Kind()
	Kind2 := rTyp.Kind()

	fmt.Printf("kind=%v kind=%v\n", Kind1 , Kind2)

	// 我们将reflect.Value转成interface
	iV := rVal.Interface()
	fmt.Printf("iV=%v, type=%T\n ", iV, iV)
	
	// 将interface通过断言转成需要的类型
	stu,ok := iV.(Student)

	
	if ok{
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

	
}

type Student struct{
	Name string
	Age int
}


func main(){
	// 先定义一个int类型
	// var num int = 100
	// reflectTest01(num)

	// 定义一个Student实例
	stu := Student{
		Name: "tom",
		Age: 20,
	}


	reflectTest02(stu)

}