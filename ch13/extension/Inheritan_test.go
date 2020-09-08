package extension

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}


type GoProgrammer struct {

}

func (p *GoProgrammer) WriteHelloWorld() Code{
	return "fmt.Print(\"hello world\")"
}


type JavaProgrammer struct {

}

func (p *JavaProgrammer) WriteHelloWorld() Code{
	return "system.out.Println(\"hello world\")"
}

func WriteFirstProgram(p Programmer){
	fmt.Printf("%T %v \n",p,p.WriteHelloWorld())
}

func TestPloymorphism(t *testing.T){
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	WriteFirstProgram(goProg)
	WriteFirstProgram(javaProg)
}
