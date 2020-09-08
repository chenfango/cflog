package _interface

import (
	"testing"
)

// 类型

// 接口
type Programmer interface {
	WriteHelloWorld() string
}


// 实现
type GoProgrammer struct {

}


func (g *GoProgrammer) WriteHelloWorld() string{
	return  "hello world"
}


func TestClient(t *testing.T){
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())

}

