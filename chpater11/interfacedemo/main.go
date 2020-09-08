package main

import (
	"fmt"
)

// 声明一个接口
type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct{

} 

// 让手机实现Usb接口的方法
func (p Phone) Start(){
	fmt.Println("手机开始工作了")
}

func (p Phone) Stop(){
	fmt.Println("手机停止工作了")
}

func (p Phone) Call(){
	fmt.Println("手机开始打电话")
}


type Camera struct{

} 

// 让相机实现Usb接口
func (c Camera) Start(){
	fmt.Println("相机开始工作了")
}

func (c Camera) Stop(){
	fmt.Println("相机停止工作了")
}


// 计算机
type Computer struct{
} 


// 编写一个方法working, 这个方法接受Usb接口类型，只要是实现了Usb接口类型的变量
// 只要实现了Usb接口，所谓实现Usb接口，就是指实现了Usb接口声明的所有方法

func (c Computer) Working(usb Usb){
	// 通过usb接口变量来调用start和stop
	usb.Start()
	// 如果usb是指向一个Phone结构体变量的,还需要使用Call方法
	// 类型断言...[注意体会好处]
	if phone,ok := usb.(Phone); ok==true{
		phone.Call()
	}

	usb.Stop()
}



func main(){

	// 测试
	// 先创建结构体变量cd 
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	// 关键点
	computer.Working(phone)
	computer.Working(camera)

	
	
}