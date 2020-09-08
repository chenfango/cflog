package main   // 包，表明代码所在的模块

/*

应用程序入口
	1. 必须是main包: package main
    2. 必须是main方法: func main()
	3. 文件名不一定是main.go
	4. 父级目录不一定是main包名字
	5. go中的main函数不支持任何返回值,通过os.Exit来返回状态


*/

import (
	"fmt"
	"os"
) // 引入代码依赖


// 功能实现
func main() {
	if len(os.Args)>1{
		fmt.Println("Hello World")
	}

}
