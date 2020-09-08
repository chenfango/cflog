/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-14 19:25:42
 * @LastEditTime: 2019-08-15 08:40:42
 * @LastEditors: Please set LastEditors
 */
package main


import (
	_ "fmt"
)

func main2(){

	intChan := make(chan int ,3)
	intChan <- 2
	intChan <- 1
	intChan <- 3
	close(intChan)

	for i:=0;i<4;i++{
		<- intChan 
	}
	// fmt.Println("ffhhf")
}


