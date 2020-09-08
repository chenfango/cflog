package testcase02

import (
	"fmt"
	"encoding/json"
	"io/ioutil"

)

type Monster struct {
	Name string
	Age int
	Skill string
}

// 给Monstart绑定Store方法，可以将一个Monster变量(对象),将序列化后保存到文件中
func (this *Monster) Store() bool{
	// 先序列化
	data, err := json.Marshal(this)
	if err !=nil{
		fmt.Println("序列化失败了", err)
		return false
	}

	// 保存到文件
	filePath :="./monster.ser"
	err=ioutil.WriteFile(filePath, data, 777)
	if err != nil{
		fmt.Println("Write file err", err)
		return false
	}
	return true	

}

func (this *Monster) Restore() bool{
	// 1.先从文件中读取序列化后的字符串
	filePath:="./monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil{
		return false
	}

	// 2. 使用读取到的data []byte, 对反序列
	err =  json.Unmarshal(data, this)
	if err != nil{
		fmt.Println("UnMarshal err=", err)
		return false
	}
	return true

}



