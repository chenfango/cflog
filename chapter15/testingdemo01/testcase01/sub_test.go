package testcase01

import (
	_ "fmt"
	"testing"  // 引入testing测试框架
)

// 编写测试用例，去测试addUpper函数是否正确
func TestGetSub(t *testing.T){
	// 调用
	res := getSub(10,3)
	if res!=7{
		// fmt.Printf("addUpper错误 返回值=%v 期望值=%v\n", res, 55)
		t.Fatalf("getSub错误 返回值=%v 期望值=%v\n", res, 7)
	}else{
		// 如果正确，输出日志
		t.Logf("getSub正确 返回值=%v 期望值=%v\n", res, 7)
	}
}