/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-14 09:45:07
 * @LastEditTime: 2019-08-14 14:13:34
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"
)

func main(){
	intChan := make(chan int ,3)
	intChan <-100
	intChan <-200
	// 关闭管道
	close(intChan)
	// 这时不能够再写入数据到这个chan
	
	// intChan<- 300

	// 当管道关闭后，读取数据是可以的
	n1:= <-intChan

	// fmt.Println("close chan")
	fmt.Println(n1)

	// 遍历管道
	intChan2 := make(chan int, 100)
	for i:=0; i<100;i++{
		intChan2<-i*2 // 放入100个数据
		
	}
	
	// 遍历管道不能使用普通的for循环结构
	// for i:=0;i<len(intChan2);i++{

	// }
	
	// 在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
	close(intChan2)
	for v:= range intChan2{
		fmt.Println("v=", v)
	}
	
}
