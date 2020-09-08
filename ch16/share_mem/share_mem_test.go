/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-07-28 11:16:07
 * @LastEditTime: 2019-08-15 15:33:02
 * @LastEditors: Please set LastEditors
 */
package share_mem

import (
	"sync"
	"testing"
	"time"
)



func TestCounter(t *testing.T){
	counter := 0

	// counter在协程中竞争

	for i:=0; i<5000; i++{
		go func(){
			counter++
		}()
	}

	time.Sleep(1*time.Second)
	t.Logf("counter = %d" , counter)
}


func TestCounterThreadSafe(t *testing.T){
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	// counter在协程中竞争

	for i:=0; i<5000; i++{
		wg.Add(1)
		go func(){
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()

		}()
	}

	// 阻塞在这里
	wg.Wait()

	// 等所有协程执行完再退出
	//time.Sleep(1*time.Second)
	t.Logf("counter = %d" , counter)
}
