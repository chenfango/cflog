package main

import (
	"fmt"
)

// 编写一个学生考试系统
type Student struct{
	Name string
	Age int
	Score int
}

// 将Pupil和Graduate共有的方法也绑定到Student类型
func (stu *Student) ShowInfo(){
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int){
	// 业务判断
	stu.Score = score
}

// 给 *Student增加一个方法，那么 Pupil和Graduate都可以使用该方法
func (stu *Student) getSum(n1 int, n2 int) int{ 
	// 业务判断
	return n1+n2
}

// 小学生
type Pupil struct{
	// 匿名结构体字段, 继承属性和方法
	Student
}

// 特有的方法保留
func (p *Pupil) testing(){
	fmt.Println("小学生正在考试中...")
}



func main(){
	// 当我们对结构体嵌入了匿名的结构体后，使用的方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom~"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()
	fmt.Println("res=", pupil.getSum(1,2))

}
