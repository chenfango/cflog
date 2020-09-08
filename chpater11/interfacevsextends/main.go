package main

import (
	"fmt"
)

// Monkey结构体
type Monkey struct {
	Name string
}

func (this *Monkey) climbing(){
	fmt.Println(this.Name, "生来会爬树")

}

// 声明接口
type BirdAble interface{
	Flying()
}

type FishAble interface{
	Swimming()
}



// LittleMonkey

type LittleMonkey struct{
	Monkey // 继承
}

// 让LitteMoney实现
func (this *LittleMonkey) Flying(){
	fmt.Println(this.Name, "通过学习会飞翔")
}

func (this *LittleMonkey) Swimming(){
	fmt.Println(this.Name, "通过学习会游泳")
}


func main(){
	// 创建一个LitterMonkey实例
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}


	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}