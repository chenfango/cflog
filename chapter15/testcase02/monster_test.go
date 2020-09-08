package testcase02

import (
	"testing"
	
)

// 测试用例
func TestStore(t *testing.T){
	// 先创建monster对象
	monster := Monster{
		Name: "红孩儿",
		Age: 11,
		Skill: "吐火", 

	}
	res := monster.Store()
	if !res{
		t.Fatalf("这个函数测试错误了, 希望为=%v, 实际=%v", true, res)
	}else{
		t.Logf("monster Store() 测试成功")
	}
}


func TestRestore(t *testing.T){
	// 先创建monster对象, 不指定字段的值
	var monster Monster
	res := monster.Restore()
	t.Log(res)
	if !res{
		t.Fatalf("monster Restore 错误了, 希望为=%v, 实际=%v", true, res)

	}
	if monster.Name != "红孩儿"{
		t.Fatalf("monster Restore 错误了, 希望为=%v, 实际=%v", "红孩儿", monster.Name)
	}
	
	t.Logf("monster ReStore() 测试成功")

}