package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	_ "github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"logagent/kafka"
	"logagent/tailfile"
	"strings"
	"time"
)


type Config struct {
	KafkaConfig `ini:"kafka"`
	CollectConfig `ini:"collect"`
}


type KafkaConfig struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
	ChanSize int64 `ini:"chan_size"`

}

type CollectConfig struct {
	LogFilePath string `ini:"logfile_path"`

}



// filebeat

// flument

// 真正的业务逻辑




// run 真正的业务逻辑
func run()(err error){

	for {
		// 循环读数据
		line, ok := <- tailfile.TailObj.Lines// chan tail.Line
		if !ok {
			fmt.Println("hahahh")
			logrus.Warn("tail file close reopen, filename:%s\n", tailfile.TailObj.Filename)
			time.Sleep(time.Second) // 读取出错等一秒
			continue
		}

		if len(strings.Trim(line.Text, "\r")) == 0 {
			logrus.Info("出现空行拉,直接跳过...")
			continue
		}
		// 利用通道将同步的代码改为异步的
		// 把line 读出的日志，包装成kafka中的msg类型，丢到通道中
		msg := &sarama.ProducerMessage{}
		msg.Topic = "web_log"
		msg.Value = sarama.StringEncoder(line.Text)

		// 丢到管道中
		kafka.MsgChan <- msg




	}

}


// 从MsgChan中读取消息,发送给kafka



func main()  {

	// 0. 读取配置
	var configObj = new(Config)  // 得到结构体指针

	err := ini.MapTo(configObj, "./conf/conf.ini")


	if err!=nil{
		logrus.Error("load config failed, err:%v", err)

	}

	fmt.Printf("%#v\n", configObj)



	// 1. 初始化连接kafka
	kafka.Init([]string{configObj.KafkaConfig.Address}, configObj.KafkaConfig.ChanSize)


	// 2. 根据配置中的日志路径初始化tail
	err = tailfile.Init(configObj.CollectConfig.LogFilePath)

	if err != nil {
		logrus.Error("init tailfile failed, err:%v", err)
		return
	}


	// 3. 把日志通过sarama发往kafka
	err = run()


	if err!=nil{
		logrus.Error("run failed")
	}



}
