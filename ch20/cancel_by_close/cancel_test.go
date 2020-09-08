/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-16 09:09:45
 * @LastEditTime: 2019-08-16 09:50:36
 * @LastEditors: Please set LastEditors
 */
package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
	"context"
)

func isCancelled(ctx context.Context) bool{
	select{
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

// func isCancelled(cancelChan chan struct{}) bool{
// 	select{
// 	case <-cancelChan:
// 		return true
// 	default:
// 		return false
// 	}
// }

// func cancel_1(cancelChan chan struct{}){
// 	cancelChan <- struct{}{}
// }


// func cancel_2(cancelChan chan struct{}){
// 	close(cancelChan)
// }

func TestCancel(t *testing.T){
	// cancelChan := make(chan struct{},0)
	ctx, cancel := context.WithCancel(context.Background())
	for i:=0;i<5;i++ {
		go func(i int, ctx context.Context){
			for {
				if isCancelled(ctx){
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cacelled")
			
		}(i,ctx)
	}
	
	cancel()
	// cancel_1(cancelChan)
	time.Sleep(time.Second *1)
}