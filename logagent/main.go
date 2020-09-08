package main

import (
	"fmt"
	_ "github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"logagent/etcd"
	"logagent/kafka"
	"logagent/tailfile"
)

type Config struct {
	KafkaConfig   `ini:"kafka"`
	CollectConfig `ini:"collect"`
	EtcdConfig    `ini:"etcd"`
}

type KafkaConfig struct {
	Address  string `ini:"address"`
	Topic    string `ini:"topic"`
	ChanSize int64  `ini:"chan_size"`
}

type EtcdConfig struct {
	Address    string `ini:"address"`
	CollectKey string `ini:"collect_key"`
}

type CollectConfig struct {
	LogFilePath string `ini:"logfile_path"`
}

// filebeat

// flument

// 真正的业务逻辑

func run() {
	select {}
}

// 从MsgChan中读取消息,发送给kafka

func main() {

	// 0.读取配置文件 `go-ini`
	var configObj = new(Config) // 得到结构体指针

	err := ini.MapTo(configObj, "./conf/conf.ini")

	if err != nil {
		logrus.Error("load config failed, err:%v", err)
		return

	}

	// 1. 初始化连接kafka
	err = kafka.Init([]string{configObj.KafkaConfig.Address}, configObj.KafkaConfig.ChanSize)

	if err != nil {
		logrus.Error("init kafka failed,  err:%v", err)
	}

	logrus.Info("init kafka success")

	// 2. 初始化etcd连接
	err = etcd.Init([]string{configObj.EtcdConfig.Address})

	if err != nil {
		logrus.Error("init etcd failed, err:%v", err)

	}

	logrus.Info("init etcd success")

	// 3. 从etcd中拉取要收集的日志的配置项, 发送到哪个topic
	allConf, err := etcd.GetConf(configObj.EtcdConfig.CollectKey)

	if err != nil {
		fmt.Println("get conf from etcd failed, err:%v", err)
		return

	}
	fmt.Println(allConf)

	// 派一个小弟去监控etcd configObj.EtcdConfig.CollectKey 对应值得变化
	go etcd.WatchConf(configObj.EtcdConfig.CollectKey)

	// 2. 根据配置中的日志路径初始化tail
	err = tailfile.Init(allConf) // 把从etcd 获取的配置项传到init中

	if err != nil {
		logrus.Error("init tailfile failed, err:%v", err)
		return
	}

	// 3. 把日志通过sarama发往kafka

	run()

}
