package testcase01

// 测试单个
// go test -v -test.run TestAddUpper



import (
	_ "fmt"
)

// 一个被测试的函数
func addUpper(n int) int {
	res := 0 
	for i:=1;i<=n;i++{
		res += i
	}
	return res
}

func getSub(n1 int, n2 int) int{
	return n1 - n2
}