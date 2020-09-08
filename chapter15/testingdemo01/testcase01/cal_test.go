package testcase01

import (
	_ "fmt"
	"testing"  // 引入testing测试框架
)

// 编写测试用例，去测试addUpper函数是否正确
func TestAddUpper(t *testing.T){
	// 调用
	res := addUpper(10)
	if res!=55{
		// fmt.Printf("addUpper错误 返回值=%v 期望值=%v\n", res, 55)
		t.Fatalf("addUpper错误 返回值=%v 期望值=%v\n", res, 55)
	}else{
		// 如果正确，输出日志
		t.Logf("addUpper正确 返回值=%v 期望值=%v\n", res, 55)
	}
}