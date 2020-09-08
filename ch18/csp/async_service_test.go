/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 15:46:01
 * @LastEditTime: 2019-08-15 16:00:33
 * @LastEditors: Please set LastEditors
 */

 package csp
 
 import (
	 "fmt"
	 "time"
	 "testing"
 )


 func service() string {
	 time.Sleep(time.Millisecond *50)
	 return "Done"
 }



 func otherTask(){
	 fmt.Println("working on something else")
	 time.Sleep(time.Millisecond *100)
	 fmt.Println("test is done")
 }


 func TestService(t *testing.T){
	 fmt.Println(service())
	 otherTask()
 }


 func AsyncService() chan string {

	 retCh := make(chan string, 1)
	 //retCh := make(chan string)

	 return retCh
 }



 func TestAsyncService(t *testing.T){
	 retCh := AsyncService()
	 t.Logf()
	 
 }