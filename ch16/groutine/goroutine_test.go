/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 15:10:44
 * @LastEditTime: 2019-08-15 15:16:29
 * @LastEditors: Please set LastEditors
 */
package groutine_test
import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T){
	for i:=0;i<10;i++{
		go func(){
			fmt.Println(i) 
		}()
	}

	time.Sleep(time.Millisecond*5)
}