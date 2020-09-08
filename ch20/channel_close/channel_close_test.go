/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-16 08:50:14
 * @LastEditTime: 2019-08-16 09:05:51
 * @LastEditors: Please set LastEditors
 */
package channel_close

import (
	"fmt"
	"sync"
	"testing"
)


func dataProducer(ch chan int, wg *sync.WaitGroup){
	go func(){
		for i:=0;i<10;i++{
			ch <-i
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup){
	go func(){
		for {
			// channel数据如果接受完了,会返回flase
			if data,ok := <-ch;ok{
				fmt.Println(data)
			}else{
				break
			}
		}
		wg.Done()
	}()

}

func TestCloseChannel(t *testing.T){
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()

	
	
}
