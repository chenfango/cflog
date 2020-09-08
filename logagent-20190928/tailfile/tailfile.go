package tailfile

import (
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)


var (
	TailObj *tail.Tail
)
// tail相关

func Init(filename string)(err error){
	config := tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll: true,
	}
	// 打开文件开始读取数据
	TailObj, err =  tail.TailFile(filename, config)
	fmt.Println("创建tail对象", TailObj.Filename )
	if err != nil {
		logrus.Error("tailfile: create tailObj for path:%s failed, err:%v\n", filename, err)
		return
	}

	fmt.Println("init tailfile success")
	return
}
