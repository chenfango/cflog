/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-13 21:50:11
 * @LastEditTime: 2019-08-14 09:04:04
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"
	"time"
	"sync"
)

// 思路
// 1. 编写一个函数，来计算各个数的阶层，并放到一个map中
// 2. 我们启动的协程是多个，将统计的结果放入到map中
// 3. map应该做成一个全局的


var (
	myMap = make(map[int]int, 10)
	// 声明一个全局的互斥锁
	// lock是一个全局的互斥锁(写锁), 结构体变量
	// sync是包: 同步的意思
	// Mutex: 互斥的意思，是一个结构体

	lock sync.Mutex
)

// test函数是计算n的阶层,将这个结果放入到map中
func test(n int){
	res:=1
	for i:=1; i<=n; i++{
		res *= i 
	}
	// 这里我们将res放入myMap中
	// 加锁
	lock.Lock()
	myMap[n]=res
	// 解锁
	lock.Unlock()

}




func main(){

	// 我们这里开启多个协层完成这个任务,启动200个
	for i:=1;i<=20;i++{
		go test(i)
	}

	// 休眠10s钟
	time.Sleep(time.Second*10)
	
	// 这里我们输出结果, 遍历这个结果
	lock.Lock()
	for i,v := range myMap{
		fmt.Printf("map[%d]=%d\n", i,v )
	}
	lock.Unlock()
	
}