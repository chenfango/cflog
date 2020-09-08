/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 16:56:39
 * @LastEditTime: 2019-08-15 16:56:44
 * @LastEditors: Please set LastEditors
 */
package select_test

/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 15:46:01
 * @LastEditTime: 2019-08-16 08:46:34
 * @LastEditors: Please set LastEditors
 */

 
 import (
	 "fmt"
	 "time"
	 "testing"
 )


 func service() string {
	 time.Sleep(time.Millisecond *500)
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
	 // 多路选择
	 select {
	 	case retCh := <-AsyncService():
			t.Log(retCh)
		case <- time.After(time.Millisecond*100):
			t.Error("timeout")

	//  otherTask()
	//  fmt.Println(<-retCh)
	 }
	 
 }