/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-14 09:24:53
 * @LastEditTime: 2019-08-14 09:38:45
 * @LastEditors: Please set LastEditors
 */

 package main
 
 import (
	 "fmt"
 )
 
 type Cat struct {
	 Name string 
	 Age int
 }


 func main(){
	// 定义一个可以存放任意类型的管道
	// var allChan chan interface{}
	// allChan = make(chan interface{}, 2)
	
	// 类型推导
	allChan := make(chan interface{}, 3)
	allChan<- 10
	allChan<- "tom cat"
	cat := Cat{
		"小花猫",
		4,
	} 
	allChan<- cat

	// 我们希望获取管道中的第三个元素, 则先将前两个推出
	<-allChan
	<-allChan
	
	newCat := <-allChan  // 从管道中取出的猫

	a := newCat.(Cat)
	fmt.Printf("newCat=%T, newCat=%v\n", newCat, newCat)
	fmt.Printf("newCat.Name=%v\n", a.Name)


	




	 
 }