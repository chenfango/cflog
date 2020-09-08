/*
 * @Description: goroutine
 * @Author: chenfan
 * @Date: 2019-08-13 19:50:39
 * @LastEditTime: 2019-08-13 20:04:52
 * @LastEditors: Please set LastEditors
 */

package main

import (
	"fmt"
	"strconv"
	"time"
)

// 编写一个函数,该函数每隔一秒输出"Hello World"
func test(){
	for i:=1;i<=10;i++{
		fmt.Println("test() hello, World "+ strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main(){
	go test()
	for i:=1;i<=10;i++{
		fmt.Println("main() hello, World "+ strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}



