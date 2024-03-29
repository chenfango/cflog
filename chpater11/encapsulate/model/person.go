package model

import (
	"fmt"
)
type person struct{
	Name string
	age int // 其他包不可导出
	sla float64

}

// 写一个工厂模式的函数, 相当于构造函数
func NewPerson(name string ) *person{
	return &person{
		Name: name,
	}
}

// 为了访问age和sal 我们编写一对SetXxx方法和GetXxx方法
func (p *person) SetAge(age int){
	if age >0 && age <150{
		p.age=age
	} else{
		fmt.Println("年龄范围不正确")
		// 给程序员一个默认值

	}
}


func (p *person) GetAge() int{
	return p.age
}


func (p *person) SetSal(sal float64){
	if sal  >3000 && sal <= 30000{
		p.sla =sal
	} else{
		fmt.Println("薪水范围不正确")
		// 给程序员一个默认值

	}  
}

func (p *person) GetSal() float64{
	return p.sla
}


