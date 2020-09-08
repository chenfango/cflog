/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-16 09:52:06
 * @LastEditTime: 2019-08-16 10:21:05
 * @LastEditors: Please set LastEditors
 */
package singleton

import (
	"sync"
	"fmt"
	"testing"
	"unsafe"
)

type Singleton struct {
	
}

var singleInstance *Singleton

var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(
		func(){
			fmt.Println("create obj")
			singleInstance = new(Singleton)
		})
		
	return singleInstance
}


func TestGetSingletonObj(t *testing.T){
	var wg sync.WaitGroup
	for i:=0;i<10;i++{
		wg.Add(1)
		go func(){
			obj:=GetSingletonObj()
			fmt.Printf("%d\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}


