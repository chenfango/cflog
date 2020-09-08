/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-14 16:28:18
 * @LastEditTime: 2019-08-14 20:01:19
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"
	"time"
)

// 向intChan放入8000个数字
func putNum(intChan chan int){
	for i:=0;i<80000;i++{
		intChan<-i
	}

	// 关闭intChan
	close(intChan)

}

// 开启四个协程从intChan取出数据,并判断是否为素数,如果是就放入到primeChan里面去
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool){
	// 使用for循环
	// var num int
	var flag bool
	
	for {
		num, ok := <-intChan
		if !ok{
			break
		}
		flag=true
		// 判断这个num是不是素数
		for i:=2;i<num;i++{
			if num %i == 0{ // 说明该num不是素数
				flag=false
				break
			}
		}
		
		if flag {
			// 将这个数放入到primeChan
			primeChan<-num
		}
	}

	fmt.Printf("有一个协程因为取不到数据,退出了\n")
	// 这里我们还不能关闭管道
	// 想exitChan写入true
	exitChan <- true
	
	
}

func main(){
	intChan:=make(chan int, 1000)
	primeChan:=make(chan int, 20000)
	// 标识退出的Chan
	exitChan:=make(chan bool, 4)  // 4个

	start := time.Now().Unix()
	// 开启一个协程想intChan放入数据
	go putNum(intChan)

	// 开启四个协程从intChan取出数据,并判断是否为素数,如果是就放入到primeChan里面去
	// 放入到primeChan
	for i:=0;i<4;i++{
		go primeNum(intChan, primeChan, exitChan)
	}
	
	// 这里我们主线程要进行处理
	go func(){

		for i:=0;i<4;i++{
			<-exitChan
		}
		
		// 当我们从这个管道取出了4个结果就可以放心的关闭primeChan
		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end-start)
		close(primeChan)
		

		
	}()

	// 把结果取出
	for {
		_,ok := <-primeChan
		if !ok{
			break
		}
		// 将结果输出
		// fmt.Printf("素数=%d\n", res)
		
	}
	

}