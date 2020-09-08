package constant_test

import "testing"

const(
	Monday = iota+1
	Tuesday
	Wednesday
)

const(
	Readable =1 <<iota
	Writeable
	Executable
)

func TestContantTry(t *testing.T)  {
	t.Log(Monday,Tuesday)

}

func TestConstanTry(t *testing.T){
	a:=1 // 0001
	t.Log(a&Readable==Readable, a&Writeable==Writeable, a&Executable==Executable)
}
