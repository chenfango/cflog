package main

import (
	"fmt"
)

// 一个被测试的函数
func addUpper(n int) int {
	res := 0 
	for i:=1;i<=n;i++{
		res += i
	}
	return res
}

func main(){
	// 传统方法测试，就是在main看看结果是否正确
	res := addUpper(10)
	if res != 55{
		fmt.Printf("addUpper错误 返回值=%v 期望值=%v\n", res, 55)
	}else{
		fmt.Printf("addUpper正确 返回值=%v 期望值=%v\n", res, 55)
	}
}