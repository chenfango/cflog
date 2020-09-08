/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-13 20:33:18
 * @LastEditTime: 2019-08-13 21:48:06
 * @LastEditors: Please set LastEditors
 */
 package main
 
 import (
	 "fmt"
	 "runtime"
 )
 
 
 func main(){
	 cpuNum := runtime.NumCPU()
	 fmt.Println("cpuNum=", cpuNum)
	 // 可以设置自己使用多少个cpu
	 runtime.GOMAXPROCS(cpuNum-1)
	 
 }
 
