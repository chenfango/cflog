/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-13 23:21:23
 * @LastEditTime: 2019-08-14 09:06:22
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"
)

func main(){
	// 管道的使用
	// 1. 创建一个可以存放三个int类型的管道
	var intChan chan int
	intChan = make(chan int ,3)
	// 2. 看intchan是什么, chan是引用类型，可以传给函数使用
	fmt.Printf("intChan的值=%v intChan本身的地址=%p\n", intChan, &intChan)

	// 3. 向管道写入数据
	intChan<- 10
	num:=210
	intChan<- num

	
	intChan<-50  // 管道不能超过其容量


	
	// 4. 看看管道的长度和容量
	fmt.Printf("channel len=%v, channel cap=%v\n", len(intChan), cap(intChan))

	// 5. 从管道中读取数据
	
	var num2 int
	num2 = <-intChan
	fmt.Printf("num2=%v\n", num2)

	// 6. 在没有使用协程的情况下，我们管道数据已经全部取出来了，如果再取会报错的

	var num3 int 
	num3 = <-intChan
	fmt.Printf("num3=%v\n", num3)


	var num4 int 
	num4 = <-intChan
	fmt.Printf("num4=%v\n", num4)


	// var num5 int 
	// num5 = <-intChan

	intChan <-num

	
	// fmt.Printf("num4=%v\n", num5)





	
	
	


}
